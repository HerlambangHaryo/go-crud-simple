package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/go-crud-simple/app/models" 
		// "strconv" 
	) 

	func GetMatchDone(c *fiber.Ctx) error { 
		// -------------------------------------------------------------- 
			b := new(models.FootballFixture)   
		// --------------------------------------------------------------   
			data, err := b.GetMatchDone() 
			if err != nil {
				return c.Status(500).SendString(err.Error())
			} 
		// --------------------------------------------------------------
			return c.Render("contents/matchdone/GetMatchDone", fiber.Map{
				"Title": "Match Done",
				"Content": "MatchDone",
				"Data": data,     
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}