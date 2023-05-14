package middlewares

import (
	"net/http"

	"github.com/abjelosevic88/go-fiber-postgres/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Auth(c *fiber.Ctx, db *gorm.DB) error {
	if c.Query("username") == "" || c.Query("password") == "" {
		return fiber.NewError(http.StatusBadRequest, "Missing credentials!")
	}

	userModel := models.User{}

	if err := db.Where("username = ?", c.Query("username")).First(&userModel).Error; err != nil {
		return fiber.NewError(http.StatusBadRequest, "User not found!")
	}

	match := checkPasswordHash(c.Query("password"), *userModel.Password)

	if !match {
		return fiber.NewError(http.StatusBadRequest, "Wrong password!")
	}

	c.Locals("user", &userModel)

	return c.Next()
}
