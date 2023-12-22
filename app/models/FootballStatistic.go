package models

	import ( 
        "github.com/HerlambangHaryo/go-crud-simple/platform/database" 
        "time"  
        "gorm.io/gorm"
	)

	type FootballStatistic struct {
		gorm.Model
		ID               		int       `gorm:"column:id;NOT NULL"`
		Date                  	time.Time `gorm:"column:date"`
		FixtureAPIID          	int       `gorm:"column:fixtureapi_id"`
		LeagueapiID           	int       `gorm:"column:leagueapi_id"`
		Season                	int `gorm:"column:season"`
		FixtureStatus         	string    `gorm:"column:fixture_status"`
		TeamsHomeapiID        	int       `gorm:"column:teams_homeapi_id"`
		TeamsAwayapiID        int       `gorm:"column:teams_awayapi_id"`
		GoalsHome             int       `gorm:"column:goals_home"`
		GoalsAway             int       `gorm:"column:goals_away"`
		ScoreHalftimeHome     int       `gorm:"column:score_halftime_home"`
		ScoreHalftimeAway     int       `gorm:"column:score_halftime_away"`
		ScoreSecondtimeHome   int       `gorm:"column:score_secondtime_home"`
		ScoreSecondtimeAway   int       `gorm:"column:score_secondtime_away"`
		ScoreFulltimeHome     int       `gorm:"column:score_fulltime_home"`
		ScoreFulltimeAway     int       `gorm:"column:score_fulltime_away"`
		ScoreExtratimeHome    int       `gorm:"column:score_extratime_home"`
		ScoreExtratimeAway    int       `gorm:"column:score_extratime_away"`
		ScorePenaltyHome      int       `gorm:"column:score_penalty_home"`
		ScorePenaltyAway      int       `gorm:"column:score_penalty_away"`
		LineupsCoachHomeapiID int       `gorm:"column:lineups_coach_homeapi_id"`
		LineupsCoachAwayapiID int       `gorm:"column:lineups_coach_awayapi_id"`
		LineupsFormationHome  string    `gorm:"column:lineups_formation_home"`
		LineupsFormationAway  string    `gorm:"column:lineups_formation_away"`
		ShotsOnGoalHome       int       `gorm:"column:shots_on_goal_home"`
		ShotsOnGoalAway       int       `gorm:"column:shots_on_goal_away"`
		ShotsOffGoalHome      int       `gorm:"column:shots_off_goal_home"`
		ShotsOffGoalAway      int       `gorm:"column:shots_off_goal_away"`
		TotalShotsHome        int       `gorm:"column:total_shots_home"`
		TotalShotsAway        int       `gorm:"column:total_shots_away"`
		BlockedShotsHome      int       `gorm:"column:blocked_shots_home"`
		BlockedShotsAway      int       `gorm:"column:blocked_shots_away"`
		ShotsInsideBoxHome    int       `gorm:"column:shots_inside_box_home"`
		ShotsInsideBoxAway    int       `gorm:"column:shots_inside_box_away"`
		ShotsOutsideBoxHome   int       `gorm:"column:shots_outside_box_home"`
		ShotsOutsideBoxAway   int       `gorm:"column:shots_outside_box_away"`
		FoulsHome             int       `gorm:"column:fouls_home"`
		FoulsAway             int       `gorm:"column:fouls_away"`
		CornerKicksHome       int       `gorm:"column:corner_kicks_home"`
		CornerKicksAway       int       `gorm:"column:corner_kicks_away"`
		OffsidesHome          int       `gorm:"column:offsides_home"`
		OffsidesAway          int       `gorm:"column:offsides_away"`
		BallPossessionHome    int       `gorm:"column:ball_possession_home"`
		BallPossessionAway    int       `gorm:"column:ball_possession_away"`
		YellowCardsHome       int       `gorm:"column:yellow_cards_home;default:0"`
		YellowCardsAway       int       `gorm:"column:yellow_cards_away;default:0"`
		RedCardsHome          int       `gorm:"column:red_cards_home;default:0"`
		RedCardsAway          int       `gorm:"column:red_cards_away;default:0"`
		GoalkeeperSavesHome   int       `gorm:"column:goalkeeper_saves_home"`
		GoalkeeperSavesAway   int       `gorm:"column:goalkeeper_saves_away"`
		TotalPassessHome      int       `gorm:"column:total_passess_home"`
		TotalPassessAway      int       `gorm:"column:total_passess_away"`
		PassessAccurateHome   int       `gorm:"column:passess_accurate_home"`
		PassessAccurateAway   int       `gorm:"column:passess_accurate_away"`
		PassessHome           int       `gorm:"column:passess_home"`
		PassessAway           int       `gorm:"column:passess_away"`
		ExpectedGoalsHome     float64   `gorm:"column:expected_goals_home"`
		ExpectedGoalsAway     float64   `gorm:"column:expected_goals_away"`
		CornerKicksTotalHome  int       `gorm:"column:corner_kicks_total_home"`
		CornerKicksAvgHome    float64   `gorm:"column:corner_kicks_avg_home"`
		CornerKicksMaxHome    int       `gorm:"column:corner_kicks_max_home"`
		CornerKicksMinHome    int       `gorm:"column:corner_kicks_min_home"`
		TotalShotsTotalHome   int       `gorm:"column:total_shots_total_home"`
		TotalShotsAvgHome     float64   `gorm:"column:total_shots_avg_home"`
		TotalShotsMaxHome     int       `gorm:"column:total_shots_max_home"`
		TotalShotsMinHome     int       `gorm:"column:total_shots_min_home"`
		ShotsOnGoalTotalHome  int       `gorm:"column:shots_on_goal_total_home"`
		ShotsOnGoalAvgHome    float64   `gorm:"column:shots_on_goal_avg_home"`
		ShotsOnGoalMaxHome    int       `gorm:"column:shots_on_goal_max_home"`
		ShotsOnGoalMinHome    int       `gorm:"column:shots_on_goal_min_home"`
		FoulsTotalHome        int       `gorm:"column:fouls_total_home"`
		FoulsAvgHome          float64   `gorm:"column:fouls_avg_home"`
		FoulsMaxHome          int       `gorm:"column:fouls_max_home"`
		FoulsMinHome          int       `gorm:"column:fouls_min_home"`
		OffsidesTotalHome     int       `gorm:"column:offsides_total_home"`
		OffsidesAvgHome       float64   `gorm:"column:offsides_avg_home"`
		OffsidesMaxHome       int       `gorm:"column:offsides_max_home"`
		OffsidesMinHome       int       `gorm:"column:offsides_min_home"`
		CornerKicksTotalAway  int       `gorm:"column:corner_kicks_total_away"`
		CornerKicksAvgAway    float64   `gorm:"column:corner_kicks_avg_away"`
		CornerKicksMaxAway    int       `gorm:"column:corner_kicks_max_away"`
		CornerKicksMinAway    int       `gorm:"column:corner_kicks_min_away"`
		TotalShotsTotalAway   int       `gorm:"column:total_shots_total_away"`
		TotalShotsAvgAway     float64   `gorm:"column:total_shots_avg_away"`
		TotalShotsMaxAway     int       `gorm:"column:total_shots_max_away"`
		TotalShotsMinAway     int       `gorm:"column:total_shots_min_away"`
		ShotsOnGoalTotalAway  int       `gorm:"column:shots_on_goal_total_away"`
		ShotsOnGoalAvgAway    float64   `gorm:"column:shots_on_goal_avg_away"`
		ShotsOnGoalMaxAway    int       `gorm:"column:shots_on_goal_max_away"`
		ShotsOnGoalMinAway    int       `gorm:"column:shots_on_goal_min_away"`
		FoulsTotalAway        int       `gorm:"column:fouls_total_away"`
		FoulsAvgAway          float64   `gorm:"column:fouls_avg_away"`
		FoulsMaxAway          int       `gorm:"column:fouls_max_away"`
		FoulsMinAway          int       `gorm:"column:fouls_min_away"`
		OffsidesTotalAway     int       `gorm:"column:offsides_total_away"`
		OffsidesAvgAway       float64   `gorm:"column:offsides_avg_away"`
		OffsidesMaxAway       int       `gorm:"column:offsides_max_away"`
		OffsidesMinAway       int       `gorm:"column:offsides_min_away"`
		CreatedAt             time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
		UpdatedAt             time.Time `gorm:"column:updated_at"`
		DeletedAt             time.Time `gorm:"column:deleted_at"`

		
		CornerKicks   	int       `gorm:"column:corner_kicks"` 
		TotalShots     	int       `gorm:"column:total_shots"`
		ShotsOnGoal    	int       `gorm:"column:shots_on_goal"`
		Fouls       	int       `gorm:"column:fouls"`
		Offsides       	int       `gorm:"column:offsides"`
	}

	func (m *FootballStatistic) TableName() string {
		return "football_statistics"
	}

	func (b *FootballStatistic) GetStatsPreGames(teamapiID uint) ([]FootballStatistic, error) {
		var footballStatistics []FootballStatistic
	
		err := database.DB.
			Select("fixtureapi_id, leagueapi_id, season, date, "+
				"corner_kicks_home as corner_kicks, "+
				"total_shots_home as total_shots, "+
				"shots_on_goal_home as shots_on_goal, "+
				"fouls_home as fouls, "+
				"offsides_home as offsides").
			Where("teams_homeapi_id = ? AND fixture_status LIKE ?", teamapiID, "%Match Finished%").
			Order("date desc").
			Limit(10).
			Find(&footballStatistics).
			Error
	
		if err != nil {
			return nil, err
		}
	
		// Fetch data for the away team
		err = database.DB.
			Select("fixtureapi_id, leagueapi_id, season, date, "+
				"corner_kicks_away as corner_kicks, "+
				"total_shots_away as total_shots, "+
				"shots_on_goal_away as shots_on_goal, "+
				"fouls_away as fouls, "+
				"offsides_away as offsides").
			Where("teams_awayapi_id = ? AND fixture_status LIKE ?", teamapiID, "%Match Finished%").
			Order("date desc").
			Limit(10).
			Find(&footballStatistics).
			Error
	
		return footballStatistics, err
	}
	