// rapidapi.go
package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/HerlambangHaryo/go-crud-simple/app/models"
	"strconv"
) 

// Get all rapidapis
func GetRapidapis(c *fiber.Ctx) error {
	b := new(models.Rapidapi)
	rapidapis, err := b.GetRapidapis()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	// return c.JSON(rapidapis) 
	 
    return c.Render("contents/rapidapi/index", fiber.Map{
        "Rapidapis": rapidapis,
    },"templates/studiov30/datatable")
}

// Get a rapidapi by id
func GetRapidapi(c *fiber.Ctx) error {
	id := c.Params("id")
	b := new(models.Rapidapi)
	// convert id from string to uint
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(400).SendString("Invalid id")
	}
	idUint := uint(idInt)
	rapidapi, err := b.GetRapidapi(idUint)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if rapidapi.ID == 0 {
		return c.Status(404).SendString("No rapidapi found with given ID")
	}
	return c.JSON(rapidapi)
}

// Create a rapidapi
func CreateRapidapi(c *fiber.Ctx) error {
	b := new(models.Rapidapi)
	if err := c.BodyParser(b); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err := b.CreateRapidapi()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(b)
}

// Update a rapidapi by id
func UpdateRapidapi(c *fiber.Ctx) error {
	id := c.Params("id")
	b := new(models.Rapidapi)
	// convert id from string to uint
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(400).SendString("Invalid id")
	}
	idUint := uint(idInt)
	err = b.UpdateRapidapi(idUint)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if b.ID == 0 {
		return c.Status(404).SendString("No rapidapi found with given ID")
	}
	if err := c.BodyParser(&b); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = b.UpdateRapidapi(idUint) // change SaveRapidapi to UpdateRapidapi
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(b)
}

// Delete a rapidapi by id
func DeleteRapidapi(c *fiber.Ctx) error {
	id := c.Params("id")
	b := new(models.Rapidapi)
	// convert id from string to uint
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(400).SendString("Invalid id")
	}
	idUint := uint(idInt)
	err = b.DeleteRapidapi(idUint)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if b.ID == 0 {
		return c.Status(404).SendString("No rapidapi found with given ID")
	}
	return c.SendString("Rapidapi successfully deleted")
}
