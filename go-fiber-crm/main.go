package main

import "github.com/gofiber/fiber/v2"

func setupRoutes(app *fiber.App) {
	app.Get(GetLeads)
	app.Post(CreateLead)
	app.Get(GetLeadByID)
	app.Put(UpdateLead)
	app.Delete(DeleteLead)
	
}

func main() {
	app := fiber.New()
	setupRoutes(app)

	app.Listen(":3000")
	

}
