package main

import (
	"fmt"
	"go-fiber-crm/database"
	"go-fiber-crm/lead"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Post("/api/v1/lead", lead.CreateLead)
	app.Get("/api/v1/lead/:id", lead.GetLeadByID)
	app.Put("/api/v1/lead/:id", lead.UpdateLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	err := database.Connect()
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Database connection successfully opened")

	// Auto-migrate the Lead model
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()

	initDatabase()
	setupRoutes(app)

	app.Listen(":3000")
}