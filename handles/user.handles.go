package handles

import (
	"fmt"

	"github.com/eakarach/AuthJWT/models"

	"github.com/gofiber/fiber/v2"
)


func GetAllUser(c *fiber.Ctx) error {
	
	return c.JSON(fiber.Map{
		"status": "GetAllUser Success !!",
	})
}

func GetUserProfile(c *fiber.Ctx) error {
	usrId := c.Params("usrId")
	return c.JSON(fiber.Map{
		"status": fmt.Sprintf("GetUserProfile %s Success !!", usrId),
	})
}

func CreateUser(c *fiber.Ctx) error {
	input := new(models.UserData)

	// Parse the request body
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"resCode": "40000", "resMessage": "Error on login request", "error": err})
	}

	return c.JSON(fiber.Map{
		"status": "CreateUser Success !!",
	})
}

func DeleteUser(c *fiber.Ctx) error {
	usrId := c.Params("usrId")
	return c.JSON(fiber.Map{
		"status": fmt.Sprintf("DeleteUser %s Success !!", usrId),
	})
}
