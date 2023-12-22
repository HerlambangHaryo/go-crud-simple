// book.go
package controllers

    import (
        "github.com/gofiber/fiber/v2" 
		"github.com/HerlambangHaryo/go-crud-simple/app/models" 
    ) 
  
    func GetDashboard(c *fiber.Ctx) error { 
        // --------------------------------------------------------------  
			MFO := new(models.FootballOdd)   
        // --------------------------------------------------------------    
            CountPre, err := MFO.GetPrePatternNonPatternPercentage() 
            if err != nil {
                return c.Status(500).SendString(err.Error())
            } 
        // --------------------------------------------------------------    
            CountOtherPre, err := MFO.GetOtherPrePatternNonPatternPercentage() 
            if err != nil {
                return c.Status(500).SendString(err.Error())
            }   
        // --------------------------------------------------------------
            return c.Render("contents/dashboard/GetDashboard", fiber.Map{
                "Title": "Dashboard", 
				"Content": "Dashboard",
                "CountPre": CountPre, 
                "CountOtherPre": CountOtherPre, 
            },"templates/studiov30/pageblank")
        // --------------------------------------------------------------
    }