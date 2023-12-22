package models
 
    import ( 
        "github.com/HerlambangHaryo/go-crud-simple/platform/database"  
        // "time"  
        "gorm.io/gorm"
    )

    type FootballCoach struct {
        gorm.Model 
        CoachAPIID 	uint   `gorm:"column:coachapi_id"` 
        Name      	string `gorm:"column:name"` 
        Image      	string `gorm:"column:image"` 

		
        TeamAPIID      	string `gorm:"column:teamapi_id"` 
    } 

	
	func (b *FootballCoach) GetCoach(teamapiID uint) ([]FootballCoach, error) {
		var footballcoach []FootballCoach
	
		err := database.DB.Raw(`
			SELECT DISTINCT fc.*
			, fs.teamapi_id
			FROM football_coaches fc
			JOIN (
				SELECT lineups_coach_homeapi_id AS coach_id,
				teams_homeapi_id as teamapi_id
				FROM football_statistics
				WHERE teams_homeapi_id = ?
				GROUP BY lineups_coach_homeapi_id
	
				UNION
	
				SELECT lineups_coach_awayapi_id AS coach_id,
				teams_awayapi_id as teamapi_id
				FROM football_statistics
				WHERE teams_awayapi_id = ?
				GROUP BY lineups_coach_awayapi_id
			) fs ON fc.coachapi_id = fs.coach_id
		`, teamapiID, teamapiID).
		Scan(&footballcoach).Error
	
		return footballcoach, err
	}
	