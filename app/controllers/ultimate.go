package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/go-crud-simple/app/models"  
	) 

	func GetUltimate(c *fiber.Ctx) error { 
		// -------------------------------------------------------------- 
			b := new(models.FootballFixture)   
		// --------------------------------------------------------------   
			data, err := b.GetUltimate() 
			if err != nil {
				return c.Status(500).SendString(err.Error())
			} 
		// --------------------------------------------------------------
			return c.Render("contents/ultimate/GetUltimate", fiber.Map{
				"Title": "Ultimate",
				"Content": "Ultimate",
				"Data": data,     
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}
 