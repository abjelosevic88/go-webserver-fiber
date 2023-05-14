package repositories

import (
	"net/http"

	"github.com/abjelosevic88/go-fiber-postgres/middlewares"
	"github.com/gofiber/fiber/v2"
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
	api := rootAPI.Group("/users", func(c *fiber.Ctx) error {
		return middlewares.Auth(c, ur.DB)
	})
	api.Get("/", ur.GetUserData)
}

func (ur *UserRepository) GetUserData(context *fiber.Ctx) error {
	user := context.Locals("user")

	context.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User found!",
		"data":    user,
	})

	return nil
}
