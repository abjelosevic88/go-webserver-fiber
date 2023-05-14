package main

import (
	"log"
	"os"

	"github.com/abjelosevic88/go-fiber-postgres/models"
	"github.com/abjelosevic88/go-fiber-postgres/repositories"
	"github.com/abjelosevic88/go-fiber-postgres/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes), err
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := storage.NewConnection(&storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	})
	if err != nil {
		log.Fatal(err)
	}

	err = models.Migrate(db)
	if err != nil {
		log.Fatal("Could not migrate db!")
	}

	app := fiber.New()

	app.Use(requestid.New())
	repositories.SetupRoutes(app, db)

	app.Listen(":8080")
}
