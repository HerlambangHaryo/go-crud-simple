package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/go-crud-simple/app/models"  
	) 

	func GetOne(c *fiber.Ctx) error { 
		// -------------------------------------------------------------- 
			b := new(models.FootballFixture)   
		// --------------------------------------------------------------   
			data, err := b.GetOne() 
			if err != nil {
				return c.Status(500).SendString(err.Error())
			} 
		// --------------------------------------------------------------
			return c.Render("contents/one/GetOne", fiber.Map{
				"Title": "One",
				"Content": "One",
				"Data": data,     
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}