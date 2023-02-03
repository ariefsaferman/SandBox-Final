package repository

import (
	"database/sql"
	"errors"
	"go-basic-server/entity"
)

type ProductRepository interface {
	GetOneProductById(id int) (*entity.Product, error)
	CreateProduct(p *entity.Product) error
	GetAllProduct() []*entity.Product
	UpdateProduct(id int, p *entity.Product) error
	DeleteProduct(id int) error
}

type ProductRepositoryImpl struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &ProductRepositoryImpl{
		db: db,
	}
}

func (r *ProductRepositoryImpl) CreateProduct(p *entity.Product) error {
	q := `INSERT INTO product(
		name, description)
		VALUES ($1, $2);`
	err := r.db.QueryRow(q, p.Name, p.Description).Err()
	if err != nil {
		return errors.New("something wrong in create product")
	}

	return nil
}

func (r *ProductRepositoryImpl) GetAllProduct() []*entity.Product {
	q := `SELECT id, name, description FROM product;`
	rows, err := r.db.Query(q)
	var plist []*entity.Product
	for rows.Next() {
		var p entity.Product
		err := rows.Scan(&p.Id, &p.Name, &p.Description)
		if err != nil {
			panic("error: " + err.Error())
		}
		plist = append(plist, &p)
	}
	if err != nil {
		panic("error: " + err.Error())
	}

	return plist
}

func (r *ProductRepositoryImpl) GetOneProductById(id int) (*entity.Product, error) {
	var p entity.Product
	q := `SELECT pc.id, pc.name, pc.description FROM product pc WHERE pc.id =$1;`
	err := r.db.QueryRow(q, id).Scan(&p.Id, &p.Name, &p.Description)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProductRepositoryImpl) UpdateProduct(id int, p *entity.Product) error {
	q := `UPDATE product SET name = $2, description = $3, quantity = $4 WHERE id = $1;`
	err := r.db.QueryRow(q, id, p.Name, p.Description, p.Quantity).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepositoryImpl) DeleteProduct(id int) error {
	q := `DELETE FROM product WHERE id = $1;`
	err := r.db.QueryRow(q, id).Err()
	if err != nil {
		return err
	}
	return err
}
