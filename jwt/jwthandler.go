package jwt

import (
	"log"
	"time"

	"mm-api/common"
	"mm-api/database"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func JWTHandler(c *fiber.Ctx) error {
	hardwareID := c.FormValue("hardwareID")
	timestamp := c.FormValue("timestamp")
	claims := jwt.MapClaims{
		"admin": false,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(common.JWTKey))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	data := database.UserInfo{
		HardwareID:     hardwareID,
		LoginTimestamp: timestamp,
	}

	err = database.AddUserUpdate(data)
	if err != nil {
		log.Println(err)
		return c.Status(501).JSON(fiber.Map{"success": false, "message": "Error updating user in database"})
	}

	return c.Status(200).JSON(fiber.Map{"success": true, "jwtToken": t})
}