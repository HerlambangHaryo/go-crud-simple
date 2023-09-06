// book.go
package controllers

import (
	"github.com/gofiber/fiber/v2" 
) 

// Get all dashboard
func GetDashboard(c *fiber.Ctx) error { 
    return c.Render("contents/dashboard/index", fiber.Map{
        "dashboard": "",
    },"templates/studiov30/pageblank")
}