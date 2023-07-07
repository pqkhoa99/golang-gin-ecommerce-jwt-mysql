package service

import (
	"golang-jwttoken/data/request"
	"golang-jwttoken/model"
	"golang-jwttoken/repository"

	"github.com/go-playground/validator/v10"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	Validate          *validator.Validate
}

func NewProductServiceImpl(productRepository repository.ProductRepository, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		Validate:          validate,
	}
}

// CreateNewProduct implements ProductService.
func (p *ProductServiceImpl) CreateNewProduct(product request.CreateProductRequest) (string, error) {
	newproduct := model.Product{
		Name:  product.Name,
		Price: product.Price,
		Stock: product.Stock,
	}
	p.ProductRepository.Save(newproduct)
	return "", nil
}

// GetAllProduct implements ProductService.
func (p *ProductServiceImpl) GetAllProduct() ([]model.Product, error) {
	var products []model.Product = p.ProductRepository.FindAll()
	return products, nil
}

// GetProductById implements ProductService.
func (p *ProductServiceImpl) GetProductById(productId int) (model.Product, error) {
	product, err := p.ProductRepository.FindById(productId)
	if err != nil {
		return product, err
	}
	return product, nil
}

// GetProductByProductName implements ProductService.
func (p *ProductServiceImpl) GetProductByProductName(productName string) ([]model.Product, error) {
	products, err := p.ProductRepository.FindByProductName(productName)
	if err != nil {
		return products, err
	}
	return products, nil
}
