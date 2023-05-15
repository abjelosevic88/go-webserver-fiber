package repositories

import (
	"net/http"

	"github.com/abjelosevic88/go-fiber-postgres/middlewares"
	"github.com/abjelosevic88/go-fiber-postgres/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func InitUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (ur *UserRepository) SetupRoutes(rootAPI fiber.Router) {
	api := rootAPI.Group("/users", middlewares.Auth)
	api.Get("/", ur.GetUserData)
	api.Get("/books", ur.GetAllBooks)
}

func (ur *UserRepository) GetUserData(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)

	ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User found!",
		"data":    user,
	})

	return nil
}

func (ur *UserRepository) GetAllBooks(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	var userModel models.User

	if err := ur.DB.Model(&models.User{}).Preload("Books").Find(&userModel, claims["subID"]).Error; err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error getting books!",
		})
	}

	books := make([]interface{}, len(userModel.Books))

	for i, val := range userModel.Books {
		books[i] = map[string]interface{}{
			"id":        val.ID,
			"author":    *val.Author,
			"title":     *val.Title,
			"publisher": *val.Publisher,
		}
	}

	ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Books found!",
		"data":    books,
	})

	return nil
}
