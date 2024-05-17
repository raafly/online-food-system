package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raafly/online-food-service/helper"
	"gorm.io/gorm"
)

func NewCustomerRoutes(app *fiber.App, db *gorm.DB) {
	pass := helper.NewPassword()

	repo := NewCustomerRepository(db)
	serv := NewCustomerService(repo, db, pass)
	handler := NewCustomerHandler(serv)

	cs := app.Group("/customer")
	cs.Post("/register", handler.Register)
	cs.Post("/login", handler.Login)
	cs.Get("/information", handler.FindById)
}