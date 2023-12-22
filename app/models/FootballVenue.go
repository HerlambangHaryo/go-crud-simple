package models

// Mengimpor paket yang dibutuhkan   
import ( 
	// "github.com/HerlambangHaryo/go-crud-simple/platform/database"  
	// "time"  
	"gorm.io/gorm"
)

type FootballVenue struct {
    gorm.Model 
    Name      	string `gorm:"column:name"`
    City      	string `gorm:"column:city"`
    Image      	string `gorm:"column:image"`
    VenueAPIID 	uint   `gorm:"column:venueapi_id"`
    // ... other fields ...
} 