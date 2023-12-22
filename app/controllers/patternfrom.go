package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/go-crud-simple/app/models" 
		"strconv" 
	) 

	func GetPatternFromOdd(c *fiber.Ctx) error {
		// -------------------------------------------------------------- 
			leagueapi_id 		:= c.Params("leagueapi_id")
			pre_ah 				:= c.Params("pre_ah")
			pre_gou 			:= c.Params("pre_gou") 
			end_ah 				:= c.Params("end_ah")
			end_gou 			:= c.Params("end_gou") 
			view 				:= c.Params("view") 

			leagueapi_idInt, err 	:= strconv.Atoi(leagueapi_id) 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			}
		// --------------------------------------------------------------   
			ba 		:= new(models.FootballOdd) 
			fpf 	:= new(models.FootballPatternFrom) 
		// --------------------------------------------------------------    
			data_row, err := ba.GetFixtureOddPatternOdd(uint(leagueapi_idInt),  
												pre_ah, 
												pre_gou, 
												end_ah, 
												end_gou) 

			data_pattern, err := fpf.GetPatternFrom(uint(leagueapi_idInt),  
												pre_ah, 
												pre_gou, 
												end_ah, 
												end_gou) 
		// --------------------------------------------------------------
			return c.Render("contents/patternfrom/" + view, fiber.Map{
				"Title"	: "Pattern From",
				"Content": "PrePre",
				"Data"	: data_row,   
				"Data2" : data_pattern,   
				"View" : view,   
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}

	func GetPatternFromPrePre(c *fiber.Ctx) error {
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
			fpf 	:= new(models.FootballPatternFrom) 
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

			data_row, err := ba.GetFixtureOddPatternPrePre(uint(leagueapi_idInt), 
												uint(fixtureapi_idInt), 
												preAhPattern, 
												preGouPattern, 
												preAhPatternMirror) 

			data_pattern, err := fpf.GetPatternFrom(uint(leagueapi_idInt),  
												preAhPattern, 
												preGouPattern, 
												preAhPattern, 
												preGouPattern)
		// --------------------------------------------------------------
			return c.Render("contents/patternfrom/GetOddPattern", fiber.Map{
				"Title"	: "Pattern From",
				"Content": "PrePre",
				"Data"	: data,   
				"Data2" : data_row,   
				"Data3" : data_pattern,  
				"leagueapi_id" : leagueapi_id,   
				"season" : season,   
				"fixtureapi_id" : fixtureapi_id,   
				"pattern_region" : "League",   
				"pattern_type" : "PrePre",   
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}

	func GetPatternFromPreEnd(c *fiber.Ctx) error {
		// --------------------------------------------------------------
			var preAhPattern, preGouPattern, preAhPatternMirror, endAhPattern, endGouPattern, endAhPatternMirror string
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
			fpf 	:= new(models.FootballPatternFrom) 
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

				endAhPattern 		= data[0].OddTeam.EndAhPattern
				endGouPattern 		= data[0].OddTeam.EndGouPattern  
				endAhPatternMirror 	= data[0].OddTeam.EndAhPatternMirror 
			} else { 
				preAhPattern 		= ""
				preGouPattern 		= "" 
				preAhPatternMirror 	= "" 

				endAhPattern 		= "" 
				endGouPattern 		= ""   
				endAhPatternMirror 	= "" 
			}

			data_row, err := ba.GetFixtureOddPatternPreEnd(uint(leagueapi_idInt), 
												uint(fixtureapi_idInt), 
												preAhPattern, 
												preGouPattern, 
												preAhPatternMirror, 
												endAhPattern, 
												endGouPattern, 
												endAhPatternMirror) 

			data_pattern, err := fpf.GetPatternFrom(uint(leagueapi_idInt),  
												preAhPattern, 
												preGouPattern, 
												endAhPattern, 
												endGouPattern)
		// --------------------------------------------------------------
			return c.Render("contents/patternfrom/GetOddPattern", fiber.Map{
				"Title"	: "Pattern From",
				"Content": "PreEnd",
				"Data"	: data,   
				"Data2" : data_row,   
				"Data3" : data_pattern,   
				"leagueapi_id" : leagueapi_id,   
				"season" : season,   
				"fixtureapi_id" : fixtureapi_id,    
				"pattern_region" : "League",   
				"pattern_type" : "PreEnd",   
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	} 

	func GetPatternFromOnlyPre(c *fiber.Ctx) error {
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
			b 	:= new(models.FootballFixture)
			ba 	:= new(models.FootballOdd) 
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

			data_row, err := ba.GetFixtureOddPatternOnlyPre(uint(leagueapi_idInt), 
												uint(fixtureapi_idInt), 
												preAhPattern, 
												preGouPattern, 
												preAhPatternMirror) 
		// --------------------------------------------------------------
			return c.Render("contents/patternfrom/GetOddPattern", fiber.Map{
				"Title"	: "Pattern From",
				"Content": "OnlyPre",
				"Data"	: data,   
				"Data2" : data_row,   
				"leagueapi_id" : leagueapi_id,   
				"season" : season,   
				"fixtureapi_id" : fixtureapi_id,    
				"pattern_region" : "League",   
				"pattern_type" : "OnlyPre",    
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}

	func GetPatternFromOnlyEnd(c *fiber.Ctx) error {
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
			b 	:= new(models.FootballFixture)
			ba 	:= new(models.FootballOdd) 
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

			data_row, err := ba.GetFixtureOddPatternOnlyEnd(uint(leagueapi_idInt), 
												uint(fixtureapi_idInt), 
												preAhPattern, 
												preGouPattern, 
												preAhPatternMirror) 
		// --------------------------------------------------------------
			return c.Render("contents/patternfrom/GetOddPattern", fiber.Map{
				"Title"	: "Pattern From",
				"Content": "OnlyEnd",
				"Data"	: data,   
				"Data2" : data_row,   
				"leagueapi_id" : leagueapi_id,   
				"season" : season,   
				"fixtureapi_id" : fixtureapi_id,     
				"pattern_region" : "League",   
				"pattern_type" : "OnlyEnd",   
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	} 