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

	userRepo := repositories.NewRepo(database.DB.Db)

	userService := &services.UserService{
		Repo: userRepo,
	}

	userHandler := &handlers.UserHandler{
		Service: userService,
	}

	// root route
	app.Get("/", func(c fiber.Ctx) error { return c.SendString("Hello, world!") })
	// app health verification handler
	app.Get("/health", func(c fiber.Ctx) error { return c.SendString("ok") })

	app.Post("/createUser", userHandler.CreateUser)
	app.Post("/login", userHandler.LoginUser)
	app.Get("/getAllUsers", middleware.Protected(), userHandler.GetAllUsers)
	app.Get("/getUserById/:id", middleware.Protected(), userHandler.GetUserByID)
	app.Patch("/updateUser/:id", middleware.Protected(), userHandler.UpdateUser)
	app.Delete("/deleteUser/:id", middleware.Protected(), userHandler.DeleteUser)

	// 404 handler
	app.Use(func(c fiber.Ctx) error { return c.SendStatus(fiber.StatusNotFound) })
}
