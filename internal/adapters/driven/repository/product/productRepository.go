package product

import (
	"context"
	l "fiap-hf-src/infrastructure/logger"
	"fiap-hf-src/internal/core/db"
	"fiap-hf-src/internal/core/domain/entity"
	"fiap-hf-src/internal/core/domain/valueObject"
)

var (
	queryGetProductByID       = `SELECT * FROM product where id = $1`
	queryGetProductByCategory = `SELECT * FROM product where category = $1`
	queryDeleteProduct        = `DELETE FROM product where id = $1 RETURNING id`
	querySaveProduct          = `INSERT INTO product (id, name, category, image, description, price, created_at, deactivated_at) VALUES (DEFAULT, $1, $2, $3, $4, $5, now(), NULL) RETURNING id, created_at`
	queryUpdateProduct        = `UPDATE product SET name = $1, category = $2, image = $3, description = $4, price = $5, deactivated_at = $6 WHERE id = $7 RETURNING id, created_at`
)

type ProductRepository interface {
	SaveProduct(product entity.Product) (*entity.Product, error)
	UpdateProductByID(id int64, product entity.Product) (*entity.Product, error)
	GetProductByID(id int64) (*entity.Product, error)
	GetProductByCategory(category string) ([]entity.Product, error)
	DeleteProductByID(id int64) error
}

type productRepository struct {
	Ctx      context.Context
	Database db.SQLDatabase
}

func NewProductRepository(ctx context.Context, db db.SQLDatabase) ProductRepository {
	return productRepository{Ctx: ctx, Database: db}
}

func (p productRepository) SaveProduct(product entity.Product) (*entity.Product, error) {
	l.Infof("SaveProduct received input: ", " | ", product)
	if err := p.Database.Connect(); err != nil {
		l.Errorf("SaveProduct connect error: ", " | ", err)
		return nil, err
	}

	defer p.Database.Close()

	if err := p.Database.PrepareStmt(querySaveProduct); err != nil {
		l.Infof("SaveProduct prepare error: ", " | ", err)
		return nil, err
	}

	defer p.Database.CloseStmt()

	var outProduct = &entity.Product{
		Name: product.Name,
		Category: valueObject.Category{
			Value: product.Category.Value,
		},
		Image:       product.Image,
		Description: product.Description,
		Price:       product.Price,
	}

	p.Database.QueryRow(product.Name, product.Category.Value, product.Image, product.Description, product.Price)

	if err := p.Database.ScanStmt(&outProduct.ID, &outProduct.CreatedAt.Value); err != nil {
		l.Infof("SaveProduct scan error: ", " | ", err)
		return nil, err
	}
	l.Infof("SaveProduct output: ", " | ", outProduct.MarshalString())
	return outProduct, nil
}

func (p productRepository) UpdateProductByID(id int64, product entity.Product) (*entity.Product, error) {
	l.Infof("UpdateProductByID received input: ", " | ", product)
	if err := p.Database.Connect(); err != nil {
		l.Errorf("UpdateProductByID connect error: ", " | ", err)
		return nil, err
	}

	defer p.Database.Close()

	if err := p.Database.PrepareStmt(queryUpdateProduct); err != nil {
		l.Errorf("UpdateProductByID prepare error: ", " | ", err)
		return nil, err
	}

	defer p.Database.CloseStmt()

	var outProduct = &entity.Product{
		Name: product.Name,
		Category: valueObject.Category{
			Value: product.Category.Value,
		},
		Image:       product.Image,
		Description: product.Description,
		Price:       product.Price,
		DeactivatedAt: valueObject.DeactivatedAt{
			Value: product.DeactivatedAt.Value,
		},
	}

	p.Database.QueryRow(product.Name, product.Category.Value, product.Image, product.Description, product.Price, product.DeactivatedAt.Value, id)

	if err := p.Database.ScanStmt(&outProduct.ID, &outProduct.CreatedAt.Value); err != nil {
		l.Errorf("UpdateProductByID scan error: ", " | ", err)
		return nil, err
	}

	l.Infof("UpdateProductByID output: ", " | ", outProduct.MarshalString())
	return outProduct, nil
}

func (p productRepository) GetProductByID(id int64) (*entity.Product, error) {
	l.Infof("GetProductByID received input: ", " | ", id)
	if err := p.Database.Connect(); err != nil {
		l.Errorf("GetProductByID connect error: ", " | ", err)
		return nil, err
	}

	defer p.Database.Close()

	var outProduct = new(entity.Product)

	if err := p.Database.Query(queryGetProductByID, id); err != nil {
		l.Errorf("GetProductByID error to connect database: ", " | ", err)
		return nil, err
	}

	for p.Database.GetNextRows() {
		err := p.Database.Scan(
			&outProduct.ID,
			&outProduct.Name,
			&outProduct.Category.Value,
			&outProduct.Image,
			&outProduct.Description,
			&outProduct.Price,
			&outProduct.CreatedAt.Value,
			&outProduct.DeactivatedAt.Value,
		)
		if err != nil {
			l.Errorf("GetProductByID error to scan database: ", " | ", err)
			return nil, err
		}
	}

	if *outProduct == (entity.Product{}) {
		return nil, nil
	}

	return outProduct, nil
}

func (p productRepository) GetProductByCategory(category string) ([]entity.Product, error) {
	l.Infof("GetProductByCategory received input: ", " | ", category)
	if err := p.Database.Connect(); err != nil {
		l.Errorf("GetProductByCategory connect error: ", " | ", err)
		return nil, err
	}

	defer p.Database.Close()

	var (
		product     = new(entity.Product)
		productList = make([]entity.Product, 0)
	)

	if err := p.Database.Query(queryGetProductByCategory, category); err != nil {
		l.Errorf("GetProductByCategory error to connect database: ", " | ", err)
		return nil, err
	}

	for p.Database.GetNextRows() {
		var productItem entity.Product

		err := p.Database.Scan(
			&product.ID,
			&product.Name,
			&product.Category.Value,
			&product.Image,
			&product.Description,
			&product.Price,
			&product.CreatedAt.Value,
			&product.DeactivatedAt.Value,
		)

		if err != nil {
			l.Errorf("GetProductByCategory error to scan database: ", " | ", err)
			return nil, err
		}

		productItem = *product
		productList = append(productList, productItem)
	}

	return productList, nil
}

func (p productRepository) DeleteProductByID(id int64) error {
	l.Infof("DeleteProductByID received input: ", " | ", id)
	if err := p.Database.Connect(); err != nil {
		l.Errorf("DeleteProductByID connect error: ", " | ", err)
		return err
	}

	defer p.Database.Close()

	if err := p.Database.PrepareStmt(queryDeleteProduct); err != nil {
		l.Errorf("DeleteProductByID prepare error: ", " | ", err)
		return err
	}

	defer p.Database.CloseStmt()

	p.Database.QueryRow(id)

	var returnID int

	if err := p.Database.ScanStmt(&returnID); err != nil {
		l.Errorf("DeleteProductByID scan error: ", " | ", err)
		return err
	}

	return nil
}
