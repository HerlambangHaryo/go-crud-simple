package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/HerlambangHaryo/go-crud-simple/app/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/v1/books", controllers.GetBooks)
	app.Get("/api/v1/books/:id", controllers.GetBook)
	app.Post("/api/v1/books", controllers.CreateBook)
	app.Put("/api/v1/books/:id", controllers.UpdateBook)
	app.Delete("/api/v1/books/:id", controllers.DeleteBook)
}
