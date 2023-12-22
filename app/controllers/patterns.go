package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/go-crud-simple/app/models" 
		"strconv" 
	) 

	func GetPatternOther(c *fiber.Ctx) error {
		// --------------------------------------------------------------
			var preAhPattern, preGouPattern, preAhPatternMirror string
		// --------------------------------------------------------------
			leagueapi_id 	:= c.Params("leagueapi_id")
			season 			:= c.Params("season")
			fixtureapi_id 	:= c.Params("fixtureapi_id")

			leagueapi_idInt, err 	:= strconv.Atoi(leagueapi_id)
			seasonInt, err 			:= strconv.Atoi(season)
			fixtureapi_idInt, err 	:= strconv.Atoi(fixtureapi_id)

			if err != nil {
				return c.Status(500).SendString(err.Error())
			}
		// --------------------------------------------------------------  
			b 		:= new(models.FootballFixture)
			ba 		:= new(models.FootballOdd) 
			// fpf 	:= new(models.FootballPatternFrom) 
		// --------------------------------------------------------------  
			data, err := b.GetFixtureGeneral(uint(leagueapi_idInt), 
												uint(seasonInt), 
												uint(fixtureapi_idInt)) 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			}

			if len(data) > 0 { 
				preAhPattern 		= data[0].OddTeam.PreAhPattern
				preGouPattern 		= data[0].OddTeam.PreGouPattern  
				preAhPatternMirror 	= data[0].OddTeam.PreAhPatternMirror 
			} else { 
				preAhPattern 		= ""
				preGouPattern 		= "" 
				preAhPatternMirror 	= "" 
			}

			data_row, err := ba.GetPatternOther(uint(leagueapi_idInt),  
												preAhPattern, 
												preGouPattern, 
												preAhPatternMirror) 
 
		// --------------------------------------------------------------
			return c.Render("contents/patterns/other", fiber.Map{
				"Title"	: "League Not Started",
				"Content": "Fixtures",
				"Data"	: data,   
				"Data2" : data_row,   
				// "Data3" : data_pattern,  
				// "preAhPattern" : preAhPattern,   
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}
	 
	func GetPatternResponse(c *fiber.Ctx) error {
		// --------------------------------------------------------------  
			leagueapi_id 	:= c.Params("leagueapi_id")
			season 			:= c.Params("season") 

			leagueapi_idInt, err 	:= strconv.Atoi(leagueapi_id)
			seasonInt, err 			:= strconv.Atoi(season) 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			}
		// --------------------------------------------------------------   
			MFO 		:= new(models.FootballOdd)  
		// --------------------------------------------------------------  
			data, err := MFO.GetPatternResponse(uint(leagueapi_idInt), 
												uint(seasonInt)) 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			}  
		// --------------------------------------------------------------
			return c.Render("contents/patterns/GetPatternResponse", fiber.Map{
				"Title"	: "Pattern",
				"Content": "Response",
				"Data"	: data,     
				"leagueapi_id": leagueapi_id,
				"season": season,
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}

	func GetPatternFrom(c *fiber.Ctx) error {
		// --------------------------------------------------------------
			var preAhPattern, preGouPattern, preAhPatternMirror string
		// --------------------------------------------------------------
			leagueapi_id 	:= c.Params("leagueapi_id")
			season 			:= c.Params("season")
			fixtureapi_id 	:= c.Params("fixtureapi_id")

			leagueapi_idInt, err 	:= strconv.Atoi(leagueapi_id)
			seasonInt, err 			:= strconv.Atoi(season)
			fixtureapi_idInt, err 	:= strconv.Atoi(fixtureapi_id)

			if err != nil {
				return c.Status(500).SendString(err.Error())
			}
		// --------------------------------------------------------------  
			b 		:= new(models.FootballFixture)
			ba 		:= new(models.FootballOdd) 
			// fpf 	:= new(models.FootballPatternFrom) 
		// --------------------------------------------------------------  
			data, err := b.GetFixtureGeneral(uint(leagueapi_idInt), 
												uint(seasonInt), 
												uint(fixtureapi_idInt)) 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			}

			if len(data) > 0 { 
				preAhPattern 		= data[0].OddTeam.PreAhPattern
				preGouPattern 		= data[0].OddTeam.PreGouPattern  
				preAhPatternMirror 	= data[0].OddTeam.PreAhPatternMirror 
			} else { 
				preAhPattern 		= ""
				preGouPattern 		= "" 
				preAhPatternMirror 	= "" 
			}

			data_row, err := ba.GetPatternOther(uint(leagueapi_idInt),  
												preAhPattern, 
												preGouPattern, 
												preAhPatternMirror) 
 
		// --------------------------------------------------------------
			return c.Render("contents/patterns/other", fiber.Map{
				"Title"	: "League Not Started",
				"Content": "Fixtures",
				"Data"	: data,   
				"Data2" : data_row,   
				// "Data3" : data_pattern,  
				// "preAhPattern" : preAhPattern,   
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}
 
	func GetPatternResponseLeague(c *fiber.Ctx) error { 
		// --------------------------------------------------------------
			leagueapi_id 	:= c.Params("leagueapi_id")
			season 			:= c.Params("season") 
			pre_response 	:= c.Params("pre_response")
			end_response 	:= c.Params("end_response") 

			leagueapi_idInt, err 	:= strconv.Atoi(leagueapi_id)
			seasonInt, err 			:= strconv.Atoi(season) 
			pre_responseInt, err 	:= strconv.Atoi(pre_response)
			end_responseInt, err 	:= strconv.Atoi(end_response) 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			}
		// -------------------------------------------------------------- 
			b := new(models.FootballOdd)   
		// --------------------------------------------------------------   
			data, err := b.GetPatternResponseLeague(uint(leagueapi_idInt), 
											uint(seasonInt), 
											uint(pre_responseInt), 
											uint(end_responseInt)) 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			}   
		// --------------------------------------------------------------
			return c.Render("contents/patterns/GetPatternResponseLeague", fiber.Map{
				"Title": "Pattern Response", 
				"Content": "Dashboard",
				"Data": data,     
				"leagueapi_id": leagueapi_id,
				"season": season,
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}