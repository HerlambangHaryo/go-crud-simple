package book

import (
	"fmt"
	"github.com/HerlambangHaryo/go-crud-simple/platform/database" 
)
 
type Book struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"size:255"`
}


// Migrate the schema of database if needed
func AutoMigrate() {
	err := database.DB.AutoMigrate(&Book{})
	if err != nil {
		panic("Failed to migrate database")
	}
	fmt.Println("Database migrated")
}

// Get all books
func (b *Book) GetBooks() ([]Book, error) {
	var books []Book
	err := database.DB.Find(&books).Error
	return books, err
}

// Get a book by id
func (b *Book) GetBook(id uint) (Book, error) {
	var book Book
	err := database.DB.First(&book, id).Error
	return book, err
}

// Create a book
func (b *Book) CreateBook() error {
	err := database.DB.Create(&b).Error
	return err
}

// Update a book by id
func (b *Book) UpdateBook(id uint) error {
	err := database.DB.Model(&b).Where("id = ?", id).Updates(b).Error
	return err
}

// Delete a book by id
func (b *Book) DeleteBook(id uint) error {
	err := database.DB.Delete(&b, id).Error
	return err
}
