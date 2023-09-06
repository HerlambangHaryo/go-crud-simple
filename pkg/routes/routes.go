package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/HerlambangHaryo/go-crud-simple/app/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", controllers.GetDashboard)
	// app.Get("/api/v1/books", controllers.GetBooks)
	// app.Get("/api/v1/books/:id", controllers.GetBook)
	// app.Post("/api/v1/books", controllers.CreateBook)
	// app.Put("/api/v1/books/:id", controllers.UpdateBook)
	// app.Delete("/api/v1/books/:id", controllers.DeleteBook)
 

	app.Get("/Books", controllers.GetBooks)
	app.Get("/Books/:id", controllers.GetBook)
	app.Post("/Books", controllers.CreateBook)
	app.Put("/Books/:id", controllers.UpdateBook)
	app.Delete("/Books/:id", controllers.DeleteBook)

	
	app.Get("/Rapidapis", controllers.GetRapidapis)
	app.Get("/Rapidapis/:id", controllers.GetRapidapi)
	app.Post("/Rapidapis", controllers.CreateRapidapi)
	app.Put("/Rapidapis/:id", controllers.UpdateRapidapi)
	app.Delete("/Rapidapis/:id", controllers.DeleteRapidapi)
}
