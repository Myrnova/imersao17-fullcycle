package service

import (
	"github.com/devfullcycle/imersao17/goapi/internal/database"
	"github.com/devfullcycle/imersao17/goapi/internal/entity"
)

type ProductService struct {
	ProductDB database.ProductDB
}

func NewProductService(productDB database.ProductDB) *ProductService {
	return &ProductService{ProductDB: productDB}
}

func (productService *ProductService) GetProducts() ([]*entity.Product, error) {
	products, err := productService.ProductDB.GetProducts()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (productService *ProductService) CreateProduct(name, description, categoryID, imageURL string, price float64) (*entity.Product, error) {
	product := entity.NewProduct(name, description, categoryID, imageURL, price)

	_, err := productService.ProductDB.CreateProduct(product)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (productService *ProductService) GetProduct(id string) (*entity.Product, error) {
	product, err := productService.ProductDB.GetProductByID(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (productService *ProductService) GetProductByCategoryID(categoryID string) ([]*entity.Product, error) {
	products, err := productService.ProductDB.GetProductByCategoryID(categoryID)

	if err != nil {
		return nil, err
	}

	return products, nil
}
