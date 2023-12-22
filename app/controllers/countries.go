package controllers

	import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/go-crud-simple/app/models" 
		// "strconv"
	) 
  
	func GetCountries(c *fiber.Ctx) error {
		// --------------------------------------------------------------  
			b := new(models.Country)
		// --------------------------------------------------------------  
			countries, err := b.GetCountries()
			if err != nil {
				return c.Status(500).SendString(err.Error())
			}  
		// --------------------------------------------------------------  
			return c.Render("contents/countries/GetCountries", fiber.Map{
				"Title": "Countries",
				"Content": "Countries",
				"Countries": countries,
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------  
	}
 
	func GetCountry(c *fiber.Ctx) error {
		// --------------------------------------------------------------  
			id := c.Params("id")
		// --------------------------------------------------------------  
			b 		:= new(models.Country) 
			MFO 	:= new(models.FootballOdd)
		// --------------------------------------------------------------  
			country, err := b.GetCountry(id)

			if err != nil {
				return c.Status(500).SendString(err.Error())
			}  
		// --------------------------------------------------------------  
			DataResponse, err := MFO.GetPatternResponseCountry(id)

			if err != nil {
				return c.Status(500).SendString(err.Error())
			}  
		// --------------------------------------------------------------  
			return c.Render("contents/countries/GetCountry", fiber.Map{
				"Title": id,
				"Content": "Countries",
				"Data": country,
				"DataResponse": DataResponse,
			},"templates/studiov30/datatable")
		// --------------------------------------------------------------  
	}