// book.go
package controllers

    import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/go-crud-simple/app/models" 
		"strconv" 
	) 

	func GetCompareOdds(c *fiber.Ctx) error {  
		// --------------------------------------------------------------
			var preAhPattern, preGouPattern, preAhPatternMirror string
		// --------------------------------------------------------------
			bet_type 		:= c.Params("bet_type")
				
			pattern_region 		:= c.Params("pattern_region")
			pattern_type 		:= c.Params("pattern_type")
			
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
            return c.Render("contents/compare/" + bet_type, fiber.Map{
                "Title": "Anytime Goal Scorer", 
				"Content": "Dashboard",
				"Data"	: data,   
				"Data2" : data_row,   
				"Data3" : data_pattern,  
				"leagueapi_id" : leagueapi_id,   
				"season" : season,   
				"fixtureapi_id" : fixtureapi_id,   
				"pattern_region" : pattern_region,   
				"pattern_type" : pattern_type,   
			},"templates/studiov30/datatable")
        // --------------------------------------------------------------
    }