package home

import (
	"github.com/gofiber/fiber/v2"
)

type HomeController struct{}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (h *HomeController) GetIndex(c *fiber.Ctx) error {
	return c.Render("pages/index", fiber.Map{
		"Title":      "Welcome to Go + Fiber + HTMX!",
		"CSRFToken":  c.Locals("CSRFToken"),
		"IsLoggedIn": c.Locals("IsLoggedIn"),
		"Username":   c.Locals("Username"),
	}, "layouts/base")
}
