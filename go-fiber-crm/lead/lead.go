package lead

import (
	"go-fiber-crm/database"

	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)
type Lead struct {
	gorm.Model
	Name string `json:"name"`
	Company string `json:"company"`

	Email string `json:"email"`

	Phone int `json:"phone"`
}

func Getleads(c, *fiber.ctx){

	db := database.DBConn
	var leads []Lead
	db.Find(&leads, id)
	c.JSON(lead)

}


func GetLeadByID(c, *fiber.ctx){

	id := c.params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)	

}

func DeleteLead(){

	id := c.Params("id")
	
}


func CreateLead(c, *fiber.Ctx){
	db := database.DBConn
	lead := new(Lead)
	if err:= c.BodyParser(lead); err !=nil {

		c.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	c.JSON(lead)

}

