package auth

import (
	"log"

	"github.com/ahsanmubariz/go_htmx_fiber_boilerplate/internal/modules/users"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

type AuthController struct {
	db    *gorm.DB
	store *session.Store
}

func NewAuthController(db *gorm.DB, store *session.Store) *AuthController {
	return &AuthController{db: db, store: store}
}

func (ac *AuthController) ShowLoginForm(c *fiber.Ctx) error {
	return c.Render("pages/auth/login", fiber.Map{
		"Title":      "Login",
		"CSRFToken":  c.Locals("CSRFToken"),
		"IsLoggedIn": c.Locals("IsLoggedIn"),
		"Username":   c.Locals("Username"),
	}, "layouts/base")
}

func (ac *AuthController) HandleLogin(c *fiber.Ctx) error {
	type LoginRequest struct {
		Email string `form:"email"`
	}
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).SendString("Invalid request")
	}

	var user users.User
	if err := ac.db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return c.Status(401).SendString("Invalid credentials")
	}

	sess, err := ac.store.Get(c)
	if err != nil {
		return c.Status(500).SendString("Session error")
	}

	sess.Set("user_id", user.ID)
	sess.Set("username", user.Name)

	if err := sess.Save(); err != nil {
		log.Printf("ERROR: Failed to save session: %v", err)
		return c.Status(500).SendString("Failed to save session")
	}

	c.Set("HX-Redirect", "/profile")
	return c.SendStatus(200)
}

func (ac *AuthController) ShowProfile(c *fiber.Ctx) error {
	sess, err := ac.store.Get(c)
	if err != nil {
		return c.Status(500).SendString("Session error")
	}

	userID := sess.Get("user_id")
	if userID == nil {
		return c.Redirect("/login")
	}

	return c.Render("pages/auth/profile", fiber.Map{
		"Title":      "Your Profile",
		"CSRFToken":  c.Locals("CSRFToken"),
		"IsLoggedIn": c.Locals("IsLoggedIn"),
		"Username":   c.Locals("Username"),
	}, "layouts/base")
}

func (ac *AuthController) HandleLogout(c *fiber.Ctx) error {
	sess, err := ac.store.Get(c)
	if err != nil {
		return c.Status(500).SendString("Session error")
	}

	if err := sess.Destroy(); err != nil {
		log.Printf("ERROR: Failed to destroy session: %v", err)
		return c.Status(500).SendString("Failed to destroy session")
	}

	return c.Redirect("/")
}
