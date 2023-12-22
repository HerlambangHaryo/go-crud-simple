package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/go-crud-simple/app/models" 
		"strconv" 
	) 

	func GetMatchFinished(c *fiber.Ctx) error { 
		// -------------------------------------------------------------- 
			b := new(models.FootballFixture)   
		// --------------------------------------------------------------   
			data, err := b.GetMatchFinished() 
			if err != nil {
				return c.Status(500).SendString(err.Error())
			} 
		// --------------------------------------------------------------
			return c.Render("contents/matchfinished/GetMatchFinished", fiber.Map{
				"Title": "Match Finished",
				"Content": "MatchFinished",
				"Data": data,     
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}

	func GetMatchFinishedTimegroup(c *fiber.Ctx) error { 
		// -------------------------------------------------------------- 
			b := new(models.FootballFixture)   
		// --------------------------------------------------------------   
			data, err := b.GetMatchFinishedTimegroup() 
			if err != nil {
				return c.Status(500).SendString(err.Error())
			} 
		// --------------------------------------------------------------
			return c.Render("contents/MatchFinished/GetMatchFinishedTimegroup", fiber.Map{
				"Title": "MatchFinished",
				"Content": "MatchFinished",
				"Data": data,     
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}

	func GetMatchFinishedLeague(c *fiber.Ctx) error { 
		// --------------------------------------------------------------
			leagueapi_id := c.Params("leagueapi_id")
			season := c.Params("season") 
		// -------------------------------------------------------------- 
			b := new(models.FootballFixture)   

			leagueapi_idInt, err := strconv.Atoi(leagueapi_id)
			seasonInt, err := strconv.Atoi(season)
		// --------------------------------------------------------------   
			data, err := b.GetMatchFinishedLeague(uint(leagueapi_idInt), uint(seasonInt)) 
			if err != nil {
				return c.Status(500).SendString(err.Error())
			} 
		// --------------------------------------------------------------
			return c.Render("contents/MatchFinished/GetMatchFinishedLeague", fiber.Map{
				"Title": "MatchFinished",
				"Content": "MatchFinished",
				"Data": data,     
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}

	func GetMatchFinishedTime(c *fiber.Ctx) error { 
		// --------------------------------------------------------------
			year 	:= c.Params("year")
			month 	:= c.Params("month") 
			day 	:= c.Params("day") 
			hour 	:= c.Params("hour") 
			minute 	:= c.Params("minute")  
		// -------------------------------------------------------------- 
			b := new(models.FootballFixture)   

			yearInt, err 	:= strconv.Atoi(year)
			monthInt, err 	:= strconv.Atoi(month)
			dayInt, err 	:= strconv.Atoi(day)
			hourInt, err 	:= strconv.Atoi(hour)
			minuteInt, err 	:= strconv.Atoi(minute)
		// --------------------------------------------------------------   
			data, err := b.GetMatchFinishedTime(uint(yearInt), uint(monthInt), uint(dayInt), uint(hourInt), uint(minuteInt)) 
			if err != nil {
				return c.Status(500).SendString(err.Error())
			} 
		// --------------------------------------------------------------
			return c.Render("contents/MatchFinished/GetMatchFinishedTime", fiber.Map{
				"Title": "MatchFinished",
				"Content": "MatchFinished",
				"Data": data,     
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}
