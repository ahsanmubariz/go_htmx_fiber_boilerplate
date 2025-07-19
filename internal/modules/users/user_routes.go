package users

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(router fiber.Router, db *gorm.DB) {
	service := NewUsersService(db)
	controller := NewUsersController(service)

	usersGroup := router.Group("/users")

	usersGroup.Get("/", controller.ListUsers)
	usersGroup.Post("/", controller.AddUser)
	usersGroup.Delete("/:id", controller.DeleteUser)
}
