package repository

import (
	"errors"
	"golang-jwttoken/data/request"
	"golang-jwttoken/helper"
	"golang-jwttoken/model"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	Db *gorm.DB
}

func NewProductRepositoryImpl(Db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{Db: Db}
}

// Delete implements ProductRepository.
func (p *ProductRepositoryImpl) Delete(productId int) {
	var product model.Product
	result := p.Db.Where("id = ?", productId).Delete(&product)
	helper.ErrorPanic(result.Error)
}

// FindAll implements ProductRepository.
func (p *ProductRepositoryImpl) FindAll() []model.Product {
	var products []model.Product
	result := p.Db.Find(&products)
	helper.ErrorPanic(result.Error)
	return products
}

// FindById implements ProductRepository.
func (p *ProductRepositoryImpl) FindById(productId int) (model.Product, error) {
	var product model.Product
	result := p.Db.Find(&product, productId)
	if result != nil {
		return product, nil
	} else {
		return product, errors.New("product is not found")
	}
}

// FindByProductName implements ProductRepository.
func (p *ProductRepositoryImpl) FindByProductName(productName string) ([]model.Product, error) {
	var product []model.Product
	result := p.Db.Find(&product, productName)
	if result != nil {
		return product, nil
	} else {
		return product, errors.New("product is not found")
	}
}

// Save implements ProductRepository.
func (p *ProductRepositoryImpl) Save(product model.Product) {
	result := p.Db.Create(&product)
	helper.ErrorPanic(result.Error)
}

// Update implements ProductRepository.
func (p *ProductRepositoryImpl) Update(product model.Product) {
	var updateProduct = request.UpdateProductRequest{
		Id: product.Id,
		Name: product.Name,
		Price: product.Price,
		Stock: product.Stock,
	}
	result := p.Db.Model(&product).Updates(updateProduct)
	helper.ErrorPanic(result.Error)
}


