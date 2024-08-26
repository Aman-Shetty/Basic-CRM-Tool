package lead

import (
	"github.com/Aman-Shetty/Basic-CRM-Tool/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// The structure of the data to be stored on the database.
type Lead struct {
	gorm.Model
	Name    string
	Company string
	Email   string
	Phone   int
}

// Function returns all the values
// that are on the lead database
func GetLeads(c *fiber.Ctx) {
	db := database.DBConnection
	var lead []Lead
	db.Find(&lead)
	c.JSON(lead)
}

// Function returns the value from the
// database based on the provided id
func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConnection
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)
}

// Function to create a new lead in the database
// based on the user input.
func NewLead(c *fiber.Ctx) {
	db := database.DBConnection
	lead := new(Lead)

	// If there is any error in parsing the data,
	// this if statement should handle it.
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
		return
	}

	// This creates the new lead on the database
	db.Create(&lead)

	// This is confirm that correct data is sent.
	c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConnection

	var lead Lead

	// Find the value based on the given id
	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(500).Send("No lead found with ID")
		return
	}

	// Delete the data on the index of the id in the database
	db.Delete(&lead)
	c.Send("Lead successfully Deleted")
}
