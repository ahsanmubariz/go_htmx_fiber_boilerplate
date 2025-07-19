package users

import (
	"github.com/ahsanmubariz/go_htmx_fiber_boilerplate/internal/validator"

	"github.com/gofiber/fiber/v2"
)

type UsersController struct {
	service *UsersService
}

func NewUsersController(service *UsersService) *UsersController {
	return &UsersController{service: service}
}

func (uc *UsersController) ListUsers(c *fiber.Ctx) error {
	users, err := uc.service.GetAllUsers()
	if err != nil {
		return c.Status(500).SendString("Error fetching users")
	}

	return c.Render("pages/users/index", fiber.Map{
		"Title":      "Users Management",
		"Users":      users,
		"CSRFToken":  c.Locals("CSRFToken"),
		"IsLoggedIn": c.Locals("IsLoggedIn"),
		"Username":   c.Locals("Username"),
	}, "layouts/base")
}

func (uc *UsersController) AddUser(c *fiber.Ctx) error {
	var user User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).SendString("Invalid user data")
	}

	if err := validator.ValidateStruct(user); err != nil {
		return c.Status(400).SendString("Validation failed: " + err.Error())
	}

	if err := uc.service.CreateUser(&user); err != nil {
		return c.Status(500).SendString("Error creating user")
	}

	return c.Render("partials/users/user-row", user)
}

func (uc *UsersController) DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString("Invalid ID")
	}

	if err := uc.service.DeleteUser(uint(id)); err != nil {
		return c.Status(500).SendString("Error deleting user")
	}

	return c.SendString("")
}
