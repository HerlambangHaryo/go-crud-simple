package models

    import ( 
        "github.com/HerlambangHaryo/go-crud-simple/platform/database" 
        // "time" 
    ) 
	
	type ResearchOnAnytimeGoalScorer struct { 
        ID               	uint   `json:"id" gorm:"primaryKey"`  
        Season           	string `json:"season"` 
		Value        		string `gorm:"default:null"`
		PreOdd       		float64
		EndOdd       		float64
		Status       		int `gorm:"default:null"` 
  
        LeagueapiID           	uint `gorm:"column:leagueapi_id"` 
        LeagueTeam 		*FootballLeague 		`gorm:"foreignKey:LeagueapiID;references:LeagueAPIID"` 
 
		FixtureAPIID uint `gorm:"column:fixtureapi_id"` 
		FixtureTeam *FootballFixture `gorm:"foreignKey:FixtureAPIID;references:FixtureAPIID"` // note the pointer here 

		
        ValueStr    		string `gorm:"column:value_str"` 

		
		PreOddStr       		string
		EndOddStr       		string
		
	}
  
	func (b *ResearchOnAnytimeGoalScorer) GetAgsByToday() ([]ResearchOnAnytimeGoalScorer, error) {
		var scorers []ResearchOnAnytimeGoalScorer
	
		// Subquery to get fixtureapi_ids from football_fixtures
		var fixtureAPIIDs []int
		if err := database.DB.Table("football_fixtures").
			Select("fixtureapi_id").
			Where("date BETWEEN CURRENT_DATE AND CURDATE() + INTERVAL 1 DAY").
			Pluck("fixtureapi_id", &fixtureAPIIDs).Error; err != nil {
			return nil, err
		}
	
		// Query to get ResearchOnAnytimeGoalScorers based on the fixtureapi_ids
		if err := database.DB.
			Select("*", "REPLACE(value, ' ', '_') AS value_str").   
			Preload("LeagueTeam"). 
            Preload("LeagueTeam.CountryTeam"). 
            Preload("FixtureTeam.HomeTeam").
            Preload("FixtureTeam.AwayTeam").
			Where("fixtureapi_id IN ?", fixtureAPIIDs).Find(&scorers).Error; err != nil {
			return nil, err
		}
	
		return scorers, nil
	} 

	func (b *ResearchOnAnytimeGoalScorer) GetAgsByLeagues(leagueAPIID uint, season uint) ([]ResearchOnAnytimeGoalScorer, error) {
		var researchOnAnytimeGoalScorers []ResearchOnAnytimeGoalScorer
	
		err := database.DB.
			Select("*", "REPLACE(value, ' ', '_') AS value_str").
			Where("leagueapi_id = ?", leagueAPIID).
			Preload("LeagueTeam"). 
            Preload("LeagueTeam.CountryTeam").  
            Preload("FixtureTeam"). 
            Preload("FixtureTeam.HomeTeam").
            Preload("FixtureTeam.AwayTeam").
			Find(&researchOnAnytimeGoalScorers).Error
	
		return researchOnAnytimeGoalScorers, err
	}
	
	// func (b *FootballOdd) GetPatternResponse(leagueapiID uint, season uint) ([]FootballOdd, error) {
	// 	var footballodd []FootballOdd
	
	// 	err := database.DB.
	// 		Select("pre_response, end_response, count(*) as counter").
	// 		Where("leagueapi_id = ?", leagueapiID). 
	// 		Group("pre_response, end_response").
	// 		Find(&footballodd).Error
	
	// 	return footballodd, err
	// }
	
	
	