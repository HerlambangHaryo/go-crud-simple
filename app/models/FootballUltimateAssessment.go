package models
    
    import ( 
        "github.com/HerlambangHaryo/go-crud-simple/platform/database" 
        "time"  
        "gorm.io/gorm" 
    ) 
 
    type FootballUltimateAssessment struct {
        gorm.Model
        ID                                       int       `gorm:"column:id;NOT NULL"`
        Date                                     string `gorm:"column:date"`
        Referee                                  string    `gorm:"column:referee"`
        Season                                   int 	   `gorm:"column:season"`
        Round                                    string    `gorm:"column:round"`
        FixtureStatus                            string    `gorm:"column:fixture_status"`
        FixtureElapsed                           int       `gorm:"column:fixture_elapsed"`
        GoalsHome                                string       `gorm:"column:goals_home"`
        GoalsAway                                string       `gorm:"column:goals_away"`
        ScoreHalftimeHome                        string       `gorm:"column:score_halftime_home"`
        ScoreHalftimeAway                        string       `gorm:"column:score_halftime_away"`
        ScoreSecondtimeHome                      string       `gorm:"column:score_secondtime_home"`
        ScoreSecondtimeAway                      string       `gorm:"column:score_secondtime_away"`
        ScoreFulltimeHome                        string       `gorm:"column:score_fulltime_home"`
        ScoreFulltimeAway                        string       `gorm:"column:score_fulltime_away"`
        ScoreExtratimeHome                       string       `gorm:"column:score_extratime_home"`
        ScoreExtratimeAway                       string       `gorm:"column:score_extratime_away"`
        ScorePenaltyHome                         string       `gorm:"column:score_penalty_home"`
        ScorePenaltyAway                         string       `gorm:"column:score_penalty_away"`
        TypeOfPattern                            int       `gorm:"column:type_of_pattern"`
        FixtureUpdatedAt                         *time.Time `gorm:"column:fixture_updated_at"`
        PreOddUpdatedAt                          *time.Time `gorm:"column:pre_odd_updated_at"`
        EndOddUpdatedAt                          *time.Time `gorm:"column:end_odd_updated_at"`
        TeamStatisticsUpdatedAt                  *time.Time `gorm:"column:team_statistics_updated_at"`
        FixtureStatisticsUpdatedAt               *time.Time `gorm:"column:fixture_statistics_updated_at"`
        ApiFixturesPlayersUpdatedAt              *time.Time `gorm:"column:api_fixtures_players_updated_at"`
        ResearchOnAnytimeGoalScorersUpdatedAt    *time.Time `gorm:"column:research_on_anytime_goal_scorers_updated_at"`
        ResearchOnToScoreTwoOrMoreGoalsUpdatedAt *time.Time `gorm:"column:research_on_to_score_two_or_more_goals_updated_at"`
        ResearchOnPlayerToBeBookedUpdatedAt      *time.Time `gorm:"column:research_on_player_to_be_booked_updated_at"`
        CreatedAt                                *time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
        UpdatedAt                                *time.Time `gorm:"column:updated_at"`
        DeletedAt                                *time.Time `gorm:"column:deleted_at"`
        BehavioralBmAssessment                  string       `gorm:"column:behavioral_bm_assessment"`
 
        // Define associations  
        VenueapiID             	uint `gorm:"column:venueapi_id"`
        TeamsHomeapiID        	uint `gorm:"column:teams_homeapi_id"`
        TeamsAwayapiID      	uint `gorm:"column:teams_awayapi_id"`
        LeagueapiID           	uint `gorm:"column:leagueapi_id"` 
        StandingHomeID        	string `gorm:"column:standing_home_id"`
        StandingAwayID      	string `gorm:"column:standing_away_id"`
        
        VenueTeam 		*FootballVenue 		`gorm:"foreignKey:VenueapiID;references:VenueAPIID"`
        HomeTeam 		*FootballTeam 		`gorm:"foreignKey:TeamsHomeapiID;references:TeamAPIID"`
        AwayTeam 		*FootballTeam 		`gorm:"foreignKey:TeamsAwayapiID;references:TeamAPIID"` 
        LeagueTeam 		*FootballLeague 		`gorm:"foreignKey:LeagueapiID;references:LeagueAPIID"`
        StatTeam 		*FootballStatistic   `gorm:"foreignKey:FixtureAPIID;references:FixtureAPIID"`  
        HomeStandingTeam 	*ApiFootballLeagueStanding   `gorm:"foreignKey:StandingHomeID;references:StandingTeamID"`  
        AwayStandingTeam 	*ApiFootballLeagueStanding   `gorm:"foreignKey:StandingAwayID;references:StandingTeamID"`  

        
		// CountryTeam      Country `gorm:"foreignKey:CountryName;references:Name"`

        
        FixtureAPIID uint `gorm:"column:fixtureapi_id"` 
        OddTeam *FootballOdd `gorm:"foreignKey:FixtureAPIID;references:FixtureAPIID"` 

        Counter 	int       `gorm:"column:counter"`
        Tanggal    	string `gorm:"column:tanggal"`
        Jam    		string `gorm:"column:jam"`
        RoundUS     string `gorm:"column:round_us"`

        TanggalPre    	string `gorm:"column:tanggal_pre"`
        JamPre    		string `gorm:"column:jam_pre"`
        TanggalEnd    	string `gorm:"column:tanggal_end"`
        JamEnd    		string `gorm:"column:jam_end"`
        
        Yeary    		string `gorm:"column:yeary"`
        Monthy    		string `gorm:"column:monthy"`
        Dayy    		string `gorm:"column:dayy"`
        Houry    		string `gorm:"column:houry"`
        Minutey    		string `gorm:"column:minutey"` 
    }

	func (m *FootballUltimateAssessment) TableName() string {
		return "football_ultimate_assessments"
	}
  
    func (b *FootballUltimateAssessment) GetUltimate() ([]FootballUltimateAssessment, error) {
        var footballultimateassessment []FootballUltimateAssessment
    
        err := database.DB.
            Select("leagueapi_id", "season", "count(*) as Counter").  
            Where("deleted_at IS NULL").
            Where("date BETWEEN CURRENT_DATE AND CURDATE() + INTERVAL 1 DAY").
            Group("leagueapi_id, season").
            Order("date asc").
            Preload("LeagueTeam"). 
            Preload("LeagueTeam.CountryTeam"). 
            Find(&footballultimateassessment).Error

        return footballultimateassessment, err
    } 
 
	func (b *FootballUltimateAssessment) GetUltimateTimegroup() ([]FootballUltimateAssessment, error) {
		var footballultimateassessment []FootballUltimateAssessment
    
        err := database.DB.
            Select("DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y') as yeary", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%m') as monthy", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%d') as dayy", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H') as houry", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%i') as minutey", "count(*) as counter").  
            Where("deleted_at IS NULL").
            Where("date BETWEEN CURRENT_DATE AND CURDATE() + INTERVAL 1 DAY").
            Group("DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y'), DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%m'), DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%d'), DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H'), DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%i')"). 
            Order("date asc").
            Find(&footballultimateassessment).Error

        return footballultimateassessment, err  
	}
    
    func (b *FootballUltimateAssessment) GetUltimateTime(yearInt uint, monthInt uint, dayInt uint, hourInt uint, minuteInt uint) ([]FootballUltimateAssessment, error) {
        var footballultimateassessment []FootballUltimateAssessment  

        err := database.DB.
            Select("*", "DATE_FORMAT(DATE_ADD(football_fixtures.date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal", "DATE_FORMAT(DATE_ADD(football_fixtures.date, INTERVAL 7 HOUR), '%H:%i:%s') as jam"). 
            Where("fixture_status LIKE ?", "Not Started%").
            Where("year( DATE_ADD(football_fixtures.date, INTERVAL 7 HOUR) ) = ?", yearInt). 
            Where("month( DATE_ADD(football_fixtures.date, INTERVAL 7 HOUR) ) = ?", monthInt). 
            Where("day( DATE_ADD(football_fixtures.date, INTERVAL 7 HOUR) ) = ?", dayInt). 
            Where("hour( DATE_ADD(football_fixtures.date, INTERVAL 7 HOUR) ) = ?", hourInt). 
            Where("minute( DATE_ADD(football_fixtures.date, INTERVAL 7 HOUR) ) = ?", minuteInt). 
            Where("deleted_at IS NULL").
            Where("football_fixtures.date < DATE_ADD(CURDATE(), INTERVAL 1 WEEK)").
            Order("football_fixtures.date asc").
            Joins("HomeTeam").
            Joins("AwayTeam").
            Joins("VenueTeam").
            Joins("LeagueTeam"). 
            Joins("OddTeam"). // Join all the associations
        
            Find(&footballultimateassessment).Error

        return footballultimateassessment, err
    } 

    func (b *FootballUltimateAssessment) GetUltimateLeague(leagueapiID uint, season uint) ([]FootballUltimateAssessment, error) {
        var footballultimateassessment []FootballUltimateAssessment 

        err := database.DB.
            Select("*", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H:%i:%s') as jam"). 
            Where("leagueapi_id = ?", leagueapiID).
            Where("season = ?", season).
            Where("deleted_at IS NULL").
            Where("date BETWEEN CURRENT_DATE AND CURDATE() + INTERVAL 1 DAY"). 
            Order("date asc").
            Preload("HomeTeam").
            Preload("AwayTeam").
            Preload("VenueTeam").
            Preload("LeagueTeam").
            Preload("OddTeam").
            Preload("LeagueTeam.CountryTeam"). 
            Find(&footballultimateassessment).Error

        return footballultimateassessment, err
    } 