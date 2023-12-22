package models

    import ( 
        "github.com/HerlambangHaryo/go-crud-simple/platform/database" 
        "time" 
    )
    
    type Country struct {
        ID                              uint   `json:"id" gorm:"primaryKey"` 
        Name                            string 
        Logo                            string 
        Code                            string `json:"code"`
        Football_pattern_from_country   int `json:"football_pattern_from_country"`
        Api_league_updated_at           *time.Time `gorm:"column:api_league_updated_at"`
        Season                          string `json:"season"`
        Leagueapi_id                    string `json:"leagueapi_id"`
    }
    
 
    func (b *Country) GetCountries() ([]Country, error) {
        var countries []Country
    
        err := database.DB.
            Select("*").    
            Order("football_star asc"). 
            Order("name asc"). 
            Find(&countries).Error

        return countries, err
    } 

    // func (b *Country) GetCountries() ([]Country, error) {
    //     var countries []Country
    //     err := database.DB.Select("ID", "Name", "Code", "Football_pattern_from_country", "Api_league_updated_at").Find(&countries).Error
    //     return countries, err
    // }

 
func (b *Country) GetCountry(id string) ([]Country, error) {
	var country []Country 

    query := `
        SELECT football_fixtures.leagueapi_id, 
                football_fixtures.season, 
                football_leagues.name, 
                football_leagues.type, 
                football_leagues.bookmakersapi_id, 
                football_leagues.bookmakers_name, 
                football_leagues.logo 
        FROM football_leagues 
        INNER JOIN football_fixtures 
        ON football_fixtures.leagueapi_id = football_leagues.leagueapi_id 
        WHERE football_leagues.country_name = ?
        AND football_leagues.deleted_at IS NULL 
        GROUP BY football_fixtures.leagueapi_id, 
                football_fixtures.season, 
                football_leagues.name, 
                football_leagues.type, 
                football_leagues.bookmakersapi_id, 
                football_leagues.bookmakers_name
    `  
    err := database.DB.Raw(query, id).Scan(&country).Error

	return country, err
}