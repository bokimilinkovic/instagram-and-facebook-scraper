package service

import (
	"holycode-task/controller/dto"
	"holycode-task/model"
	"holycode-task/repository/postgres"
)

type ProductService struct {
	store *postgres.Store
}

func NewProductService(store *postgres.Store) *ProductService {
	return &ProductService{store: store}
}

func (ps *ProductService) FindAll() ([]model.Product, error) {
	return ps.store.FindAllProducts()
}

func (ps *ProductService) CreateNew(product *dto.ProductDto) (*model.Product, error) {
	toCreate := &model.Product{
		Price:       product.Price,
		Name:        product.Name,
		Description: product.Description,
		Sponsor:     product.Sponsor,
	}

	err := ps.store.CreateProduct(toCreate)
	if err != nil {
		return nil, err
	}

	return toCreate, nil
}
