// book.go
package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/HerlambangHaryo/go-crud-simple/app/models/book"
)

// Get all books
func GetBooks(c *fiber.Ctx) error {
	b := new(book.Book)
	books, err := b.GetBooks()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(books)
}

// Get a book by id
func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	b := new(book.Book)
	book, err := b.GetBook(id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if book.ID == 0 {
		return c.Status(404).SendString("No book found with given ID")
	}
	return c.JSON(book)
}

// Create a book
func CreateBook(c *fiber.Ctx) error {
	b := new(book.Book)
	if err := c.BodyParser(b); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err := b.CreateBook()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(b)
}

// Update a book by id
func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	b := new(book.Book)
	err := b.UpdateBook(id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if b.ID == 0 {
		return c.Status(404).SendString("No book found with given ID")
	}
	if err := c.BodyParser(&b); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = b.SaveBook()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(b)
}

// Delete a book by id
func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	b := new(book.Book)
	err := b.DeleteBook(id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if b.ID == 0 {
		return c.Status(404).SendString("No book found with given ID")
	}
	return c.SendString("Book successfully deleted")
}
