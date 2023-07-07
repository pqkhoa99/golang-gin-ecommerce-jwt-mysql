package service

import (
	"golang-jwttoken/data/request"
	"golang-jwttoken/model"
)

type ProductService interface {
	CreateNewProduct(product request.CreateProductRequest) (string, error)
	GetAllProduct() ([]model.Product, error)
	GetProductById(productId int) (model.Product, error)
	GetProductByProductName(productName string) ([]model.Product, error)
}