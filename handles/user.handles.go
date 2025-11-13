package handles

import (
	"fmt"
	"strconv"

	"github.com/eakarach/AuthJWT/database"
	"github.com/eakarach/AuthJWT/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllUser(c *fiber.Ctx) error {

	userModel, err := new([]models.User), *new(error)

	// get User by id
	userModel, err = database.GetAllUser()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"resCode": "50000", "resMessage": "Internal Server Error", "error": err})
	} else if userModel == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"resCode": "40400", "resMessage": "User not found", "error": err})
	} else {
		var userDatas []models.UserData
		for _, p := range *userModel {
			userData := models.UserData{
				ID:       p.ID,
				Name:     p.Name,
				Email:    p.Email,
				Username: p.Username,
			}
			userDatas = append(userDatas, userData)
		}
		return c.JSON(fiber.Map{
			"resCode":    "20000",
			"resMessage": "Get All User Success !!",
			"data":       userDatas})
	}

}

func GetUserProfile(c *fiber.Ctx) error {
	// validate usrId is int
	usrId, err := strconv.Atoi(c.Params("usrId"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"resCode": "40400", "resMessage": "User not found", "error": err})
	}

	// get User by id
	userModel, err := database.GetUserByUserId(usrId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"resCode": "50000", "resMessage": "Internal Server Error", "error": err})
	} else if userModel == nil || userModel.ID == 0  {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"resCode": "40400", "resMessage": "User not found"})
	} else {
		userData := models.UserData{
			ID:       userModel.ID,
			Username: userModel.Username,
			Name:     userModel.Name,
			Email:    userModel.Email,
		}
		return c.JSON(fiber.Map{
			"resCode":    "20000",
			"resMessage": fmt.Sprintf("Get User Profile %d Success !!", usrId),
			"data":       userData})
	}
}

func CreateUser(c *fiber.Ctx) error {
	input := new(models.User)

	// Parse the request body
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"resCode": "40000", "resMessage": "Error on login request", "error": err})
	}

	userModel, err := database.CreateUser(input)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"resCode": "50000", "resMessage": "Internal Server Error", "error": err})
	} else if userModel == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"resCode": "40000", "resMessage": "Cannot create user", "error": err})
	} else {
		userData := models.UserData{
			ID:       userModel.ID,
			Username: userModel.Username,
			Name:     userModel.Name,
			Email:    userModel.Email,
		}
		return c.JSON(fiber.Map{
			"resCode":    "20000",
			"resMessage": "Create User Success !!",
			"data":       userData})
	}
}

func DeleteUser(c *fiber.Ctx) error {
	// validate usrId is int
	usrId, err := strconv.Atoi(c.Params("usrId"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"resCode": "40400", "resMessage": "User not found", "error": err})
	}

	// get User by id
	userModel, err := database.GetUserByUserId(usrId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"resCode": "50000", "resMessage": "Internal Server Error", "error": err})
	} else if userModel == nil || userModel.ID == 0  {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"resCode": "40400", "resMessage": "User not found"})
	} else {
		// Delete User by id
		err := database.DeleteUserByUserId(usrId)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"resCode": "50000", "resMessage": "Internal Server Error", "error": err})
		} else {
			return c.JSON(fiber.Map{
				"resCode":    "20000",
				"resMessage": fmt.Sprintf("Delete User %d Success !!", usrId)})
		}
	}

	
}

func UpdateUserProfile(c *fiber.Ctx) error {
	// validate usrId is int
	usrId, err := strconv.Atoi(c.Params("usrId"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"resCode": "40400", "resMessage": "User not found", "error": err})
	}

	// get User by id
	userModel, err := database.GetUserByUserId(usrId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"resCode": "50000", "resMessage": "Internal Server Error", "error": err})
	} else if userModel == nil || userModel.ID == 0  {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"resCode": "40400", "resMessage": "User not found"})
	} else {
		input := new(models.User)

		// Parse the request body
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"resCode": "40000", "resMessage": "Error on login request", "error": err})
		}

		if input.Name != "" {
			userModel.Name = input.Name
		}
		if input.Email != "" {
			userModel.Email = input.Email
		}

		userModel, err := database.UpdateUserByUserId(userModel)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"resCode": "50000", "resMessage": "Internal Server Error", "error": err})
		} else if userModel == nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"resCode": "40000", "resMessage": "Cannot create user", "error": err})
		} else {
			userData := models.UserData{
				ID:       userModel.ID,
				Username: userModel.Username,
				Name:     userModel.Name,
				Email:    userModel.Email,
			}
			return c.JSON(fiber.Map{
				"resCode":    "20000",
				"resMessage": "Update User Success !!",
				"data":       userData})
		}
	}
}
