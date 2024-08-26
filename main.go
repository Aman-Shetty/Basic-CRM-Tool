package main

import (
	"fmt"

	"github.com/Aman-Shetty/Basic-CRM-Tool/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get(GetLeads)
	app.Get(GetLead)
	app.Post(NewLead)
	app.Delete(DeleteLead)
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
