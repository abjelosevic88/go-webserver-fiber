package repositories

import (
	"fmt"
	"net/http"

	"github.com/abjelosevic88/go-fiber-postgres/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

var store = session.New(session.Config{
	CookieSecure: true,
})

type BookRepository struct {
	DB *gorm.DB
}

func InitBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		DB: db,
	}
}

func (br *BookRepository) SetupRoutes(rootAPI fiber.Router) {
	api := rootAPI.Group("/books")
	api.Post("/", br.CreateBook)
	api.Delete("/", br.DeleteBook)
	api.Get("/:id", br.GetBookByID)
	api.Get("/", br.GetBooks)
}

func (br *BookRepository) CreateBook(context *fiber.Ctx) error {
	book := models.Book{}

	if err := context.BodyParser(&book); err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Request failed!",
		})
		return err
	}

	if errors := book.Validate(); errors != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errors)
	}

	if err := br.DB.Create(&book).Error; err != nil {
		context.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Could not create book!",
		})
		return err
	}

	context.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Book has been added!",
	})

	return nil
}

func (br *BookRepository) DeleteBook(context *fiber.Ctx) error {
	bookModel := models.Book{}
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "ID cannot be empty!",
		})
		return nil
	}

	err := br.DB.Delete(&bookModel, id)

	if err.Error != nil {
		context.Status(http.StatusBadGateway).JSON(fiber.Map{
			"message": "Could not delete book!",
		})
		return err.Error
	}

	context.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Book successfully deleted!",
	})

	return nil
}

func (br *BookRepository) GetBooks(context *fiber.Ctx) error {
	sess, err := store.Get(context)

	if err != nil {
		panic(err)
	}

	sess.Set("name", "Alex123")

	bookModels := []models.Book{}
	err = br.DB.Find(&bookModels).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error getting books!",
		})
		return err
	}

	context.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Books fetched successfully!",
		"data":    bookModels,
	})

	// Save session
	if err := sess.Save(); err != nil {
		panic(err)
	}

	return nil
}

func (br *BookRepository) GetBookByID(context *fiber.Ctx) error {
	sess, err := store.Get(context)

	if err != nil {
		panic(err)
	}

	name := sess.Get("name")

	fmt.Println(name)

	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Id cannot be empty!",
		})
		return nil
	}

	bookModel := models.Book{}

	err = br.DB.Where("id = ?", id).First(&bookModel).Error

	if err != nil {
		context.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Book not found!",
		})
		return err
	}

	context.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Book fetched successfuly",
		"data":    bookModel,
	})

	return nil
}
