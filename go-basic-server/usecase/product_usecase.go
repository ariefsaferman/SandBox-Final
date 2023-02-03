package usecase

import (
	"errors"
	"go-basic-server/entity"
	"go-basic-server/repository"
	"go-basic-server/util"
)

type ProductUsecase interface {
	GetOneProductById(id int) (*entity.Product, error)
	CreateProduct(p *entity.Product) error
	GetAllProduct() ([]*entity.Product, error)
	UpdateProduct(id int, p *entity.Product) error
	DeleteProduct(id int) error
}

type ProductUsecaseImpl struct {
	repository repository.ProductRepository
}

func NewProductUsecase(r repository.ProductRepository) ProductUsecase {
	return &ProductUsecaseImpl{
		repository: r,
	}
}

func (u *ProductUsecaseImpl) GetOneProductById(id int) (*entity.Product, error) {
	p, err := u.repository.GetOneProductById(id)
	if err != nil {
		return nil, err
	}

	if p == nil {
		return nil, util.ErrProductNotFound
	}

	return p, nil
}

func (u *ProductUsecaseImpl) CreateProduct(p *entity.Product) error {
	err := u.repository.CreateProduct(p)
	return err
}

func (u *ProductUsecaseImpl) GetAllProduct() ([]*entity.Product, error) {
	p := u.repository.GetAllProduct()
	if p == nil {
		return []*entity.Product{}, errors.New("product is empty")
	}
	return p, nil
}

func (u *ProductUsecaseImpl) UpdateProduct(id int, p *entity.Product) error {
	err := u.repository.UpdateProduct(id, p)
	return err
}

func (u *ProductUsecaseImpl) DeleteProduct(id int) error {
	err := u.repository.DeleteProduct(id)
	return err
}
