package models
 
    import ( 
        "github.com/HerlambangHaryo/go-crud-simple/platform/database"  
        // "time"  
        "gorm.io/gorm"
    )

    type FootballTeam struct {
        gorm.Model 
        Name      	string `gorm:"column:name"`
        Logo      	string `gorm:"column:logo"`
        TeamAPIID 	uint   `gorm:"column:teamapi_id"`
        Code 	    string   `gorm:"column:code"` 
        Country 	    string   `gorm:"column:country"`
        Founded 	    string   `gorm:"column:founded"`
        VenueAPIID 	    uint   `gorm:"column:venueapi_id"`
        // ... other fields ...
    } 


    
    func (b *FootballTeam) GetTeam(teamapiID uint) ([]FootballTeam, error) {
        var footballteam []FootballTeam

        err := database.DB.
            Select("*").
            Where("teamapi_id = ?", teamapiID). 
            Find(&footballteam).Error

        return footballteam, err
    } 

// func (b *Team) GetTeam(id uint) (Team, error) {
// 	var team Team
// 	err := database.DB.First(&team, id).Error
// 	return team, err
// }