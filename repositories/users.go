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
	api := rootAPI.Group("/users", middlewares.Auth)
	api.Get("/", ur.GetUserData)
}

func (ur *UserRepository) GetUserData(ctx *fiber.Ctx) error {
	user := ctx.Locals("user")

	ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User found!",
		"data":    user,
	})

	return nil
}
