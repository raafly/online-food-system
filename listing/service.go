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
	ProductRepository ProductRepository
	DB                *gorm.DB
}

func NewProductService(productRepository ProductRepository, DB *gorm.DB) *ProductServiceImpl {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		DB:                DB,
	}
}

func (ser *ProductServiceImpl) GetAll() ([]Products, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := ser.ProductRepository.GetAllProducts(ctx)
	if err != nil {
		return nil, helper.NewInternalServerError()
	}

	return response, nil
}

func (ser *ProductServiceImpl) Create(req *Products) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := ser.ProductRepository.Create(ctx, req)
	if err != nil {
		return helper.NewInternalServerError()
	}
	return nil
}

func (ser *ProductServiceImpl) GetById(productId string) (*Products, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := ser.ProductRepository.GetById(ctx, productId)
	if err != nil {
		return nil, helper.NewNotFoundError()
	}

	return res, nil
}

func (ser *ProductServiceImpl) Delete(productId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := ser.ProductRepository.Delete(ctx, productId)
	if err != nil {
		return helper.NewNotFoundError()
	}

	return nil
}