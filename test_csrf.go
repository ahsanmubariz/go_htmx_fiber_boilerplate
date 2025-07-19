// FILE: test_csrf.go
package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// 1. Create a new Fiber app with the template engine
	app := fiber.New(fiber.Config{
		Views: html.New("./web/template", ".html"),
	})

	// 2. Use ONLY the CSRF middleware with its simplest configuration
	app.Use(csrf.New())

	// 3. Create a single route to test the token
	app.Get("/", func(c *fiber.Ctx) error {
		// 4. Get the token from locals
		token := c.Locals("csrf")

		// 5. Log what we found
		log.Printf("CSRF Token from locals is: '%v'", token)
		if token == nil {
			log.Println("ERROR: Token is nil!")
		} else {
			log.Println("SUCCESS: Token was generated!")
		}

		// 6. Render a simple page
		return c.Render("pages/index", fiber.Map{
			"Title":     "CSRF Test",
			"CSRFToken": token,
		})
	})

	log.Println("Starting CSRF test server on :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
