package order

import (
	"context"

	"github.com/raafly/online-food-service/helper"
	"gorm.io/gorm"
)

type OrderRepository interface {
	create(ctx context.Context, req *Order) error
	getOrderByID(ctx context.Context, id int) (*Order, error)
	updateQuantity(ctx context.Context, id, quantity int) error
	cancel(ctx context.Context, id int) error 
}

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{db: db}
}

func (r *OrderRepositoryImpl) create(ctx context.Context, req *Order) error {
	return r.db.WithContext(ctx).Create(&req).Error
}

func (r *OrderRepositoryImpl) getOrderByID(ctx context.Context, id int) (*Order, error) {
	var order Order
	err := r.db.WithContext(ctx).Where("id = ?", id).Take(&order).Error
	if err != nil {
		return nil, helper.NewNotFoundError("id order not found")
	}

	return &order, nil
}
func (r *OrderRepositoryImpl) updateQuantity(ctx context.Context, id, quantity int) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Update("quantity", quantity).Error
}
func (r *OrderRepositoryImpl) cancel(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(id).Error
}
