package models
	
	import ( 
		"github.com/HerlambangHaryo/go-crud-simple/platform/database"  
		"time"  
		"gorm.io/gorm" 
	)

	type FootballOdd struct {
		gorm.Model  
		ID                                      int       `gorm:"column:id;NOT NULL"`
		Date                                    time.Time `gorm:"column:date"`
		LeagueapiID                             int       `gorm:"column:leagueapi_id"`
		Season                                  int 	  `gorm:"column:season"`
		FixtureStatus                           string    `gorm:"column:fixture_status"`
		PreMatchWinnerHome                      float64   `gorm:"column:pre_match_winner_home"`
		PreMatchWinnerDraw                      float64   `gorm:"column:pre_match_winner_draw"`
		PreMatchWinnerAway                      float64   `gorm:"column:pre_match_winner_away"`
		EndMatchWinnerHome                      float64   `gorm:"column:end_match_winner_home"`
		EndMatchWinnerDraw                      float64   `gorm:"column:end_match_winner_draw"`
		EndMatchWinnerAway                      float64   `gorm:"column:end_match_winner_away"`
		PreHomeAwayHome                         float64   `gorm:"column:pre_homeaway_home"`
		PreHomeAwayAway                         float64   `gorm:"column:pre_homeaway_away"`
		EndHomeAwayHome                         float64   `gorm:"column:end_homeaway_home"`
		EndHomeAwayAway                         float64   `gorm:"column:end_homeaway_away"`
		PreFirstHalfWinnerHome                  float64   `gorm:"column:pre_first_half_winner_home"`
		PreFirstHalfWinnerDraw                  float64   `gorm:"column:pre_first_half_winner_draw"`
		PreFirstHalfWinnerAway                  float64   `gorm:"column:pre_first_half_winner_away"`
		EndFirstHalfWinnerHome                  float64   `gorm:"column:end_first_half_winner_home"`
		EndFirstHalfWinnerDraw                  float64   `gorm:"column:end_first_half_winner_draw"`
		EndFirstHalfWinnerAway                  float64   `gorm:"column:end_first_half_winner_away"`
		PreSecondHalfWinnerHome                 float64   `gorm:"column:pre_second_half_winner_home"`
		PreSecondHalfWinnerDraw                 float64   `gorm:"column:pre_second_half_winner_draw"`
		PreSecondHalfWinnerAway                 float64   `gorm:"column:pre_second_half_winner_away"`
		EndSecondHalfWinnerHome                 float64   `gorm:"column:end_second_half_winner_home"`
		EndSecondHalfWinnerDraw                 float64   `gorm:"column:end_second_half_winner_draw"`
		EndSecondHalfWinnerAway                 float64   `gorm:"column:end_second_half_winner_away"`
		PreAsianHandicapHomeMin675              float64   `gorm:"column:pre_asian_handicap_home_min_675"`
		PreAsianHandicapAwayMin675              float64   `gorm:"column:pre_asian_handicap_away_min_675"`
		PreAsianHandicapHomeMin65               float64   `gorm:"column:pre_asian_handicap_home_min_65"`
		PreAsianHandicapAwayMin65               float64   `gorm:"column:pre_asian_handicap_away_min_65"`
		PreAsianHandicapHomeMin625              float64   `gorm:"column:pre_asian_handicap_home_min_625"`
		PreAsianHandicapAwayMin625              float64   `gorm:"column:pre_asian_handicap_away_min_625"`
		PreAsianHandicapHomeMin6                float64   `gorm:"column:pre_asian_handicap_home_min_6"`
		PreAsianHandicapAwayMin6                float64   `gorm:"column:pre_asian_handicap_away_min_6"`
		PreAsianHandicapHomeMin575              float64   `gorm:"column:pre_asian_handicap_home_min_575"`
		PreAsianHandicapAwayMin575              float64   `gorm:"column:pre_asian_handicap_away_min_575"`
		PreAsianHandicapHomeMin55               float64   `gorm:"column:pre_asian_handicap_home_min_55"`
		PreAsianHandicapAwayMin55               float64   `gorm:"column:pre_asian_handicap_away_min_55"`
		PreAsianHandicapHomeMin525              float64   `gorm:"column:pre_asian_handicap_home_min_525"`
		PreAsianHandicapAwayMin525              float64   `gorm:"column:pre_asian_handicap_away_min_525"`
		PreAsianHandicapHomeMin5                float64   `gorm:"column:pre_asian_handicap_home_min_5"`
		PreAsianHandicapAwayMin5                float64   `gorm:"column:pre_asian_handicap_away_min_5"`
		PreAsianHandicapHomeMin475              float64   `gorm:"column:pre_asian_handicap_home_min_475"`
		PreAsianHandicapAwayMin475              float64   `gorm:"column:pre_asian_handicap_away_min_475"`
		PreAsianHandicapHomeMin45               float64   `gorm:"column:pre_asian_handicap_home_min_45"`
		PreAsianHandicapAwayMin45               float64   `gorm:"column:pre_asian_handicap_away_min_45"`
		PreAsianHandicapHomeMin425              float64   `gorm:"column:pre_asian_handicap_home_min_425"`
		PreAsianHandicapAwayMin425              float64   `gorm:"column:pre_asian_handicap_away_min_425"`
		PreAsianHandicapHomeMin4                float64   `gorm:"column:pre_asian_handicap_home_min_4"`
		PreAsianHandicapAwayMin4                float64   `gorm:"column:pre_asian_handicap_away_min_4"`
		PreAsianHandicapHomeMin375              float64   `gorm:"column:pre_asian_handicap_home_min_375"`
		PreAsianHandicapAwayMin375              float64   `gorm:"column:pre_asian_handicap_away_min_375"`
		PreAsianHandicapHomeMin35               float64   `gorm:"column:pre_asian_handicap_home_min_35"`
		PreAsianHandicapAwayMin35               float64   `gorm:"column:pre_asian_handicap_away_min_35"`
		PreAsianHandicapHomeMin325              float64   `gorm:"column:pre_asian_handicap_home_min_325"`
		PreAsianHandicapAwayMin325              float64   `gorm:"column:pre_asian_handicap_away_min_325"`
		PreAsianHandicapHomeMin3                float64   `gorm:"column:pre_asian_handicap_home_min_3"`
		PreAsianHandicapAwayMin3                float64   `gorm:"column:pre_asian_handicap_away_min_3"`
		PreAsianHandicapHomeMin275              float64   `gorm:"column:pre_asian_handicap_home_min_275"`
		PreAsianHandicapAwayMin275              float64   `gorm:"column:pre_asian_handicap_away_min_275"`
		PreAsianHandicapHomeMin25               float64   `gorm:"column:pre_asian_handicap_home_min_25"`
		PreAsianHandicapAwayMin25               float64   `gorm:"column:pre_asian_handicap_away_min_25"`
		PreAsianHandicapHomeMin225              float64   `gorm:"column:pre_asian_handicap_home_min_225"`
		PreAsianHandicapAwayMin225              float64   `gorm:"column:pre_asian_handicap_away_min_225"`
		PreAsianHandicapHomeMin2                float64   `gorm:"column:pre_asian_handicap_home_min_2"`
		PreAsianHandicapAwayMin2                float64   `gorm:"column:pre_asian_handicap_away_min_2"`
		PreAsianHandicapHomeMin175              float64   `gorm:"column:pre_asian_handicap_home_min_175"`
		PreAsianHandicapAwayMin175              float64   `gorm:"column:pre_asian_handicap_away_min_175"`
		PreAsianHandicapHomeMin15               float64   `gorm:"column:pre_asian_handicap_home_min_15"`
		PreAsianHandicapAwayMin15               float64   `gorm:"column:pre_asian_handicap_away_min_15"`
		PreAsianHandicapHomeMin125              float64   `gorm:"column:pre_asian_handicap_home_min_125"`
		PreAsianHandicapAwayMin125              float64   `gorm:"column:pre_asian_handicap_away_min_125"`
		PreAsianHandicapHomeMin1                float64   `gorm:"column:pre_asian_handicap_home_min_1"`
		PreAsianHandicapAwayMin1                float64   `gorm:"column:pre_asian_handicap_away_min_1"`
		PreAsianHandicapHomeMin075              float64   `gorm:"column:pre_asian_handicap_home_min_075"`
		PreAsianHandicapAwayMin075              float64   `gorm:"column:pre_asian_handicap_away_min_075"`
		PreAsianHandicapHomeMin05               float64   `gorm:"column:pre_asian_handicap_home_min_05"`
		PreAsianHandicapAwayMin05               float64   `gorm:"column:pre_asian_handicap_away_min_05"`
		PreAsianHandicapHomeMin025              float64   `gorm:"column:pre_asian_handicap_home_min_025"`
		PreAsianHandicapAwayMin025              float64   `gorm:"column:pre_asian_handicap_away_min_025"`
		PreAsianHandicapHomePlus0               float64   `gorm:"column:pre_asian_handicap_home_plus_0"`
		PreAsianHandicapAwayPlus0               float64   `gorm:"column:pre_asian_handicap_away_plus_0"`
		PreAsianHandicapHomePlus025             float64   `gorm:"column:pre_asian_handicap_home_plus_025"`
		PreAsianHandicapAwayPlus025             float64   `gorm:"column:pre_asian_handicap_away_plus_025"`
		PreAsianHandicapHomePlus05              float64   `gorm:"column:pre_asian_handicap_home_plus_05"`
		PreAsianHandicapAwayPlus05              float64   `gorm:"column:pre_asian_handicap_away_plus_05"`
		PreAsianHandicapHomePlus075             float64   `gorm:"column:pre_asian_handicap_home_plus_075"`
		PreAsianHandicapAwayPlus075             float64   `gorm:"column:pre_asian_handicap_away_plus_075"`
		PreAsianHandicapHomePlus1               float64   `gorm:"column:pre_asian_handicap_home_plus_1"`
		PreAsianHandicapAwayPlus1               float64   `gorm:"column:pre_asian_handicap_away_plus_1"`
		PreAsianHandicapHomePlus125             float64   `gorm:"column:pre_asian_handicap_home_plus_125"`
		PreAsianHandicapAwayPlus125             float64   `gorm:"column:pre_asian_handicap_away_plus_125"`
		PreAsianHandicapHomePlus15              float64   `gorm:"column:pre_asian_handicap_home_plus_15"`
		PreAsianHandicapAwayPlus15              float64   `gorm:"column:pre_asian_handicap_away_plus_15"`
		PreAsianHandicapHomePlus175             float64   `gorm:"column:pre_asian_handicap_home_plus_175"`
		PreAsianHandicapAwayPlus175             float64   `gorm:"column:pre_asian_handicap_away_plus_175"`
		PreAsianHandicapHomePlus2               float64   `gorm:"column:pre_asian_handicap_home_plus_2"`
		PreAsianHandicapAwayPlus2               float64   `gorm:"column:pre_asian_handicap_away_plus_2"`
		PreAsianHandicapHomePlus225             float64   `gorm:"column:pre_asian_handicap_home_plus_225"`
		PreAsianHandicapAwayPlus225             float64   `gorm:"column:pre_asian_handicap_away_plus_225"`
		PreAsianHandicapHomePlus25              float64   `gorm:"column:pre_asian_handicap_home_plus_25"`
		PreAsianHandicapAwayPlus25              float64   `gorm:"column:pre_asian_handicap_away_plus_25"`
		PreAsianHandicapHomePlus275             float64   `gorm:"column:pre_asian_handicap_home_plus_275"`
		PreAsianHandicapAwayPlus275             float64   `gorm:"column:pre_asian_handicap_away_plus_275"`
		PreAsianHandicapHomePlus3               float64   `gorm:"column:pre_asian_handicap_home_plus_3"`
		PreAsianHandicapAwayPlus3               float64   `gorm:"column:pre_asian_handicap_away_plus_3"`
		PreAsianHandicapHomePlus325             float64   `gorm:"column:pre_asian_handicap_home_plus_325"`
		PreAsianHandicapAwayPlus325             float64   `gorm:"column:pre_asian_handicap_away_plus_325"`
		PreAsianHandicapHomePlus35              float64   `gorm:"column:pre_asian_handicap_home_plus_35"`
		PreAsianHandicapAwayPlus35              float64   `gorm:"column:pre_asian_handicap_away_plus_35"`
		PreAsianHandicapHomePlus375             float64   `gorm:"column:pre_asian_handicap_home_plus_375"`
		PreAsianHandicapAwayPlus375             float64   `gorm:"column:pre_asian_handicap_away_plus_375"`
		PreAsianHandicapHomePlus4               float64   `gorm:"column:pre_asian_handicap_home_plus_4"`
		PreAsianHandicapAwayPlus4               float64   `gorm:"column:pre_asian_handicap_away_plus_4"`
		PreAsianHandicapHomePlus425             float64   `gorm:"column:pre_asian_handicap_home_plus_425"`
		PreAsianHandicapAwayPlus425             float64   `gorm:"column:pre_asian_handicap_away_plus_425"`
		PreAsianHandicapHomePlus45              float64   `gorm:"column:pre_asian_handicap_home_plus_45"`
		PreAsianHandicapAwayPlus45              float64   `gorm:"column:pre_asian_handicap_away_plus_45"`
		PreAsianHandicapHomePlus475             float64   `gorm:"column:pre_asian_handicap_home_plus_475"`
		PreAsianHandicapAwayPlus475             float64   `gorm:"column:pre_asian_handicap_away_plus_475"`
		PreAsianHandicapHomePlus5               float64   `gorm:"column:pre_asian_handicap_home_plus_5"`
		PreAsianHandicapAwayPlus5               float64   `gorm:"column:pre_asian_handicap_away_plus_5"`
		PreAsianHandicapHomePlus525             float64   `gorm:"column:pre_asian_handicap_home_plus_525"`
		PreAsianHandicapAwayPlus525             float64   `gorm:"column:pre_asian_handicap_away_plus_525"`
		PreAsianHandicapHomePlus55              float64   `gorm:"column:pre_asian_handicap_home_plus_55"`
		PreAsianHandicapAwayPlus55              float64   `gorm:"column:pre_asian_handicap_away_plus_55"`
		PreAsianHandicapHomePlus575             float64   `gorm:"column:pre_asian_handicap_home_plus_575"`
		PreAsianHandicapAwayPlus575             float64   `gorm:"column:pre_asian_handicap_away_plus_575"`
		PreAsianHandicapHomePlus6               float64   `gorm:"column:pre_asian_handicap_home_plus_6"`
		PreAsianHandicapAwayPlus6               float64   `gorm:"column:pre_asian_handicap_away_plus_6"`
		PreAsianHandicapHomePlus625             float64   `gorm:"column:pre_asian_handicap_home_plus_625"`
		PreAsianHandicapAwayPlus625             float64   `gorm:"column:pre_asian_handicap_away_plus_625"`
		PreAsianHandicapHomePlus65              float64   `gorm:"column:pre_asian_handicap_home_plus_65"`
		PreAsianHandicapAwayPlus65              float64   `gorm:"column:pre_asian_handicap_away_plus_65"`
		PreAsianHandicapHomePlus675             float64   `gorm:"column:pre_asian_handicap_home_plus_675"`
		PreAsianHandicapAwayPlus675             float64   `gorm:"column:pre_asian_handicap_away_plus_675"`
		EndAsianHandicapHomeMin675              float64   `gorm:"column:end_asian_handicap_home_min_675"`
		EndAsianHandicapAwayMin675              float64   `gorm:"column:end_asian_handicap_away_min_675"`
		EndAsianHandicapHomeMin65               float64   `gorm:"column:end_asian_handicap_home_min_65"`
		EndAsianHandicapAwayMin65               float64   `gorm:"column:end_asian_handicap_away_min_65"`
		EndAsianHandicapHomeMin625              float64   `gorm:"column:end_asian_handicap_home_min_625"`
		EndAsianHandicapAwayMin625              float64   `gorm:"column:end_asian_handicap_away_min_625"`
		EndAsianHandicapHomeMin6                float64   `gorm:"column:end_asian_handicap_home_min_6"`
		EndAsianHandicapAwayMin6                float64   `gorm:"column:end_asian_handicap_away_min_6"`
		EndAsianHandicapHomeMin575              float64   `gorm:"column:end_asian_handicap_home_min_575"`
		EndAsianHandicapAwayMin575              float64   `gorm:"column:end_asian_handicap_away_min_575"`
		EndAsianHandicapHomeMin55               float64   `gorm:"column:end_asian_handicap_home_min_55"`
		EndAsianHandicapAwayMin55               float64   `gorm:"column:end_asian_handicap_away_min_55"`
		EndAsianHandicapHomeMin525              float64   `gorm:"column:end_asian_handicap_home_min_525"`
		EndAsianHandicapAwayMin525              float64   `gorm:"column:end_asian_handicap_away_min_525"`
		EndAsianHandicapHomeMin5                float64   `gorm:"column:end_asian_handicap_home_min_5"`
		EndAsianHandicapAwayMin5                float64   `gorm:"column:end_asian_handicap_away_min_5"`
		EndAsianHandicapHomeMin475              float64   `gorm:"column:end_asian_handicap_home_min_475"`
		EndAsianHandicapAwayMin475              float64   `gorm:"column:end_asian_handicap_away_min_475"`
		EndAsianHandicapHomeMin45               float64   `gorm:"column:end_asian_handicap_home_min_45"`
		EndAsianHandicapAwayMin45               float64   `gorm:"column:end_asian_handicap_away_min_45"`
		EndAsianHandicapHomeMin425              float64   `gorm:"column:end_asian_handicap_home_min_425"`
		EndAsianHandicapAwayMin425              float64   `gorm:"column:end_asian_handicap_away_min_425"`
		EndAsianHandicapHomeMin4                float64   `gorm:"column:end_asian_handicap_home_min_4"`
		EndAsianHandicapAwayMin4                float64   `gorm:"column:end_asian_handicap_away_min_4"`
		EndAsianHandicapHomeMin375              float64   `gorm:"column:end_asian_handicap_home_min_375"`
		EndAsianHandicapAwayMin375              float64   `gorm:"column:end_asian_handicap_away_min_375"`
		EndAsianHandicapHomeMin35               float64   `gorm:"column:end_asian_handicap_home_min_35"`
		EndAsianHandicapAwayMin35               float64   `gorm:"column:end_asian_handicap_away_min_35"`
		EndAsianHandicapHomeMin325              float64   `gorm:"column:end_asian_handicap_home_min_325"`
		EndAsianHandicapAwayMin325              float64   `gorm:"column:end_asian_handicap_away_min_325"`
		EndAsianHandicapHomeMin3                float64   `gorm:"column:end_asian_handicap_home_min_3"`
		EndAsianHandicapAwayMin3                float64   `gorm:"column:end_asian_handicap_away_min_3"`
		EndAsianHandicapHomeMin275              float64   `gorm:"column:end_asian_handicap_home_min_275"`
		EndAsianHandicapAwayMin275              float64   `gorm:"column:end_asian_handicap_away_min_275"`
		EndAsianHandicapHomeMin25               float64   `gorm:"column:end_asian_handicap_home_min_25"`
		EndAsianHandicapAwayMin25               float64   `gorm:"column:end_asian_handicap_away_min_25"`
		EndAsianHandicapHomeMin225              float64   `gorm:"column:end_asian_handicap_home_min_225"`
		EndAsianHandicapAwayMin225              float64   `gorm:"column:end_asian_handicap_away_min_225"`
		EndAsianHandicapHomeMin2                float64   `gorm:"column:end_asian_handicap_home_min_2"`
		EndAsianHandicapAwayMin2                float64   `gorm:"column:end_asian_handicap_away_min_2"`
		EndAsianHandicapHomeMin175              float64   `gorm:"column:end_asian_handicap_home_min_175"`
		EndAsianHandicapAwayMin175              float64   `gorm:"column:end_asian_handicap_away_min_175"`
		EndAsianHandicapHomeMin15               float64   `gorm:"column:end_asian_handicap_home_min_15"`
		EndAsianHandicapAwayMin15               float64   `gorm:"column:end_asian_handicap_away_min_15"`
		EndAsianHandicapHomeMin125              float64   `gorm:"column:end_asian_handicap_home_min_125"`
		EndAsianHandicapAwayMin125              float64   `gorm:"column:end_asian_handicap_away_min_125"`
		EndAsianHandicapHomeMin1                float64   `gorm:"column:end_asian_handicap_home_min_1"`
		EndAsianHandicapAwayMin1                float64   `gorm:"column:end_asian_handicap_away_min_1"`
		EndAsianHandicapHomeMin075              float64   `gorm:"column:end_asian_handicap_home_min_075"`
		EndAsianHandicapAwayMin075              float64   `gorm:"column:end_asian_handicap_away_min_075"`
		EndAsianHandicapHomeMin05               float64   `gorm:"column:end_asian_handicap_home_min_05"`
		EndAsianHandicapAwayMin05               float64   `gorm:"column:end_asian_handicap_away_min_05"`
		EndAsianHandicapHomeMin025              float64   `gorm:"column:end_asian_handicap_home_min_025"`
		EndAsianHandicapAwayMin025              float64   `gorm:"column:end_asian_handicap_away_min_025"`
		EndAsianHandicapHomePlus0               float64   `gorm:"column:end_asian_handicap_home_plus_0"`
		EndAsianHandicapAwayPlus0               float64   `gorm:"column:end_asian_handicap_away_plus_0"`
		EndAsianHandicapHomePlus025             float64   `gorm:"column:end_asian_handicap_home_plus_025"`
		EndAsianHandicapAwayPlus025             float64   `gorm:"column:end_asian_handicap_away_plus_025"`
		EndAsianHandicapHomePlus05              float64   `gorm:"column:end_asian_handicap_home_plus_05"`
		EndAsianHandicapAwayPlus05              float64   `gorm:"column:end_asian_handicap_away_plus_05"`
		EndAsianHandicapHomePlus075             float64   `gorm:"column:end_asian_handicap_home_plus_075"`
		EndAsianHandicapAwayPlus075             float64   `gorm:"column:end_asian_handicap_away_plus_075"`
		EndAsianHandicapHomePlus1               float64   `gorm:"column:end_asian_handicap_home_plus_1"`
		EndAsianHandicapAwayPlus1               float64   `gorm:"column:end_asian_handicap_away_plus_1"`
		EndAsianHandicapHomePlus125             float64   `gorm:"column:end_asian_handicap_home_plus_125"`
		EndAsianHandicapAwayPlus125             float64   `gorm:"column:end_asian_handicap_away_plus_125"`
		EndAsianHandicapHomePlus15              float64   `gorm:"column:end_asian_handicap_home_plus_15"`
		EndAsianHandicapAwayPlus15              float64   `gorm:"column:end_asian_handicap_away_plus_15"`
		EndAsianHandicapHomePlus175             float64   `gorm:"column:end_asian_handicap_home_plus_175"`
		EndAsianHandicapAwayPlus175             float64   `gorm:"column:end_asian_handicap_away_plus_175"`
		EndAsianHandicapHomePlus2               float64   `gorm:"column:end_asian_handicap_home_plus_2"`
		EndAsianHandicapAwayPlus2               float64   `gorm:"column:end_asian_handicap_away_plus_2"`
		EndAsianHandicapHomePlus225             float64   `gorm:"column:end_asian_handicap_home_plus_225"`
		EndAsianHandicapAwayPlus225             float64   `gorm:"column:end_asian_handicap_away_plus_225"`
		EndAsianHandicapHomePlus25              float64   `gorm:"column:end_asian_handicap_home_plus_25"`
		EndAsianHandicapAwayPlus25              float64   `gorm:"column:end_asian_handicap_away_plus_25"`
		EndAsianHandicapHomePlus275             float64   `gorm:"column:end_asian_handicap_home_plus_275"`
		EndAsianHandicapAwayPlus275             float64   `gorm:"column:end_asian_handicap_away_plus_275"`
		EndAsianHandicapHomePlus3               float64   `gorm:"column:end_asian_handicap_home_plus_3"`
		EndAsianHandicapAwayPlus3               float64   `gorm:"column:end_asian_handicap_away_plus_3"`
		EndAsianHandicapHomePlus325             float64   `gorm:"column:end_asian_handicap_home_plus_325"`
		EndAsianHandicapAwayPlus325             float64   `gorm:"column:end_asian_handicap_away_plus_325"`
		EndAsianHandicapHomePlus35              float64   `gorm:"column:end_asian_handicap_home_plus_35"`
		EndAsianHandicapAwayPlus35              float64   `gorm:"column:end_asian_handicap_away_plus_35"`
		EndAsianHandicapHomePlus375             float64   `gorm:"column:end_asian_handicap_home_plus_375"`
		EndAsianHandicapAwayPlus375             float64   `gorm:"column:end_asian_handicap_away_plus_375"`
		EndAsianHandicapHomePlus4               float64   `gorm:"column:end_asian_handicap_home_plus_4"`
		EndAsianHandicapAwayPlus4               float64   `gorm:"column:end_asian_handicap_away_plus_4"`
		EndAsianHandicapHomePlus425             float64   `gorm:"column:end_asian_handicap_home_plus_425"`
		EndAsianHandicapAwayPlus425             float64   `gorm:"column:end_asian_handicap_away_plus_425"`
		EndAsianHandicapHomePlus45              float64   `gorm:"column:end_asian_handicap_home_plus_45"`
		EndAsianHandicapAwayPlus45              float64   `gorm:"column:end_asian_handicap_away_plus_45"`
		EndAsianHandicapHomePlus475             float64   `gorm:"column:end_asian_handicap_home_plus_475"`
		EndAsianHandicapAwayPlus475             float64   `gorm:"column:end_asian_handicap_away_plus_475"`
		EndAsianHandicapHomePlus5               float64   `gorm:"column:end_asian_handicap_home_plus_5"`
		EndAsianHandicapAwayPlus5               float64   `gorm:"column:end_asian_handicap_away_plus_5"`
		EndAsianHandicapHomePlus525             float64   `gorm:"column:end_asian_handicap_home_plus_525"`
		EndAsianHandicapAwayPlus525             float64   `gorm:"column:end_asian_handicap_away_plus_525"`
		EndAsianHandicapHomePlus55              float64   `gorm:"column:end_asian_handicap_home_plus_55"`
		EndAsianHandicapAwayPlus55              float64   `gorm:"column:end_asian_handicap_away_plus_55"`
		EndAsianHandicapHomePlus575             float64   `gorm:"column:end_asian_handicap_home_plus_575"`
		EndAsianHandicapAwayPlus575             float64   `gorm:"column:end_asian_handicap_away_plus_575"`
		EndAsianHandicapHomePlus6               float64   `gorm:"column:end_asian_handicap_home_plus_6"`
		EndAsianHandicapAwayPlus6               float64   `gorm:"column:end_asian_handicap_away_plus_6"`
		EndAsianHandicapHomePlus625             float64   `gorm:"column:end_asian_handicap_home_plus_625"`
		EndAsianHandicapAwayPlus625             float64   `gorm:"column:end_asian_handicap_away_plus_625"`
		EndAsianHandicapHomePlus65              float64   `gorm:"column:end_asian_handicap_home_plus_65"`
		EndAsianHandicapAwayPlus65              float64   `gorm:"column:end_asian_handicap_away_plus_65"`
		EndAsianHandicapHomePlus675             float64   `gorm:"column:end_asian_handicap_home_plus_675"`
		EndAsianHandicapAwayPlus675             float64   `gorm:"column:end_asian_handicap_away_plus_675"`
		PreGoalsOverunderOver05                 float64   `gorm:"column:pre_goals_overunder_over_05"`
		PreGoalsOverunderUnder05                float64   `gorm:"column:pre_goals_overunder_under_05"`
		PreGoalsOverunderOver075                float64   `gorm:"column:pre_goals_overunder_over_075"`
		PreGoalsOverunderUnder075               float64   `gorm:"column:pre_goals_overunder_under_075"`
		PreGoalsOverunderOver10                 float64   `gorm:"column:pre_goals_overunder_over_10"`
		PreGoalsOverunderUnder10                float64   `gorm:"column:pre_goals_overunder_under_10"`
		PreGoalsOverunderOver125                float64   `gorm:"column:pre_goals_overunder_over_125"`
		PreGoalsOverunderUnder125               float64   `gorm:"column:pre_goals_overunder_under_125"`
		PreGoalsOverunderOver15                 float64   `gorm:"column:pre_goals_overunder_over_15"`
		PreGoalsOverunderUnder15                float64   `gorm:"column:pre_goals_overunder_under_15"`
		PreGoalsOverunderOver175                float64   `gorm:"column:pre_goals_overunder_over_175"`
		PreGoalsOverunderUnder175               float64   `gorm:"column:pre_goals_overunder_under_175"`
		PreGoalsOverunderOver20                 float64   `gorm:"column:pre_goals_overunder_over_20"`
		PreGoalsOverunderUnder20                float64   `gorm:"column:pre_goals_overunder_under_20"`
		PreGoalsOverunderOver225                float64   `gorm:"column:pre_goals_overunder_over_225"`
		PreGoalsOverunderUnder225               float64   `gorm:"column:pre_goals_overunder_under_225"`
		PreGoalsOverunderOver25                 float64   `gorm:"column:pre_goals_overunder_over_25"`
		PreGoalsOverunderUnder25                float64   `gorm:"column:pre_goals_overunder_under_25"`
		PreGoalsOverunderOver275                float64   `gorm:"column:pre_goals_overunder_over_275"`
		PreGoalsOverunderUnder275               float64   `gorm:"column:pre_goals_overunder_under_275"`
		PreGoalsOverunderOver30                 float64   `gorm:"column:pre_goals_overunder_over_30"`
		PreGoalsOverunderUnder30                float64   `gorm:"column:pre_goals_overunder_under_30"`
		PreGoalsOverunderOver325                float64   `gorm:"column:pre_goals_overunder_over_325"`
		PreGoalsOverunderUnder325               float64   `gorm:"column:pre_goals_overunder_under_325"`
		PreGoalsOverunderOver35                 float64   `gorm:"column:pre_goals_overunder_over_35"`
		PreGoalsOverunderUnder35                float64   `gorm:"column:pre_goals_overunder_under_35"`
		PreGoalsOverunderOver375                float64   `gorm:"column:pre_goals_overunder_over_375"`
		PreGoalsOverunderUnder375               float64   `gorm:"column:pre_goals_overunder_under_375"`
		PreGoalsOverunderOver40                 float64   `gorm:"column:pre_goals_overunder_over_40"`
		PreGoalsOverunderUnder40                float64   `gorm:"column:pre_goals_overunder_under_40"`
		PreGoalsOverunderOver425                float64   `gorm:"column:pre_goals_overunder_over_425"`
		PreGoalsOverunderUnder425               float64   `gorm:"column:pre_goals_overunder_under_425"`
		PreGoalsOverunderOver45                 float64   `gorm:"column:pre_goals_overunder_over_45"`
		PreGoalsOverunderUnder45                float64   `gorm:"column:pre_goals_overunder_under_45"`
		PreGoalsOverunderOver475                float64   `gorm:"column:pre_goals_overunder_over_475"`
		PreGoalsOverunderUnder475               float64   `gorm:"column:pre_goals_overunder_under_475"`
		PreGoalsOverunderOver50                 float64   `gorm:"column:pre_goals_overunder_over_50"`
		PreGoalsOverunderUnder50                float64   `gorm:"column:pre_goals_overunder_under_50"`
		PreGoalsOverunderOver525                float64   `gorm:"column:pre_goals_overunder_over_525"`
		PreGoalsOverunderUnder525               float64   `gorm:"column:pre_goals_overunder_under_525"`
		PreGoalsOverunderOver55                 float64   `gorm:"column:pre_goals_overunder_over_55"`
		PreGoalsOverunderUnder55                float64   `gorm:"column:pre_goals_overunder_under_55"`
		PreGoalsOverunderOver575                float64   `gorm:"column:pre_goals_overunder_over_575"`
		PreGoalsOverunderUnder575               float64   `gorm:"column:pre_goals_overunder_under_575"`
		PreGoalsOverunderOver60                 float64   `gorm:"column:pre_goals_overunder_over_60"`
		PreGoalsOverunderUnder60                float64   `gorm:"column:pre_goals_overunder_under_60"`
		PreGoalsOverunderOver625                float64   `gorm:"column:pre_goals_overunder_over_625"`
		PreGoalsOverunderUnder625               float64   `gorm:"column:pre_goals_overunder_under_625"`
		PreGoalsOverunderOver65                 float64   `gorm:"column:pre_goals_overunder_over_65"`
		PreGoalsOverunderUnder65                float64   `gorm:"column:pre_goals_overunder_under_65"`
		PreGoalsOverunderOver675                float64   `gorm:"column:pre_goals_overunder_over_675"`
		PreGoalsOverunderUnder675               float64   `gorm:"column:pre_goals_overunder_under_675"`
		PreGoalsOverunderOver70                 float64   `gorm:"column:pre_goals_overunder_over_70"`
		PreGoalsOverunderUnder70                float64   `gorm:"column:pre_goals_overunder_under_70"`
		PreGoalsOverunderOver75                 float64   `gorm:"column:pre_goals_overunder_over_75"`
		PreGoalsOverunderUnder75                float64   `gorm:"column:pre_goals_overunder_under_75"`
		PreGoalsOverunderOver85                 float64   `gorm:"column:pre_goals_overunder_over_85"`
		PreGoalsOverunderUnder85                float64   `gorm:"column:pre_goals_overunder_under_85"`
		PreGoalsOverunderOver95                 float64   `gorm:"column:pre_goals_overunder_over_95"`
		PreGoalsOverunderUnder95                float64   `gorm:"column:pre_goals_overunder_under_95"`
		EndGoalsOverunderOver05                 float64   `gorm:"column:end_goals_overunder_over_05"`
		EndGoalsOverunderUnder05                float64   `gorm:"column:end_goals_overunder_under_05"`
		EndGoalsOverunderOver075                float64   `gorm:"column:end_goals_overunder_over_075"`
		EndGoalsOverunderUnder075               float64   `gorm:"column:end_goals_overunder_under_075"`
		EndGoalsOverunderOver10                 float64   `gorm:"column:end_goals_overunder_over_10"`
		EndGoalsOverunderUnder10                float64   `gorm:"column:end_goals_overunder_under_10"`
		EndGoalsOverunderOver125                float64   `gorm:"column:end_goals_overunder_over_125"`
		EndGoalsOverunderUnder125               float64   `gorm:"column:end_goals_overunder_under_125"`
		EndGoalsOverunderOver15                 float64   `gorm:"column:end_goals_overunder_over_15"`
		EndGoalsOverunderUnder15                float64   `gorm:"column:end_goals_overunder_under_15"`
		EndGoalsOverunderOver175                float64   `gorm:"column:end_goals_overunder_over_175"`
		EndGoalsOverunderUnder175               float64   `gorm:"column:end_goals_overunder_under_175"`
		EndGoalsOverunderOver20                 float64   `gorm:"column:end_goals_overunder_over_20"`
		EndGoalsOverunderUnder20                float64   `gorm:"column:end_goals_overunder_under_20"`
		EndGoalsOverunderOver225                float64   `gorm:"column:end_goals_overunder_over_225"`
		EndGoalsOverunderUnder225               float64   `gorm:"column:end_goals_overunder_under_225"`
		EndGoalsOverunderOver25                 float64   `gorm:"column:end_goals_overunder_over_25"`
		EndGoalsOverunderUnder25                float64   `gorm:"column:end_goals_overunder_under_25"`
		EndGoalsOverunderOver275                float64   `gorm:"column:end_goals_overunder_over_275"`
		EndGoalsOverunderUnder275               float64   `gorm:"column:end_goals_overunder_under_275"`
		EndGoalsOverunderOver30                 float64   `gorm:"column:end_goals_overunder_over_30"`
		EndGoalsOverunderUnder30                float64   `gorm:"column:end_goals_overunder_under_30"`
		EndGoalsOverunderOver325                float64   `gorm:"column:end_goals_overunder_over_325"`
		EndGoalsOverunderUnder325               float64   `gorm:"column:end_goals_overunder_under_325"`
		EndGoalsOverunderOver35                 float64   `gorm:"column:end_goals_overunder_over_35"`
		EndGoalsOverunderUnder35                float64   `gorm:"column:end_goals_overunder_under_35"`
		EndGoalsOverunderOver375                float64   `gorm:"column:end_goals_overunder_over_375"`
		EndGoalsOverunderUnder375               float64   `gorm:"column:end_goals_overunder_under_375"`
		EndGoalsOverunderOver40                 float64   `gorm:"column:end_goals_overunder_over_40"`
		EndGoalsOverunderUnder40                float64   `gorm:"column:end_goals_overunder_under_40"`
		EndGoalsOverunderOver425                float64   `gorm:"column:end_goals_overunder_over_425"`
		EndGoalsOverunderUnder425               float64   `gorm:"column:end_goals_overunder_under_425"`
		EndGoalsOverunderOver45                 float64   `gorm:"column:end_goals_overunder_over_45"`
		EndGoalsOverunderUnder45                float64   `gorm:"column:end_goals_overunder_under_45"`
		EndGoalsOverunderOver475                float64   `gorm:"column:end_goals_overunder_over_475"`
		EndGoalsOverunderUnder475               float64   `gorm:"column:end_goals_overunder_under_475"`
		EndGoalsOverunderOver50                 float64   `gorm:"column:end_goals_overunder_over_50"`
		EndGoalsOverunderUnder50                float64   `gorm:"column:end_goals_overunder_under_50"`
		EndGoalsOverunderOver525                float64   `gorm:"column:end_goals_overunder_over_525"`
		EndGoalsOverunderUnder525               float64   `gorm:"column:end_goals_overunder_under_525"`
		EndGoalsOverunderOver55                 float64   `gorm:"column:end_goals_overunder_over_55"`
		EndGoalsOverunderUnder55                float64   `gorm:"column:end_goals_overunder_under_55"`
		EndGoalsOverunderOver575                float64   `gorm:"column:end_goals_overunder_over_575"`
		EndGoalsOverunderUnder575               float64   `gorm:"column:end_goals_overunder_under_575"`
		EndGoalsOverunderOver60                 float64   `gorm:"column:end_goals_overunder_over_60"`
		EndGoalsOverunderUnder60                float64   `gorm:"column:end_goals_overunder_under_60"`
		EndGoalsOverunderOver625                float64   `gorm:"column:end_goals_overunder_over_625"`
		EndGoalsOverunderUnder625               float64   `gorm:"column:end_goals_overunder_under_625"`
		EndGoalsOverunderOver65                 float64   `gorm:"column:end_goals_overunder_over_65"`
		EndGoalsOverunderUnder65                float64   `gorm:"column:end_goals_overunder_under_65"`
		EndGoalsOverunderOver675                float64   `gorm:"column:end_goals_overunder_over_675"`
		EndGoalsOverunderUnder675               float64   `gorm:"column:end_goals_overunder_under_675"`
		EndGoalsOverunderOver70                 float64   `gorm:"column:end_goals_overunder_over_70"`
		EndGoalsOverunderUnder70                float64   `gorm:"column:end_goals_overunder_under_70"`
		EndGoalsOverunderOver75                 float64   `gorm:"column:end_goals_overunder_over_75"`
		EndGoalsOverunderUnder75                float64   `gorm:"column:end_goals_overunder_under_75"`
		EndGoalsOverunderOver85                 float64   `gorm:"column:end_goals_overunder_over_85"`
		EndGoalsOverunderUnder85                float64   `gorm:"column:end_goals_overunder_under_85"`
		EndGoalsOverunderOver95                 float64   `gorm:"column:end_goals_overunder_over_95"`
		EndGoalsOverunderUnder95                float64   `gorm:"column:end_goals_overunder_under_95"`
		PreAsianHandicapFirstHalfHomeMin175     float64   `gorm:"column:pre_asian_handicap_first_half_home_min_175"`
		PreAsianHandicapFirstHalfAwayMin175     float64   `gorm:"column:pre_asian_handicap_first_half_away_min_175"`
		PreAsianHandicapFirstHalfHomeMin15      float64   `gorm:"column:pre_asian_handicap_first_half_home_min_15"`
		PreAsianHandicapFirstHalfAwayMin15      float64   `gorm:"column:pre_asian_handicap_first_half_away_min_15"`
		PreAsianHandicapFirstHalfHomeMin125     float64   `gorm:"column:pre_asian_handicap_first_half_home_min_125"`
		PreAsianHandicapFirstHalfAwayMin125     float64   `gorm:"column:pre_asian_handicap_first_half_away_min_125"`
		PreAsianHandicapFirstHalfHomeMin1       float64   `gorm:"column:pre_asian_handicap_first_half_home_min_1"`
		PreAsianHandicapFirstHalfAwayMin1       float64   `gorm:"column:pre_asian_handicap_first_half_away_min_1"`
		PreAsianHandicapFirstHalfHomeMin075     float64   `gorm:"column:pre_asian_handicap_first_half_home_min_075"`
		PreAsianHandicapFirstHalfAwayMin075     float64   `gorm:"column:pre_asian_handicap_first_half_away_min_075"`
		PreAsianHandicapFirstHalfHomeMin05      float64   `gorm:"column:pre_asian_handicap_first_half_home_min_05"`
		PreAsianHandicapFirstHalfAwayMin05      float64   `gorm:"column:pre_asian_handicap_first_half_away_min_05"`
		PreAsianHandicapFirstHalfHomeMin025     float64   `gorm:"column:pre_asian_handicap_first_half_home_min_025"`
		PreAsianHandicapFirstHalfAwayMin025     float64   `gorm:"column:pre_asian_handicap_first_half_away_min_025"`
		PreAsianHandicapFirstHalfHomePlus0      float64   `gorm:"column:pre_asian_handicap_first_half_home_plus_0"`
		PreAsianHandicapFirstHalfAwayPlus0      float64   `gorm:"column:pre_asian_handicap_first_half_away_plus_0"`
		PreAsianHandicapFirstHalfHomePlus025    float64   `gorm:"column:pre_asian_handicap_first_half_home_plus_025"`
		PreAsianHandicapFirstHalfAwayPlus025    float64   `gorm:"column:pre_asian_handicap_first_half_away_plus_025"`
		PreAsianHandicapFirstHalfHomePlus05     float64   `gorm:"column:pre_asian_handicap_first_half_home_plus_05"`
		PreAsianHandicapFirstHalfAwayPlus05     float64   `gorm:"column:pre_asian_handicap_first_half_away_plus_05"`
		PreAsianHandicapFirstHalfHomePlus075    float64   `gorm:"column:pre_asian_handicap_first_half_home_plus_075"`
		PreAsianHandicapFirstHalfAwayPlus075    float64   `gorm:"column:pre_asian_handicap_first_half_away_plus_075"`
		PreAsianHandicapFirstHalfHomePlus1      float64   `gorm:"column:pre_asian_handicap_first_half_home_plus_1"`
		PreAsianHandicapFirstHalfAwayPlus1      float64   `gorm:"column:pre_asian_handicap_first_half_away_plus_1"`
		PreAsianHandicapFirstHalfHomePlus125    float64   `gorm:"column:pre_asian_handicap_first_half_home_plus_125"`
		PreAsianHandicapFirstHalfAwayPlus125    float64   `gorm:"column:pre_asian_handicap_first_half_away_plus_125"`
		PreAsianHandicapFirstHalfHomePlus15     float64   `gorm:"column:pre_asian_handicap_first_half_home_plus_15"`
		PreAsianHandicapFirstHalfAwayPlus15     float64   `gorm:"column:pre_asian_handicap_first_half_away_plus_15"`
		PreAsianHandicapFirstHalfHomePlus175    float64   `gorm:"column:pre_asian_handicap_first_half_home_plus_175"`
		PreAsianHandicapFirstHalfAwayPlus175    float64   `gorm:"column:pre_asian_handicap_first_half_away_plus_175"`
		EndAsianHandicapFirstHalfHomeMin175     float64   `gorm:"column:end_asian_handicap_first_half_home_min_175"`
		EndAsianHandicapFirstHalfAwayMin175     float64   `gorm:"column:end_asian_handicap_first_half_away_min_175"`
		EndAsianHandicapFirstHalfHomeMin15      float64   `gorm:"column:end_asian_handicap_first_half_home_min_15"`
		EndAsianHandicapFirstHalfAwayMin15      float64   `gorm:"column:end_asian_handicap_first_half_away_min_15"`
		EndAsianHandicapFirstHalfHomeMin125     float64   `gorm:"column:end_asian_handicap_first_half_home_min_125"`
		EndAsianHandicapFirstHalfAwayMin125     float64   `gorm:"column:end_asian_handicap_first_half_away_min_125"`
		EndAsianHandicapFirstHalfHomeMin1       float64   `gorm:"column:end_asian_handicap_first_half_home_min_1"`
		EndAsianHandicapFirstHalfAwayMin1       float64   `gorm:"column:end_asian_handicap_first_half_away_min_1"`
		EndAsianHandicapFirstHalfHomeMin075     float64   `gorm:"column:end_asian_handicap_first_half_home_min_075"`
		EndAsianHandicapFirstHalfAwayMin075     float64   `gorm:"column:end_asian_handicap_first_half_away_min_075"`
		EndAsianHandicapFirstHalfHomeMin05      float64   `gorm:"column:end_asian_handicap_first_half_home_min_05"`
		EndAsianHandicapFirstHalfAwayMin05      float64   `gorm:"column:end_asian_handicap_first_half_away_min_05"`
		EndAsianHandicapFirstHalfHomeMin025     float64   `gorm:"column:end_asian_handicap_first_half_home_min_025"`
		EndAsianHandicapFirstHalfAwayMin025     float64   `gorm:"column:end_asian_handicap_first_half_away_min_025"`
		EndAsianHandicapFirstHalfHomePlus0      float64   `gorm:"column:end_asian_handicap_first_half_home_plus_0"`
		EndAsianHandicapFirstHalfAwayPlus0      float64   `gorm:"column:end_asian_handicap_first_half_away_plus_0"`
		EndAsianHandicapFirstHalfHomePlus025    float64   `gorm:"column:end_asian_handicap_first_half_home_plus_025"`
		EndAsianHandicapFirstHalfAwayPlus025    float64   `gorm:"column:end_asian_handicap_first_half_away_plus_025"`
		EndAsianHandicapFirstHalfHomePlus05     float64   `gorm:"column:end_asian_handicap_first_half_home_plus_05"`
		EndAsianHandicapFirstHalfAwayPlus05     float64   `gorm:"column:end_asian_handicap_first_half_away_plus_05"`
		EndAsianHandicapFirstHalfHomePlus075    float64   `gorm:"column:end_asian_handicap_first_half_home_plus_075"`
		EndAsianHandicapFirstHalfAwayPlus075    float64   `gorm:"column:end_asian_handicap_first_half_away_plus_075"`
		EndAsianHandicapFirstHalfHomePlus1      float64   `gorm:"column:end_asian_handicap_first_half_home_plus_1"`
		EndAsianHandicapFirstHalfAwayPlus1      float64   `gorm:"column:end_asian_handicap_first_half_away_plus_1"`
		EndAsianHandicapFirstHalfHomePlus125    float64   `gorm:"column:end_asian_handicap_first_half_home_plus_125"`
		EndAsianHandicapFirstHalfAwayPlus125    float64   `gorm:"column:end_asian_handicap_first_half_away_plus_125"`
		EndAsianHandicapFirstHalfHomePlus15     float64   `gorm:"column:end_asian_handicap_first_half_home_plus_15"`
		EndAsianHandicapFirstHalfAwayPlus15     float64   `gorm:"column:end_asian_handicap_first_half_away_plus_15"`
		EndAsianHandicapFirstHalfHomePlus175    float64   `gorm:"column:end_asian_handicap_first_half_home_plus_175"`
		EndAsianHandicapFirstHalfAwayPlus175    float64   `gorm:"column:end_asian_handicap_first_half_away_plus_175"`
		PreGoalsOverunderFirstHalfOver05        float64   `gorm:"column:pre_goals_overunder_first_half_over_05"`
		PreGoalsOverunderFirstHalfUnder05       float64   `gorm:"column:pre_goals_overunder_first_half_under_05"`
		PreGoalsOverunderFirstHalfOver075       float64   `gorm:"column:pre_goals_overunder_first_half_over_075"`
		PreGoalsOverunderFirstHalfUnder075      float64   `gorm:"column:pre_goals_overunder_first_half_under_075"`
		PreGoalsOverunderFirstHalfOver10        float64   `gorm:"column:pre_goals_overunder_first_half_over_10"`
		PreGoalsOverunderFirstHalfUnder10       float64   `gorm:"column:pre_goals_overunder_first_half_under_10"`
		PreGoalsOverunderFirstHalfOver125       float64   `gorm:"column:pre_goals_overunder_first_half_over_125"`
		PreGoalsOverunderFirstHalfUnder125      float64   `gorm:"column:pre_goals_overunder_first_half_under_125"`
		PreGoalsOverunderFirstHalfOver15        float64   `gorm:"column:pre_goals_overunder_first_half_over_15"`
		PreGoalsOverunderFirstHalfUnder15       float64   `gorm:"column:pre_goals_overunder_first_half_under_15"`
		PreGoalsOverunderFirstHalfOver175       float64   `gorm:"column:pre_goals_overunder_first_half_over_175"`
		PreGoalsOverunderFirstHalfUnder175      float64   `gorm:"column:pre_goals_overunder_first_half_under_175"`
		PreGoalsOverunderFirstHalfOver20        float64   `gorm:"column:pre_goals_overunder_first_half_over_20"`
		PreGoalsOverunderFirstHalfUnder20       float64   `gorm:"column:pre_goals_overunder_first_half_under_20"`
		PreGoalsOverunderFirstHalfOver225       float64   `gorm:"column:pre_goals_overunder_first_half_over_225"`
		PreGoalsOverunderFirstHalfUnder225      float64   `gorm:"column:pre_goals_overunder_first_half_under_225"`
		PreGoalsOverunderFirstHalfOver25        float64   `gorm:"column:pre_goals_overunder_first_half_over_25"`
		PreGoalsOverunderFirstHalfUnder25       float64   `gorm:"column:pre_goals_overunder_first_half_under_25"`
		PreGoalsOverunderFirstHalfOver275       float64   `gorm:"column:pre_goals_overunder_first_half_over_275"`
		PreGoalsOverunderFirstHalfUnder275      float64   `gorm:"column:pre_goals_overunder_first_half_under_275"`
		PreGoalsOverunderFirstHalfOver30        float64   `gorm:"column:pre_goals_overunder_first_half_over_30"`
		PreGoalsOverunderFirstHalfUnder30       float64   `gorm:"column:pre_goals_overunder_first_half_under_30"`
		PreGoalsOverunderFirstHalfOver325       float64   `gorm:"column:pre_goals_overunder_first_half_over_325"`
		PreGoalsOverunderFirstHalfUnder325      float64   `gorm:"column:pre_goals_overunder_first_half_under_325"`
		PreGoalsOverunderFirstHalfOver35        float64   `gorm:"column:pre_goals_overunder_first_half_over_35"`
		PreGoalsOverunderFirstHalfUnder35       float64   `gorm:"column:pre_goals_overunder_first_half_under_35"`
		PreGoalsOverunderFirstHalfOver375       float64   `gorm:"column:pre_goals_overunder_first_half_over_375"`
		PreGoalsOverunderFirstHalfUnder375      float64   `gorm:"column:pre_goals_overunder_first_half_under_375"`
		EndGoalsOverunderFirstHalfOver05        float64   `gorm:"column:end_goals_overunder_first_half_over_05"`
		EndGoalsOverunderFirstHalfUnder05       float64   `gorm:"column:end_goals_overunder_first_half_under_05"`
		EndGoalsOverunderFirstHalfOver075       float64   `gorm:"column:end_goals_overunder_first_half_over_075"`
		EndGoalsOverunderFirstHalfUnder075      float64   `gorm:"column:end_goals_overunder_first_half_under_075"`
		EndGoalsOverunderFirstHalfOver10        float64   `gorm:"column:end_goals_overunder_first_half_over_10"`
		EndGoalsOverunderFirstHalfUnder10       float64   `gorm:"column:end_goals_overunder_first_half_under_10"`
		EndGoalsOverunderFirstHalfOver125       float64   `gorm:"column:end_goals_overunder_first_half_over_125"`
		EndGoalsOverunderFirstHalfUnder125      float64   `gorm:"column:end_goals_overunder_first_half_under_125"`
		EndGoalsOverunderFirstHalfOver15        float64   `gorm:"column:end_goals_overunder_first_half_over_15"`
		EndGoalsOverunderFirstHalfUnder15       float64   `gorm:"column:end_goals_overunder_first_half_under_15"`
		EndGoalsOverunderFirstHalfOver175       float64   `gorm:"column:end_goals_overunder_first_half_over_175"`
		EndGoalsOverunderFirstHalfUnder175      float64   `gorm:"column:end_goals_overunder_first_half_under_175"`
		EndGoalsOverunderFirstHalfOver20        float64   `gorm:"column:end_goals_overunder_first_half_over_20"`
		EndGoalsOverunderFirstHalfUnder20       float64   `gorm:"column:end_goals_overunder_first_half_under_20"`
		EndGoalsOverunderFirstHalfOver225       float64   `gorm:"column:end_goals_overunder_first_half_over_225"`
		EndGoalsOverunderFirstHalfUnder225      float64   `gorm:"column:end_goals_overunder_first_half_under_225"`
		EndGoalsOverunderFirstHalfOver25        float64   `gorm:"column:end_goals_overunder_first_half_over_25"`
		EndGoalsOverunderFirstHalfUnder25       float64   `gorm:"column:end_goals_overunder_first_half_under_25"`
		EndGoalsOverunderFirstHalfOver275       float64   `gorm:"column:end_goals_overunder_first_half_over_275"`
		EndGoalsOverunderFirstHalfUnder275      float64   `gorm:"column:end_goals_overunder_first_half_under_275"`
		EndGoalsOverunderFirstHalfOver30        float64   `gorm:"column:end_goals_overunder_first_half_over_30"`
		EndGoalsOverunderFirstHalfUnder30       float64   `gorm:"column:end_goals_overunder_first_half_under_30"`
		EndGoalsOverunderFirstHalfOver325       float64   `gorm:"column:end_goals_overunder_first_half_over_325"`
		EndGoalsOverunderFirstHalfUnder325      float64   `gorm:"column:end_goals_overunder_first_half_under_325"`
		EndGoalsOverunderFirstHalfOver35        float64   `gorm:"column:end_goals_overunder_first_half_over_35"`
		EndGoalsOverunderFirstHalfUnder35       float64   `gorm:"column:end_goals_overunder_first_half_under_35"`
		EndGoalsOverunderFirstHalfOver375       float64   `gorm:"column:end_goals_overunder_first_half_over_375"`
		EndGoalsOverunderFirstHalfUnder375      float64   `gorm:"column:end_goals_overunder_first_half_under_375"`
		PreGoalsOverunderSecondHalfOver05       float64   `gorm:"column:pre_goals_overunder__second_half_over_05"`
		PreGoalsOverunderSecondHalfUnder05      float64   `gorm:"column:pre_goals_overunder__second_half_under_05"`
		PreGoalsOverunderSecondHalfOver075      float64   `gorm:"column:pre_goals_overunder__second_half_over_075"`
		PreGoalsOverunderSecondHalfUnder075     float64   `gorm:"column:pre_goals_overunder__second_half_under_075"`
		PreGoalsOverunderSecondHalfOver10       float64   `gorm:"column:pre_goals_overunder__second_half_over_10"`
		PreGoalsOverunderSecondHalfUnder10      float64   `gorm:"column:pre_goals_overunder__second_half_under_10"`
		PreGoalsOverunderSecondHalfOver125      float64   `gorm:"column:pre_goals_overunder__second_half_over_125"`
		PreGoalsOverunderSecondHalfUnder125     float64   `gorm:"column:pre_goals_overunder__second_half_under_125"`
		PreGoalsOverunderSecondHalfOver15       float64   `gorm:"column:pre_goals_overunder__second_half_over_15"`
		PreGoalsOverunderSecondHalfUnder15      float64   `gorm:"column:pre_goals_overunder__second_half_under_15"`
		PreGoalsOverunderSecondHalfOver175      float64   `gorm:"column:pre_goals_overunder__second_half_over_175"`
		PreGoalsOverunderSecondHalfUnder175     float64   `gorm:"column:pre_goals_overunder__second_half_under_175"`
		PreGoalsOverunderSecondHalfOver20       float64   `gorm:"column:pre_goals_overunder__second_half_over_20"`
		PreGoalsOverunderSecondHalfUnder20      float64   `gorm:"column:pre_goals_overunder__second_half_under_20"`
		PreGoalsOverunderSecondHalfOver225      float64   `gorm:"column:pre_goals_overunder__second_half_over_225"`
		PreGoalsOverunderSecondHalfUnder225     float64   `gorm:"column:pre_goals_overunder__second_half_under_225"`
		PreGoalsOverunderSecondHalfOver25       float64   `gorm:"column:pre_goals_overunder__second_half_over_25"`
		PreGoalsOverunderSecondHalfUnder25      float64   `gorm:"column:pre_goals_overunder__second_half_under_25"`
		PreGoalsOverunderSecondHalfOver275      float64   `gorm:"column:pre_goals_overunder__second_half_over_275"`
		PreGoalsOverunderSecondHalfUnder275     float64   `gorm:"column:pre_goals_overunder__second_half_under_275"`
		PreGoalsOverunderSecondHalfOver30       float64   `gorm:"column:pre_goals_overunder__second_half_over_30"`
		PreGoalsOverunderSecondHalfUnder30      float64   `gorm:"column:pre_goals_overunder__second_half_under_30"`
		PreGoalsOverunderSecondHalfOver325      float64   `gorm:"column:pre_goals_overunder__second_half_over_325"`
		PreGoalsOverunderSecondHalfUnder325     float64   `gorm:"column:pre_goals_overunder__second_half_under_325"`
		PreGoalsOverunderSecondHalfOver35       float64   `gorm:"column:pre_goals_overunder__second_half_over_35"`
		PreGoalsOverunderSecondHalfUnder35      float64   `gorm:"column:pre_goals_overunder__second_half_under_35"`
		PreGoalsOverunderSecondHalfOver375      float64   `gorm:"column:pre_goals_overunder__second_half_over_375"`
		PreGoalsOverunderSecondHalfUnder375     float64   `gorm:"column:pre_goals_overunder__second_half_under_375"`
		EndGoalsOverunderSecondHalfOver05       float64   `gorm:"column:end_goals_overunder__second_half_over_05"`
		EndGoalsOverunderSecondHalfUnder05      float64   `gorm:"column:end_goals_overunder__second_half_under_05"`
		EndGoalsOverunderSecondHalfOver075      float64   `gorm:"column:end_goals_overunder__second_half_over_075"`
		EndGoalsOverunderSecondHalfUnder075     float64   `gorm:"column:end_goals_overunder__second_half_under_075"`
		EndGoalsOverunderSecondHalfOver10       float64   `gorm:"column:end_goals_overunder__second_half_over_10"`
		EndGoalsOverunderSecondHalfUnder10      float64   `gorm:"column:end_goals_overunder__second_half_under_10"`
		EndGoalsOverunderSecondHalfOver125      float64   `gorm:"column:end_goals_overunder__second_half_over_125"`
		EndGoalsOverunderSecondHalfUnder125     float64   `gorm:"column:end_goals_overunder__second_half_under_125"`
		EndGoalsOverunderSecondHalfOver15       float64   `gorm:"column:end_goals_overunder__second_half_over_15"`
		EndGoalsOverunderSecondHalfUnder15      float64   `gorm:"column:end_goals_overunder__second_half_under_15"`
		EndGoalsOverunderSecondHalfOver175      float64   `gorm:"column:end_goals_overunder__second_half_over_175"`
		EndGoalsOverunderSecondHalfUnder175     float64   `gorm:"column:end_goals_overunder__second_half_under_175"`
		EndGoalsOverunderSecondHalfOver20       float64   `gorm:"column:end_goals_overunder__second_half_over_20"`
		EndGoalsOverunderSecondHalfUnder20      float64   `gorm:"column:end_goals_overunder__second_half_under_20"`
		EndGoalsOverunderSecondHalfOver225      float64   `gorm:"column:end_goals_overunder__second_half_over_225"`
		EndGoalsOverunderSecondHalfUnder225     float64   `gorm:"column:end_goals_overunder__second_half_under_225"`
		EndGoalsOverunderSecondHalfOver25       float64   `gorm:"column:end_goals_overunder__second_half_over_25"`
		EndGoalsOverunderSecondHalfUnder25      float64   `gorm:"column:end_goals_overunder__second_half_under_25"`
		EndGoalsOverunderSecondHalfOver275      float64   `gorm:"column:end_goals_overunder__second_half_over_275"`
		EndGoalsOverunderSecondHalfUnder275     float64   `gorm:"column:end_goals_overunder__second_half_under_275"`
		EndGoalsOverunderSecondHalfOver30       float64   `gorm:"column:end_goals_overunder__second_half_over_30"`
		EndGoalsOverunderSecondHalfUnder30      float64   `gorm:"column:end_goals_overunder__second_half_under_30"`
		EndGoalsOverunderSecondHalfOver325      float64   `gorm:"column:end_goals_overunder__second_half_over_325"`
		EndGoalsOverunderSecondHalfUnder325     float64   `gorm:"column:end_goals_overunder__second_half_under_325"`
		EndGoalsOverunderSecondHalfOver35       float64   `gorm:"column:end_goals_overunder__second_half_over_35"`
		EndGoalsOverunderSecondHalfUnder35      float64   `gorm:"column:end_goals_overunder__second_half_under_35"`
		EndGoalsOverunderSecondHalfOver375      float64   `gorm:"column:end_goals_overunder__second_half_over_375"`
		EndGoalsOverunderSecondHalfUnder375     float64   `gorm:"column:end_goals_overunder__second_half_under_375"`
		PreHtftDoubleHomeHome                   float64   `gorm:"column:pre_htft_double_home_home"`
		PreHtftDoubleHomeDraw                   float64   `gorm:"column:pre_htft_double_home_draw"`
		PreHtftDoubleHomeAway                   float64   `gorm:"column:pre_htft_double_home_away"`
		PreHtftDoubleDrawHome                   float64   `gorm:"column:pre_htft_double_draw_home"`
		PreHtftDoubleDrawDraw                   float64   `gorm:"column:pre_htft_double_draw_draw"`
		PreHtftDoubleDrawAway                   float64   `gorm:"column:pre_htft_double_draw_away"`
		PreHtftDoubleAwayHome                   float64   `gorm:"column:pre_htft_double_away_home"`
		PreHtftDoubleAwayDraw                   float64   `gorm:"column:pre_htft_double_away_draw"`
		PreHtftDoubleAwayAway                   float64   `gorm:"column:pre_htft_double_away_away"`
		EndHtftDoubleHomeHome                   float64   `gorm:"column:end_htft_double_home_home"`
		EndHtftDoubleHomeDraw                   float64   `gorm:"column:end_htft_double_home_draw"`
		EndHtftDoubleHomeAway                   float64   `gorm:"column:end_htft_double_home_away"`
		EndHtftDoubleDrawHome                   float64   `gorm:"column:end_htft_double_draw_home"`
		EndHtftDoubleDrawDraw                   float64   `gorm:"column:end_htft_double_draw_draw"`
		EndHtftDoubleDrawAway                   float64   `gorm:"column:end_htft_double_draw_away"`
		EndHtftDoubleAwayHome                   float64   `gorm:"column:end_htft_double_away_home"`
		EndHtftDoubleAwayDraw                   float64   `gorm:"column:end_htft_double_away_draw"`
		EndHtftDoubleAwayAway                   float64   `gorm:"column:end_htft_double_away_away"`
		PreBothTeamsScoreYes                    float64   `gorm:"column:pre_both_teams_score_yes"`
		PreBothTeamsScoreNo                     float64   `gorm:"column:pre_both_teams_score_no"`
		EndBothTeamsScoreYes                    float64   `gorm:"column:end_both_teams_score_yes"`
		EndBothTeamsScoreNo                     float64   `gorm:"column:end_both_teams_score_no"`
		PreBothTeamsScoreFirstHalfYes           float64   `gorm:"column:pre_both_teams_score__first_half_yes"`
		PreBothTeamsScoreFirstHalfNo            float64   `gorm:"column:pre_both_teams_score__first_half_no"`
		EndBothTeamsScoreFirstHalfYes           float64   `gorm:"column:end_both_teams_score__first_half_yes"`
		EndBothTeamsScoreFirstHalfNo            float64   `gorm:"column:end_both_teams_score__first_half_no"`
		PreBothTeamsToScoreSecondHalfYes        float64   `gorm:"column:pre_both_teams_to_score__second_half_yes"`
		PreBothTeamsToScoreSecondHalfNo         float64   `gorm:"column:pre_both_teams_to_score__second_half_no"`
		EndBothTeamsToScoreSecondHalfYes        float64   `gorm:"column:end_both_teams_to_score__second_half_yes"`
		EndBothTeamsToScoreSecondHalfNo         float64   `gorm:"column:end_both_teams_to_score__second_half_no"`
		PreResultsBothTeamsScoreHomeYes         float64   `gorm:"column:pre_results_both_teams_score_home_yes"`
		PreResultsBothTeamsScoreDrawYes         float64   `gorm:"column:pre_results_both_teams_score_draw_yes"`
		PreResultsBothTeamsScoreAwayYes         float64   `gorm:"column:pre_results_both_teams_score_away_yes"`
		PreResultsBothTeamsScoreHomeNo          float64   `gorm:"column:pre_results_both_teams_score_home_no"`
		PreResultsBothTeamsScoreDrawNo          float64   `gorm:"column:pre_results_both_teams_score_draw_no"`
		PreResultsBothTeamsScoreAwayNo          float64   `gorm:"column:pre_results_both_teams_score_away_no"`
		EndResultsBothTeamsScoreHomeYes         float64   `gorm:"column:end_results_both_teams_score_home_yes"`
		EndResultsBothTeamsScoreDrawYes         float64   `gorm:"column:end_results_both_teams_score_draw_yes"`
		EndResultsBothTeamsScoreAwayYes         float64   `gorm:"column:end_results_both_teams_score_away_yes"`
		EndResultsBothTeamsScoreHomeNo          float64   `gorm:"column:end_results_both_teams_score_home_no"`
		EndResultsBothTeamsScoreDrawNo          float64   `gorm:"column:end_results_both_teams_score_draw_no"`
		EndResultsBothTeamsScoreAwayNo          float64   `gorm:"column:end_results_both_teams_score_away_no"`
		PreToScoreInBothHalvesByTeamsHome       float64   `gorm:"column:pre_to_score_in_both_halves_by_teams_home"`
		PreToScoreInBothHalvesByTeamsAway       float64   `gorm:"column:pre_to_score_in_both_halves_by_teams_away"`
		EndToScoreInBothHalvesByTeamsHome       float64   `gorm:"column:end_to_score_in_both_halves_by_teams_home"`
		EndToScoreInBothHalvesByTeamsAway       float64   `gorm:"column:end_to_score_in_both_halves_by_teams_away"`
		PreTotalGoalsBothTeamsToScoreOverYes25  float64   `gorm:"column:pre_total_goals_both_teams_to_score_over_yes_25"`
		PreTotalGoalsBothTeamsToScoreOverNo25   float64   `gorm:"column:pre_total_goals_both_teams_to_score_over_no_25"`
		PreTotalGoalsBothTeamsToScoreUnderYes25 float64   `gorm:"column:pre_total_goals_both_teams_to_score_under_yes_25"`
		PreTotalGoalsBothTeamsToScoreUnderNo25  float64   `gorm:"column:pre_total_goals_both_teams_to_score_under_no_25"`
		EndTotalGoalsBothTeamsToScoreOverYes25  float64   `gorm:"column:end_total_goals_both_teams_to_score_over_yes_25"`
		EndTotalGoalsBothTeamsToScoreOverNo25   float64   `gorm:"column:end_total_goals_both_teams_to_score_over_no_25"`
		EndTotalGoalsBothTeamsToScoreUnderYes25 float64   `gorm:"column:end_total_goals_both_teams_to_score_under_yes_25"`
		EndTotalGoalsBothTeamsToScoreUnderNo25  float64   `gorm:"column:end_total_goals_both_teams_to_score_under_no_25"`
		PreBothTeamsToScore1StHalf2NdHalfYesYes float64   `gorm:"column:pre_both_teams_to_score_1st_half__2nd_half_yes_yes"`
		PreBothTeamsToScore1StHalf2NdHalfYesNo  float64   `gorm:"column:pre_both_teams_to_score_1st_half__2nd_half_yes_no"`
		PreBothTeamsToScore1StHalf2NdHalfNoYes  float64   `gorm:"column:pre_both_teams_to_score_1st_half__2nd_half_no_yes"`
		PreBothTeamsToScore1StHalf2NdHalfNoNo   float64   `gorm:"column:pre_both_teams_to_score_1st_half__2nd_half_no_no"`
		EndBothTeamsToScore1StHalf2NdHalfYesYes float64   `gorm:"column:end_both_teams_to_score_1st_half__2nd_half_yes_yes"`
		EndBothTeamsToScore1StHalf2NdHalfYesNo  float64   `gorm:"column:end_both_teams_to_score_1st_half__2nd_half_yes_no"`
		EndBothTeamsToScore1StHalf2NdHalfNoYes  float64   `gorm:"column:end_both_teams_to_score_1st_half__2nd_half_no_yes"`
		EndBothTeamsToScore1StHalf2NdHalfNoNo   float64   `gorm:"column:end_both_teams_to_score_1st_half__2nd_half_no_no"`
		PreHighestScoringHalfFirst              float64   `gorm:"column:pre_highest_scoring_half_first"`
		PreHighestScoringHalfDraw               float64   `gorm:"column:pre_highest_scoring_half_draw"`
		PreHighestScoringHalfSecond             float64   `gorm:"column:pre_highest_scoring_half_second"`
		EndHighestScoringHalfFirst              float64   `gorm:"column:end_highest_scoring_half_first"`
		EndHighestScoringHalfDraw               float64   `gorm:"column:end_highest_scoring_half_draw"`
		EndHighestScoringHalfSecond             float64   `gorm:"column:end_highest_scoring_half_second"`
		PreDoubleChanceHomeDraw                 float64   `gorm:"column:pre_double_chance_home_draw"`
		PreDoubleChanceHomeAway                 float64   `gorm:"column:pre_double_chance_home_away"`
		PreDoubleChanceDrawAway                 float64   `gorm:"column:pre_double_chance_draw_away"`
		EndDoubleChanceHomeDraw                 float64   `gorm:"column:end_double_chance_home_draw"`
		EndDoubleChanceHomeAway                 float64   `gorm:"column:end_double_chance_home_away"`
		EndDoubleChanceDrawAway                 float64   `gorm:"column:end_double_chance_draw_away"`
		PreDoubleChanceFirstHalfHomeDraw        float64   `gorm:"column:pre_double_chance__first_half_home_draw"`
		PreDoubleChanceFirstHalfHomeAway        float64   `gorm:"column:pre_double_chance__first_half_home_away"`
		PreDoubleChanceFirstHalfDrawAway        float64   `gorm:"column:pre_double_chance__first_half_draw_away"`
		EndDoubleChanceFirstHalfHomeDraw        float64   `gorm:"column:end_double_chance__first_half_home_draw"`
		EndDoubleChanceFirstHalfHomeAway        float64   `gorm:"column:end_double_chance__first_half_home_away"`
		EndDoubleChanceFirstHalfDrawAway        float64   `gorm:"column:end_double_chance__first_half_draw_away"`
		PreOddEvenOdd                           float64   `gorm:"column:pre_oddeven_odd"`
		PreOddEvenEven                          float64   `gorm:"column:pre_oddeven_even"`
		EndOddEvenOdd                           float64   `gorm:"column:end_oddeven_odd"`
		EndOddEvenEven                          float64   `gorm:"column:end_oddeven_even"`
		PreResultTotalGoalsHomeOver35           float64   `gorm:"column:pre_result_total_goals_home_over_35"`
		PreResultTotalGoalsDrawOver35           float64   `gorm:"column:pre_result_total_goals_draw_over_35"`
		PreResultTotalGoalsAwayOver35           float64   `gorm:"column:pre_result_total_goals_away_over_35"`
		PreResultTotalGoalsHomeUnder35          float64   `gorm:"column:pre_result_total_goals_home_under_35"`
		PreResultTotalGoalsDrawUnder35          float64   `gorm:"column:pre_result_total_goals_draw_under_35"`
		PreResultTotalGoalsAwayUnder35          float64   `gorm:"column:pre_result_total_goals_away_under_35"`
		PreResultTotalGoalsHomeOver25           float64   `gorm:"column:pre_result_total_goals_home_over_25"`
		PreResultTotalGoalsDrawOver25           float64   `gorm:"column:pre_result_total_goals_draw_over_25"`
		PreResultTotalGoalsAwayOver25           float64   `gorm:"column:pre_result_total_goals_away_over_25"`
		PreResultTotalGoalsHomeUnder25          float64   `gorm:"column:pre_result_total_goals_home_under_25"`
		PreResultTotalGoalsDrawUnder25          float64   `gorm:"column:pre_result_total_goals_draw_under_25"`
		PreResultTotalGoalsAwayUnder25          float64   `gorm:"column:pre_result_total_goals_away_under_25"`
		EndResultTotalGoalsHomeOver35           float64   `gorm:"column:end_result_total_goals_home_over_35"`
		EndResultTotalGoalsDrawOver35           float64   `gorm:"column:end_result_total_goals_draw_over_35"`
		EndResultTotalGoalsAwayOver35           float64   `gorm:"column:end_result_total_goals_away_over_35"`
		EndResultTotalGoalsHomeUnder35          float64   `gorm:"column:end_result_total_goals_home_under_35"`
		EndResultTotalGoalsDrawUnder35          float64   `gorm:"column:end_result_total_goals_draw_under_35"`
		EndResultTotalGoalsAwayUnder35          float64   `gorm:"column:end_result_total_goals_away_under_35"`
		EndResultTotalGoalsHomeOver25           float64   `gorm:"column:end_result_total_goals_home_over_25"`
		EndResultTotalGoalsDrawOver25           float64   `gorm:"column:end_result_total_goals_draw_over_25"`
		EndResultTotalGoalsAwayOver25           float64   `gorm:"column:end_result_total_goals_away_over_25"`
		EndResultTotalGoalsHomeUnder25          float64   `gorm:"column:end_result_total_goals_home_under_25"`
		EndResultTotalGoalsDrawUnder25          float64   `gorm:"column:end_result_total_goals_draw_under_25"`
		EndResultTotalGoalsAwayUnder25          float64   `gorm:"column:end_result_total_goals_away_under_25"`
		PreCleanSheetHomeYes                    float64   `gorm:"column:pre_clean_sheet__home_yes"`
		PreCleanSheetHomeNo                     float64   `gorm:"column:pre_clean_sheet__home_no"`
		EndCleanSheetHomeYes                    float64   `gorm:"column:end_clean_sheet__home_yes"`
		EndCleanSheetHomeNo                     float64   `gorm:"column:end_clean_sheet__home_no"`
		PreCleanSheetAwayYes                    float64   `gorm:"column:pre_clean_sheet__away_yes"`
		PreCleanSheetAwayNo                     float64   `gorm:"column:pre_clean_sheet__away_no"`
		EndCleanSheetAwayYes                    float64   `gorm:"column:end_clean_sheet__away_yes"`
		EndCleanSheetAwayNo                     float64   `gorm:"column:end_clean_sheet__away_no"`
		PreWinBothHalvesHome                    float64   `gorm:"column:pre_win_both_halves_home"`
		PreWinBothHalvesAway                    float64   `gorm:"column:pre_win_both_halves_away"`
		EndWinBothHalvesHome                    float64   `gorm:"column:end_win_both_halves_home"`
		EndWinBothHalvesAway                    float64   `gorm:"column:end_win_both_halves_away"`
		PreWinToNilHome                         float64   `gorm:"column:pre_win_to_nil_home"`
		PreWinToNilAway                         float64   `gorm:"column:pre_win_to_nil_away"`
		EndWinToNilHome                         float64   `gorm:"column:end_win_to_nil_home"`
		EndWinToNilAway                         float64   `gorm:"column:end_win_to_nil_away"`
		PreToWinEitherHalfHome                  float64   `gorm:"column:pre_to_win_either_half_home"`
		PreToWinEitherHalfAway                  float64   `gorm:"column:pre_to_win_either_half_away"`
		EndToWinEitherHalfHome                  float64   `gorm:"column:end_to_win_either_half_home"`
		EndToWinEitherHalfAway                  float64   `gorm:"column:end_to_win_either_half_away"`
		PreHalftimeResultBothTeamsScoreHomeYes  float64   `gorm:"column:pre_halftime_result_both_teams_score_home_yes"`
		PreHalftimeResultBothTeamsScoreDrawYes  float64   `gorm:"column:pre_halftime_result_both_teams_score_draw_yes"`
		PreHalftimeResultBothTeamsScoreAwayYes  float64   `gorm:"column:pre_halftime_result_both_teams_score_away_yes"`
		PreHalftimeResultBothTeamsScoreHomeNo   float64   `gorm:"column:pre_halftime_result_both_teams_score_home_no"`
		PreHalftimeResultBothTeamsScoreDrawNo   float64   `gorm:"column:pre_halftime_result_both_teams_score_draw_no"`
		PreHalftimeResultBothTeamsScoreAwayNo   float64   `gorm:"column:pre_halftime_result_both_teams_score_away_no"`
		EndHalftimeResultBothTeamsScoreHomeYes  float64   `gorm:"column:end_halftime_result_both_teams_score_home_yes"`
		EndHalftimeResultBothTeamsScoreDrawYes  float64   `gorm:"column:end_halftime_result_both_teams_score_draw_yes"`
		EndHalftimeResultBothTeamsScoreAwayYes  float64   `gorm:"column:end_halftime_result_both_teams_score_away_yes"`
		EndHalftimeResultBothTeamsScoreHomeNo   float64   `gorm:"column:end_halftime_result_both_teams_score_home_no"`
		EndHalftimeResultBothTeamsScoreDrawNo   float64   `gorm:"column:end_halftime_result_both_teams_score_draw_no"`
		EndHalftimeResultBothTeamsScoreAwayNo   float64   `gorm:"column:end_halftime_result_both_teams_score_away_no"`
		PreExactGoalsNumber0                    float64   `gorm:"column:pre_exact_goals_number_0"`
		PreExactGoalsNumber1                    float64   `gorm:"column:pre_exact_goals_number_1"`
		PreExactGoalsNumber2                    float64   `gorm:"column:pre_exact_goals_number_2"`
		PreExactGoalsNumber3                    float64   `gorm:"column:pre_exact_goals_number_3"`
		PreExactGoalsNumber4                    float64   `gorm:"column:pre_exact_goals_number_4"`
		PreExactGoalsNumber5                    float64   `gorm:"column:pre_exact_goals_number_5"`
		PreExactGoalsNumber6                    float64   `gorm:"column:pre_exact_goals_number_6"`
		PreExactGoalsNumberMore7                float64   `gorm:"column:pre_exact_goals_number_more_7"`
		EndExactGoalsNumber0                    float64   `gorm:"column:end_exact_goals_number_0"`
		EndExactGoalsNumber1                    float64   `gorm:"column:end_exact_goals_number_1"`
		EndExactGoalsNumber2                    float64   `gorm:"column:end_exact_goals_number_2"`
		EndExactGoalsNumber3                    float64   `gorm:"column:end_exact_goals_number_3"`
		EndExactGoalsNumber4                    float64   `gorm:"column:end_exact_goals_number_4"`
		EndExactGoalsNumber5                    float64   `gorm:"column:end_exact_goals_number_5"`
		EndExactGoalsNumber6                    float64   `gorm:"column:end_exact_goals_number_6"`
		EndExactGoalsNumberMore7                float64   `gorm:"column:end_exact_goals_number_more_7"`
		PreExactGoalsNumberFirstHalf0           float64   `gorm:"column:pre_exact_goals_number__first_half_0"`
		PreExactGoalsNumberFirstHalf1           float64   `gorm:"column:pre_exact_goals_number__first_half_1"`
		PreExactGoalsNumberFirstHalf2           float64   `gorm:"column:pre_exact_goals_number__first_half_2"`
		PreExactGoalsNumberFirstHalf3           float64   `gorm:"column:pre_exact_goals_number__first_half_3"`
		PreExactGoalsNumberFirstHalf4           float64   `gorm:"column:pre_exact_goals_number__first_half_4"`
		PreExactGoalsNumberFirstHalfMore5       float64   `gorm:"column:pre_exact_goals_number__first_half_more_5"`
		EndExactGoalsNumberFirstHalf0           float64   `gorm:"column:end_exact_goals_number__first_half_0"`
		EndExactGoalsNumberFirstHalf1           float64   `gorm:"column:end_exact_goals_number__first_half_1"`
		EndExactGoalsNumberFirstHalf2           float64   `gorm:"column:end_exact_goals_number__first_half_2"`
		EndExactGoalsNumberFirstHalf3           float64   `gorm:"column:end_exact_goals_number__first_half_3"`
		EndExactGoalsNumberFirstHalf4           float64   `gorm:"column:end_exact_goals_number__first_half_4"`
		EndExactGoalsNumberFirstHalfMore5       float64   `gorm:"column:end_exact_goals_number__first_half_more_5"`
		PreSecondHalfExactGoalsNumber0          float64   `gorm:"column:pre_second_half_exact_goals_number_0"`
		PreSecondHalfExactGoalsNumber1          float64   `gorm:"column:pre_second_half_exact_goals_number_1"`
		PreSecondHalfExactGoalsNumber2          float64   `gorm:"column:pre_second_half_exact_goals_number_2"`
		PreSecondHalfExactGoalsNumber3          float64   `gorm:"column:pre_second_half_exact_goals_number_3"`
		PreSecondHalfExactGoalsNumber4          float64   `gorm:"column:pre_second_half_exact_goals_number_4"`
		PreSecondHalfExactGoalsNumberMore5      float64   `gorm:"column:pre_second_half_exact_goals_number_more_5"`
		EndSecondHalfExactGoalsNumber0          float64   `gorm:"column:end_second_half_exact_goals_number_0"`
		EndSecondHalfExactGoalsNumber1          float64   `gorm:"column:end_second_half_exact_goals_number_1"`
		EndSecondHalfExactGoalsNumber2          float64   `gorm:"column:end_second_half_exact_goals_number_2"`
		EndSecondHalfExactGoalsNumber3          float64   `gorm:"column:end_second_half_exact_goals_number_3"`
		EndSecondHalfExactGoalsNumber4          float64   `gorm:"column:end_second_half_exact_goals_number_4"`
		EndSecondHalfExactGoalsNumberMore5      float64   `gorm:"column:end_second_half_exact_goals_number_more_5"`
		PreHomeTeamExactGoalsNumber0            float64   `gorm:"column:pre_home_team_exact_goals_number_0"`
		PreHomeTeamExactGoalsNumber1            float64   `gorm:"column:pre_home_team_exact_goals_number_1"`
		PreHomeTeamExactGoalsNumber2            float64   `gorm:"column:pre_home_team_exact_goals_number_2"`
		PreHomeTeamExactGoalsNumberMore3        float64   `gorm:"column:pre_home_team_exact_goals_number_more_3"`
		EndHomeTeamExactGoalsNumber0            float64   `gorm:"column:end_home_team_exact_goals_number_0"`
		EndHomeTeamExactGoalsNumber1            float64   `gorm:"column:end_home_team_exact_goals_number_1"`
		EndHomeTeamExactGoalsNumber2            float64   `gorm:"column:end_home_team_exact_goals_number_2"`
		EndHomeTeamExactGoalsNumberMore3        float64   `gorm:"column:end_home_team_exact_goals_number_more_3"`
		PreAwayTeamExactGoalsNumber0            float64   `gorm:"column:pre_away_team_exact_goals_number_0"`
		PreAwayTeamExactGoalsNumber1            float64   `gorm:"column:pre_away_team_exact_goals_number_1"`
		PreAwayTeamExactGoalsNumber2            float64   `gorm:"column:pre_away_team_exact_goals_number_2"`
		PreAwayTeamExactGoalsNumberMore3        float64   `gorm:"column:pre_away_team_exact_goals_number_more_3"`
		EndAwayTeamExactGoalsNumber0            float64   `gorm:"column:end_away_team_exact_goals_number_0"`
		EndAwayTeamExactGoalsNumber1            float64   `gorm:"column:end_away_team_exact_goals_number_1"`
		EndAwayTeamExactGoalsNumber2            float64   `gorm:"column:end_away_team_exact_goals_number_2"`
		EndAwayTeamExactGoalsNumberMore3        float64   `gorm:"column:end_away_team_exact_goals_number_more_3"`
		PreCornersWinnerHome                    float64   `gorm:"column:pre_corners_winner_home"`
		PreCornersWinnerDraw                    float64   `gorm:"column:pre_corners_winner_draw"`
		PreCornersWinnerAway                    float64   `gorm:"column:pre_corners_winner_away"`
		EndCornersWinnerHome                    float64   `gorm:"column:end_corners_winner_home"`
		EndCornersWinnerDraw                    float64   `gorm:"column:end_corners_winner_draw"`
		EndCornersWinnerAway                    float64   `gorm:"column:end_corners_winner_away"`
		Pre10OverunderOver05                    float64   `gorm:"column:pre_10_overunder_over_05"`
		Pre10OverunderUnder05                   float64   `gorm:"column:pre_10_overunder_under_05"`
		End10OverunderOver05                    float64   `gorm:"column:end_10_overunder_over_05"`
		End10OverunderUnder05                   float64   `gorm:"column:end_10_overunder_under_05"`
		PreCardsEuropeanHandicapHomeMin2        float64   `gorm:"column:pre_cards_european_handicap_home_min_2"`
		PreCardsEuropeanHandicapDrawMin2        float64   `gorm:"column:pre_cards_european_handicap_draw_min_2"`
		PreCardsEuropeanHandicapAwayMin2        float64   `gorm:"column:pre_cards_european_handicap_away_min_2"`
		PreCardsEuropeanHandicapHomeMin1        float64   `gorm:"column:pre_cards_european_handicap_home_min_1"`
		PreCardsEuropeanHandicapDrawMin1        float64   `gorm:"column:pre_cards_european_handicap_draw_min_1"`
		PreCardsEuropeanHandicapAwayMin1        float64   `gorm:"column:pre_cards_european_handicap_away_min_1"`
		PreCardsEuropeanHandicapHomePlus0       float64   `gorm:"column:pre_cards_european_handicap_home_plus_0"`
		PreCardsEuropeanHandicapDrawPlus0       float64   `gorm:"column:pre_cards_european_handicap_draw_plus_0"`
		PreCardsEuropeanHandicapAwayPlus0       float64   `gorm:"column:pre_cards_european_handicap_away_plus_0"`
		PreCardsEuropeanHandicapHomePlus1       float64   `gorm:"column:pre_cards_european_handicap_home_plus_1"`
		PreCardsEuropeanHandicapDrawPlus1       float64   `gorm:"column:pre_cards_european_handicap_draw_plus_1"`
		PreCardsEuropeanHandicapAwayPlus1       float64   `gorm:"column:pre_cards_european_handicap_away_plus_1"`
		PreCardsEuropeanHandicapHomePlus2       float64   `gorm:"column:pre_cards_european_handicap_home_plus_2"`
		PreCardsEuropeanHandicapDrawPlus2       float64   `gorm:"column:pre_cards_european_handicap_draw_plus_2"`
		PreCardsEuropeanHandicapAwayPlus2       float64   `gorm:"column:pre_cards_european_handicap_away_plus_2"`
		EndCardsEuropeanHandicapHomeMin2        float64   `gorm:"column:end_cards_european_handicap_home_min_2"`
		EndCardsEuropeanHandicapDrawMin2        float64   `gorm:"column:end_cards_european_handicap_draw_min_2"`
		EndCardsEuropeanHandicapAwayMin2        float64   `gorm:"column:end_cards_european_handicap_away_min_2"`
		EndCardsEuropeanHandicapHomeMin1        float64   `gorm:"column:end_cards_european_handicap_home_min_1"`
		EndCardsEuropeanHandicapDrawMin1        float64   `gorm:"column:end_cards_european_handicap_draw_min_1"`
		EndCardsEuropeanHandicapAwayMin1        float64   `gorm:"column:end_cards_european_handicap_away_min_1"`
		EndCardsEuropeanHandicapHomePlus0       float64   `gorm:"column:end_cards_european_handicap_home_plus_0"`
		EndCardsEuropeanHandicapDrawPlus0       float64   `gorm:"column:end_cards_european_handicap_draw_plus_0"`
		EndCardsEuropeanHandicapAwayPlus0       float64   `gorm:"column:end_cards_european_handicap_away_plus_0"`
		EndCardsEuropeanHandicapHomePlus1       float64   `gorm:"column:end_cards_european_handicap_home_plus_1"`
		EndCardsEuropeanHandicapDrawPlus1       float64   `gorm:"column:end_cards_european_handicap_draw_plus_1"`
		EndCardsEuropeanHandicapAwayPlus1       float64   `gorm:"column:end_cards_european_handicap_away_plus_1"`
		EndCardsEuropeanHandicapHomePlus2       float64   `gorm:"column:end_cards_european_handicap_home_plus_2"`
		EndCardsEuropeanHandicapDrawPlus2       float64   `gorm:"column:end_cards_european_handicap_draw_plus_2"`
		EndCardsEuropeanHandicapAwayPlus2       float64   `gorm:"column:end_cards_european_handicap_away_plus_2"`
		PreRcardYes                             float64   `gorm:"column:pre_rcard_yes"`
		PreRcardNo                              float64   `gorm:"column:pre_rcard_no"`
		EndRcardYes                             float64   `gorm:"column:end_rcard_yes"`
		EndRcardNo                              float64   `gorm:"column:end_rcard_no"`
		PreTeamToScoreFirstHome                 float64   `gorm:"column:pre_team_to_score_first_home"`
		PreTeamToScoreFirstDraw                 float64   `gorm:"column:pre_team_to_score_first_draw"`
		PreTeamToScoreFirstAway                 float64   `gorm:"column:pre_team_to_score_first_away"`
		EndTeamToScoreFirstHome                 float64   `gorm:"column:end_team_to_score_first_home"`
		EndTeamToScoreFirstDraw                 float64   `gorm:"column:end_team_to_score_first_draw"`
		EndTeamToScoreFirstAway                 float64   `gorm:"column:end_team_to_score_first_away"`
		PreTeamToScoreLastHome                  float64   `gorm:"column:pre_team_to_score_last_home"`
		PreTeamToScoreLastDraw                  float64   `gorm:"column:pre_team_to_score_last_draw"`
		PreTeamToScoreLastAway                  float64   `gorm:"column:pre_team_to_score_last_away"`
		EndTeamToScoreLastHome                  float64   `gorm:"column:end_team_to_score_last_home"`
		EndTeamToScoreLastDraw                  float64   `gorm:"column:end_team_to_score_last_draw"`
		EndTeamToScoreLastAway                  float64   `gorm:"column:end_team_to_score_last_away"`
		PreTotalHomeOver15                      float64   `gorm:"column:pre_total_home_over_15"`
		PreTotalHomeUnder15                     float64   `gorm:"column:pre_total_home_under_15"`
		PreTotalHomeOver25                      float64   `gorm:"column:pre_total_home_over_25"`
		PreTotalHomeUnder25                     float64   `gorm:"column:pre_total_home_under_25"`
		PreTotalHomeOver35                      float64   `gorm:"column:pre_total_home_over_35"`
		PreTotalHomeUnder35                     float64   `gorm:"column:pre_total_home_under_35"`
		PreTotalHomeOver45                      float64   `gorm:"column:pre_total_home_over_45"`
		PreTotalHomeUnder45                     float64   `gorm:"column:pre_total_home_under_45"`
		PreTotalHomeOver55                      float64   `gorm:"column:pre_total_home_over_55"`
		PreTotalHomeUnder55                     float64   `gorm:"column:pre_total_home_under_55"`
		PreTotalHomeOver65                      float64   `gorm:"column:pre_total_home_over_65"`
		PreTotalHomeUnder65                     float64   `gorm:"column:pre_total_home_under_65"`
		EndTotalHomeOver15                      float64   `gorm:"column:end_total_home_over_15"`
		EndTotalHomeUnder15                     float64   `gorm:"column:end_total_home_under_15"`
		EndTotalHomeOver25                      float64   `gorm:"column:end_total_home_over_25"`
		EndTotalHomeUnder25                     float64   `gorm:"column:end_total_home_under_25"`
		EndTotalHomeOver35                      float64   `gorm:"column:end_total_home_over_35"`
		EndTotalHomeUnder35                     float64   `gorm:"column:end_total_home_under_35"`
		EndTotalHomeOver45                      float64   `gorm:"column:end_total_home_over_45"`
		EndTotalHomeUnder45                     float64   `gorm:"column:end_total_home_under_45"`
		EndTotalHomeOver55                      float64   `gorm:"column:end_total_home_over_55"`
		EndTotalHomeUnder55                     float64   `gorm:"column:end_total_home_under_55"`
		EndTotalHomeOver65                      float64   `gorm:"column:end_total_home_over_65"`
		EndTotalHomeUnder65                     float64   `gorm:"column:end_total_home_under_65"`
		PreTotalAwayOver15                      float64   `gorm:"column:pre_total_away_over_15"`
		PreTotalAwayUnder15                     float64   `gorm:"column:pre_total_away_under_15"`
		PreTotalAwayOver25                      float64   `gorm:"column:pre_total_away_over_25"`
		PreTotalAwayUnder25                     float64   `gorm:"column:pre_total_away_under_25"`
		PreTotalAwayOver35                      float64   `gorm:"column:pre_total_away_over_35"`
		PreTotalAwayUnder35                     float64   `gorm:"column:pre_total_away_under_35"`
		PreTotalAwayOver45                      float64   `gorm:"column:pre_total_away_over_45"`
		PreTotalAwayUnder45                     float64   `gorm:"column:pre_total_away_under_45"`
		PreTotalAwayOver55                      float64   `gorm:"column:pre_total_away_over_55"`
		PreTotalAwayUnder55                     float64   `gorm:"column:pre_total_away_under_55"`
		PreTotalAwayOver65                      float64   `gorm:"column:pre_total_away_over_65"`
		PreTotalAwayUnder65                     float64   `gorm:"column:pre_total_away_under_65"`
		EndTotalAwayOver15                      float64   `gorm:"column:end_total_away_over_15"`
		EndTotalAwayUnder15                     float64   `gorm:"column:end_total_away_under_15"`
		EndTotalAwayOver25                      float64   `gorm:"column:end_total_away_over_25"`
		EndTotalAwayUnder25                     float64   `gorm:"column:end_total_away_under_25"`
		EndTotalAwayOver35                      float64   `gorm:"column:end_total_away_over_35"`
		EndTotalAwayUnder35                     float64   `gorm:"column:end_total_away_under_35"`
		EndTotalAwayOver45                      float64   `gorm:"column:end_total_away_over_45"`
		EndTotalAwayUnder45                     float64   `gorm:"column:end_total_away_under_45"`
		EndTotalAwayOver55                      float64   `gorm:"column:end_total_away_over_55"`
		EndTotalAwayUnder55                     float64   `gorm:"column:end_total_away_under_55"`
		EndTotalAwayOver65                      float64   `gorm:"column:end_total_away_over_65"`
		EndTotalAwayUnder65                     float64   `gorm:"column:end_total_away_under_65"`
		PreCornersAsianHandicapHomeMin65        float64   `gorm:"column:pre_corners_asian_handicap_home_min_65"`
		PreCornersAsianHandicapAwayMin65        float64   `gorm:"column:pre_corners_asian_handicap_away_min_65"`
		PreCornersAsianHandicapHomeMin6         float64   `gorm:"column:pre_corners_asian_handicap_home_min_6"`
		PreCornersAsianHandicapAwayMin6         float64   `gorm:"column:pre_corners_asian_handicap_away_min_6"`
		PreCornersAsianHandicapHomeMin55        float64   `gorm:"column:pre_corners_asian_handicap_home_min_55"`
		PreCornersAsianHandicapAwayMin55        float64   `gorm:"column:pre_corners_asian_handicap_away_min_55"`
		PreCornersAsianHandicapHomeMin5         float64   `gorm:"column:pre_corners_asian_handicap_home_min_5"`
		PreCornersAsianHandicapAwayMin5         float64   `gorm:"column:pre_corners_asian_handicap_away_min_5"`
		PreCornersAsianHandicapHomeMin45        float64   `gorm:"column:pre_corners_asian_handicap_home_min_45"`
		PreCornersAsianHandicapAwayMin45        float64   `gorm:"column:pre_corners_asian_handicap_away_min_45"`
		PreCornersAsianHandicapHomeMin4         float64   `gorm:"column:pre_corners_asian_handicap_home_min_4"`
		PreCornersAsianHandicapAwayMin4         float64   `gorm:"column:pre_corners_asian_handicap_away_min_4"`
		PreCornersAsianHandicapHomeMin35        float64   `gorm:"column:pre_corners_asian_handicap_home_min_35"`
		PreCornersAsianHandicapAwayMin35        float64   `gorm:"column:pre_corners_asian_handicap_away_min_35"`
		PreCornersAsianHandicapHomeMin3         float64   `gorm:"column:pre_corners_asian_handicap_home_min_3"`
		PreCornersAsianHandicapAwayMin3         float64   `gorm:"column:pre_corners_asian_handicap_away_min_3"`
		PreCornersAsianHandicapHomeMin25        float64   `gorm:"column:pre_corners_asian_handicap_home_min_25"`
		PreCornersAsianHandicapAwayMin25        float64   `gorm:"column:pre_corners_asian_handicap_away_min_25"`
		PreCornersAsianHandicapHomeMin2         float64   `gorm:"column:pre_corners_asian_handicap_home_min_2"`
		PreCornersAsianHandicapAwayMin2         float64   `gorm:"column:pre_corners_asian_handicap_away_min_2"`
		PreCornersAsianHandicapHomeMin15        float64   `gorm:"column:pre_corners_asian_handicap_home_min_15"`
		PreCornersAsianHandicapAwayMin15        float64   `gorm:"column:pre_corners_asian_handicap_away_min_15"`
		PreCornersAsianHandicapHomeMin1         float64   `gorm:"column:pre_corners_asian_handicap_home_min_1"`
		PreCornersAsianHandicapHomeMin05        float64   `gorm:"column:pre_corners_asian_handicap_home_min_05"`
		PreCornersAsianHandicapAwayMin1         float64   `gorm:"column:pre_corners_asian_handicap_away_min_1"`
		PreCornersAsianHandicapAwayMin05        float64   `gorm:"column:pre_corners_asian_handicap_away_min_05"`
		PreCornersAsianHandicapHomePlus0        float64   `gorm:"column:pre_corners_asian_handicap_home_plus_0"`
		PreCornersAsianHandicapAwayPlus0        float64   `gorm:"column:pre_corners_asian_handicap_away_plus_0"`
		PreCornersAsianHandicapHomePlus05       float64   `gorm:"column:pre_corners_asian_handicap_home_plus_05"`
		PreCornersAsianHandicapAwayPlus05       float64   `gorm:"column:pre_corners_asian_handicap_away_plus_05"`
		PreCornersAsianHandicapHomePlus1        float64   `gorm:"column:pre_corners_asian_handicap_home_plus_1"`
		PreCornersAsianHandicapAwayPlus1        float64   `gorm:"column:pre_corners_asian_handicap_away_plus_1"`
		PreCornersAsianHandicapHomePlus15       float64   `gorm:"column:pre_corners_asian_handicap_home_plus_15"`
		PreCornersAsianHandicapAwayPlus15       float64   `gorm:"column:pre_corners_asian_handicap_away_plus_15"`
		PreCornersAsianHandicapHomePlus2        float64   `gorm:"column:pre_corners_asian_handicap_home_plus_2"`
		PreCornersAsianHandicapAwayPlus2        float64   `gorm:"column:pre_corners_asian_handicap_away_plus_2"`
		PreCornersAsianHandicapHomePlus25       float64   `gorm:"column:pre_corners_asian_handicap_home_plus_25"`
		PreCornersAsianHandicapAwayPlus25       float64   `gorm:"column:pre_corners_asian_handicap_away_plus_25"`
		PreCornersAsianHandicapHomePlus3        float64   `gorm:"column:pre_corners_asian_handicap_home_plus_3"`
		PreCornersAsianHandicapAwayPlus3        float64   `gorm:"column:pre_corners_asian_handicap_away_plus_3"`
		PreCornersAsianHandicapHomePlus35       float64   `gorm:"column:pre_corners_asian_handicap_home_plus_35"`
		PreCornersAsianHandicapAwayPlus35       float64   `gorm:"column:pre_corners_asian_handicap_away_plus_35"`
		PreCornersAsianHandicapHomePlus4        float64   `gorm:"column:pre_corners_asian_handicap_home_plus_4"`
		PreCornersAsianHandicapAwayPlus4        float64   `gorm:"column:pre_corners_asian_handicap_away_plus_4"`
		PreCornersAsianHandicapHomePlus45       float64   `gorm:"column:pre_corners_asian_handicap_home_plus_45"`
		PreCornersAsianHandicapAwayPlus45       float64   `gorm:"column:pre_corners_asian_handicap_away_plus_45"`
		PreCornersAsianHandicapHomePlus5        float64   `gorm:"column:pre_corners_asian_handicap_home_plus_5"`
		PreCornersAsianHandicapAwayPlus5        float64   `gorm:"column:pre_corners_asian_handicap_away_plus_5"`
		PreCornersAsianHandicapHomePlus55       float64   `gorm:"column:pre_corners_asian_handicap_home_plus_55"`
		PreCornersAsianHandicapAwayPlus55       float64   `gorm:"column:pre_corners_asian_handicap_away_plus_55"`
		PreCornersAsianHandicapHomePlus6        float64   `gorm:"column:pre_corners_asian_handicap_home_plus_6"`
		PreCornersAsianHandicapAwayPlus6        float64   `gorm:"column:pre_corners_asian_handicap_away_plus_6"`
		PreCornersAsianHandicapHomePlus65       float64   `gorm:"column:pre_corners_asian_handicap_home_plus_65"`
		PreCornersAsianHandicapAwayPlus65       float64   `gorm:"column:pre_corners_asian_handicap_away_plus_65"`
		EndCornersAsianHandicapHomeMin65        float64   `gorm:"column:end_corners_asian_handicap_home_min_65"`
		EndCornersAsianHandicapAwayMin65        float64   `gorm:"column:end_corners_asian_handicap_away_min_65"`
		EndCornersAsianHandicapHomeMin6         float64   `gorm:"column:end_corners_asian_handicap_home_min_6"`
		EndCornersAsianHandicapAwayMin6         float64   `gorm:"column:end_corners_asian_handicap_away_min_6"`
		EndCornersAsianHandicapHomeMin55        float64   `gorm:"column:end_corners_asian_handicap_home_min_55"`
		EndCornersAsianHandicapAwayMin55        float64   `gorm:"column:end_corners_asian_handicap_away_min_55"`
		EndCornersAsianHandicapHomeMin5         float64   `gorm:"column:end_corners_asian_handicap_home_min_5"`
		EndCornersAsianHandicapAwayMin5         float64   `gorm:"column:end_corners_asian_handicap_away_min_5"`
		EndCornersAsianHandicapHomeMin45        float64   `gorm:"column:end_corners_asian_handicap_home_min_45"`
		EndCornersAsianHandicapAwayMin45        float64   `gorm:"column:end_corners_asian_handicap_away_min_45"`
		EndCornersAsianHandicapHomeMin4         float64   `gorm:"column:end_corners_asian_handicap_home_min_4"`
		EndCornersAsianHandicapAwayMin4         float64   `gorm:"column:end_corners_asian_handicap_away_min_4"`
		EndCornersAsianHandicapHomeMin35        float64   `gorm:"column:end_corners_asian_handicap_home_min_35"`
		EndCornersAsianHandicapAwayMin35        float64   `gorm:"column:end_corners_asian_handicap_away_min_35"`
		EndCornersAsianHandicapHomeMin3         float64   `gorm:"column:end_corners_asian_handicap_home_min_3"`
		EndCornersAsianHandicapAwayMin3         float64   `gorm:"column:end_corners_asian_handicap_away_min_3"`
		EndCornersAsianHandicapHomeMin25        float64   `gorm:"column:end_corners_asian_handicap_home_min_25"`
		EndCornersAsianHandicapAwayMin25        float64   `gorm:"column:end_corners_asian_handicap_away_min_25"`
		EndCornersAsianHandicapHomeMin2         float64   `gorm:"column:end_corners_asian_handicap_home_min_2"`
		EndCornersAsianHandicapAwayMin2         float64   `gorm:"column:end_corners_asian_handicap_away_min_2"`
		EndCornersAsianHandicapHomeMin15        float64   `gorm:"column:end_corners_asian_handicap_home_min_15"`
		EndCornersAsianHandicapAwayMin15        float64   `gorm:"column:end_corners_asian_handicap_away_min_15"`
		EndCornersAsianHandicapHomeMin1         float64   `gorm:"column:end_corners_asian_handicap_home_min_1"`
		EndCornersAsianHandicapHomeMin05        float64   `gorm:"column:end_corners_asian_handicap_home_min_05"`
		EndCornersAsianHandicapAwayMin1         float64   `gorm:"column:end_corners_asian_handicap_away_min_1"`
		EndCornersAsianHandicapAwayMin05        float64   `gorm:"column:end_corners_asian_handicap_away_min_05"`
		EndCornersAsianHandicapHomePlus0        float64   `gorm:"column:end_corners_asian_handicap_home_plus_0"`
		EndCornersAsianHandicapAwayPlus0        float64   `gorm:"column:end_corners_asian_handicap_away_plus_0"`
		EndCornersAsianHandicapHomePlus05       float64   `gorm:"column:end_corners_asian_handicap_home_plus_05"`
		EndCornersAsianHandicapAwayPlus05       float64   `gorm:"column:end_corners_asian_handicap_away_plus_05"`
		EndCornersAsianHandicapHomePlus1        float64   `gorm:"column:end_corners_asian_handicap_home_plus_1"`
		EndCornersAsianHandicapAwayPlus1        float64   `gorm:"column:end_corners_asian_handicap_away_plus_1"`
		EndCornersAsianHandicapHomePlus15       float64   `gorm:"column:end_corners_asian_handicap_home_plus_15"`
		EndCornersAsianHandicapAwayPlus15       float64   `gorm:"column:end_corners_asian_handicap_away_plus_15"`
		EndCornersAsianHandicapHomePlus2        float64   `gorm:"column:end_corners_asian_handicap_home_plus_2"`
		EndCornersAsianHandicapAwayPlus2        float64   `gorm:"column:end_corners_asian_handicap_away_plus_2"`
		EndCornersAsianHandicapHomePlus25       float64   `gorm:"column:end_corners_asian_handicap_home_plus_25"`
		EndCornersAsianHandicapAwayPlus25       float64   `gorm:"column:end_corners_asian_handicap_away_plus_25"`
		EndCornersAsianHandicapHomePlus3        float64   `gorm:"column:end_corners_asian_handicap_home_plus_3"`
		EndCornersAsianHandicapAwayPlus3        float64   `gorm:"column:end_corners_asian_handicap_away_plus_3"`
		EndCornersAsianHandicapHomePlus35       float64   `gorm:"column:end_corners_asian_handicap_home_plus_35"`
		EndCornersAsianHandicapAwayPlus35       float64   `gorm:"column:end_corners_asian_handicap_away_plus_35"`
		EndCornersAsianHandicapHomePlus4        float64   `gorm:"column:end_corners_asian_handicap_home_plus_4"`
		EndCornersAsianHandicapAwayPlus4        float64   `gorm:"column:end_corners_asian_handicap_away_plus_4"`
		EndCornersAsianHandicapHomePlus45       float64   `gorm:"column:end_corners_asian_handicap_home_plus_45"`
		EndCornersAsianHandicapAwayPlus45       float64   `gorm:"column:end_corners_asian_handicap_away_plus_45"`
		EndCornersAsianHandicapHomePlus5        float64   `gorm:"column:end_corners_asian_handicap_home_plus_5"`
		EndCornersAsianHandicapAwayPlus5        float64   `gorm:"column:end_corners_asian_handicap_away_plus_5"`
		EndCornersAsianHandicapHomePlus55       float64   `gorm:"column:end_corners_asian_handicap_home_plus_55"`
		EndCornersAsianHandicapAwayPlus55       float64   `gorm:"column:end_corners_asian_handicap_away_plus_55"`
		EndCornersAsianHandicapHomePlus6        float64   `gorm:"column:end_corners_asian_handicap_home_plus_6"`
		EndCornersAsianHandicapAwayPlus6        float64   `gorm:"column:end_corners_asian_handicap_away_plus_6"`
		EndCornersAsianHandicapHomePlus65       float64   `gorm:"column:end_corners_asian_handicap_home_plus_65"`
		EndCornersAsianHandicapAwayPlus65       float64   `gorm:"column:end_corners_asian_handicap_away_plus_65"`
		PreCornersOverUnderOver700              float64   `gorm:"column:pre_corners_over_under_over_700"`
		PreCornersOverUnderUnder700             float64   `gorm:"column:pre_corners_over_under_under_700"`
		PreCornersOverUnderOver725              float64   `gorm:"column:pre_corners_over_under_over_725"`
		PreCornersOverUnderUnder725             float64   `gorm:"column:pre_corners_over_under_under_725"`
		PreCornersOverUnderOver750              float64   `gorm:"column:pre_corners_over_under_over_750"`
		PreCornersOverUnderUnder750             float64   `gorm:"column:pre_corners_over_under_under_750"`
		PreCornersOverUnderOver775              float64   `gorm:"column:pre_corners_over_under_over_775"`
		PreCornersOverUnderUnder775             float64   `gorm:"column:pre_corners_over_under_under_775"`
		PreCornersOverUnderOver800              float64   `gorm:"column:pre_corners_over_under_over_800"`
		PreCornersOverUnderUnder800             float64   `gorm:"column:pre_corners_over_under_under_800"`
		PreCornersOverUnderOver825              float64   `gorm:"column:pre_corners_over_under_over_825"`
		PreCornersOverUnderUnder825             float64   `gorm:"column:pre_corners_over_under_under_825"`
		PreCornersOverUnderOver850              float64   `gorm:"column:pre_corners_over_under_over_850"`
		PreCornersOverUnderUnder850             float64   `gorm:"column:pre_corners_over_under_under_850"`
		PreCornersOverUnderOver875              float64   `gorm:"column:pre_corners_over_under_over_875"`
		PreCornersOverUnderUnder875             float64   `gorm:"column:pre_corners_over_under_under_875"`
		PreCornersOverUnderOver900              float64   `gorm:"column:pre_corners_over_under_over_900"`
		PreCornersOverUnderUnder900             float64   `gorm:"column:pre_corners_over_under_under_900"`
		PreCornersOverUnderOver925              float64   `gorm:"column:pre_corners_over_under_over_925"`
		PreCornersOverUnderUnder925             float64   `gorm:"column:pre_corners_over_under_under_925"`
		PreCornersOverUnderOver950              float64   `gorm:"column:pre_corners_over_under_over_950"`
		PreCornersOverUnderUnder950             float64   `gorm:"column:pre_corners_over_under_under_950"`
		PreCornersOverUnderOver975              float64   `gorm:"column:pre_corners_over_under_over_975"`
		PreCornersOverUnderUnder975             float64   `gorm:"column:pre_corners_over_under_under_975"`
		PreCornersOverUnderOver1000             float64   `gorm:"column:pre_corners_over_under_over_1000"`
		PreCornersOverUnderUnder1000            float64   `gorm:"column:pre_corners_over_under_under_1000"`
		PreCornersOverUnderOver1025             float64   `gorm:"column:pre_corners_over_under_over_1025"`
		PreCornersOverUnderUnder1025            float64   `gorm:"column:pre_corners_over_under_under_1025"`
		PreCornersOverUnderOver1050             float64   `gorm:"column:pre_corners_over_under_over_1050"`
		PreCornersOverUnderUnder1050            float64   `gorm:"column:pre_corners_over_under_under_1050"`
		PreCornersOverUnderOver1075             float64   `gorm:"column:pre_corners_over_under_over_1075"`
		PreCornersOverUnderUnder1075            float64   `gorm:"column:pre_corners_over_under_under_1075"`
		PreCornersOverUnderOver1100             float64   `gorm:"column:pre_corners_over_under_over_1100"`
		PreCornersOverUnderUnder1100            float64   `gorm:"column:pre_corners_over_under_under_1100"`
		PreCornersOverUnderOver1125             float64   `gorm:"column:pre_corners_over_under_over_1125"`
		PreCornersOverUnderUnder1125            float64   `gorm:"column:pre_corners_over_under_under_1125"`
		PreCornersOverUnderOver1150             float64   `gorm:"column:pre_corners_over_under_over_1150"`
		PreCornersOverUnderUnder1150            float64   `gorm:"column:pre_corners_over_under_under_1150"`
		PreCornersOverUnderOver1175             float64   `gorm:"column:pre_corners_over_under_over_1175"`
		PreCornersOverUnderUnder1175            float64   `gorm:"column:pre_corners_over_under_under_1175"`
		PreCornersOverUnderOver1200             float64   `gorm:"column:pre_corners_over_under_over_1200"`
		PreCornersOverUnderUnder1200            float64   `gorm:"column:pre_corners_over_under_under_1200"`
		PreCornersOverUnderOver1225             float64   `gorm:"column:pre_corners_over_under_over_1225"`
		PreCornersOverUnderUnder1225            float64   `gorm:"column:pre_corners_over_under_under_1225"`
		PreCornersOverUnderOver1250             float64   `gorm:"column:pre_corners_over_under_over_1250"`
		PreCornersOverUnderUnder1250            float64   `gorm:"column:pre_corners_over_under_under_1250"`
		PreCornersOverUnderOver1275             float64   `gorm:"column:pre_corners_over_under_over_1275"`
		PreCornersOverUnderUnder1275            float64   `gorm:"column:pre_corners_over_under_under_1275"`
		EndCornersOverUnderOver700              float64   `gorm:"column:end_corners_over_under_over_700"`
		EndCornersOverUnderUnder700             float64   `gorm:"column:end_corners_over_under_under_700"`
		EndCornersOverUnderOver725              float64   `gorm:"column:end_corners_over_under_over_725"`
		EndCornersOverUnderUnder725             float64   `gorm:"column:end_corners_over_under_under_725"`
		EndCornersOverUnderOver750              float64   `gorm:"column:end_corners_over_under_over_750"`
		EndCornersOverUnderUnder750             float64   `gorm:"column:end_corners_over_under_under_750"`
		EndCornersOverUnderOver775              float64   `gorm:"column:end_corners_over_under_over_775"`
		EndCornersOverUnderUnder775             float64   `gorm:"column:end_corners_over_under_under_775"`
		EndCornersOverUnderOver800              float64   `gorm:"column:end_corners_over_under_over_800"`
		EndCornersOverUnderUnder800             float64   `gorm:"column:end_corners_over_under_under_800"`
		EndCornersOverUnderOver825              float64   `gorm:"column:end_corners_over_under_over_825"`
		EndCornersOverUnderUnder825             float64   `gorm:"column:end_corners_over_under_under_825"`
		EndCornersOverUnderOver850              float64   `gorm:"column:end_corners_over_under_over_850"`
		EndCornersOverUnderUnder850             float64   `gorm:"column:end_corners_over_under_under_850"`
		EndCornersOverUnderOver875              float64   `gorm:"column:end_corners_over_under_over_875"`
		EndCornersOverUnderUnder875             float64   `gorm:"column:end_corners_over_under_under_875"`
		EndCornersOverUnderOver900              float64   `gorm:"column:end_corners_over_under_over_900"`
		EndCornersOverUnderUnder900             float64   `gorm:"column:end_corners_over_under_under_900"`
		EndCornersOverUnderOver925              float64   `gorm:"column:end_corners_over_under_over_925"`
		EndCornersOverUnderUnder925             float64   `gorm:"column:end_corners_over_under_under_925"`
		EndCornersOverUnderOver950              float64   `gorm:"column:end_corners_over_under_over_950"`
		EndCornersOverUnderUnder950             float64   `gorm:"column:end_corners_over_under_under_950"`
		EndCornersOverUnderOver975              float64   `gorm:"column:end_corners_over_under_over_975"`
		EndCornersOverUnderUnder975             float64   `gorm:"column:end_corners_over_under_under_975"`
		EndCornersOverUnderOver1000             float64   `gorm:"column:end_corners_over_under_over_1000"`
		EndCornersOverUnderUnder1000            float64   `gorm:"column:end_corners_over_under_under_1000"`
		EndCornersOverUnderOver1025             float64   `gorm:"column:end_corners_over_under_over_1025"`
		EndCornersOverUnderUnder1025            float64   `gorm:"column:end_corners_over_under_under_1025"`
		EndCornersOverUnderOver1050             float64   `gorm:"column:end_corners_over_under_over_1050"`
		EndCornersOverUnderUnder1050            float64   `gorm:"column:end_corners_over_under_under_1050"`
		EndCornersOverUnderOver1075             float64   `gorm:"column:end_corners_over_under_over_1075"`
		EndCornersOverUnderUnder1075            float64   `gorm:"column:end_corners_over_under_under_1075"`
		EndCornersOverUnderOver1100             float64   `gorm:"column:end_corners_over_under_over_1100"`
		EndCornersOverUnderUnder1100            float64   `gorm:"column:end_corners_over_under_under_1100"`
		EndCornersOverUnderOver1125             float64   `gorm:"column:end_corners_over_under_over_1125"`
		EndCornersOverUnderUnder1125            float64   `gorm:"column:end_corners_over_under_under_1125"`
		EndCornersOverUnderOver1150             float64   `gorm:"column:end_corners_over_under_over_1150"`
		EndCornersOverUnderUnder1150            float64   `gorm:"column:end_corners_over_under_under_1150"`
		EndCornersOverUnderOver1175             float64   `gorm:"column:end_corners_over_under_over_1175"`
		EndCornersOverUnderUnder1175            float64   `gorm:"column:end_corners_over_under_under_1175"`
		EndCornersOverUnderOver1200             float64   `gorm:"column:end_corners_over_under_over_1200"`
		EndCornersOverUnderUnder1200            float64   `gorm:"column:end_corners_over_under_under_1200"`
		EndCornersOverUnderOver1225             float64   `gorm:"column:end_corners_over_under_over_1225"`
		EndCornersOverUnderUnder1225            float64   `gorm:"column:end_corners_over_under_under_1225"`
		EndCornersOverUnderOver1250             float64   `gorm:"column:end_corners_over_under_over_1250"`
		EndCornersOverUnderUnder1250            float64   `gorm:"column:end_corners_over_under_under_1250"`
		EndCornersOverUnderOver1275             float64   `gorm:"column:end_corners_over_under_over_1275"`
		EndCornersOverUnderUnder1275            float64   `gorm:"column:end_corners_over_under_under_1275"`
		PreHomeCornersOverunderOver25           float64   `gorm:"column:pre_home_corners_overunder_over_25"`
		PreHomeCornersOverunderUnder25          float64   `gorm:"column:pre_home_corners_overunder_under_25"`
		PreHomeCornersOverunderOver35           float64   `gorm:"column:pre_home_corners_overunder_over_35"`
		PreHomeCornersOverunderUnder35          float64   `gorm:"column:pre_home_corners_overunder_under_35"`
		PreHomeCornersOverunderOver45           float64   `gorm:"column:pre_home_corners_overunder_over_45"`
		PreHomeCornersOverunderUnder45          float64   `gorm:"column:pre_home_corners_overunder_under_45"`
		PreHomeCornersOverunderOver55           float64   `gorm:"column:pre_home_corners_overunder_over_55"`
		PreHomeCornersOverunderUnder55          float64   `gorm:"column:pre_home_corners_overunder_under_55"`
		PreHomeCornersOverunderOver65           float64   `gorm:"column:pre_home_corners_overunder_over_65"`
		PreHomeCornersOverunderUnder65          float64   `gorm:"column:pre_home_corners_overunder_under_65"`
		PreHomeCornersOverunderOver75           float64   `gorm:"column:pre_home_corners_overunder_over_75"`
		PreHomeCornersOverunderUnder75          float64   `gorm:"column:pre_home_corners_overunder_under_75"`
		PreHomeCornersOverunderOver85           float64   `gorm:"column:pre_home_corners_overunder_over_85"`
		PreHomeCornersOverunderUnder85          float64   `gorm:"column:pre_home_corners_overunder_under_85"`
		EndHomeCornersOverunderOver25           float64   `gorm:"column:end_home_corners_overunder_over_25"`
		EndHomeCornersOverunderUnder25          float64   `gorm:"column:end_home_corners_overunder_under_25"`
		EndHomeCornersOverunderOver35           float64   `gorm:"column:end_home_corners_overunder_over_35"`
		EndHomeCornersOverunderUnder35          float64   `gorm:"column:end_home_corners_overunder_under_35"`
		EndHomeCornersOverunderOver45           float64   `gorm:"column:end_home_corners_overunder_over_45"`
		EndHomeCornersOverunderUnder45          float64   `gorm:"column:end_home_corners_overunder_under_45"`
		EndHomeCornersOverunderOver55           float64   `gorm:"column:end_home_corners_overunder_over_55"`
		EndHomeCornersOverunderUnder55          float64   `gorm:"column:end_home_corners_overunder_under_55"`
		EndHomeCornersOverunderOver65           float64   `gorm:"column:end_home_corners_overunder_over_65"`
		EndHomeCornersOverunderUnder65          float64   `gorm:"column:end_home_corners_overunder_under_65"`
		EndHomeCornersOverunderOver75           float64   `gorm:"column:end_home_corners_overunder_over_75"`
		EndHomeCornersOverunderUnder75          float64   `gorm:"column:end_home_corners_overunder_under_75"`
		EndHomeCornersOverunderOver85           float64   `gorm:"column:end_home_corners_overunder_over_85"`
		EndHomeCornersOverunderUnder85          float64   `gorm:"column:end_home_corners_overunder_under_85"`
		PreAwayCornersOverunderOver25           float64   `gorm:"column:pre_away_corners_overunder_over_25"`
		PreAwayCornersOverunderUnder25          float64   `gorm:"column:pre_away_corners_overunder_under_25"`
		PreAwayCornersOverunderOver35           float64   `gorm:"column:pre_away_corners_overunder_over_35"`
		PreAwayCornersOverunderUnder35          float64   `gorm:"column:pre_away_corners_overunder_under_35"`
		PreAwayCornersOverunderOver45           float64   `gorm:"column:pre_away_corners_overunder_over_45"`
		PreAwayCornersOverunderUnder45          float64   `gorm:"column:pre_away_corners_overunder_under_45"`
		PreAwayCornersOverunderOver55           float64   `gorm:"column:pre_away_corners_overunder_over_55"`
		PreAwayCornersOverunderUnder55          float64   `gorm:"column:pre_away_corners_overunder_under_55"`
		PreAwayCornersOverunderOver65           float64   `gorm:"column:pre_away_corners_overunder_over_65"`
		PreAwayCornersOverunderUnder65          float64   `gorm:"column:pre_away_corners_overunder_under_65"`
		PreAwayCornersOverunderOver75           float64   `gorm:"column:pre_away_corners_overunder_over_75"`
		PreAwayCornersOverunderUnder75          float64   `gorm:"column:pre_away_corners_overunder_under_75"`
		PreAwayCornersOverunderOver85           float64   `gorm:"column:pre_away_corners_overunder_over_85"`
		PreAwayCornersOverunderUnder85          float64   `gorm:"column:pre_away_corners_overunder_under_85"`
		EndAwayCornersOverunderOver25           float64   `gorm:"column:end_away_corners_overunder_over_25"`
		EndAwayCornersOverunderUnder25          float64   `gorm:"column:end_away_corners_overunder_under_25"`
		EndAwayCornersOverunderOver35           float64   `gorm:"column:end_away_corners_overunder_over_35"`
		EndAwayCornersOverunderUnder35          float64   `gorm:"column:end_away_corners_overunder_under_35"`
		EndAwayCornersOverunderOver45           float64   `gorm:"column:end_away_corners_overunder_over_45"`
		EndAwayCornersOverunderUnder45          float64   `gorm:"column:end_away_corners_overunder_under_45"`
		EndAwayCornersOverunderOver55           float64   `gorm:"column:end_away_corners_overunder_over_55"`
		EndAwayCornersOverunderUnder55          float64   `gorm:"column:end_away_corners_overunder_under_55"`
		EndAwayCornersOverunderOver65           float64   `gorm:"column:end_away_corners_overunder_over_65"`
		EndAwayCornersOverunderUnder65          float64   `gorm:"column:end_away_corners_overunder_under_65"`
		EndAwayCornersOverunderOver75           float64   `gorm:"column:end_away_corners_overunder_over_75"`
		EndAwayCornersOverunderUnder75          float64   `gorm:"column:end_away_corners_overunder_under_75"`
		EndAwayCornersOverunderOver85           float64   `gorm:"column:end_away_corners_overunder_over_85"`
		EndAwayCornersOverunderUnder85          float64   `gorm:"column:end_away_corners_overunder_under_85"`
		PreFirstCornerHome                      float64   `gorm:"column:pre_first_corner_home"`
		PreFirstCornerAway                      float64   `gorm:"column:pre_first_corner_away"`
		EndFirstCornerHome                      float64   `gorm:"column:end_first_corner_home"`
		EndFirstCornerAway                      float64   `gorm:"column:end_first_corner_away"`
		PreLastCornerHome                       float64   `gorm:"column:pre_last_corner_home"`
		PreLastCornerAway                       float64   `gorm:"column:pre_last_corner_away"`
		EndLastCornerHome                       float64   `gorm:"column:end_last_corner_home"`
		EndLastCornerAway                       float64   `gorm:"column:end_last_corner_away"`
		PreCardsAsianHandicapHomeMin05          float64   `gorm:"column:pre_cards_asian_handicap_home_min_05"`
		PreCardsAsianHandicapAwayMin05          float64   `gorm:"column:pre_cards_asian_handicap_away_min_05"`
		PreCardsAsianHandicapHomePlus0          float64   `gorm:"column:pre_cards_asian_handicap_home_plus_0"`
		PreCardsAsianHandicapAwayPlus0          float64   `gorm:"column:pre_cards_asian_handicap_away_plus_0"`
		PreCardsAsianHandicapHomePlus05         float64   `gorm:"column:pre_cards_asian_handicap_home_plus_05"`
		PreCardsAsianHandicapAwayPlus05         float64   `gorm:"column:pre_cards_asian_handicap_away_plus_05"`
		EndCardsAsianHandicapHomeMin05          float64   `gorm:"column:end_cards_asian_handicap_home_min_05"`
		EndCardsAsianHandicapAwayMin05          float64   `gorm:"column:end_cards_asian_handicap_away_min_05"`
		EndCardsAsianHandicapHomePlus0          float64   `gorm:"column:end_cards_asian_handicap_home_plus_0"`
		EndCardsAsianHandicapAwayPlus0          float64   `gorm:"column:end_cards_asian_handicap_away_plus_0"`
		EndCardsAsianHandicapHomePlus05         float64   `gorm:"column:end_cards_asian_handicap_home_plus_05"`
		EndCardsAsianHandicapAwayPlus05         float64   `gorm:"column:end_cards_asian_handicap_away_plus_05"`
		PreCardsOverunderOver25                 float64   `gorm:"column:pre_cards_overunder_over_25"`
		PreCardsOverunderUnder25                float64   `gorm:"column:pre_cards_overunder_under_25"`
		PreCardsOverunderOver35                 float64   `gorm:"column:pre_cards_overunder_over_35"`
		PreCardsOverunderUnder35                float64   `gorm:"column:pre_cards_overunder_under_35"`
		PreCardsOverunderOver45                 float64   `gorm:"column:pre_cards_overunder_over_45"`
		PreCardsOverunderUnder45                float64   `gorm:"column:pre_cards_overunder_under_45"`
		PreCardsOverunderOver55                 float64   `gorm:"column:pre_cards_overunder_over_55"`
		PreCardsOverunderUnder55                float64   `gorm:"column:pre_cards_overunder_under_55"`
		PreCardsOverunderOver65                 float64   `gorm:"column:pre_cards_overunder_over_65"`
		PreCardsOverunderUnder65                float64   `gorm:"column:pre_cards_overunder_under_65"`
		PreCardsOverunderOver75                 float64   `gorm:"column:pre_cards_overunder_over_75"`
		PreCardsOverunderUnder75                float64   `gorm:"column:pre_cards_overunder_under_75"`
		EndCardsOverunderOver25                 float64   `gorm:"column:end_cards_overunder_over_25"`
		EndCardsOverunderUnder25                float64   `gorm:"column:end_cards_overunder_under_25"`
		EndCardsOverunderOver35                 float64   `gorm:"column:end_cards_overunder_over_35"`
		EndCardsOverunderUnder35                float64   `gorm:"column:end_cards_overunder_under_35"`
		EndCardsOverunderOver45                 float64   `gorm:"column:end_cards_overunder_over_45"`
		EndCardsOverunderUnder45                float64   `gorm:"column:end_cards_overunder_under_45"`
		EndCardsOverunderOver55                 float64   `gorm:"column:end_cards_overunder_over_55"`
		EndCardsOverunderUnder55                float64   `gorm:"column:end_cards_overunder_under_55"`
		EndCardsOverunderOver65                 float64   `gorm:"column:end_cards_overunder_over_65"`
		EndCardsOverunderUnder65                float64   `gorm:"column:end_cards_overunder_under_65"`
		EndCardsOverunderOver75                 float64   `gorm:"column:end_cards_overunder_over_75"`
		EndCardsOverunderUnder75                float64   `gorm:"column:end_cards_overunder_under_75"`
		PreHomeTeamTotalCardsOver15             float64   `gorm:"column:pre_home_team_total_cards_over_15"`
		PreHomeTeamTotalCardsUnder15            float64   `gorm:"column:pre_home_team_total_cards_under_15"`
		PreHomeTeamTotalCardsOver25             float64   `gorm:"column:pre_home_team_total_cards_over_25"`
		PreHomeTeamTotalCardsUnder25            float64   `gorm:"column:pre_home_team_total_cards_under_25"`
		PreHomeTeamTotalCardsOver35             float64   `gorm:"column:pre_home_team_total_cards_over_35"`
		PreHomeTeamTotalCardsUnder35            float64   `gorm:"column:pre_home_team_total_cards_under_35"`
		EndHomeTeamTotalCardsOver15             float64   `gorm:"column:end_home_team_total_cards_over_15"`
		EndHomeTeamTotalCardsUnder15            float64   `gorm:"column:end_home_team_total_cards_under_15"`
		EndHomeTeamTotalCardsOver25             float64   `gorm:"column:end_home_team_total_cards_over_25"`
		EndHomeTeamTotalCardsUnder25            float64   `gorm:"column:end_home_team_total_cards_under_25"`
		EndHomeTeamTotalCardsOver35             float64   `gorm:"column:end_home_team_total_cards_over_35"`
		EndHomeTeamTotalCardsUnder35            float64   `gorm:"column:end_home_team_total_cards_under_35"`
		PreAwayTeamTotalCardsOver15             float64   `gorm:"column:pre_away_team_total_cards_over_15"`
		PreAwayTeamTotalCardsUnder15            float64   `gorm:"column:pre_away_team_total_cards_under_15"`
		PreAwayTeamTotalCardsOver25             float64   `gorm:"column:pre_away_team_total_cards_over_25"`
		PreAwayTeamTotalCardsUnder25            float64   `gorm:"column:pre_away_team_total_cards_under_25"`
		PreAwayTeamTotalCardsOver35             float64   `gorm:"column:pre_away_team_total_cards_over_35"`
		PreAwayTeamTotalCardsUnder35            float64   `gorm:"column:pre_away_team_total_cards_under_35"`
		EndAwayTeamTotalCardsOver15             float64   `gorm:"column:end_away_team_total_cards_over_15"`
		EndAwayTeamTotalCardsUnder15            float64   `gorm:"column:end_away_team_total_cards_under_15"`
		EndAwayTeamTotalCardsOver25             float64   `gorm:"column:end_away_team_total_cards_over_25"`
		EndAwayTeamTotalCardsUnder25            float64   `gorm:"column:end_away_team_total_cards_under_25"`
		EndAwayTeamTotalCardsOver35             float64   `gorm:"column:end_away_team_total_cards_over_35"`
		EndAwayTeamTotalCardsUnder35            float64   `gorm:"column:end_away_team_total_cards_under_35"`
		PreTotalGoalsUnder2                     float64   `gorm:"column:pre_total_goals_under_2"`
		PreTotalGoals2Or3                       float64   `gorm:"column:pre_total_goals_2_or_3"`
		PreTotalGoalsOver3                      float64   `gorm:"column:pre_total_goals_over_3"`
		EndTotalGoalsUnder2                     float64   `gorm:"column:end_total_goals_under_2"`
		EndTotalGoals2Or3                       float64   `gorm:"column:end_total_goals_2_or_3"`
		EndTotalGoalsOver3                      float64   `gorm:"column:end_total_goals_over_3"`
		PreFirstGoalMethodShot                  float64   `gorm:"column:pre_first_goal_method_shot"`
		PreFirstGoalMethodHeader                float64   `gorm:"column:pre_first_goal_method_header"`
		PreFirstGoalMethodPenalty               float64   `gorm:"column:pre_first_goal_method_penalty"`
		PreFirstGoalMethodFreekick              float64   `gorm:"column:pre_first_goal_method_freekick"`
		PreFirstGoalMethodOwngoal               float64   `gorm:"column:pre_first_goal_method_owngoal"`
		PreFirstGoalMethodDraw                  float64   `gorm:"column:pre_first_goal_method_draw"`
		EndFirstGoalMethodShot                  float64   `gorm:"column:end_first_goal_method_shot"`
		EndFirstGoalMethodHeader                float64   `gorm:"column:end_first_goal_method_header"`
		EndFirstGoalMethodPenalty               float64   `gorm:"column:end_first_goal_method_penalty"`
		EndFirstGoalMethodFreekick              float64   `gorm:"column:end_first_goal_method_freekick"`
		EndFirstGoalMethodOwngoal               float64   `gorm:"column:end_first_goal_method_owngoal"`
		EndFirstGoalMethodDraw                  float64   `gorm:"column:end_first_goal_method_draw"`
		PreToScoreAPenaltyHome                  float64   `gorm:"column:pre_to_score_a_penalty_home"`
		PreToScoreAPenaltyAway                  float64   `gorm:"column:pre_to_score_a_penalty_away"`
		EndToScoreAPenaltyHome                  float64   `gorm:"column:end_to_score_a_penalty_home"`
		EndToScoreAPenaltyAway                  float64   `gorm:"column:end_to_score_a_penalty_away"`
		PreToMissAPenaltyHome                   float64   `gorm:"column:pre_to_miss_a_penalty_home"`
		PreToMissAPenaltyAway                   float64   `gorm:"column:pre_to_miss_a_penalty_away"`
		EndToMissAPenaltyHome                   float64   `gorm:"column:end_to_miss_a_penalty_home"`
		EndToMissAPenaltyAway                   float64   `gorm:"column:end_to_miss_a_penalty_away"`
		PreAnytimeGoalScorer                    string    `gorm:"column:pre_anytime_goal_scorer"`
		EndAnytimeGoalScorer                    string    `gorm:"column:end_anytime_goal_scorer"`
		PreFirstGoalScorer                      string    `gorm:"column:pre_first_goal_scorer"`
		EndFirstGoalScorer                      string    `gorm:"column:end_first_goal_scorer"`
		PreLastGoalScorer                       string    `gorm:"column:pre_last_goal_scorer"`
		EndLastGoalScorer                       string    `gorm:"column:end_last_goal_scorer"`
		PreToScoreTwoOrMoreGoals                string    `gorm:"column:pre_to_score_two_or_more_goals"`
		EndToScoreTwoOrMoreGoals                string    `gorm:"column:end_to_score_two_or_more_goals"`
		PreLastGoalScorer2                      string    `gorm:"column:pre_last_goal_scorer2"`
		EndLastGoalScorer2                      string    `gorm:"column:end_last_goal_scorer2"`
		PreExactScore                           string    `gorm:"column:pre_exact_score"`
		EndExactScore                           string    `gorm:"column:end_exact_score"`
		PrePlayerToBeBooked                     string    `gorm:"column:pre_player_to_be_booked"`
		EndPlayerToBeBooked                     string    `gorm:"column:end_player_to_be_booked"`
		PreAhPattern                            string    `gorm:"column:pre_ah_pattern"`
		PreAhPatternMirror                      string    `gorm:"column:pre_ah_pattern_mirror"`
		PreGouPattern                           string    `gorm:"column:pre_gou_pattern"`
		EndAhPattern                            string    `gorm:"column:end_ah_pattern"`
		EndAhPatternMirror                      string    `gorm:"column:end_ah_pattern_mirror"`
		EndGouPattern                           string    `gorm:"column:end_gou_pattern"`
		PreResponse                             int       `gorm:"column:pre_response;default:0"`
		OneResponse                             int       `gorm:"column:one_response"`
		EndResponse                             int       `gorm:"column:end_response;default:0"`
		PreOddUpdatedAt                         time.Time `gorm:"column:pre_odd_updated_at"`
		OneOddUpdatedAt                         time.Time `gorm:"column:one_odd_updated_at"`
		EndOddUpdatedAt                         time.Time `gorm:"column:end_odd_updated_at"`
		PreAhPattern4                           string    `gorm:"column:pre_ah_pattern_4"`
		PreAhPatternMirror4                     string    `gorm:"column:pre_ah_pattern_mirror_4"`
		PreGouPattern4                          string    `gorm:"column:pre_gou_pattern_4"`
		EndAhPattern4                           string    `gorm:"column:end_ah_pattern_4"`
		EndAhPatternMirror4                     string    `gorm:"column:end_ah_pattern_mirror_4"`
		EndGouPattern4                          string    `gorm:"column:end_gou_pattern_4"`
		PreSpecialOdds                          int       `gorm:"column:pre_special_odds"`
		OneSpecialOdds                          int       `gorm:"column:one_special_odds"`
		EndSpecialOdds                          int       `gorm:"column:end_special_odds"`
		FootballPatternFromOnlyUpdatedAt        time.Time `gorm:"column:football_pattern_from_only_updated_at"`
		CreatedAt                               time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
		UpdatedAt                               time.Time `gorm:"column:updated_at"`
		DeletedAt                               time.Time `gorm:"column:deleted_at"`
	
		
		FixtureAPIID uint `gorm:"column:fixtureapi_id"` 
		FixtureTeam *FootballFixture `gorm:"foreignKey:FixtureAPIID;references:FixtureAPIID"` // note the pointer here 
		
		
		Tanggal    	string `gorm:"column:tanggal"`
		Jam    		string `gorm:"column:jam"`
		Counter    		string `gorm:"column:counter"`
		// ... other fields ...
		
		PreAh                   	string    `gorm:"column:pre_ah"` 
		EndAh                     	string    `gorm:"column:end_ah"` 

		QueryStatus                     	string    `gorm:"column:query_status"` 
	} 

	// ---------------------------------------------------

	// ---------------------------------------------------
	
	func (b *FootballOdd) GetFixtureInformation(leagueapiID uint, season uint, fixtureapi_id uint) ([]FootballOdd, error) {
		var footballodd []FootballOdd

		err := database.DB.
			Select("*").
			Where("leagueapi_id = ?", leagueapiID).
			Where("season = ?", season).
			Where("fixtureapi_id = ?", fixtureapi_id). 
			First(&footballodd).Error
			

		return footballodd, err
	} 

	func (b *FootballOdd) GetPatternOther(leagueapiID uint, preAhPattern string, preGouPattern string, preAhPatternMirror string) ([]FootballOdd, error) {
		var footballodd []FootballOdd

		err := database.DB.
			Raw(`SELECT pre_ah, pre_gou_pattern, end_ah, end_gou_pattern, leagueapi_id, COUNT(*) as counter
			FROM (SELECT leagueapi_id, pre_ah_pattern as pre_ah, pre_gou_pattern, end_ah_pattern as end_ah, end_gou_pattern 
				FROM football_odds 
				WHERE leagueapi_id = ? 
					AND pre_ah_pattern = ?
					AND pre_gou_pattern = ? 
					AND deleted_at IS NULL
				union all
				SELECT leagueapi_id, pre_ah_pattern as pre_ah, pre_gou_pattern, end_ah_pattern as end_ah, end_gou_pattern 
								FROM football_odds 
								WHERE leagueapi_id = ? 
									AND pre_ah_pattern_mirror = ?
									AND pre_gou_pattern = ? 
									AND deleted_at IS NULL
							) AS subquery group by leagueapi_id, pre_ah, pre_gou_pattern, end_ah, end_gou_pattern;`,
			leagueapiID, preAhPattern, preGouPattern, 
			leagueapiID, preAhPatternMirror, preGouPattern). 
			Find(&footballodd).Error
			 
		return footballodd, err
	} 
	
	// ---------------------------------------------------

	func (b *FootballOdd) GetFixtureOddPatternPrePre(leagueapiID uint, fixtureapi_idInt uint, preAhPattern string, preGouPattern string, preAhPatternMirror string) ([]FootballOdd, error) {
		var footballodd []FootballOdd
	
		db := database.DB.
			Raw(`SELECT *,
				'Mir' as query_status,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H:%i:%s') as jam
				FROM football_odds 
				WHERE leagueapi_id = ? 
				AND pre_ah_pattern = ?
				AND pre_gou_pattern = ?
				AND end_ah_pattern = ?
				AND end_gou_pattern = ? 
				AND deleted_at IS NULL
				UNION 
				SELECT *,
				'Ori' as query_status,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H:%i:%s') as jam
				FROM football_odds 
				WHERE leagueapi_id = ? 
				AND pre_ah_pattern = ?
				AND pre_gou_pattern = ?
				AND end_ah_pattern = ?
				AND end_gou_pattern = ? 
				AND deleted_at IS NULL
				UNION 
				SELECT *,
				'Ori' as query_status,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H:%i:%s') as jam
				FROM football_odds 
				WHERE fixtureapi_id = ? 
				AND deleted_at IS NULL
				`,   
				leagueapiID, preAhPatternMirror, preGouPattern, preAhPatternMirror, preGouPattern,
				leagueapiID, preAhPattern, preGouPattern, preAhPattern, preGouPattern,
				fixtureapi_idInt)

			err := db.Find(&footballodd).Error

			if err != nil {
				return nil, err
			}
		
			// Now, you can preload related data
			for i := range footballodd {
				db.Preload("FixtureTeam"). 
					Preload("FixtureTeam.HomeTeam"). 
					Preload("FixtureTeam.AwayTeam"). 
					Where("id = ?", footballodd[i].ID).Find(&footballodd[i])
			}
		
			return footballodd, err   
	}

	func (b *FootballOdd) GetFixtureOddPatternOdd(leagueapiID uint, preAhPattern string, preGouPattern string, endAhPattern string, endGouPattern string) ([]FootballOdd, error) {
		var footballodd []FootballOdd
	
		db := database.DB.
			Raw(`SELECT *,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H:%i:%s') as jam
				FROM football_odds 
				WHERE leagueapi_id = ? 
				AND pre_ah_pattern_mirror = ?
				AND pre_gou_pattern = ?
				AND end_ah_pattern_mirror = ?
				AND end_gou_pattern = ? 
				AND deleted_at IS NULL
				UNION 
				SELECT *,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H:%i:%s') as jam
				FROM football_odds 
				WHERE leagueapi_id = ? 
				AND pre_ah_pattern = ?
				AND pre_gou_pattern = ?
				AND end_ah_pattern = ?
				AND end_gou_pattern = ? 
				AND deleted_at IS NULL 
				`,   
				leagueapiID, preAhPattern, preGouPattern, endAhPattern, endGouPattern,
				leagueapiID, preAhPattern, preGouPattern, endAhPattern, endGouPattern)

			err := db.Find(&footballodd).Error

			if err != nil {
				return nil, err
			}
		
			// Now, you can preload related data
			for i := range footballodd {
				db.Preload("FixtureTeam"). 
					Preload("FixtureTeam.HomeTeam"). 
					Preload("FixtureTeam.AwayTeam"). 
					Where("id = ?", footballodd[i].ID).Find(&footballodd[i])
			}
		
			return footballodd, err   
	}

	func (b *FootballOdd) GetFixtureOddPatternPreEnd(leagueapiID uint, fixtureapi_idInt uint, preAhPattern string, preGouPattern string, preAhPatternMirror string, endAhPattern string, endGouPattern string, endAhPatternMirror string) ([]FootballOdd, error) {
		var footballodd []FootballOdd
	
		db := database.DB.
			Raw(`SELECT *,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H:%i:%s') as jam
				FROM football_odds 
				WHERE leagueapi_id = ? 
				AND pre_ah_pattern = ?
				AND pre_gou_pattern = ?
				AND end_ah_pattern = ?
				AND end_gou_pattern = ? 
				AND deleted_at IS NULL
				UNION 
				SELECT *,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H:%i:%s') as jam
				FROM football_odds 
				WHERE leagueapi_id = ? 
				AND pre_ah_pattern = ?
				AND pre_gou_pattern = ?
				AND end_ah_pattern = ?
				AND end_gou_pattern = ? 
				AND deleted_at IS NULL
				UNION 
				SELECT *,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H:%i:%s') as jam
				FROM football_odds 
				WHERE fixtureapi_id = ? 
				AND deleted_at IS NULL
				`,   
				leagueapiID, preAhPatternMirror, preGouPattern, endAhPatternMirror, endGouPattern,
				leagueapiID, preAhPattern, preGouPattern, endAhPattern, endGouPattern,
				fixtureapi_idInt)

			err := db.Find(&footballodd).Error

			if err != nil {
				return nil, err
			}
		
			// Now, you can preload related data
			for i := range footballodd {
				db.Preload("FixtureTeam"). 
					Preload("FixtureTeam.HomeTeam"). 
					Preload("FixtureTeam.AwayTeam"). 
					Where("id = ?", footballodd[i].ID).Find(&footballodd[i])
			}
		
			return footballodd, err   
	}

	func (b *FootballOdd) GetFixtureOddPatternOnlyPre(leagueapiID uint, fixtureapi_idInt uint, preAhPattern string, preGouPattern string, preAhPatternMirror string) ([]FootballOdd, error) {
		var footballodd []FootballOdd
	
		db := database.DB.
			Raw(`SELECT *,
			DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal,
			DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H:%i:%s') as jam
				FROM football_odds 
				WHERE leagueapi_id = ? 
				AND pre_ah_pattern = ?
				AND pre_gou_pattern = ?
				AND end_ah_pattern = ?
				AND end_gou_pattern = ? 
				AND deleted_at IS NULL
				UNION
				SELECT *,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H:%i:%s') as jam
				FROM football_odds 
				WHERE leagueapi_id = ? 
				AND pre_ah_pattern = ?
				AND pre_gou_pattern = ?
				AND end_ah_pattern = ?
				AND end_gou_pattern = ? 
				AND deleted_at IS NULL
				UNION 
				SELECT *,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H:%i:%s') as jam
				FROM football_odds 
				WHERE fixtureapi_id = ? 
				AND deleted_at IS NULL
				`,   
				leagueapiID, preAhPatternMirror, preGouPattern, preAhPatternMirror, preGouPattern,
				leagueapiID, preAhPattern, preGouPattern, preAhPattern, preGouPattern,
				fixtureapi_idInt)

			err := db.Find(&footballodd).Error

			if err != nil {
				return nil, err
			}
		
			// Now, you can preload related data
			for i := range footballodd {
				db.Preload("FixtureTeam"). 
					Preload("FixtureTeam.HomeTeam"). 
					Preload("FixtureTeam.AwayTeam"). 
					Where("id = ?", footballodd[i].ID).Find(&footballodd[i])
			}
		
			return footballodd, err   
	}

	func (b *FootballOdd) GetFixtureOddPatternOnlyEnd(leagueapiID uint, fixtureapi_idInt uint, preAhPattern string, preGouPattern string, preAhPatternMirror string) ([]FootballOdd, error) {
		var footballodd []FootballOdd
	
		db := database.DB.
			Raw(`SELECT *,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H:%i:%s') as jam
				FROM football_odds 
				WHERE leagueapi_id = ? 
				AND pre_ah_pattern = ?
				AND pre_gou_pattern = ?
				AND end_ah_pattern = ?
				AND end_gou_pattern = ? 
				AND deleted_at IS NULL
				UNION
				SELECT *,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H:%i:%s') as jam
				FROM football_odds 
				WHERE leagueapi_id = ? 
				AND pre_ah_pattern = ?
				AND pre_gou_pattern = ?
				AND end_ah_pattern = ?
				AND end_gou_pattern = ? 
				AND deleted_at IS NULL
				UNION 
				SELECT *,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H:%i:%s') as jam
				FROM football_odds 
				WHERE fixtureapi_id = ? 
				AND deleted_at IS NULL
				`,   
				leagueapiID, preAhPatternMirror, preGouPattern, preAhPatternMirror, preGouPattern,
				leagueapiID, preAhPattern, preGouPattern, preAhPattern, preGouPattern,
				fixtureapi_idInt)

			err := db.Find(&footballodd).Error

			if err != nil {
				return nil, err
			}
		
			// Now, you can preload related data
			for i := range footballodd {
				db.Preload("FixtureTeam"). 
					Preload("FixtureTeam.HomeTeam"). 
					Preload("FixtureTeam.AwayTeam"). 
					Where("id = ?", footballodd[i].ID).Find(&footballodd[i])
			}
		
			return footballodd, err   
	}
	
	// ---------------------------------------------------
	
	func (b *FootballOdd) GetFixtureOddPatternPrePreCountry(leagueapiID uint, fixtureapi_idInt uint, preAhPattern string, preGouPattern string, preAhPatternMirror string) ([]FootballOdd, error) {
		var footballodd []FootballOdd
	
		db := database.DB.
			Raw(`SELECT *,
				'Mir' as query_status,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H:%i:%s') as jam
				FROM football_odds 
				WHERE leagueapi_id IN (SELECT leagueapi_id  FROM football_leagues WHERE  country_name in (SELECT country_name FROM football_leagues WHERE leagueapi_id = ?))
				AND pre_ah_pattern = ?
				AND pre_gou_pattern = ?
				AND end_ah_pattern = ?
				AND end_gou_pattern = ?  
				AND deleted_at IS NULL
				UNION 
				SELECT *,
				'Ori' as query_status,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H:%i:%s') as jam
				FROM football_odds 
				WHERE leagueapi_id IN (SELECT leagueapi_id  FROM football_leagues WHERE  country_name in (SELECT country_name FROM football_leagues WHERE leagueapi_id = ?))
				AND pre_ah_pattern = ?
				AND pre_gou_pattern = ?
				AND end_ah_pattern = ?
				AND end_gou_pattern = ? 
				AND deleted_at IS NULL
				UNION 
				SELECT *,
				'Ori' as query_status,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H:%i:%s') as jam
				FROM football_odds 
				WHERE fixtureapi_id = ? 
				AND deleted_at IS NULL
				`,   
				leagueapiID, preAhPatternMirror, preGouPattern, preAhPatternMirror, preGouPattern,
				leagueapiID, preAhPattern, preGouPattern, preAhPattern, preGouPattern,
				fixtureapi_idInt)

			err := db.Find(&footballodd).Error

			if err != nil {
				return nil, err
			}
		
			// Now, you can preload related data
			for i := range footballodd {
				db.Preload("FixtureTeam"). 
					Preload("FixtureTeam.HomeTeam"). 
					Preload("FixtureTeam.AwayTeam"). 
					Where("id = ?", footballodd[i].ID).Find(&footballodd[i])
			}
		
			return footballodd, err   
	}

	func (b *FootballOdd) GetFixtureOddPatternPreEndCountry(leagueapiID uint, fixtureapi_idInt uint, preAhPattern string, preGouPattern string, preAhPatternMirror string, endAhPattern string, endGouPattern string, endAhPatternMirror string) ([]FootballOdd, error) {
		var footballodd []FootballOdd
	
		db := database.DB.
			Raw(`SELECT *,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H:%i:%s') as jam
				FROM football_odds 
				WHERE leagueapi_id IN (SELECT leagueapi_id  FROM football_leagues WHERE  country_name in (SELECT country_name FROM football_leagues WHERE leagueapi_id = ?))
				AND pre_ah_pattern = ?
				AND pre_gou_pattern = ?
				AND end_ah_pattern = ?
				AND end_gou_pattern = ? 
				AND deleted_at IS NULL
				UNION
				SELECT *,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H:%i:%s') as jam
				FROM football_odds 
				WHERE leagueapi_id IN (SELECT leagueapi_id  FROM football_leagues WHERE  country_name in (SELECT country_name FROM football_leagues WHERE leagueapi_id = ?))
				AND pre_ah_pattern = ?
				AND pre_gou_pattern = ?
				AND end_ah_pattern = ?
				AND end_gou_pattern = ? 
				AND deleted_at IS NULL
				UNION 
				SELECT *,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H:%i:%s') as jam
				FROM football_odds 
				WHERE fixtureapi_id = ? 
				AND deleted_at IS NULL
				`,   
				leagueapiID, preAhPatternMirror, preGouPattern, endAhPatternMirror, endGouPattern,
				leagueapiID, preAhPattern, preGouPattern, endAhPattern, endGouPattern,
				fixtureapi_idInt)

			err := db.Find(&footballodd).Error

			if err != nil {
				return nil, err
			}
		
			// Now, you can preload related data
			for i := range footballodd {
				db.Preload("FixtureTeam"). 
					Preload("FixtureTeam.HomeTeam"). 
					Preload("FixtureTeam.AwayTeam"). 
					Where("id = ?", footballodd[i].ID).Find(&footballodd[i])
			}
		
			return footballodd, err   
	}

	func (b *FootballOdd) GetFixtureOddPatternOnlyPreCountry(leagueapiID uint, fixtureapi_idInt uint, preAhPattern string, preGouPattern string, preAhPatternMirror string) ([]FootballOdd, error) {
		var footballodd []FootballOdd
	
		db := database.DB.
			Raw(`SELECT *,
			DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal,
			DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H:%i:%s') as jam
				FROM football_odds 
				WHERE leagueapi_id IN (SELECT leagueapi_id  FROM football_leagues WHERE  country_name in (SELECT country_name FROM football_leagues WHERE leagueapi_id = ?))
				AND pre_ah_pattern = ?
				AND pre_gou_pattern = ?
				AND end_ah_pattern = ?
				AND end_gou_pattern = ? 
				AND deleted_at IS NULL
				UNION
				SELECT *,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H:%i:%s') as jam
				FROM football_odds 
				WHERE leagueapi_id IN (SELECT leagueapi_id  FROM football_leagues WHERE  country_name in (SELECT country_name FROM football_leagues WHERE leagueapi_id = ?))
				AND pre_ah_pattern = ?
				AND pre_gou_pattern = ?
				AND end_ah_pattern = ?
				AND end_gou_pattern = ? 
				AND deleted_at IS NULL
				UNION 
				SELECT *,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H:%i:%s') as jam
				FROM football_odds 
				WHERE fixtureapi_id = ? 
				AND deleted_at IS NULL
				`,   
				leagueapiID, preAhPatternMirror, preGouPattern, preAhPatternMirror, preGouPattern,
				leagueapiID, preAhPattern, preGouPattern, preAhPattern, preGouPattern,
				fixtureapi_idInt)

			err := db.Find(&footballodd).Error

			if err != nil {
				return nil, err
			}
		
			// Now, you can preload related data
			for i := range footballodd {
				db.Preload("FixtureTeam"). 
					Preload("FixtureTeam.HomeTeam"). 
					Preload("FixtureTeam.AwayTeam"). 
					Where("id = ?", footballodd[i].ID).Find(&footballodd[i])
			}
		
			return footballodd, err   
	}

	func (b *FootballOdd) GetFixtureOddPatternOnlyEndCountry(leagueapiID uint, fixtureapi_idInt uint, preAhPattern string, preGouPattern string, preAhPatternMirror string) ([]FootballOdd, error) {
		var footballodd []FootballOdd
	
		db := database.DB.
			Raw(`SELECT *,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H:%i:%s') as jam
				FROM football_odds 
				WHERE leagueapi_id IN (SELECT leagueapi_id  FROM football_leagues WHERE  country_name in (SELECT country_name FROM football_leagues WHERE leagueapi_id = ?))
				AND pre_ah_pattern = ?
				AND pre_gou_pattern = ?
				AND end_ah_pattern = ?
				AND end_gou_pattern = ? 
				AND deleted_at IS NULL
				UNION
				SELECT *,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H:%i:%s') as jam
				FROM football_odds 
				WHERE leagueapi_id IN (SELECT leagueapi_id  FROM football_leagues WHERE  country_name in (SELECT country_name FROM football_leagues WHERE leagueapi_id = ?))
				AND pre_ah_pattern = ?
				AND pre_gou_pattern = ?
				AND end_ah_pattern = ?
				AND end_gou_pattern = ? 
				AND deleted_at IS NULL
				UNION 
				SELECT *,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%Y-%m-%d') as tanggal,
				DATE_FORMAT(DATE_ADD(date, INTERVAL 7 HOUR), '%H:%i:%s') as jam
				FROM football_odds 
				WHERE fixtureapi_id = ? 
				AND deleted_at IS NULL
				`,   
				leagueapiID, preAhPatternMirror, preGouPattern, preAhPatternMirror, preGouPattern,
				leagueapiID, preAhPattern, preGouPattern, preAhPattern, preGouPattern,
				fixtureapi_idInt)

			err := db.Find(&footballodd).Error

			if err != nil {
				return nil, err
			}
		
			// Now, you can preload related data
			for i := range footballodd {
				db.Preload("FixtureTeam"). 
					Preload("FixtureTeam.HomeTeam"). 
					Preload("FixtureTeam.AwayTeam"). 
					Where("id = ?", footballodd[i].ID).Find(&footballodd[i])
			}
		
			return footballodd, err   
	}
 
	// ---------------------------------------------------   
	func (b *FootballOdd) GetPrePatternNonPatternPercentage() (float64, error) {
		var totalRows, specificRows int64
	
		// Count all rows
		if err := database.DB.Model(&FootballOdd{}).Count(&totalRows).Error; err != nil {
			return 0, err
		}
	
		// Count specific rows where pre_response = 4
		if err := database.DB.Model(&FootballOdd{}).Where("pre_response = 4").Count(&specificRows).Error; err != nil {
			return 0, err
		}
	
		// Calculate percentage
		percentage := float64(specificRows) / float64(totalRows) * 100
	
		return percentage, nil
	}  

	func (b *FootballOdd) GetOtherPrePatternNonPatternPercentage() (float64, error) {
		var totalRows, specificRows int64
	
		// Count all rows
		if err := database.DB.Model(&FootballOdd{}).Count(&totalRows).Error; err != nil {
			return 0, err
		}
	
		// Count specific rows where pre_response = 4
		if err := database.DB.Model(&FootballOdd{}).Where("pre_response != 4").Count(&specificRows).Error; err != nil {
			return 0, err
		}
	
		// Calculate percentage
		percentage := float64(specificRows) / float64(totalRows) * 100
	
		return percentage, nil
	}
	
	func (b *FootballOdd) GetPatternResponse(leagueapiID uint, season uint) ([]FootballOdd, error) {
		var footballodd []FootballOdd
	
		err := database.DB.
			Select("leagueapi_id, pre_response, end_response, count(*) as counter").
			Where("leagueapi_id = ?", leagueapiID). 
			Group("pre_response, end_response").
			Find(&footballodd).Error
	
		return footballodd, err
	}
	
	func (b *FootballOdd) GetPatternResponseLeague(leagueapiID uint, season uint, pre_response uint, end_response uint) ([]FootballOdd, error) {
		var footballodd []FootballOdd
	
		err := database.DB.
			Select("*").
			Where("leagueapi_id = ?", leagueapiID).  
			// Where("season = ?", season).  
			Where("pre_response = ?", pre_response).  
			Where("end_response = ?", end_response).  
			Preload("FixtureTeam").
			Find(&footballodd).Error
	
		return footballodd, err
	}

	func (b *FootballOdd) GetPatternResponseCountry(countryName string) ([]FootballOdd, error) {
		var footballOdd []FootballOdd
	
		err := database.DB.
			Select("leagueapi_id, pre_response, end_response").
			Where("leagueapi_id IN (SELECT leagueapi_id FROM football_leagues WHERE country_name LIKE ?)", countryName).  
			Group("leagueapi_id, pre_response, end_response").
			Find(&footballOdd).Error
	
		return footballOdd, err
	}
	
	
	
	
	
	