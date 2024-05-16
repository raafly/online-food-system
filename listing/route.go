package listing

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewProductRoutes(app *fiber.App, db *gorm.DB) {
	repo := NewProductRepository(db)
	serv := NewProductService(repo, db)
	handler := NewProductHandler(serv)

	product := app.Group("/products")
	product.Post("/create", handler.Create)
	product.Get("/find-product", handler.GetById)
	product.Get("/all-product", handler.GetAll)
	product.Delete("/deletes", handler.Delete)
}