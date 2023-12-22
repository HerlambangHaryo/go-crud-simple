package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/go-crud-simple/app/models" 
		"strconv" 
	) 

	func GetToday(c *fiber.Ctx) error { 
		// -------------------------------------------------------------- 
			b := new(models.FootballFixture)   
		// --------------------------------------------------------------   
			data, err := b.GetToday() 
			if err != nil {
				return c.Status(500).SendString(err.Error())
			} 
		// --------------------------------------------------------------
			return c.Render("contents/today/GetToday", fiber.Map{
				"Title": "Today",
				"Content": "Today",
				"Data": data,     
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}

	func GetTodayTimegroup(c *fiber.Ctx) error { 
		// -------------------------------------------------------------- 
			b := new(models.FootballFixture)   
		// --------------------------------------------------------------   
			data, err := b.GetTodayTimegroup() 
			if err != nil {
				return c.Status(500).SendString(err.Error())
			} 
		// --------------------------------------------------------------
			return c.Render("contents/today/GetTodayTimegroup", fiber.Map{
				"Title": "Today",
				"Content": "Today",
				"Data": data,     
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}

	func GetTodayLeague(c *fiber.Ctx) error { 
		// --------------------------------------------------------------
			leagueapi_id := c.Params("leagueapi_id")
			season := c.Params("season") 
		// -------------------------------------------------------------- 
			b := new(models.FootballFixture)   

			leagueapi_idInt, err := strconv.Atoi(leagueapi_id)
			seasonInt, err := strconv.Atoi(season)
		// --------------------------------------------------------------   
			data, err := b.GetTodayLeague(uint(leagueapi_idInt), uint(seasonInt)) 
			if err != nil {
				return c.Status(500).SendString(err.Error())
			} 
		// --------------------------------------------------------------
			return c.Render("contents/today/GetTodayLeague", fiber.Map{
				"Title": "Today",
				"Content": "Today",
				"Data": data,     
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}

	func GetTodayTime(c *fiber.Ctx) error { 
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
			data, err := b.GetTodayTime(uint(yearInt), uint(monthInt), uint(dayInt), uint(hourInt), uint(minuteInt)) 
			if err != nil {
				return c.Status(500).SendString(err.Error())
			} 
		// --------------------------------------------------------------
			return c.Render("contents/today/GetTodayTime", fiber.Map{
				"Title": "Today",
				"Content": "Today",
				"Data": data,     
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}