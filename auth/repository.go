package auth

import (
	"context"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	create(ctx context.Context, req *Customers) error
	findByEmail(ctx context.Context, req *Customers) (*Customers, error)
	findById(ctx context.Context, id string) (*Customers, error)
}

type CustomerRepositoryImpl struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepositoryImpl {
	return &CustomerRepositoryImpl{db: db}
}

func (r *CustomerRepositoryImpl) create(ctx context.Context, req *Customers) error {
	err := r.db.WithContext(ctx).Create(&req).Error
	return err
}

func (r *CustomerRepositoryImpl) findByEmail(ctx context.Context, req *Customers) (*Customers, error) {
	var customer Customers
	err := r.db.WithContext(ctx).Where("email = ?", req.Email).Take(&customer).Error
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (r *CustomerRepositoryImpl) findById(ctx context.Context, id string) (*Customers, error) {
	var customer Customers
	err := r.db.WithContext(ctx).Where("id = ?", id).Take(&customer).Error
	if err != nil {
		return nil, err
	}

	return &customer, nil
}