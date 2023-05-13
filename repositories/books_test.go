package repositories

import (
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/abjelosevic88/go-fiber-postgres/storage"
	"github.com/gofiber/fiber/v2"
)

func TestHelloName(t *testing.T) {
	// http.Request
	req := httptest.NewRequest("GET", "http://localhost/api/books", nil)
	req.Header.Set("X-Custom-Header", "hi")

	app := fiber.New()

	db, _ := storage.NewConnection(&storage.Config{
		Host:     "localhost",
		Port:     "5432",
		Password: "root",
		User:     "root",
		SSLMode:  "disable",
		DBName:   "postgres",
	})

	SetupRoutes(app, db)

	// http.Response
	resp, _ := app.Test(req)

	// Do something with results:
	if resp.StatusCode == fiber.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body)) // => Hello, World!
	}
}
