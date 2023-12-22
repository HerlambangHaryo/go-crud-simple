package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/go-crud-simple/app/models" 
		"strconv" 
	) 
	
	func GetFixtureInformation(c *fiber.Ctx) error {
		// --------------------------------------------------------------
			leagueapi_id := c.Params("leagueapi_id")
			season := c.Params("season")
			fixtureapi_id := c.Params("fixtureapi_id")
	
			leagueapi_idInt, err 	:= strconv.Atoi(leagueapi_id)
			seasonInt, err 			:= strconv.Atoi(season)
			fixtureapi_idInt, err 	:= strconv.Atoi(fixtureapi_id)

			if err != nil {
				return c.Status(500).SendString(err.Error())
			}
		// -------------------------------------------------------------- 
			b 		:= new(models.FootballFixture)
		// --------------------------------------------------------------
			data, err := b.GetFixtureGeneral(uint(leagueapi_idInt), 
						uint(seasonInt), 
						uint(fixtureapi_idInt)) 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			}
		// --------------------------------------------------------------
			return c.Render("contents/fixtures/GetFixtureInformation", fiber.Map{
				"Title": "League Not Started",
				"Content": "Fixtures",
				"Data"	: data,    
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}
	
	func GetFixtureOdd(c *fiber.Ctx) error {
		// --------------------------------------------------------------
			leagueapi_id := c.Params("leagueapi_id")
			season := c.Params("season")
			fixtureapi_id := c.Params("fixtureapi_id")
	
			leagueapi_idInt, err 	:= strconv.Atoi(leagueapi_id)
			seasonInt, err 			:= strconv.Atoi(season)
			fixtureapi_idInt, err 	:= strconv.Atoi(fixtureapi_id)

			if err != nil {
				return c.Status(500).SendString(err.Error())
			}
		// -------------------------------------------------------------- 
			b 		:= new(models.FootballFixture)
		// --------------------------------------------------------------
			// data, err := b.GetFixtureInfoOdd(uint(leagueapi_idInt), uint(seasonInt), uint(fixtureapi_idInt)) 
			// if err != nil {
			// 	return c.Status(500).SendString(err.Error())
			// }    

			
			data, err := b.GetFixtureGeneral(uint(leagueapi_idInt), 
						uint(seasonInt), 
						uint(fixtureapi_idInt)) 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			}
		// --------------------------------------------------------------
			return c.Render("contents/fixtures/GetFixtureOdd", fiber.Map{
				"Title": "League Not Started",
				"Content": "Fixtures",
				"Data": data,  
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}

	func GetFixturePrePre(c *fiber.Ctx) error {
		// --------------------------------------------------------------
			// var preAhPattern, preGouPattern, preAhPatternMirror string
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
			// ba 		:= new(models.FootballOdd) 
			// fpf 	:= new(models.FootballPatternFrom) 
		// --------------------------------------------------------------  
			data, err := b.GetFixtureGeneral(uint(leagueapi_idInt), 
												uint(seasonInt), 
												uint(fixtureapi_idInt)) 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			}
	
			// if len(data) > 0 { 
			// 	preAhPattern 		= data[0].OddTeam.PreAhPattern
			// 	preGouPattern 		= data[0].OddTeam.PreGouPattern  
			// 	preAhPatternMirror 	= data[0].OddTeam.PreAhPatternMirror 
			// } else { 
			// 	preAhPattern 		= ""
			// 	preGouPattern 		= "" 
			// 	preAhPatternMirror 	= "" 
			// }

			// data_row, err := ba.GetFixtureOddPatternPrePre(uint(leagueapi_idInt), 
			// 									uint(fixtureapi_idInt), 
			// 									preAhPattern, 
			// 									preGouPattern, 
			// 									preAhPatternMirror) 

			// data_pattern, err := fpf.GetPatternFrom(uint(leagueapi_idInt),  
			// 									preAhPattern, 
			// 									preGouPattern, 
			// 									preAhPattern, 
			// 									preGouPattern)
		// --------------------------------------------------------------
			return c.Render("contents/fixtures/GetOddPattern", fiber.Map{
				"Title"	: "League Not Started",
				"Content": "Fixtures",
				"Data"	: data,   
				// "Data2" : data_row,   
				// "Data3" : data_pattern,  
				// "preAhPattern" : preAhPattern,   
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}

	func GetFixturePreEnd(c *fiber.Ctx) error {
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
			return c.Render("contents/fixtures/GetOddPattern", fiber.Map{
				"Title"	: "League Not Started",
				"Content": "Fixtures",
				"Data"	: data,   
				"Data2" : data_row,   
				"Data3" : data_pattern,   
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	} 

	func GetFixtureOnlyPre(c *fiber.Ctx) error {
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
			return c.Render("contents/fixtures/GetOddPattern", fiber.Map{
				"Title"	: "League Not Started",
				"Content": "Fixtures",
				"Data"	: data,   
				"Data2" : data_row,   
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}

	func GetFixtureOnlyEnd(c *fiber.Ctx) error {
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
			return c.Render("contents/fixtures/GetOddPattern", fiber.Map{
				"Title"	: "League Not Started",
				"Content": "Fixtures",
				"Data"	: data,   
				"Data2" : data_row,   
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	} 

	func UpdateFixtureBehavioral(c *fiber.Ctx) error { 
		// --------------------------------------------------------------
			id 				:= c.Params("id")
			val 			:= c.Params("val") 
	
			idInt, err 		:= strconv.Atoi(id)
			valInt, err 	:= strconv.Atoi(val) 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			}
		// --------------------------------------------------------------   
			ff := &models.FootballFixture{}
		// --------------------------------------------------------------  
			if err := ff.UpdateSingleColumnByID(uint(idInt), "BehavioralBmAssessment", uint(valInt)); err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}
		// --------------------------------------------------------------  
			previousPage := c.Get("Referer")
			if previousPage == "" {
				previousPage = "/default/page"
			}

			return c.Redirect(previousPage)
		// --------------------------------------------------------------
	} 

	func UpdateSetOne(c *fiber.Ctx) error { 
		// --------------------------------------------------------------
			id 				:= c.Params("id") 
	
			idInt, err 		:= strconv.Atoi(id) 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			}
		// --------------------------------------------------------------   
			ff := &models.FootballFixture{}
		// --------------------------------------------------------------  
			if err := ff.UpdateSingleColumnByID(uint(idInt), "set_one", 1); err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}
		// --------------------------------------------------------------  
			previousPage := c.Get("Referer")
			if previousPage == "" {
				previousPage = "/default/page"
			}

			return c.Redirect(previousPage)
		// --------------------------------------------------------------
	} 

	func UpdateClearSetOne(c *fiber.Ctx) error { 
		// --------------------------------------------------------------   
			ff := &models.FootballFixture{}
		// --------------------------------------------------------------  
			if err := ff.UpdateClearSetOne(1, "set_one", 2); err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			}
		// --------------------------------------------------------------  
			previousPage := c.Get("Referer")
			if previousPage == "" {
				previousPage = "/default/page"
			}

			return c.Redirect(previousPage)
		// --------------------------------------------------------------
	} 
 
	func GetFixturePreGames(c *fiber.Ctx) error {
		// --------------------------------------------------------------
			leagueapi_id := c.Params("leagueapi_id")
			season := c.Params("season")
			fixtureapi_id := c.Params("fixtureapi_id")
			teams_homeapi_id := c.Params("teams_homeapi_id")
			teams_awayapi_id := c.Params("teams_awayapi_id")
	
			leagueapi_idInt, err 	:= strconv.Atoi(leagueapi_id)
			seasonInt, err 			:= strconv.Atoi(season)
			fixtureapi_idInt, err 	:= strconv.Atoi(fixtureapi_id)

			teams_homeapi_idInt, err 	:= strconv.Atoi(teams_homeapi_id)
			teams_awayapi_idInt, err 	:= strconv.Atoi(teams_awayapi_id)

			if err != nil {
				return c.Status(500).SendString(err.Error())
			}
		// -------------------------------------------------------------- 
			b 		:= new(models.FootballFixture)
			MFS 		:= new(models.FootballStatistic)
		// --------------------------------------------------------------
			data, err := b.GetFixtureGeneral(uint(leagueapi_idInt), 
						uint(seasonInt), 
						uint(fixtureapi_idInt)) 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			}
		// --------------------------------------------------------------
			dataPreGamesHome, err := MFS.GetStatsPreGames(uint(teams_homeapi_idInt)) 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			}
		// --------------------------------------------------------------
			dataPreGamesAway, err := MFS.GetStatsPreGames(uint(teams_awayapi_idInt)) 

			if err != nil {
				return c.Status(500).SendString(err.Error())
			}
		// --------------------------------------------------------------
			return c.Render("contents/fixtures/GetFixturePreGames", fiber.Map{
				"Title": "League Not Started",
				"Content": "Fixtures",
				"Data"	: data,    
				"dataPreGamesHome"	: dataPreGamesHome,    
				"dataPreGamesAway"	: dataPreGamesAway,    
			},"templates/studiov30/apexchart")
		// --------------------------------------------------------------
	}
 