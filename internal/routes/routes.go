// internal/routes/routes.go: setup all the protected or not routes, and its connections with handlers, services and repository
package routes

import (
	"github.com/gofiber/fiber/v3"

	"github.com/kauanpecanha/odsquiz-auth/internal/handlers"
	"github.com/kauanpecanha/odsquiz-auth/internal/middleware"
	"github.com/kauanpecanha/odsquiz-auth/internal/repositories"
	"github.com/kauanpecanha/odsquiz-auth/internal/services"
	"github.com/kauanpecanha/odsquiz-auth/pkg/database"
)

func Setup(app *fiber.App) {

	// setup User Repository to interact with database
	userRepo := repositories.NewRepo(database.DB.Db)

	// setup User Service to manage the users business rules
	userService := &services.UserService{
		Repo: userRepo,
	}

	// setup User Handler to manage the users HTTP requests and responses
	userHandler := &handlers.UserHandler{
		Service: userService,
	}

	app.Get("/", func(c fiber.Ctx) error { return c.SendString("Hello, world!") })
	app.Get("/health", func(c fiber.Ctx) error { return c.SendString("ok") })

	// unprotected routes (can be used without bearer token)
	app.Post("/createUser", userHandler.CreateUser)
	app.Post("/login", userHandler.LoginUser)
	// protected routes (can be used ONLY with bearer token)
	app.Get("/getAllUsers", middleware.Protected(), userHandler.GetAllUsers)
	app.Get("/getUserById/:id", middleware.Protected(), userHandler.GetUserByID)
	app.Patch("/updateUser/:id", middleware.Protected(), userHandler.UpdateUser)
	app.Delete("/deleteUser/:id", middleware.Protected(), userHandler.DeleteUser)

	app.Use(func(c fiber.Ctx) error { return c.SendStatus(fiber.StatusNotFound) })
}
