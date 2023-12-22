package models

	import ( 
		"github.com/HerlambangHaryo/go-crud-simple/platform/database" 
	)

	type Rapidapi struct {
		ID   			uint   `json:"id" gorm:"primaryKey"`
		Account 		string `json:"account"`
		Apikey 			string `json:"apikey"`
		Active 			string `json:"active"`
		Counter 		string `json:"counter"` 
        UpdatedAt   	string `gorm:"column:updated_at"`
        UpdatedAtDIFF   	string `gorm:"column:updated_atDIFF"`
	} 

	func (Rapidapi) TableName() string {
		return "apiaccounts"
	}
	 
	func (b *Rapidapi) GetRapidapis() ([]Rapidapi, error) {
		var rapidapis []Rapidapi 

		err := database.DB.
			Select("*", "CONCAT(TIMESTAMPDIFF(DAY, updated_at, CURDATE()), ' days, ', TIMESTAMPDIFF(HOUR, updated_at, CURDATE()) - TIMESTAMPDIFF(DAY, updated_at, CURDATE()) * 24, ' hours, ', TIMESTAMPDIFF(MINUTE, updated_at, CURDATE()) - TIMESTAMPDIFF(HOUR, updated_at, CURDATE()) * 60, ' minutes') AS updated_atDIFF").  
			Where("deleted_at IS NULL").  
			Find(&rapidapis).Error

		return rapidapis, err
	}
  
	func (b *Rapidapi) GetRapidapi(id uint) (Rapidapi, error) {
		var rapidapi Rapidapi
		err := database.DB.First(&rapidapi, id).Error
		return rapidapi, err
	}
 
	func (b *Rapidapi) CreateRapidapi() error {
		err := database.DB.Create(&b).Error
		return err
	}
 
	func (b *Rapidapi) UpdateRapidapi(id uint) error {
		err := database.DB.Model(&b).Where("id = ?", id).Updates(b).Error
		return err
	}
 
	func (b *Rapidapi) DeleteRapidapi(id uint) error {
		err := database.DB.Delete(&b, id).Error
		return err
	}
