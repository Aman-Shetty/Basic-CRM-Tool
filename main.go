package main

import (
	"fmt"

	"github.com/Aman-Shetty/Basic-CRM-Tool/database"
	"github.com/Aman-Shetty/Basic-CRM-Tool/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConnection, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("Failed to Connect to Database.")
	}
	fmt.Println("Successfully connected to the Database.")
	database.DBConnection.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBConnection.Close()
}
