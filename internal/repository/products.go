package repository

import (
	"github.com/midedickson/instashop/constants"
	"github.com/midedickson/instashop/database/models"
	"github.com/midedickson/instashop/internal/dto"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) GetAllProducts() ([]*models.Product, error) {
	var products []*models.Product
	err := repo.db.Find(&products).Error
	return products, err
}

func (repo *ProductRepository) GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	err := repo.db.Where("id =?", id).First(&product).Error
	return &product, err
}

func (repo *ProductRepository) CreateProduct(createProductPayload dto.CreateUpdateDBProduct) (*models.Product, error) {
	product := &models.Product{
		Name:     createProductPayload.Name,
		Price:    constants.Money(createProductPayload.Price),
		Quantity: createProductPayload.Quantity,
	}

	err := repo.db.Create(product).Error
	return product, err
}

func (repo *ProductRepository) UpdateProduct(updateProductPayload dto.CreateUpdateDBProduct, id uint) (*models.Product, error) {
	var product models.Product
	err := repo.db.Where("id =?", id).First(&product).Error
	if err != nil {
		return nil, err
	}

	product.Name = updateProductPayload.Name
	product.Price = constants.Money(updateProductPayload.Price)
	product.Quantity = updateProductPayload.Quantity

	err = repo.db.Save(&product).Error
	return &product, err
}

func (repo *ProductRepository) DeleteProduct(id uint) error {
	return repo.db.Delete(&models.Product{}, id).Error
}
