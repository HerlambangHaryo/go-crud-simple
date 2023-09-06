package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	
	"github.com/HerlambangHaryo/go-crud-simple/platform/database"
	"github.com/HerlambangHaryo/go-crud-simple/app/controllers"
)

func main() {
	connectToDatabase()
	app := fiber.New()
	app.Use(logger.New())
	registerRoutes(app)
	log.Fatal(app.Listen(":8000"))
}

func connectToDatabase() {
	database.Connect()
}

func registerRoutes(app *fiber.App) {
	// ...
}
		
	
		
