package listing

import (
	"context"
	"time"

	"github.com/raafly/online-food-service/helper"
	"gorm.io/gorm"
)

type ProductService interface {
	GetAll() ([]Products, error)
	Create(req *Products) error
	GetById(productId string) (*Products, error)
	Delete(string) error
}

type ProductServiceImpl struct {
	port ProductRepository
	DB   *gorm.DB
}

func NewProductService(port ProductRepository, DB *gorm.DB) *ProductServiceImpl {
	return &ProductServiceImpl{
		port: port,
		DB:   DB,
	}
}

func (s *ProductServiceImpl) GetAll() ([]Products, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := s.port.GetAllProducts(ctx)
	if err != nil {
		return nil, helper.NewInternalServerError()
	}

	return response, nil
}

func (s *ProductServiceImpl) Create(req *Products) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := s.port.Create(ctx, req)
	if err != nil {
		return helper.NewInternalServerError()
	}
	return nil
}

func (s *ProductServiceImpl) GetById(productId string) (*Products, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.port.GetById(ctx, productId)
	if err != nil {
		return nil, helper.NewNotFoundError("id not found")
	}

	return res, nil
}

func (s *ProductServiceImpl) Delete(productId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := s.port.Delete(ctx, productId)
	if err != nil {
		return helper.NewNotFoundError("id not found")
	}

	return nil
}
