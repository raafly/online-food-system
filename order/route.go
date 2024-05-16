package order

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewOrderRoutes(app *fiber.App, db *gorm.DB) {
	repo := NewOrderRepository(db)
	serv := NewOrderService(repo)
	handler := NewOrderHandler(serv)

	or := app.Group("/order")
	or.Post("/new", handler.Create)
	or.Get("/find", handler.GetOrderByID)
	or.Delete("/cancel", handler.Cancel)
	or.Put("/update", handler.UpdateQuantity)
}