package services

import (
	"github.com/midedickson/instashop/database/models"
	"github.com/midedickson/instashop/internal/dto"
	"github.com/midedickson/instashop/internal/entity"
	"github.com/midedickson/instashop/utils"
)

type ProductRepository interface {
	GetAllProducts() ([]*models.Product, error)
	GetProductByID(id uint) (*models.Product, error)
	CreateProduct(createProductPayload dto.CreateUpdateDBProduct) (*models.Product, error)
	UpdateProduct(updateProductPayload dto.CreateUpdateDBProduct, id uint) (*models.Product, error)
	DeleteProduct(id uint) error
}

type ProductService struct {
	productRepository ProductRepository
}

func NewProductService(productRepository ProductRepository) *ProductService {
	return &ProductService{productRepository: productRepository}
}

func (p *ProductService) CreateProduct(createProductPayload dto.CreateProductPayload) (*entity.Product, error) {
	createDBProduct := dto.CreateUpdateDBProduct(createProductPayload)
	product, err := p.productRepository.CreateProduct(createDBProduct)
	if err != nil {
		return nil, err
	}

	return product.ToEntity(), nil
}

func (p *ProductService) GetAllProducts() ([]*entity.Product, error) {
	// retrieve all products from the database
	products, err := p.productRepository.GetAllProducts()
	if err != nil {
		return nil, err
	}

	productEntities := utils.MapConcurrent(products, func(product *models.Product) *entity.Product {
		return product.ToEntity()
	})

	return productEntities, nil
}
func (p *ProductService) GetProductByID(id uint) (*entity.Product, error) {
	// retrieve a product by its ID from the database
	product, err := p.productRepository.GetProductByID(id)
	if err != nil {
		return nil, err
	}

	return product.ToEntity(), nil
}

func (p *ProductService) UpdateProduct(id uint, updateProductPayload dto.UpdateProductPayload) (*entity.Product, error) {
	// update a product by its ID with the provided data
	createDBProduct := dto.CreateUpdateDBProduct(updateProductPayload)
	product, err := p.productRepository.UpdateProduct(createDBProduct, id)
	if err != nil {
		return nil, err
	}

	return product.ToEntity(), nil
}

func (p *ProductService) DeleteProduct(id uint) error {
	return p.productRepository.DeleteProduct(id)
}
