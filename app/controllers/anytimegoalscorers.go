// book.go
package controllers

    import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/go-crud-simple/app/models" 
		"strconv" 
	) 

	func GetAgsToday(c *fiber.Ctx) error { 
		// -------------------------------------------------------------- 
			b := new(models.ResearchOnAnytimeGoalScorer)   
		// --------------------------------------------------------------   
			data, err := b.GetAgsByToday() 
			if err != nil {
				return c.Status(500).SendString(err.Error())
			}   
        // --------------------------------------------------------------
            return c.Render("contents/anytimegoalscorers/GetAgsToday", fiber.Map{
                "Title": "Anytime Goal Scorer", 
				"Content": "Dashboard",
				"Data": data,     
			},"templates/studiov30/datatable")
        // --------------------------------------------------------------
    }

	func GetAgsLeagues(c *fiber.Ctx) error { 
		// --------------------------------------------------------------
			leagueapi_id 	:= c.Params("leagueapi_id")
			season 			:= c.Params("season") 

			leagueapi_idInt, err 	:= strconv.Atoi(leagueapi_id)
			seasonInt, err 			:= strconv.Atoi(season) 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			}
		// -------------------------------------------------------------- 
			b := new(models.ResearchOnAnytimeGoalScorer)   
		// --------------------------------------------------------------   
			data, err := b.GetAgsByLeagues(uint(leagueapi_idInt), 
											uint(seasonInt)) 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			}   
        // --------------------------------------------------------------
            return c.Render("contents/anytimegoalscorers/GetAgsLeagues", fiber.Map{
                "Title": "Anytime Goal Scorer", 
				"Content": "Dashboard",
				"Data": data,     
				"leagueapi_id": leagueapi_id,
				"season": season,
			},"templates/studiov30/datatable")
        // --------------------------------------------------------------
    }