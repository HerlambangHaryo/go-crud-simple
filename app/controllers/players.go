package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/go-crud-simple/app/models" 
		"strconv"
	) 
  
	func GetPlayersAnytimegoalscorers(c *fiber.Ctx) error {
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
			return c.Render("contents/players/GetPlayersAnytimegoalscorers", fiber.Map{
				"Title": "League Not Started",
				"Content": "Players",
				"Data": data,
				"leagueapi_id": leagueapi_id,
				"season": season,
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------
	}