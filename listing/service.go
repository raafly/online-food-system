package listing

import (
	"context"
	"database/sql"
	"time"

	"github.com/raafly/online-food-service/helper"
)

type ProductService interface {
	GetAll(ctx context.Context) ([]Products, error)
	Create(ctx context.Context, req *RequestProduct) error
	GetById(ctx context.Context, productId int) (*ResponseProduct, error)
	Delete(ctx context.Context, productInt int) error
}

type ProductServiceImpl struct {
	ProductRepository ProductRepository
	DB                *sql.DB
}

func NewProductService(productRepository ProductRepository, DB *sql.DB) *ProductServiceImpl {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		DB:                DB,
	}
}

func (ser *ProductServiceImpl) GetAll(ctx context.Context) ([]Products, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := ser.ProductRepository.GetAllProducts(ctx)
	if err != nil {
		return nil, helper.NewInternalServerError()
	}

	return response, nil
}

func (ser *ProductServiceImpl) Create(ctx context.Context, req *Products) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := ser.ProductRepository.Create(ctx, req)
	if err != nil {
		return helper.NewInternalServerError()
	}
	return nil
}

func (ser *ProductServiceImpl) GetById(ctx context.Context, productId int) (*Products, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := ser.ProductRepository.GetById(ctx, productId)
	if err != nil {
		return nil, helper.NewNotFoundError()
	}

	return res, nil
}

func (ser *ProductServiceImpl) Delete(ctx context.Context, productId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := ser.ProductRepository.Delete(ctx, productId)
	if err != nil {
		return helper.NewNotFoundError()
	}

	return nil

}
