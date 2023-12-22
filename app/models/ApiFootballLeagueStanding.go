package models
 
	import ( 
		"github.com/HerlambangHaryo/pr_go_buzz_2023_11/platform/database"  
		"time"  
		"gorm.io/gorm"
	)

	type ApiFootballLeagueStanding struct {
		gorm.Model   
		StandingTeamID                         string    `gorm:"column:standing_team_id"`
		LeagueapiID                            int       `gorm:"column:leagueapi_id;NOT NULL"`
		Season                                 int       `gorm:"column:season;NOT NULL"`
		Rank                                   int       `gorm:"column:rank;NOT NULL"`
		TeamapiID                              uint       `gorm:"column:teamapi_id;NOT NULL"`
		Points                                 int       `gorm:"column:points"`
		GoalsDiff                              int       `gorm:"column:goals_diff"`
		GroupStatus                            string    `gorm:"column:group_status"`
		Form                                   string    `gorm:"column:form"`
		LongForm                               string    `gorm:"column:long_form"`
		Status                                 string    `gorm:"column:status"`
		Description                            string    `gorm:"column:description"`
		Played                                 int       `gorm:"column:played"`
		Win                                    int       `gorm:"column:win"`
		Draw                                   int       `gorm:"column:draw"`
		Lose                                   int       `gorm:"column:lose"`
		GoalsFor                               int       `gorm:"column:goals_for"`
		GoalsAgainst                           int       `gorm:"column:goals_against"`
		
		HomePlayed                             int       `gorm:"column:home_played"`
		HomeWin                                int       `gorm:"column:home_win"`
		HomeDraw                               int       `gorm:"column:home_draw"`
		HomeLose                               int       `gorm:"column:home_lose"`
		HomeGoalsFor                           int       `gorm:"column:home_goals_for"`
		HomeGoalsAgainst                       int       `gorm:"column:home_goals_against"`

		AwayPlayed                             int       `gorm:"column:away_played"`
		AwayWin                                int       `gorm:"column:away_win"`
		AwayDraw                               int       `gorm:"column:away_draw"`
		AwayLose                               int       `gorm:"column:away_lose"`
		AwayGoalsFor                           int       `gorm:"column:away_goals_for"`
		AwayGoalsAgainst                       int       `gorm:"column:away_goals_against"`
		
		HomeGoalsForAverage                    float64   `gorm:"column:home_goals_for_average"`
		AwayGoalsForAverage                    float64   `gorm:"column:away_goals_for_average"`

		GoalsForAverage                        float64   `gorm:"column:goals_for_average"`


		HomeGoalsForMinute015Total             int       `gorm:"column:home_goals_for_minute_0_15_total"`
		HomeGoalsForMinute015Percentage        float64   `gorm:"column:home_goals_for_minute_0_15_percentage"`
		HomeGoalsForMinute1630Total            int       `gorm:"column:home_goals_for_minute_16_30_total"`
		HomeGoalsForMinute1630Percentage       float64   `gorm:"column:home_goals_for_minute_16_30_percentage"`
		HomeGoalsForMinute3145Total            int       `gorm:"column:home_goals_for_minute_31_45_total"`
		HomeGoalsForMinute3145Percentage       float64   `gorm:"column:home_goals_for_minute_31_45_percentage"`
		HomeGoalsForMinute4660Total            int       `gorm:"column:home_goals_for_minute_46_60_total"`
		HomeGoalsForMinute4660Percentage       float64   `gorm:"column:home_goals_for_minute_46_60_percentage"`
		HomeGoalsForMinute6175Total            int       `gorm:"column:home_goals_for_minute_61_75_total"`
		HomeGoalsForMinute6175Percentage       float64   `gorm:"column:home_goals_for_minute_61_75_percentage"`
		HomeGoalsForMinute7690Total            int       `gorm:"column:home_goals_for_minute_76_90_total"`
		HomeGoalsForMinute7690Percentage       float64   `gorm:"column:home_goals_for_minute_76_90_percentage"`
		HomeGoalsForMinute91105Total           int       `gorm:"column:home_goals_for_minute_91_105_total"`
		HomeGoalsForMinute91105Percentage      float64   `gorm:"column:home_goals_for_minute_91_105_percentage"`
		HomeGoalsForMinute106120Total          int       `gorm:"column:home_goals_for_minute_106_120_total"`
		HomeGoalsForMinute106120Percentage     float64   `gorm:"column:home_goals_for_minute_106_120_percentage"`
		AwayGoalsForMinute015Total             int       `gorm:"column:away_goals_for_minute_0_15_total"`
		AwayGoalsForMinute015Percentage        float64   `gorm:"column:away_goals_for_minute_0_15_percentage"`
		AwayGoalsForMinute1630Total            int       `gorm:"column:away_goals_for_minute_16_30_total"`
		AwayGoalsForMinute1630Percentage       float64   `gorm:"column:away_goals_for_minute_16_30_percentage"`
		AwayGoalsForMinute3145Total            int       `gorm:"column:away_goals_for_minute_31_45_total"`
		AwayGoalsForMinute3145Percentage       float64   `gorm:"column:away_goals_for_minute_31_45_percentage"`
		AwayGoalsForMinute4660Total            int       `gorm:"column:away_goals_for_minute_46_60_total"`
		AwayGoalsForMinute4660Percentage       float64   `gorm:"column:away_goals_for_minute_46_60_percentage"`
		AwayGoalsForMinute6175Total            int       `gorm:"column:away_goals_for_minute_61_75_total"`
		AwayGoalsForMinute6175Percentage       float64   `gorm:"column:away_goals_for_minute_61_75_percentage"`
		AwayGoalsForMinute7690Total            int       `gorm:"column:away_goals_for_minute_76_90_total"`
		AwayGoalsForMinute7690Percentage       float64   `gorm:"column:away_goals_for_minute_76_90_percentage"`
		AwayGoalsForMinute91105Total           int       `gorm:"column:away_goals_for_minute_91_105_total"`
		AwayGoalsForMinute91105Percentage      float64   `gorm:"column:away_goals_for_minute_91_105_percentage"`
		AwayGoalsForMinute106120Total          int       `gorm:"column:away_goals_for_minute_106_120_total"`
		AwayGoalsForMinute106120Percentage     float64   `gorm:"column:away_goals_for_minute_106_120_percentage"`
		HomeGoalsAgainstAverage                float64   `gorm:"column:home_goals_against_average"`
		AwayGoalsAgainstAverage                float64   `gorm:"column:away_goals_against_average"`

		GoalsAgainstAverage                    float64   `gorm:"column:goals_against_average"`
		HomeGoalsAgainstMinute015Total         int       `gorm:"column:home_goals_against_minute_0_15_total"`
		HomeGoalsAgainstMinute015Percentage    float64   `gorm:"column:home_goals_against_minute_0_15_percentage"`
		HomeGoalsAgainstMinute1630Total        int       `gorm:"column:home_goals_against_minute_16_30_total"`
		HomeGoalsAgainstMinute1630Percentage   float64   `gorm:"column:home_goals_against_minute_16_30_percentage"`
		HomeGoalsAgainstMinute3145Total        int       `gorm:"column:home_goals_against_minute_31_45_total"`
		HomeGoalsAgainstMinute3145Percentage   float64   `gorm:"column:home_goals_against_minute_31_45_percentage"`
		HomeGoalsAgainstMinute4660Total        int       `gorm:"column:home_goals_against_minute_46_60_total"`
		HomeGoalsAgainstMinute4660Percentage   float64   `gorm:"column:home_goals_against_minute_46_60_percentage"`
		HomeGoalsAgainstMinute6175Total        int       `gorm:"column:home_goals_against_minute_61_75_total"`
		HomeGoalsAgainstMinute6175Percentage   float64   `gorm:"column:home_goals_against_minute_61_75_percentage"`
		HomeGoalsAgainstMinute7690Total        int       `gorm:"column:home_goals_against_minute_76_90_total"`
		HomeGoalsAgainstMinute7690Percentage   float64   `gorm:"column:home_goals_against_minute_76_90_percentage"`
		HomeGoalsAgainstMinute91105Total       int       `gorm:"column:home_goals_against_minute_91_105_total"`
		HomeGoalsAgainstMinute91105Percentage  float64   `gorm:"column:home_goals_against_minute_91_105_percentage"`
		HomeGoalsAgainstMinute106120Total      int       `gorm:"column:home_goals_against_minute_106_120_total"`
		HomeGoalsAgainstMinute106120Percentage float64   `gorm:"column:home_goals_against_minute_106_120_percentage"`
		AwayGoalsAgainstMinute015Total         int       `gorm:"column:away_goals_against_minute_0_15_total"`
		AwayGoalsAgainstMinute015Percentage    float64   `gorm:"column:away_goals_against_minute_0_15_percentage"`
		AwayGoalsAgainstMinute1630Total        int       `gorm:"column:away_goals_against_minute_16_30_total"`
		AwayGoalsAgainstMinute1630Percentage   float64   `gorm:"column:away_goals_against_minute_16_30_percentage"`
		AwayGoalsAgainstMinute3145Total        int       `gorm:"column:away_goals_against_minute_31_45_total"`
		AwayGoalsAgainstMinute3145Percentage   float64   `gorm:"column:away_goals_against_minute_31_45_percentage"`
		AwayGoalsAgainstMinute4660Total        int       `gorm:"column:away_goals_against_minute_46_60_total"`
		AwayGoalsAgainstMinute4660Percentage   float64   `gorm:"column:away_goals_against_minute_46_60_percentage"`
		AwayGoalsAgainstMinute6175Total        int       `gorm:"column:away_goals_against_minute_61_75_total"`
		AwayGoalsAgainstMinute6175Percentage   float64   `gorm:"column:away_goals_against_minute_61_75_percentage"`
		AwayGoalsAgainstMinute7690Total        int       `gorm:"column:away_goals_against_minute_76_90_total"`
		AwayGoalsAgainstMinute7690Percentage   float64   `gorm:"column:away_goals_against_minute_76_90_percentage"`
		AwayGoalsAgainstMinute91105Total       int       `gorm:"column:away_goals_against_minute_91_105_total"`
		AwayGoalsAgainstMinute91105Percentage  float64   `gorm:"column:away_goals_against_minute_91_105_percentage"`
		AwayGoalsAgainstMinute106120Total      int       `gorm:"column:away_goals_against_minute_106_120_total"`
		AwayGoalsAgainstMinute106120Percentage float64   `gorm:"column:away_goals_against_minute_106_120_percentage"`
		
		BiggestStreakWins                      int       `gorm:"column:biggest_streak_wins"`
		BiggestStreakDraws                     int       `gorm:"column:biggest_streak_draws"`
		BiggestStreakLoses                     int       `gorm:"column:biggest_streak_loses"`
		BiggestWinsHome                        string    `gorm:"column:biggest_wins_home"`
		BiggestWinsAway                        string    `gorm:"column:biggest_wins_away"`
		BiggestLosesHome                       string    `gorm:"column:biggest_loses_home"`
		BiggestLosesAway                       string    `gorm:"column:biggest_loses_away"`
		BiggestGoalsForHome                    int       `gorm:"column:biggest_goals_for_home"`
		BiggestGoalsForAway                    int       `gorm:"column:biggest_goals_for_away"`
		BiggestGoalsAgainstHome                int       `gorm:"column:biggest_goals_against_home"`
		BiggestGoalsAgainstAway                int       `gorm:"column:biggest_goals_against_away"`
		CleanSheetHome                         int       `gorm:"column:clean_sheet_home"`
		CleanSheetAway                         int       `gorm:"column:clean_sheet_away"`
		CleanSheetTotal                        int       `gorm:"column:clean_sheet_total"`
		FailedToScoreHome                      int       `gorm:"column:failed_to_score_home"`
		FailedToScoreAway                      int       `gorm:"column:failed_to_score_away"`
		FailedToScoreTotal                     int       `gorm:"column:failed_to_score_total"`
		PenaltyScoredTotal                     int       `gorm:"column:penalty_scored_total"`
		PenaltyScoredPercentage                float64   `gorm:"column:penalty_scored_percentage"`
		PenaltyMissedTotal                     int       `gorm:"column:penalty_missed_total"`
		PenaltyMissedPercentage                float64   `gorm:"column:penalty_missed_percentage"`
		PenaltyTotal                           int       `gorm:"column:penalty_total"`

		CardsYellow015Total                    int       `gorm:"column:cards_yellow_0_15_total"`
		CardsYellow015Percentage               float64   `gorm:"column:cards_yellow_0_15_percentage"`
		CardsYellow1630Total                   int       `gorm:"column:cards_yellow_16_30_total"`
		CardsYellow1630Percentage              float64   `gorm:"column:cards_yellow_16_30_percentage"`
		CardsYellow3145Total                   int       `gorm:"column:cards_yellow_31_45_total"`
		CardsYellow3145Percentage              float64   `gorm:"column:cards_yellow_31_45_percentage"`
		CardsYellow4660Total                   int       `gorm:"column:cards_yellow_46_60_total"`
		CardsYellow4660Percentage              float64   `gorm:"column:cards_yellow_46_60_percentage"`
		CardsYellow6175Total                   int       `gorm:"column:cards_yellow_61_75_total"`
		CardsYellow6175Percentage              float64   `gorm:"column:cards_yellow_61_75_percentage"`
		CardsYellow7690Total                   int       `gorm:"column:cards_yellow_76_90_total"`
		CardsYellow7690Percentage              float64   `gorm:"column:cards_yellow_76_90_percentage"`
		CardsYellow91105Total                  int       `gorm:"column:cards_yellow_91_105_total"`
		CardsYellow91105Percentage             float64   `gorm:"column:cards_yellow_91_105_percentage"`
		CardsYellow106120Total                 int       `gorm:"column:cards_yellow_106_120_total"`
		CardsYellow106120Percentage            float64   `gorm:"column:cards_yellow_106_120_percentage"`
		CardsRed015Total                       int       `gorm:"column:cards_red_0_15_total"`
		CardsRed015Percentage                  float64   `gorm:"column:cards_red_0_15_percentage"`
		CardsRed1630Total                      int       `gorm:"column:cards_red_16_30_total"`
		CardsRed1630Percentage                 float64   `gorm:"column:cards_red_16_30_percentage"`
		CardsRed3145Total                      int       `gorm:"column:cards_red_31_45_total"`
		CardsRed3145Percentage                 float64   `gorm:"column:cards_red_31_45_percentage"`
		CardsRed4660Total                      int       `gorm:"column:cards_red_46_60_total"`
		CardsRed4660Percentage                 float64   `gorm:"column:cards_red_46_60_percentage"`
		CardsRed6175Total                      int       `gorm:"column:cards_red_61_75_total"`
		CardsRed6175Percentage                 float64   `gorm:"column:cards_red_61_75_percentage"`
		CardsRed7690Total                      int       `gorm:"column:cards_red_76_90_total"`
		CardsRed7690Percentage                 float64   `gorm:"column:cards_red_76_90_percentage"`
		CardsRed91105Total                     int       `gorm:"column:cards_red_91_105_total"`
		CardsRed91105Percentage                float64   `gorm:"column:cards_red_91_105_percentage"`
		CardsRed106120Total                    int       `gorm:"column:cards_red_106_120_total"`
		CardsRed106120Percentage               float64   `gorm:"column:cards_red_106_120_percentage"`

		CornerKicksTotal                       int       `gorm:"column:corner_kicks_total"`
		CornerKicksAvg                         float64   `gorm:"column:corner_kicks_avg"`
		CornerKicksMax                         int       `gorm:"column:corner_kicks_max"`
		CornerKicksMin                         int       `gorm:"column:corner_kicks_min"`

		TotalShotsTotal                        int       `gorm:"column:total_shots_total"`
		TotalShotsAvg                          float64   `gorm:"column:total_shots_avg"`
		TotalShotsMax                          int       `gorm:"column:total_shots_max"`
		TotalShotsMin                          int       `gorm:"column:total_shots_min"`
		ShotsOnGoalTotal                       int       `gorm:"column:shots_on_goal_total"`
		ShotsOnGoalAvg                         float64   `gorm:"column:shots_on_goal_avg"`
		ShotsOnGoalMax                         int       `gorm:"column:shots_on_goal_max"`
		ShotsOnGoalMin                         int       `gorm:"column:shots_on_goal_min"`

		FoulsTotal                             int       `gorm:"column:fouls_total"`
		FoulsAvg                               float64   `gorm:"column:fouls_avg"`
		FoulsMax                               int       `gorm:"column:fouls_max"`
		FoulsMin                               int       `gorm:"column:fouls_min"`

		OffsidesTotal                          int       `gorm:"column:offsides_total"`
		OffsidesAvg                            float64   `gorm:"column:offsides_avg"`
		OffsidesMax                            int       `gorm:"column:offsides_max"`
		OffsidesMin                            int       `gorm:"column:offsides_min"`
		CreatedAt                              time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL"`
		UpdatedAt                              time.Time `gorm:"column:updated_at"`
		DeletedAt                              time.Time `gorm:"column:deleted_at"`

		
        Team 		FootballTeam 		`gorm:"foreignKey:TeamapiID;references:TeamAPIID"`
	} 

	func (m *ApiFootballLeagueStanding) TableName() string {
		return "api_football_league_standings"
	}
 
    func (b *ApiFootballLeagueStanding) GetLeagueStanding(leagueapiID uint, season uint) ([]ApiFootballLeagueStanding, error) {
        var apifootballleaguestanding []ApiFootballLeagueStanding
    
        err := database.DB.
            Select("*").  
            Where("leagueapi_id = ?", leagueapiID).
            Where("season = ?", season).
            Where("deleted_at IS NULL"). 
            Order("rank asc"). 
            Preload("Team").
            Find(&apifootballleaguestanding).Error

        return apifootballleaguestanding, err
    } 