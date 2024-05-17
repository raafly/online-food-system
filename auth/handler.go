package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raafly/online-food-service/helper"
)

type CustomerHandler interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
}


type CustomerHandlerImpl struct {
	port CustomerService
}

func NewCustomerHandler(port CustomerService) *CustomerHandlerImpl {
	return &CustomerHandlerImpl{
		port: port,
	}
}

func (h *CustomerHandlerImpl) Register(c *fiber.Ctx) error {
	req := new(Customers)
	c.BodyParser(req)

	err := h.port.register(req)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	return c.Status(201).JSON(helper.NewCreated())
}

func (h *CustomerHandlerImpl) Login(c *fiber.Ctx) error {
	req := new(Customers)
	c.BodyParser(req)

	err := h.port.login(req)
	if err != nil {
		return c.Status(401).JSON(err)
	}

	return c.Status(200).JSON(helper.NewSuccess(nil))
}

func (h *CustomerHandlerImpl) FindById(c *fiber.Ctx) error {
	id := c.Params("userId")

	resp, err := h.port.findById(id)
	if err != nil {
		return c.Status(404).JSON(err)
	}

	return c.Status(200).JSON(helper.NewSuccess(resp))
}