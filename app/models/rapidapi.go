package models

import ( 
	"github.com/HerlambangHaryo/go-crud-simple/platform/database" 
)

type Rapidapi struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Account string `json:"account"`
	Apikey string `json:"apikey"`
	Active string `json:"active"`
	Counter string `json:"counter"` 
}

// TableName menentukan nama tabel untuk model Rapidapi
func (Rapidapi) TableName() string {
	return "apiaccounts"
}
  
// Get all rapidapis with selected fields
func (b *Rapidapi) GetRapidapis() ([]Rapidapi, error) {
	var rapidapis []Rapidapi
	err := database.DB.Select("ID", "Account", "Apikey", "Active", "Counter").Find(&rapidapis).Error
	return rapidapis, err
}


// Get a rapidapi by id
func (b *Rapidapi) GetRapidapi(id uint) (Rapidapi, error) {
	var rapidapi Rapidapi
	err := database.DB.First(&rapidapi, id).Error
	return rapidapi, err
}

// Create a rapidapi
func (b *Rapidapi) CreateRapidapi() error {
	err := database.DB.Create(&b).Error
	return err
}

// Update a rapidapi by id
func (b *Rapidapi) UpdateRapidapi(id uint) error {
	err := database.DB.Model(&b).Where("id = ?", id).Updates(b).Error
	return err
}

// Delete a rapidapi by id
func (b *Rapidapi) DeleteRapidapi(id uint) error {
	err := database.DB.Delete(&b, id).Error
	return err
}
