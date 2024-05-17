package listing

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raafly/online-food-service/helper"
)

type ProductHandler interface {
	Create(c *fiber.Ctx) error 
	GetAll(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error 
	Delete(c *fiber.Ctx) error
}

type ProductHandlerImpl struct {
	port ProductService
}

func NewProductHandler(port ProductService) *ProductHandlerImpl {
	return &ProductHandlerImpl{
		port: port,
	}
}

func (h *ProductHandlerImpl) Create(c *fiber.Ctx) error {
	req := new(Products)
	c.BodyParser(req)
	
	err := h.port.Create(req)
	if err != nil {
		return c.Status(400).JSON(err)
	}

	return c.Status(201).JSON(helper.NewCreated())
}

func (h *ProductHandlerImpl) GetById(c *fiber.Ctx) error {
	id := c.Params("id")

	response, err := h.port.GetById(id)
	if err != nil {
		return c.Status(404).JSON(helper.NewNotFoundError("id product not found"))
	}

	return c.Status(200).JSON(helper.NewSuccess(response))
}

func (h *ProductHandlerImpl) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.port.Delete(id)
	if err != nil {
		return c.Status(404).JSON(helper.NewNotFoundError("id product not found"))
	}

	return c.Status(200).JSON(helper.NewSuccess(nil))		
}

func (h *ProductHandlerImpl) GetAll(c *fiber.Ctx) error  {
	response, err := h.port.GetAll()
	if err != nil {
		return c.Status(500).JSON(helper.NewInternalServerError())
	}

	return c.Status(200).JSON(helper.NewSuccess(response))
}