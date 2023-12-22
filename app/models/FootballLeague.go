package models
 
	import ( 
		// "github.com/HerlambangHaryo/go-crud-simple/platform/database"  
		"time"  
		"gorm.io/gorm"
	)

	type FootballLeague struct {
		gorm.Model 
		Name      		string `gorm:"column:name"` 
		CountryName                           string    `gorm:"column:country_name;NOT NULL"`
		Continent                             string    `gorm:"column:continent"`  
		Type                                  string    `gorm:"column:type;NOT NULL"`
		Logo                                  string    `gorm:"column:logo"`
		BookmakersapiID                       string       `gorm:"column:bookmakersapi_id;default:0"`
		BookmakersName                        string    `gorm:"column:bookmakers_name"`
		TeamStatistics                        int       `gorm:"column:team_statistics"`
		LeagueFixtureUpdatedAt                time.Time `gorm:"column:league_fixture_updated_at"`
		TipsDate                              time.Time `gorm:"column:tips_date"`
		TdEndCountry                          time.Time `gorm:"column:td_end_country"`
		PercEndLeague                         float64   `gorm:"column:perc_end_league"`
		PercEndCountry                        float64   `gorm:"column:perc_end_country"`
		PercOneLeague                         float64   `gorm:"column:perc_one_league"`
		PercOneCountry                        float64   `gorm:"column:perc_one_country"`
		Status                                string    `gorm:"column:status"`
		TanggalUpdateFixtures                 time.Time `gorm:"column:tanggal_update_fixtures"`
		TanggalResetPrePattern                time.Time `gorm:"column:tanggal_reset_pre_pattern"`
		TanggalResetEndPattern                time.Time `gorm:"column:tanggal_reset_end_pattern"`
		TanggalUpdatePatternlists             time.Time `gorm:"column:tanggal_update_patternlists"`
		LeagueStandingUpdatedAt               time.Time `gorm:"column:league_standing_updated_at"`
		ApiLeagueUpdatedAt                    time.Time `gorm:"column:api_league_updated_at"`
		ResearchOnAnytimeGoalScorersUpdatedAt time.Time `gorm:"column:research_on_anytime_goal_scorers_updated_at"`
		Star                                  string    `gorm:"column:star"`
		TotalPatternFromPre                   int       `gorm:"column:total_pattern_from_pre"`
		TotalPatternPreEnds                   int       `gorm:"column:total_pattern_pre_ends"`
		TotalPatternOnlyEnds                  int       `gorm:"column:total_pattern_only_ends"`
		DetailStats                           int       `gorm:"column:detail_stats"`
		CreatedAt                             time.Time `gorm:"column:created_at"`
		UpdatedAt                             time.Time `gorm:"column:updated_at"`
		DeletedAt                             time.Time `gorm:"column:deleted_at"`


		LeagueAPIID		uint   `gorm:"column:leagueapi_id"`
		// ... other fields ...

		
		CountryTeam      Country `gorm:"foreignKey:CountryName;references:Name"`
		// CountryTeam Country  `gorm:"foreignKey:CountryName;references:Name"`
	} 