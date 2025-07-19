package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

func RegisterRoutes(router fiber.Router, db *gorm.DB, store *session.Store) {
	controller := NewAuthController(db, store)

	router.Get("/login", controller.ShowLoginForm)
	router.Post("/login", controller.HandleLogin)
	router.Get("/logout", controller.HandleLogout)
	router.Get("/profile", controller.ShowProfile)
}
