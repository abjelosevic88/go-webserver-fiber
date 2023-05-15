package repositories

import (
	"net/http"
	"os"
	"time"

	"github.com/abjelosevic88/go-fiber-postgres/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepository struct {
	DB *gorm.DB
}

func InitAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		DB: db,
	}
}

func (ar *AuthRepository) SetupRoutes(rootAPI fiber.Router) {
	api := rootAPI.Group("/auth")
	api.Post("/login", ar.Login)
}

func (ar *AuthRepository) Login(ctx *fiber.Ctx) error {
	body := struct {
		Username string `json:"username" form:"username"`
		Password string `json:"password" form:"password"`
	}{}

	if err := ctx.BodyParser(&body); err != nil {
		return err
	}

	if body.Username == "" || body.Password == "" {
		return fiber.NewError(http.StatusBadRequest, "Missing credentials!")
	}

	userModel := models.User{}

	if err := ar.DB.Where("username = ?", body.Username).First(&userModel).Error; err != nil {
		return fiber.NewError(http.StatusBadRequest, "User not found!")
	}

	match := checkPasswordHash(body.Password, *userModel.Password)

	if !match {
		return fiber.NewError(http.StatusBadRequest, "Wrong password!")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   userModel.Username,
		"subID": userModel.ID,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})
	jwtToken, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))

	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Successfully logged in!",
		"jwt":     jwtToken,
	})

	return nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
