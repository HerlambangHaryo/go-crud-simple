package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/HerlambangHaryo/go-crud-simple/app/controllers"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/", controllers.GetDashboard) 
 

	app.Get("/Books", controllers.GetBooks)
	app.Get("/Books/:id", controllers.GetBook)
	app.Post("/Books", controllers.CreateBook)
	app.Put("/Books/:id", controllers.UpdateBook)
	app.Delete("/Books/:id", controllers.DeleteBook)
 
	app.Get("/Dashboard", controllers.GetDashboard)
	
	app.Get("/Rapidapis", controllers.GetRapidapis)
	app.Get("/Rapidapis/:id", controllers.GetRapidapi)
	app.Post("/Rapidapis", controllers.CreateRapidapi)
	app.Put("/Rapidapis/:id", controllers.UpdateRapidapi)
	app.Delete("/Rapidapis/:id", controllers.DeleteRapidapi)

	
	app.Get("/Countries", controllers.GetCountries)
	app.Get("/Countries/:id", controllers.GetCountry)

	
	app.Get("/Compare/Odds/:bet_type/:pattern_region/:pattern_type/:leagueapi_id/:season/:fixtureapi_id", controllers.GetCompareOdds) 

	app.Get("/Leagues/Notstarted/:leagueapi_id/:season", controllers.GetLeagueNotstarted)
	app.Get("/Leagues/Matchfinished/:leagueapi_id/:season", controllers.GetLeagueMatchfinished)
	app.Get("/Leagues/Detailround/:leagueapi_id/:season/:round", controllers.GetLeagueDetailRound)
	app.Get("/Leagues/Standing/:leagueapi_id/:season/:val", controllers.GetLeagueStanding)

	app.Get("/Fixtures/Information/:leagueapi_id/:season/:fixtureapi_id", controllers.GetFixtureInformation) 
	
	app.Get("/Fixtures/Odd/:leagueapi_id/:season/:fixtureapi_id", controllers.GetFixtureOdd) 

	app.Get("/Fixtures/Behavioral/:id/:val", controllers.UpdateFixtureBehavioral) 
	
	app.Get("/Fixtures/PreGames/:leagueapi_id/:season/:fixtureapi_id/:teams_homeapi_id/:teams_awayapi_id", controllers.GetFixturePreGames) 

	app.Get("/Fixtures/SetOne/:id", controllers.UpdateSetOne) 
	app.Get("/Fixtures/ClearOne/", controllers.UpdateClearSetOne) 

	app.Get("/Today", controllers.GetToday) 
	app.Get("/Today/Timegroup", controllers.GetTodayTimegroup) 
	app.Get("/Today/League/:leagueapi_id/:season", controllers.GetTodayLeague)
	app.Get("/Today/Time/:year/:month/:day/:hour/:minute", controllers.GetTodayTime)
 
	app.Get("/Tomorrow", controllers.GetTomorrow) 
	// app.Get("/Tomorrow/Timegroup", controllers.GetTomorrowTimegroup) 
	// app.Get("/Tomorrow/League/:leagueapi_id/:season", controllers.GetTomorrowLeague)
	// app.Get("/Tomorrow/Time/:year/:month/:day/:hour/:minute", controllers.GetTomorrowTime)
 
	app.Get("/Ultimate", controllers.GetUltimate)  

	
	app.Get("/One", controllers.GetOne) 
	
	app.Get("/Players/Anytimegoalscorers/:leagueapi_id/:season/:round", controllers.GetPlayersAnytimegoalscorers)
	 
	app.Get("/Pattern/Other/:leagueapi_id/:season/:fixtureapi_id", controllers.GetPatternOther) 
	app.Get("/Pattern/Response/:leagueapi_id/:season", controllers.GetPatternResponse) 
	app.Get("/Pattern/ResponseLeague/:leagueapi_id/:pre_response/:end_response", controllers.GetPatternResponseLeague) 
 
	app.Get("/PatternFrom/Odd/:leagueapi_id/:pre_ah/:pre_gou/:end_ah/:end_gou/:view", controllers.GetPatternFromOdd) 

	app.Get("/PatternFrom/PrePre/:leagueapi_id/:season/:fixtureapi_id", controllers.GetPatternFromPrePre) 
	app.Get("/PatternFrom/PreEnd/:leagueapi_id/:season/:fixtureapi_id", controllers.GetPatternFromPreEnd) 
	app.Get("/PatternFrom/OnlyPre/:leagueapi_id/:season/:fixtureapi_id", controllers.GetPatternFromOnlyPre) 
	app.Get("/PatternFrom/OnlyEnd/:leagueapi_id/:season/:fixtureapi_id", controllers.GetPatternFromOnlyEnd) 
	
	app.Get("/PatternFromCountry/PrePre/:leagueapi_id/:season/:fixtureapi_id", controllers.GetPatternFromCountryPrePre) 
	app.Get("/PatternFromCountry/PreEnd/:leagueapi_id/:season/:fixtureapi_id", controllers.GetPatternFromCountryPreEnd) 
	app.Get("/PatternFromCountry/OnlyPre/:leagueapi_id/:season/:fixtureapi_id", controllers.GetPatternFromCountryOnlyPre) 
	app.Get("/PatternFromCountry/OnlyEnd/:leagueapi_id/:season/:fixtureapi_id", controllers.GetPatternFromCountryOnlyEnd) 

	
	app.Get("/MatchFinished", controllers.GetMatchFinished) 

	
	app.Get("/MatchDone", controllers.GetMatchDone)

	app.Get("/Teams/:id", controllers.GetTeam) 

	
	app.Get("/AnytimeGoalScorers/Today", controllers.GetAgsToday)
	app.Get("/AnytimeGoalScorers/Leagues/:leagueapi_id/:season", controllers.GetAgsLeagues)

	
	app.Get("/ModelTesting/Leagues/:leagueapi_id/:season", controllers.GetMTLeagues)

}
