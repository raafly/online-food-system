package listing

import (
	"context"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(ctx context.Context, data *Products) error
	Delete(ctx context.Context, productId string) error
	GetAllProducts(ctx context.Context) ([]Products, error)
	GetById(ctx context.Context, productId string) (*Products, error)
}

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{db: db}
}

func (r *ProductRepositoryImpl) Create(ctx context.Context, data *Products) error {
	err := r.db.WithContext(ctx).Create(&data).Error
	return err
}

func (r *ProductRepositoryImpl) GetAllProducts(ctx context.Context) ([]Products, error) {
    var products []Products
    if err := r.db.Find(&products).Error; err != nil {
        return nil, err
    }
    return products, nil
}

func (r *ProductRepositoryImpl) GetById(ctx context.Context, productId string) (*Products, error) {
	var product Products
	err := r.db.WithContext(ctx).First(&Products{}, "id", productId).Error	
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepositoryImpl) Delete(ctx context.Context, productId string) error {
	err := r.db.Where("id = ?", productId).Delete(&Products{}).Error
	return err

}
