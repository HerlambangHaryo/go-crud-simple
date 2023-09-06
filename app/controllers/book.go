// book.go
package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/<username>/go-crud-simple/platform/database"
	"github.com/<username>/go-crud-simple/app/models" 
)

// Get all books
func GetBooks(c *fiber.Ctx) error {
	var books []models.Book
	database.DB.Find(&books)
	return c.JSON(books)
}

// Get a book by id
func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	database.DB.First(&book, id)
	if book.ID == 0 {
		return c.Status(404).SendString("No book found with given ID")
	}
	return c.JSON(book)
}

// Create a book
func CreateBook(c *fiber.Ctx) error {
	book := new(models.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	database.DB.Create(&book)
	return c.JSON(book)
}

// Update a book by id
func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	database.DB.First(&book, id)
	if book.ID == 0 {
		return c.Status(404).SendString("No book found with given ID")
	}
	if err := c.BodyParser(&book); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	database.DB.Save(&book)
	return c.JSON(book)
}

// Delete a book by id
func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book
	database.DB.First(&book, id)
	if book.ID == 0 {
		return c.Status(404).SendString("No book found with given ID")
	}
	database.DB.Delete(&book)
	return c.SendString("Book successfully deleted")
}
