package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/go-crud-simple/app/models" 
		"strconv" 
	) 

	func GetPatternFromCountryPrePre(c *fiber.Ctx) error {
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
			fpf 	:= new(models.FootballPatternFromCountry) 
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

			data_row, err := ba.GetFixtureOddPatternPrePreCountry(uint(leagueapi_idInt), 
												uint(fixtureapi_idInt), 
												preAhPattern, 
												preGouPattern, 
												preAhPatternMirror) 
 

			data_pattern, err := fpf.GetPatternFromCountry(uint(leagueapi_idInt),  
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
				"pattern_region" : "Country",   
				"pattern_type" : "PrePre",    
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}

	func GetPatternFromCountryPreEnd(c *fiber.Ctx) error {
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
			fpf 	:= new(models.FootballPatternFromCountry) 
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

			data_row, err := ba.GetFixtureOddPatternPreEndCountry(uint(leagueapi_idInt), 
												uint(fixtureapi_idInt), 
												preAhPattern, 
												preGouPattern, 
												preAhPatternMirror, 
												endAhPattern, 
												endGouPattern, 
												endAhPatternMirror) 

			data_pattern, err := fpf.GetPatternFromCountry(uint(leagueapi_idInt),  
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
				"pattern_region" : "Country",   
				"pattern_type" : "PreEnd",
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	} 

	func GetPatternFromCountryOnlyPre(c *fiber.Ctx) error {
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

			data_row, err := ba.GetFixtureOddPatternOnlyPreCountry(uint(leagueapi_idInt), 
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
				"pattern_region" : "Country",   
				"pattern_type" : "OnlyPre",
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}

	func GetPatternFromCountryOnlyEnd(c *fiber.Ctx) error {
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

			data_row, err := ba.GetFixtureOddPatternOnlyEndCountry(uint(leagueapi_idInt), 
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
				"pattern_region" : "Country",   
				"pattern_type" : "OnlyEnd",
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	} 