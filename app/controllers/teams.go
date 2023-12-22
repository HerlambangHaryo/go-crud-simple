package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/go-crud-simple/app/models" 
		"strconv" 
	) 
	
	func GetTeam(c *fiber.Ctx) error {
		// --------------------------------------------------------------
			teamapi_id := c.Params("id")  

			teamapi_idInt, err 	:= strconv.Atoi(teamapi_id)

			if err != nil {
				return c.Status(500).SendString(err.Error())
			}
		// -------------------------------------------------------------- 
			MTF 		:= new(models.FootballTeam)
			MTC 		:= new(models.FootballCoach)
		// --------------------------------------------------------------
			data, err := MTF.GetTeam(uint(teamapi_idInt)) 

			if err != nil {
			return c.Status(500).SendString(err.Error())
			}
		// --------------------------------------------------------------
			dataCoach, err := MTC.GetCoach(uint(teamapi_idInt)) 

			if err != nil {
			return c.Status(500).SendString(err.Error())
			}
		// --------------------------------------------------------------
			return c.Render("contents/teams/GetTeam", fiber.Map{
				"Title": "Teams",
				"Content": "Teams",
				"teamapi_id"	: teamapi_id,    
				"Data"	: data,    
				"DataCoach"	: dataCoach,    
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}