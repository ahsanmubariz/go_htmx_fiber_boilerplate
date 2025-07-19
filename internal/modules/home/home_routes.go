package home

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router) {
	homeController := NewHomeController()
	router.Get("/", homeController.GetIndex)
}
