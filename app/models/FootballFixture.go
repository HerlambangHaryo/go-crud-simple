package models
    
    import ( 
        "github.com/HerlambangHaryo/go-crud-simple/platform/database" 
        "time"  
        "gorm.io/gorm"
        "strings"
    ) 
 
    type FootballFixture struct {
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
 
        TypeOfPattern                               int       `gorm:"column:type_of_pattern"`
        OtherPattern                                int       `gorm:"column:other_pattern"`

        TotalPrepre                                 string       `gorm:"column:total_prepre"`
        TotalPrepreCountry                          string       `gorm:"column:total_prepre_country"`
        TotalPrepreWorld                            string       `gorm:"column:total_prepre_world"`
        
        TotalPreend                                 string       `gorm:"column:total_preend"`
        TotalPreendCountry                          string       `gorm:"column:total_preend_country"`
        TotalPreendWorld                            string       `gorm:"column:total_preend_world"`

        SetOne                                 string       `gorm:"column:set_one"`

        // Define associations  
        VenueapiID             	uint `gorm:"column:venueapi_id"`
        TeamsHomeapiID        	uint `gorm:"column:teams_homeapi_id"`
        TeamsAwayapiID      	uint `gorm:"column:teams_awayapi_id"`
        // LeagueapiID           	uint `gorm:"column:leagueapi_id"` 
        StandingHomeID        	string `gorm:"column:standing_home_id"`
        StandingAwayID      	string `gorm:"column:standing_away_id"`
        
        VenueTeam 		*FootballVenue 		`gorm:"foreignKey:VenueapiID;references:VenueAPIID"`
        HomeTeam 		*FootballTeam 		`gorm:"foreignKey:TeamsHomeapiID;references:TeamAPIID"`
        AwayTeam 		*FootballTeam 		`gorm:"foreignKey:TeamsAwayapiID;references:TeamAPIID"` 

        LeagueapiID           	uint `gorm:"column:leagueapi_id"` 
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

        // ---------------------------------------- 
        
        PreAvgGoalsHome    		string `gorm:"column:pre_avg_goals_home"` 
        PreAvgGoalsAway    		string `gorm:"column:pre_avg_goals_away"` 
        PreAvgHandicap    		string `gorm:"column:pre_avg_handicap"` 
        PreAvgOverunder    		string `gorm:"column:pre_avg_overunder"` 
    }
  
    func (b *FootballFixture) GetFixtureInfoOdd(leagueapiID uint, season uint, fixtureapi_id uint) ([]FootballFixture, error) {
        var footballfixture []FootballFixture

        err := database.DB.
            Select("*").
            Where("leagueapi_id = ?", leagueapiID).
            Where("season = ?", season).
            Where("fixtureapi_id = ?", fixtureapi_id).
            Preload("HomeTeam").
            Preload("AwayTeam").
            Preload("VenueTeam").
            Preload("LeagueTeam").
            Preload("OddTeam").
            Find(&footballfixture).Error

        return footballfixture, err
    } 

    func (b *FootballFixture) GetFixtureGeneral(leagueapiID uint, season uint, fixtureapi_id uint) ([]FootballFixture, error) {
        var footballfixture []FootballFixture

        err := database.DB.
            Select("*").
            Where("leagueapi_id = ?", leagueapiID).
            Where("season = ?", season).
            Where("fixtureapi_id = ?", fixtureapi_id).
            Preload("HomeTeam").
            Preload("AwayTeam").
            Preload("VenueTeam").
            Preload("LeagueTeam").
            Preload("OddTeam"). 
            Find(&footballfixture).Error

        return footballfixture, err
    } 

    func (b *FootballFixture) GetFixturePrePre(leagueapiID uint, season uint, fixtureapi_id uint) ([]FootballFixture, error) {
        var footballfixture []FootballFixture

        err := database.DB.
            Select("*").
            Where("leagueapi_id = ?", leagueapiID).
            Where("season = ?", season).
            Where("fixtureapi_id = ?", fixtureapi_id).
            Preload("HomeTeam").
            Preload("AwayTeam").
            Preload("VenueTeam").
            Preload("LeagueTeam").
            Preload("OddTeam").
            Find(&footballfixture).Error

        return footballfixture, err
    } 
    
    func (b *FootballFixture) GetFixturePreEnd(leagueapiID uint, season uint, fixtureapi_id uint) ([]FootballFixture, error) {
        var footballfixture []FootballFixture

        err := database.DB.
            Select("*").
            Where("leagueapi_id = ?", leagueapiID).
            Where("season = ?", season).
            Where("fixtureapi_id = ?", fixtureapi_id).
            Preload("HomeTeam").
            Preload("AwayTeam").
            Preload("VenueTeam").
            Preload("LeagueTeam").
            Preload("OddTeam").
            Find(&footballfixture).Error

        return footballfixture, err
    } 

    func (b *FootballFixture) GetFixtureOnlyPre(leagueapiID uint, season uint, fixtureapi_id uint) ([]FootballFixture, error) {
        var footballfixture []FootballFixture

        err := database.DB.
            Select("*").
            Where("leagueapi_id = ?", leagueapiID).
            Where("season = ?", season).
            Where("fixtureapi_id = ?", fixtureapi_id).
            Preload("HomeTeam").
            Preload("AwayTeam").
            Preload("VenueTeam").
            Preload("LeagueTeam").
            Preload("OddTeam").
            Find(&footballfixture).Error

        return footballfixture, err
    } 
    
    func (b *FootballFixture) GetFixtureOnlyEnd(leagueapiID uint, season uint, fixtureapi_id uint) ([]FootballFixture, error) {
        var footballfixture []FootballFixture

        err := database.DB.
            Select("*").
            Where("leagueapi_id = ?", leagueapiID).
            Where("season = ?", season).
            Where("fixtureapi_id = ?", fixtureapi_id).
            Preload("HomeTeam").
            Preload("AwayTeam").
            Preload("VenueTeam").
            Preload("LeagueTeam").
            Preload("OddTeam").
            Find(&footballfixture).Error

        return footballfixture, err
    } 
    
    func (b *FootballFixture) GetLeagueNotstarted(leagueapiID uint, season uint) ([]FootballFixture, error) {
        var footballfixture []FootballFixture
    
        err := database.DB.
            Select("*", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H:%i:%s') as jam"). 
            Where("fixture_status LIKE ?", "Not Started%").
            Where("leagueapi_id = ?", leagueapiID).
            Where("season = ?", season).
            Where("deleted_at IS NULL").
            Where("date < DATE_ADD(CURDATE(), INTERVAL 1 WEEK)").
            Order("date asc").
            
            // Joins("HomeTeam").
            // Joins("AwayTeam").
            // Joins("VenueTeam").
            // Joins("LeagueTeam"). 
            // Joins("OddTeam").

            Preload("HomeTeam").
            Preload("AwayTeam").
            Preload("VenueTeam").
            Preload("LeagueTeam").
            Preload("OddTeam"). 
            Preload("LeagueTeam.CountryTeam"). 

            Preload("HomeStandingTeam"). 
            Preload("AwayStandingTeam"). 

            
            Find(&footballfixture).Error

        return footballfixture, err
    } 
 
    func (b *FootballFixture) GetLeagueMatchfinishedGroup(leagueapiID uint, season uint) ([]FootballFixture, error) {
        var footballfixture []FootballFixture
  
        err := database.DB.
            Select("leagueapi_id", "season", "round", "REPLACE(round, ' ', '_') AS round_us").
            Where("leagueapi_id = ?", leagueapiID).
            Where("season = ?", season).
            Where("fixture_status LIKE ?", "Match Finished%").
            Group("leagueapi_id, season, round ").
            Order("date asc").
            Find(&footballfixture).Error

        return footballfixture, err 
    } 
 
    func (b *FootballFixture) GetLeagueDetailRound(leagueapiID uint, season uint, round string) ([]FootballFixture, error) {
        var footballfixture []FootballFixture

        roundx := strings.ReplaceAll(round, "_", " ")
         
        err := database.DB.
            Select("*", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H:%i:%s') as jam", "DATE_FORMAT(DATE_ADD(pre_odd_updated_at, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal_pre", "DATE_FORMAT(DATE_ADD(pre_odd_updated_at, INTERVAL 7 HOUR), '%H:%i:%s') as jam_pre", "DATE_FORMAT(DATE_ADD(pre_odd_updated_at, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal_end", "DATE_FORMAT(DATE_ADD(pre_odd_updated_at, INTERVAL 7 HOUR), '%H:%i:%s') as jam_end"). 
            Where("leagueapi_id = ?", leagueapiID).
            Where("season = ?", season).
            Where("round = ?", roundx). 
            Where("deleted_at IS NULL"). 
            Order("date asc"). 
            Preload("HomeTeam").
            Preload("AwayTeam").
            Preload("VenueTeam").
            Preload("LeagueTeam").
            Preload("OddTeam").
            Preload("LeagueTeam.CountryTeam"). 
            Preload("StatTeam").
            Find(&footballfixture).Error

        return footballfixture, err
    } 
  
    func (b *FootballFixture) UpdateSingleColumnByID(id uint, columnName string, newValue interface{}) error {
        if err := database.DB.First(b, id).Error; err != nil {
            return err
        }
    
        if err := database.DB.Model(b).Update(columnName, newValue).Error; err != nil {
            return err
        } 

        return nil
    }
  
    func (b *FootballFixture) UpdateClearSetOne(id uint, columnName string, newValue interface{}) error {
     
    
        if err := database.DB.Model(b).Where("set_one = ?", id).Update("set_one", newValue).Error; err != nil {
            return err
        } 
        
        return nil
    }

     

    // ---------------------------------------------------------------------------------------
 
    func (b *FootballFixture) GetToday() ([]FootballFixture, error) {
        var footballfixture []FootballFixture
    
        err := database.DB.
            Select("leagueapi_id", "season", "count(*) as Counter"). 
            Where("fixture_status LIKE ?", "Not Started%").
            Where("deleted_at IS NULL").
            Where("date BETWEEN CURRENT_DATE AND CURDATE() + INTERVAL 1 DAY").
            Group("leagueapi_id, season").
            Order("date asc").
            Preload("LeagueTeam"). 
            Preload("LeagueTeam.CountryTeam"). 
            Preload("OddTeam"). 
            Find(&footballfixture).Error

        return footballfixture, err
    } 
 
	func (b *FootballFixture) GetTodayTimegroup() ([]FootballFixture, error) {
		var footballfixture []FootballFixture
    
        err := database.DB.
            Select("DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y') as yeary", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%m') as monthy", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%d') as dayy", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H') as houry", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%i') as minutey", "count(*) as counter").  
            Where("deleted_at IS NULL").
            Where("date BETWEEN CURRENT_DATE AND CURDATE() + INTERVAL 1 DAY").
            Group("DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y'), DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%m'), DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%d'), DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H'), DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%i')"). 
            Order("date asc").
            Find(&footballfixture).Error

        return footballfixture, err  
	}
    
    func (b *FootballFixture) GetTodayTime(yearInt uint, monthInt uint, dayInt uint, hourInt uint, minuteInt uint) ([]FootballFixture, error) {
        var footballfixture []FootballFixture 
  
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
            Preload("HomeTeam").
            Preload("AwayTeam").
            Preload("VenueTeam").
            Preload("LeagueTeam"). 
            Preload("OddTeam"). 
            Find(&footballfixture).Error

        return footballfixture, err
    } 

    func (b *FootballFixture) GetTodayLeague(leagueapiID uint, season uint) ([]FootballFixture, error) {
        var footballfixture []FootballFixture 

        err := database.DB.
            Select("*", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H:%i:%s') as jam"). 
            Where("fixture_status LIKE ?", "Not Started%").
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
            Find(&footballfixture).Error

        return footballfixture, err
    } 

    // ---------------------------------------------------------------------------------------
 
    func (b *FootballFixture) GetTomorrow() ([]FootballFixture, error) {
        var footballfixture []FootballFixture
    
        err := database.DB.
            Select("leagueapi_id", "season", "count(*) as Counter"). 
            Where("fixture_status LIKE ?", "Not Started%").
            Where("deleted_at IS NULL"). 
            Where("date BETWEEN CURDATE() + INTERVAL 1 DAY AND CURDATE() + INTERVAL 2 DAY"). 
            Group("leagueapi_id, season").
            Order("date asc").
            Preload("LeagueTeam"). 
            Preload("LeagueTeam.CountryTeam"). 
            Find(&footballfixture).Error

        return footballfixture, err
    } 
 
	func (b *FootballFixture) GetTomorrowTimegroup() ([]FootballFixture, error) {
		var footballfixture []FootballFixture
    
        err := database.DB.
            Select("DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y') as yeary", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%m') as monthy", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%d') as dayy", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H') as houry", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%i') as minutey", "count(*) as counter").  
            Where("deleted_at IS NULL").
            Where("date BETWEEN CURDATE() + INTERVAL 1 DAY AND CURDATE() + INTERVAL 2 DAY"). 
            Group("DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y'), DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%m'), DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%d'), DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H'), DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%i')"). 
            Order("date asc").
            Find(&footballfixture).Error

        return footballfixture, err  
	}
    
    func (b *FootballFixture) GetTomorrowTime(yearInt uint, monthInt uint, dayInt uint, hourInt uint, minuteInt uint) ([]FootballFixture, error) {
        var footballfixture []FootballFixture 
  
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
            Preload("HomeTeam").
            Preload("AwayTeam").
            Preload("VenueTeam").
            Preload("LeagueTeam"). 
            Preload("OddTeam"). 
            Find(&footballfixture).Error

        return footballfixture, err
    } 

    func (b *FootballFixture) GetTomorrowLeague(leagueapiID uint, season uint) ([]FootballFixture, error) {
        var footballfixture []FootballFixture 

        err := database.DB.
            Select("*", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H:%i:%s') as jam"). 
            Where("fixture_status LIKE ?", "Not Started%").
            Where("leagueapi_id = ?", leagueapiID).
            Where("season = ?", season).
            Where("deleted_at IS NULL").
            Where("date BETWEEN CURDATE() + INTERVAL 1 DAY AND CURDATE() + INTERVAL 2 DAY"). 
            Order("date asc").
            Preload("HomeTeam").
            Preload("AwayTeam").
            Preload("VenueTeam").
            Preload("LeagueTeam").
            Preload("OddTeam").
            Preload("LeagueTeam.CountryTeam"). 
            Find(&footballfixture).Error

        return footballfixture, err
    } 

    // ---------------------------------------------------------------------------------------
    
    func (b *FootballFixture) GetMatchFinished() ([]FootballFixture, error) {
        var footballfixture []FootballFixture
    
        err := database.DB.
            Select("leagueapi_id", "season", "count(*) as Counter"). 
            Where("fixture_status LIKE ?", "Match Finished").
            Where("deleted_at IS NULL").
            Where("date BETWEEN DATE_SUB(CURRENT_DATE, INTERVAL 7 DAY) AND CURRENT_DATE").
            Group("leagueapi_id, season").
            Order("date asc").
            Preload("LeagueTeam"). 
            Preload("LeagueTeam.CountryTeam"). 
            Find(&footballfixture).Error

        return footballfixture, err
    } 
 
	func (b *FootballFixture) GetMatchFinishedTimegroup() ([]FootballFixture, error) {
		var footballfixture []FootballFixture
    
        err := database.DB.
            Select("DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y') as yeary", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%m') as monthy", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%d') as dayy", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H') as houry", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%i') as minutey", "count(*) as counter").  
            Where("deleted_at IS NULL").
            Where("date BETWEEN CURRENT_DATE AND CURDATE() + INTERVAL 1 DAY").
            Group("DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y'), DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%m'), DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%d'), DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H'), DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%i')"). 
            Order("date asc").
            Find(&footballfixture).Error

        return footballfixture, err  
	}
    
    func (b *FootballFixture) GetMatchFinishedTime(yearInt uint, monthInt uint, dayInt uint, hourInt uint, minuteInt uint) ([]FootballFixture, error) {
        var footballfixture []FootballFixture  

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
        
            Find(&footballfixture).Error

        return footballfixture, err
    } 

    func (b *FootballFixture) GetMatchFinishedLeague(leagueapiID uint, season uint) ([]FootballFixture, error) {
        var footballfixture []FootballFixture 

        err := database.DB.
            Select("*", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H:%i:%s') as jam"). 
            Where("fixture_status LIKE ?", "Not Started%").
            Where("leagueapi_id = ?", leagueapiID).
            Where("season = ?", season).
            Where("deleted_at IS NULL").
            Where("date BETWEEN DATE_SUB(CURRENT_DATE, INTERVAL 7 DAY) AND CURRENT_DATE"). 
            Order("date asc").
            Preload("HomeTeam").
            Preload("AwayTeam").
            Preload("VenueTeam").
            Preload("LeagueTeam").
            Preload("OddTeam").
            Preload("LeagueTeam.CountryTeam"). 
            Find(&footballfixture).Error

        return footballfixture, err
    } 

    // ---------------------------------------------------------------------------------------
 
    func (b *FootballFixture) GetMatchFinishedEnded() ([]FootballFixture, error) {
        var footballfixture []FootballFixture
    
        err := database.DB.
            Select("leagueapi_id", "season", "count(*) as Counter"). 
            Where("fixture_status LIKE ?", "Match Finished").
            Where("deleted_at IS NULL").
            Where("date BETWEEN DATE_SUB(CURRENT_DATE, INTERVAL 7 DAY) AND CURRENT_DATE").
            Group("leagueapi_id, season").
            Order("date asc").
            Preload("LeagueTeam"). 
            Preload("LeagueTeam.CountryTeam"). 
            Find(&footballfixture).Error

        return footballfixture, err
    } 
 
	func (b *FootballFixture) GetMatchFinishedEndedTimegroup() ([]FootballFixture, error) {
		var footballfixture []FootballFixture
    
        err := database.DB.
            Select("DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y') as yeary", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%m') as monthy", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%d') as dayy", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H') as houry", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%i') as minutey", "count(*) as counter").  
            Where("deleted_at IS NULL").
            Where("date BETWEEN CURRENT_DATE AND CURDATE() + INTERVAL 1 DAY").
            Group("DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y'), DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%m'), DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%d'), DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H'), DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%i')"). 
            Order("date asc").
            Find(&footballfixture).Error

        return footballfixture, err  
	}
    
    func (b *FootballFixture) GetMatchFinishedEndedTime(yearInt uint, monthInt uint, dayInt uint, hourInt uint, minuteInt uint) ([]FootballFixture, error) {
        var footballfixture []FootballFixture  

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
        
            Find(&footballfixture).Error

        return footballfixture, err
    } 

    func (b *FootballFixture) GetMatchFinishedEndedLeague(leagueapiID uint, season uint) ([]FootballFixture, error) {
        var footballfixture []FootballFixture 

        err := database.DB.
            Select("*", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H:%i:%s') as jam"). 
            Where("fixture_status LIKE ?", "Not Started%").
            Where("leagueapi_id = ?", leagueapiID).
            Where("season = ?", season).
            Where("deleted_at IS NULL").
            Where("date BETWEEN DATE_SUB(CURRENT_DATE, INTERVAL 7 DAY) AND CURRENT_DATE"). 
            Order("date asc").
            Preload("HomeTeam").
            Preload("AwayTeam").
            Preload("VenueTeam").
            Preload("LeagueTeam").
            Preload("OddTeam").
            Preload("LeagueTeam.CountryTeam"). 
            Find(&footballfixture).Error

        return footballfixture, err
    } 

    // ---------------------------------------------------------------------------------------
    

    // ---------------------------------------------------------------------------------------
 
    func (b *FootballFixture) GetUltimate() ([]FootballFixture, error) {
        var footballfixture []FootballFixture


        
        err := database.DB.
            Select("*", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H:%i:%s') as jam"). 
            Where("fixture_status LIKE ?", "Not Started Goto"). 
            Where("deleted_at IS NULL").
            Where("date BETWEEN DATE_SUB(NOW(), INTERVAL 7 HOUR) AND CURDATE() + INTERVAL 1 DAY"). 
            Order("date asc").
            Preload("HomeTeam").
            Preload("AwayTeam"). 
            Preload("LeagueTeam"). 
            Preload("LeagueTeam.CountryTeam"). 
            Find(&footballfixture).Error

        return footballfixture, err

        // err := database.DB.
        //     Select("leagueapi_id", "season", "count(*) as Counter"). 
        //     Where("fixture_status LIKE ?", "Not Started Goto").
        //     Where("deleted_at IS NULL").
        //     Where("date BETWEEN CURRENT_DATE AND CURDATE() + INTERVAL 1 DAY").
        //     Group("leagueapi_id, season").
        //     Order("date asc").
        //     Preload("LeagueTeam"). 
        //     Preload("LeagueTeam.CountryTeam"). 
        //     Find(&footballfixture).Error

        // return footballfixture, err

        
		// db := database.DB.
        //     Raw(`SELECT *
        //         FROM football_fixtures 
        //         WHERE date BETWEEN CURRENT_DATE AND CURDATE() + INTERVAL 1 DAY
        //         AND other_pattern is not null 
        //         AND deleted_at IS NULL
        //         UNION 
        //         SELECT *
        //         FROM football_fixtures 
        //         WHERE date BETWEEN CURRENT_DATE AND CURDATE() + INTERVAL 1 DAY
        //         AND total_prepre is not null 
        //         AND deleted_at IS NULL
        //         UNION 
        //         SELECT *
        //         FROM football_fixtures 
        //         WHERE date BETWEEN CURRENT_DATE AND CURDATE() + INTERVAL 1 DAY
        //         AND total_prepre_country is not null 
        //         AND deleted_at IS NULL
        //         UNION 
        //         SELECT *
        //         FROM football_fixtures 
        //         WHERE date BETWEEN CURRENT_DATE AND CURDATE() + INTERVAL 1 DAY
        //         AND total_prepre_world is not null 
        //         AND deleted_at IS NULL
        //         `)
 
        // err := db.Find(&footballfixture).Error

        // if err != nil {
        //     return nil, err
        // }
    
        // // Now, you can preload related data
        // for i := range footballfixture {
        //     db.Order("date asc").
        //         Preload("HomeTeam").
        //         Preload("AwayTeam").
        //         Preload("VenueTeam").
        //         Preload("LeagueTeam").
        //         Preload("OddTeam").
        //         Preload("LeagueTeam.CountryTeam"). 
        //         Where("id = ?", footballfixture[i].ID).Find(&footballfixture[i])
        // }
    
        // return footballfixture, err   
    } 
 
	func (b *FootballFixture) GetUltimateTimegroup() ([]FootballFixture, error) {
		var footballfixture []FootballFixture
    
        err := database.DB.
            Select("DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y') as yeary", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%m') as monthy", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%d') as dayy", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H') as houry", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%i') as minutey", "count(*) as counter").  
            Where("deleted_at IS NULL").
            Where("date BETWEEN CURRENT_DATE AND CURDATE() + INTERVAL 1 DAY").
            Group("DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y'), DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%m'), DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%d'), DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H'), DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%i')"). 
            Order("date asc").
            Find(&footballfixture).Error

        return footballfixture, err  
	}
    
    func (b *FootballFixture) GetUltimateTime(yearInt uint, monthInt uint, dayInt uint, hourInt uint, minuteInt uint) ([]FootballFixture, error) {
        var footballfixture []FootballFixture 

        // err := database.DB.
        //     Select("*", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H:%i:%s') as jam"). 
        //     Where("fixture_status LIKE ?", "Not Started%").
        //     Where("year( DATE_ADD(date, INTERVAL 7 HOUR) ) = ?", yearInt). 
        //     Where("month( DATE_ADD(date, INTERVAL 7 HOUR) ) = ?", monthInt). 
        //     Where("day( DATE_ADD(date, INTERVAL 7 HOUR) ) = ?", dayInt). 
        //     Where("hour( DATE_ADD(date, INTERVAL 7 HOUR) ) = ?", hourInt). 
        //     Where("minute( DATE_ADD(date, INTERVAL 7 HOUR) ) = ?", minuteInt). 
        //     Where("deleted_at IS NULL"). 
        //     Order("date asc").  
        //     Preload("HomeTeam").
        //     Preload("AwayTeam").
        //     Preload("VenueTeam").
        //     Preload("LeagueTeam"). 
        //     Preload("OddTeam").Or("1=1").  
        //     Find(&footballfixture).Error

        // return footballfixture, err

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
        
            Find(&footballfixture).Error

        return footballfixture, err
    } 

    func (b *FootballFixture) GetUltimateLeague(leagueapiID uint, season uint) ([]FootballFixture, error) {
        var footballfixture []FootballFixture 

        err := database.DB.
            Select("*", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H:%i:%s') as jam"). 
            Where("fixture_status LIKE ?", "Not Started%").
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
            Find(&footballfixture).Error

        return footballfixture, err
    } 

    // ---------------------------------------------------------------------------------------
    
    func (b *FootballFixture) GetOne() ([]FootballFixture, error) {
        var footballfixture []FootballFixture 

        err := database.DB.
            Select("*", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal", "DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H:%i:%s') as jam").  
            Where("set_one = 1").
            Where("deleted_at IS NULL"). 
            Order("date asc").
            Preload("HomeTeam").
            Preload("AwayTeam").
            Preload("VenueTeam").
            Preload("LeagueTeam").
            Preload("OddTeam").
            Preload("LeagueTeam.CountryTeam"). 
            Find(&footballfixture).Error

        return footballfixture, err
    } 

    // ---------------------------------------------------------------------------------------
    
    func (b *FootballFixture) GetMatchDone() ([]FootballFixture, error) {
        var footballfixture []FootballFixture
    
        err := database.DB.
            Select("leagueapi_id", "season", "count(*) as Counter").  
            Where("deleted_at IS NULL").
            Where("date <= CURRENT_DATE").
            Where("fixture_status not in ('Match Finished', 'Match Finished Ended', 'Walkover', 'Match Cancelled', 'Match Abandoned', 'Match Postponed', 'Technical loss', 'Time to be defined') ").
            Group("leagueapi_id, season").
            Order("date asc").
            Preload("LeagueTeam"). 
            Preload("LeagueTeam.CountryTeam"). 
            Find(&footballfixture).Error

        return footballfixture, err
    } 

    // ---------------------------------------------------------------------------------------
 