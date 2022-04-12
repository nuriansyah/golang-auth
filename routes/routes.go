package routes

import (
	"golang-auth/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/v1/auth/register", controllers.Register)
	app.Post("/api/v1/auth/login", controllers.Login)
	app.Get("/api/v1/auth/user", controllers.User)
	app.Post("/api/v1/auth/logout", controllers.Logout)

}
