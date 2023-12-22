// book.go
package controllers

    import (
		"github.com/gofiber/fiber/v2"
		"github.com/HerlambangHaryo/go-crud-simple/app/models" 
		// "strconv" 
	) 

	func GetMTLeagues(c *fiber.Ctx) error { 
		// -------------------------------------------------------------- 
			b := new(models.ResearchOnAnytimeGoalScorer)   
		// --------------------------------------------------------------   
			data, err := b.GetAgsByToday() 
			if err != nil {
				return c.Status(500).SendString(err.Error())
			}   
        // --------------------------------------------------------------
            return c.Render("contents/modeltesting/GetMTLeagues", fiber.Map{
                "Title": "Anytime Goal Scorer", 
				"Content": "Dashboard",
				"Data": data,     
			},"templates/studiov30/datatable")
        // --------------------------------------------------------------
    }