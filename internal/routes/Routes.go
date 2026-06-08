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
	repository := repositories.NewRepo(database.DB.Db)
	service := &services.Service{
		Repo: repository,
	}
	handler := &handlers.Handler{
		Service: service,
	}

	app.Get("/", func(c fiber.Ctx) error { return c.SendString("Hello, world!") })
	app.Get("/health", func(c fiber.Ctx) error { return c.SendString("ok") })
	// unprotected routes (can be used without bearer token)
	app.Post("/signup", handler.CreateOne)
	app.Post("/login", handler.Login)
	// protected routes (can be used ONLY with bearer token)
	app.Get("/getAllOnes", middleware.Protected(), handler.GetAllOnes)
	app.Get("/getOneById/:id", middleware.Protected(), handler.GetOneByID)
	app.Patch("/updateOne/:id", middleware.Protected(), handler.UpdateOne)
	app.Delete("/deleteOne/:id", middleware.Protected(), handler.DeleteOne)

	app.Use(func(c fiber.Ctx) error { return c.SendStatus(fiber.StatusNotFound) })
}
