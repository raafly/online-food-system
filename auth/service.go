package auth

import (
	"context"
	"log"
	"time"

	"github.com/raafly/online-food-service/helper"
	"gorm.io/gorm"
)

type CustomerService interface {
	register(req *Customers) error
	login(req *Customers) error
	findById(id string) (*Customers, error)
}

type CustomerServiceImpl struct {
	port CustomerRepository
	DB *gorm.DB
	pass *helper.Password
}

func NewCustomerService(port CustomerRepository, DB *gorm.DB, pass *helper.Password) *CustomerServiceImpl {
	return &CustomerServiceImpl {
		port: port,
		DB: DB,
		pass: pass,
	}
}

func (s *CustomerServiceImpl) register(req *Customers ) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5 *time.Second)
	defer cancel()

	req.Password = s.pass.HashPassword(req.Password)

	err := s.port.create(ctx, req)
	if err != nil {
		log.Println(err)
		return helper.NewInternalServerError()
	}

	return nil
}

func (s *CustomerServiceImpl) login(req *Customers) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := s.port.findByEmail(ctx, req)
	if err != nil {
		return helper.NewNotFoundError("email not founded")
	}

	err = s.pass.CompareHashAndPassword(resp.Password, req.Password)
	if err != nil {
		return helper.NewBadRequestError("incorect email and password")
	}

	return nil
}

func (s *CustomerServiceImpl) findById(id string) (*Customers, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5 *time.Second)
	defer cancel()

	customer, err := s.port.findById(ctx, id)
	if err != nil {
		return nil, helper.NewNotFoundError("account not found")
	}

	return customer, nil
}