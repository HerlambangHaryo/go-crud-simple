package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/go-crud-simple/app/models" 
		// "strconv" 
	) 

	func GetTomorrow(c *fiber.Ctx) error { 
		// -------------------------------------------------------------- 
			b := new(models.FootballFixture)   
		// --------------------------------------------------------------   
			data, err := b.GetTomorrow() 
			if err != nil {
				return c.Status(500).SendString(err.Error())
			} 
		// --------------------------------------------------------------
			return c.Render("contents/tomorrow/GetTomorrow", fiber.Map{
				"Title": "Tomorrow",
				"Content": "Tomorrow",
				"Data": data,     
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}