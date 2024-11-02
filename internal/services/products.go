package services

import (
	"github.com/midedickson/instashop/internal/dto"
	"github.com/midedickson/instashop/internal/entity"
)

type IProductService interface {
	CreateProduct(createProductPayload dto.CreateProductPayload) (*entity.Product, error)
	GetAllProducts() ([]*entity.Product, error)
	GetProductByID(id uint) (*entity.Product, error)
	UpdateProduct(id uint, updateProductPayload dto.UpdateProductPayload) (*entity.Product, error)
	DeleteProduct(id uint) error
}

type ProductService struct{}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (p *ProductService) CreateProduct(createProductPayload dto.CreateProductPayload) (*entity.Product, error) {
	product := &entity.Product{
		ID:       uint(1),
		Name:     createProductPayload.Name,
		Price:    createProductPayload.DecimalPrice(),
		Quantity: createProductPayload.Quantity,
	}
	// save the product to the database
	// ...

	return product, nil
}

func (p *ProductService) GetAllProducts() ([]*entity.Product, error) {
	// retrieve all products from the database
	//...

	products := []*entity.Product{}

	return products, nil
}
func (p *ProductService) GetProductByID(id uint) (*entity.Product, error) {
	// retrieve a product by its ID from the database
	//...

	product := &entity.Product{}

	return product, nil
}

func (p *ProductService) UpdateProduct(id uint, updateProductPayload dto.UpdateProductPayload) (*entity.Product, error) {
	// update a product by its ID with the provided data
	//...

	product, err := p.GetProductByID(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductService) DeleteProduct(id uint) error {
	return nil
}
