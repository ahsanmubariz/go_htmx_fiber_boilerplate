package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/helmet/v2"
	"github.com/gofiber/storage/postgres"
	"github.com/gofiber/template/html/v2"
	"github.com/gorilla/csrf"
	"gorm.io/gorm"

	"github.com/ahsanmubariz/go_htmx_fiber_boilerplate/internal/config"
	"github.com/ahsanmubariz/go_htmx_fiber_boilerplate/internal/modules/auth"
	"github.com/ahsanmubariz/go_htmx_fiber_boilerplate/internal/modules/home"
	"github.com/ahsanmubariz/go_htmx_fiber_boilerplate/internal/modules/users"
)

type Server struct {
	App *fiber.App
	Cfg *config.Config
	DB  *gorm.DB
}

func NewServer(cfg *config.Config, db *gorm.DB) *Server {
	engine := html.New("./web/template", ".html")
	if cfg.Env == "development" {
		engine.Reload(true)
	}

	// 1. Create the "internal" Fiber app where we define our routes and logic.
	internalApp := fiber.New(fiber.Config{
		Views: engine,
	})

	storage := postgres.New(postgres.Config{
		ConnectionURI: cfg.DSN,
		Table:         "fiber_sessions",
	})
	store := session.New(session.Config{
		Storage:      storage,
		KeyLookup:    "cookie:session_id",
		CookieSecure: cfg.Env == "production",
		Expiration:   24 * time.Hour,
	})

	// 2. Register middleware on the internal app.
	// THE FIX: This middleware MUST be registered BEFORE the routes.
	internalApp.Use(func(c *fiber.Ctx) error {
		sess, _ := store.Get(c)
		c.Locals("IsLoggedIn", sess.Get("user_id") != nil)
		c.Locals("Username", sess.Get("username"))

		// Read the token from the special header set by our bridge middleware.
		token := c.Get("X-CSRF-Token-View")
		c.Locals("CSRFToken", token)
		log.Printf("DEBUG: [Fiber Middleware] CSRF Token set in locals: '%s'", token)
		return c.Next()
	})
	// Register routes AFTER the middleware.
	RegisterRoutes(internalApp, db, store)

	// 3. Create the standard net/http middleware chain.
	csrfMiddleware := csrf.Protect(
		[]byte(cfg.CSRFSecret),
		csrf.Secure(cfg.Env == "production"),
		csrf.Path("/"),
		csrf.HttpOnly(true),
		csrf.RequestHeader("X-CSRF-Token"), // For HTMX
	)

	// Convert our internal Fiber app to a standard http.Handler.
	internalHttpHandler := adaptor.FiberApp(internalApp)

	// Create the "bridge" middleware. It runs after Gorilla CSRF and before our Fiber app.
	tokenPasser := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the token that Gorilla CSRF just generated.
		token := csrf.Token(r)
		// Put it in a header so the internal Fiber app can read it.
		r.Header.Set("X-CSRF-Token-View", token)
		// Call the next handler in the chain (our Fiber app).
		internalHttpHandler.ServeHTTP(w, r)
	})

	// Chain the http handlers together: CSRF -> Token Bridge -> Fiber App
	finalHandler := csrfMiddleware(tokenPasser)

	// 4. Create the "public" Fiber app that will serve the entire chain.
	publicApp := fiber.New()
	publicApp.Use(recover.New(), logger.New(), helmet.New())
	publicApp.Static("/", "./public")
	// All requests to the public app are passed to our http handler chain.
	publicApp.All("/*", adaptor.HTTPHandler(finalHandler))

	return &Server{
		App: publicApp, // We return the public-facing app to main.go
		Cfg: cfg,
		DB:  db,
	}
}

func RegisterRoutes(router fiber.Router, db *gorm.DB, store *session.Store) {
	api := router.Group("/")
	home.RegisterRoutes(api)
	users.RegisterRoutes(api, db)
	auth.RegisterRoutes(api, db, store)
}

func (s *Server) Start() error {
	log.Println("Registering routes...")
	return s.App.Listen(fmt.Sprintf(":%s", s.Cfg.Port))
}
