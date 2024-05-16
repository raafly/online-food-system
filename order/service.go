package order

import (
	"context"
	"log"
	"time"

	"github.com/raafly/online-food-service/helper"
)

type OrderService interface {
	create(req *Order) error
	getOrderByID(id int) (*Order, error)
	updateQuantity(id, quantity int) error
	cancel(id int) error
}

type OrderServiceImpl struct {
	port OrderRepository
}

func NewOrderService(port OrderRepository) OrderService {
	return &OrderServiceImpl{port: port}
}

func (s *OrderServiceImpl) create(req *Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := s.port.create(ctx, req)
	if err != nil {
		log.Println(err)
		return helper.NewBadRequestError("input right data")
	}

	return nil
}


func (s *OrderServiceImpl) getOrderByID(id int) (*Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := s.port.getOrderByID(ctx, id)
	if err != nil {
		log.Println(err)
		return nil, helper.NewNotFoundError("order not found")
	}

	return resp, nil
}

func (s *OrderServiceImpl )updateQuantity(id, quantity int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := s.port.updateQuantity(ctx, id, quantity)
	if err != nil {
		log.Println(err)
		return helper.NewInternalServerError()
	}

	return nil
}
func (s *OrderServiceImpl) cancel(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := s.port.cancel(ctx, id)
	if err != nil {
		return helper.NewNotFoundError("order not found")
	}

	return nil
}
