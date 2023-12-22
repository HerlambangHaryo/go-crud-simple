package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/go-crud-simple/app/models" 
		"strconv"
	) 
  
	func GetLeagueNotstarted(c *fiber.Ctx) error {
		// --------------------------------------------------------------
			leagueapi_id := c.Params("leagueapi_id")
			season := c.Params("season")
		// --------------------------------------------------------------
			b := new(models.FootballFixture)

			leagueapi_idInt, err := strconv.Atoi(leagueapi_id)
			seasonInt, err := strconv.Atoi(season) 
			if err != nil {
				return c.Status(500).SendString(err.Error())
			}
		// -------------------------------------------------------------- 
			data, err := b.GetLeagueNotstarted(uint(leagueapi_idInt), uint(seasonInt)) 
			if err != nil {
				return c.Status(500).SendString(err.Error())
			}   
		// --------------------------------------------------------------
			return c.Render("contents/leagues/GetLeagueNotstarted", fiber.Map{
				"Title": "League Not Started",
				"Content": "Leagues",
				"Data": data,
				"leagueapi_id": leagueapi_id,
				"season": season,
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}
 
	func GetLeagueMatchfinished(c *fiber.Ctx) error { 
		// --------------------------------------------------------------
			leagueapi_id 	:= c.Params("leagueapi_id")
			season 			:= c.Params("season") 
	
			leagueapi_idInt, err 	:= strconv.Atoi(leagueapi_id)
			seasonInt, err 			:= strconv.Atoi(season) 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			}
		// --------------------------------------------------------------  
			ff 	:= new(models.FootballFixture) 
		// --------------------------------------------------------------  
			data, err := ff.GetLeagueMatchfinishedGroup(uint(leagueapi_idInt), 
												uint(seasonInt)) 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			} 
		// --------------------------------------------------------------
    		return c.Render("contents/leagues/GetLeagueMatchfinished", fiber.Map{
				"Title"	: "League Not Started",
				"Content": "Leagues",
				"Data"	: data,    
				"leagueapi_id": leagueapi_id,
				"season": season,
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	} 
 
	func GetLeagueDetailRound(c *fiber.Ctx) error { 
		// --------------------------------------------------------------
			leagueapi_id 	:= c.Params("leagueapi_id")
			season 			:= c.Params("season") 
			round 			:= c.Params("round")  
	
			leagueapi_idInt, err 	:= strconv.Atoi(leagueapi_id)
			seasonInt, err 			:= strconv.Atoi(season) 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			}
		// --------------------------------------------------------------  
			ff 	:= new(models.FootballFixture) 
		// --------------------------------------------------------------  
			data, err := ff.GetLeagueDetailRound(uint(leagueapi_idInt), 
												uint(seasonInt),
												round) 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			} 
		// --------------------------------------------------------------
    		return c.Render("contents/leagues/GetLeagueDetailRound", fiber.Map{
				"Title"	: "League Not Started",
				"Data"	: data,    
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	} 
 
	func GetLeagueStanding(c *fiber.Ctx) error { 
		// --------------------------------------------------------------
			leagueapi_id 	:= c.Params("leagueapi_id")
			season 			:= c.Params("season") 
			val 			:= c.Params("val") 
	
			leagueapi_idInt, err 	:= strconv.Atoi(leagueapi_id)
			seasonInt, err 			:= strconv.Atoi(season) 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			}
		// --------------------------------------------------------------  
			afls 	:= new(models.ApiFootballLeagueStanding) 
		// --------------------------------------------------------------  
			data, err := afls.GetLeagueStanding(uint(leagueapi_idInt), 
												uint(seasonInt)) 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			} 
		// --------------------------------------------------------------
    		return c.Render("contents/leagues/GetLeagueStanding" + val, fiber.Map{
				"Title"	: "League Not Started",
				"Content": "Leagues",
				"Data"	: data,    
				"leagueapi_id": leagueapi_id,
				"season": season,
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	} 
	