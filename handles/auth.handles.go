package handles

import (
	"time"

	"github.com/eakarach/AuthJWT/models"
	"github.com/eakarach/AuthJWT/config"
	"github.com/eakarach/AuthJWT/database"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

)

func Auth(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "Auth Success !!",
	})
}

func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	input := new(LoginInput)
	var userData models.UserData

	// Parse the request body
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"resCode": "40000", "resMessage": "Error on login request", "error": err})
	}

	username := input.Username
	pass := input.Password
	userModel, err := new(models.User), *new(error)

	// get User by Username
	userModel, err = database.GetUserByUsername(username)
	
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"resCode": "50000", "resMessage": "Internal Server Error", "error": err})
	} else if userModel == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"resCode": "40100", "resMessage": "Invalid username or password 1", "error": err})
	} else if pass != userModel.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"resCode": "40100", "resMessage": "Invalid username or password 2", "error": nil})
	} else {
		userData = models.UserData{
			ID:       userModel.ID,
			Username: userModel.Username,
			Email:    userModel.Email,
			Password: userModel.Password,
		}
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = userData.Username
	claims["user_id"] = userData.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"resCode": "20000", "resMessage": "Success login", "data": t})

}
