package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	
	"github.com/<username>/go-crud-simple/platform/database"
	"github.com/<username>/go-crud-simple/app/controllers"
)

	func getMySQLConnection() string {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		dbname := os.Getenv("DB_NAME")

		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
	}

	type Book struct {
		ID   uint   `json:"id" gorm:"primaryKey"`
		Name string `json:"name"`
	}

	var db *gorm.DB

	func connectToDatabase() {
		var err error
		db, err = gorm.Open(mysql.Open(getMySQLConnection()), &gorm.Config{})
		if err != nil {
			panic("Failed to connect to database")
		}
	
		fmt.Println("Database connection successfully opened")
	
		db.AutoMigrate(&Book{})
		fmt.Println("Database migrated")
	} 

	// Get all books
	func getBooks(c *fiber.Ctx) error {
		var books []Book
		db.Find(&books)
		return c.JSON(books)
	}

	// Get a book by id
	func getBook(c *fiber.Ctx) error {
		id := c.Params("id")
		var book Book
		db.First(&book, id)
		if book.ID == 0 {
			return c.Status(404).SendString("No book found with given ID")
		}
		return c.JSON(book)
	}

	// Create a book
	func createBook(c *fiber.Ctx) error {
		book := new(Book)
		if err := c.BodyParser(book); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		db.Create(&book)
		return c.JSON(book)
	}

	// Update a book by id
	func updateBook(c *fiber.Ctx) error {
		id := c.Params("id")
		var book Book
		db.First(&book, id)
		if book.ID == 0 {
			return c.Status(404).SendString("No book found with given ID")
		}
		if err := c.BodyParser(&book); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		db.Save(&book)
		return c.JSON(book)
	}

	// Delete a book by id
	func deleteBook(c *fiber.Ctx) error {
		id := c.Params("id")
		var book Book
		db.First(&book, id)
		if book.ID == 0 {
			return c.Status(404).SendString("No book found with given ID")
		}
		db.Delete(&book)
		return c.SendString("Book successfully deleted")
	}

	func registerRoutes(app *fiber.App) {
		app.Get("/api/v1/books", getBooks)
		app.Get("/api/v1/books/:id", getBook)
		app.Post("/api/v1/books", createBook)
		app.Put("/api/v1/books/:id", updateBook)
		app.Delete("/api/v1/books/:id", deleteBook)
	}

	func main() {
		connectToDatabase()
		app := fiber.New()
		app.Use(logger.New())
		registerRoutes(app)
		log.Fatal(app.Listen(":8000"))
	}

		
	
		
