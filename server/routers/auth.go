package routers

import (
	"zerodesk/handlers"

	"github.com/gofiber/fiber/v2"
)


func AuthRoutes(router fiber.Router) {
	router.Post("/sign-up",handlers.SignUp)
	router.Post("/sign-in",handlers.SignIn)
}