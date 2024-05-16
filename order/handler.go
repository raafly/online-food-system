package order

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/raafly/online-food-service/helper"
)

type OrderHandler interface {
	Create(c *fiber.Ctx) error
	GetOrderByID(c *fiber.Ctx) error
	UpdateQuantity(c *fiber.Ctx) error
	Cancel(c *fiber.Ctx) error
}

type OrderHandlerImpl struct {
	port OrderService
}

func NewOrderHandler(port OrderService) OrderHandler {
	return &OrderHandlerImpl{port: port}
}

func (h *OrderHandlerImpl) Create(c *fiber.Ctx) error {
	req := new(Order)
	c.BodyParser(req)

	err := h.port.create(req)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	return c.Status(201).JSON(helper.NewCreated())
}
func (h *OrderHandlerImpl) GetOrderByID(c *fiber.Ctx) error {
	seed := c.Params("id")
	id,_ := strconv.Atoi(seed)

	resp, err := h.port.getOrderByID(id)
	if err != nil {
		return c.Status(404).JSON(helper.NewNotFoundError("order id not found"))
	}

	return c.Status(200).JSON(helper.NewSuccess(resp))
}
func (h *OrderHandlerImpl) UpdateQuantity(c *fiber.Ctx) error {
	req := new(UpdateQuantity)
	c.BodyParser(req)

	err := h.port.updateQuantity(req.ID, req.Quantity)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	return c.Status(200).JSON(helper.NewSuccess("success update quantity"))
}

func (h *OrderHandlerImpl) Cancel(c *fiber.Ctx) error {
	seed := c.Params("id")
	id,_ := strconv.Atoi(seed)

	err := h.port.cancel(id)
	if err != nil {
		return c.Status(404).JSON(helper.NewNotFoundError("order not found"))
	}

	return c.Status(200).JSON(helper.NewSuccess("success delete order"))
}
