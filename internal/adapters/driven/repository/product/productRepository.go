package product

import (
	"context"
	psqldb "fiap-hf-src/infrastructure/db/postgres"
	"fiap-hf-src/internal/core/domain/entity"
	"fiap-hf-src/internal/core/domain/valueObject"
)

var (
	queryGetProductByID = `SELECT * FROM product where id = $1`
	querySaveProduct    = `INSERT INTO product (id, name, category, image, description, price, created_at, deactivated_at) VALUES (DEFAULT, $1, $2, $3, $4, $5, now(), NULL) RETURNING id, created_at`
	queryUpdateProduct  = `UPDATE product SET name = $1, category = $2, image = $3, description = $4, price = $5, deactivated_at = $6 WHERE id = $7 RETURNING id, created_at`
)

type ProductRepository interface {
	SaveProduct(product entity.Product) (*entity.Product, error)
	UpdateProductByID(id int64, product entity.Product) (*entity.Product, error)
	GetProductByID(id int64) (*entity.Product, error)
}

type productRepository struct {
	Ctx      context.Context
	Database psqldb.PostgresDB
}

func NewProductRepository(ctx context.Context, db psqldb.PostgresDB) ProductRepository {
	return productRepository{Ctx: ctx, Database: db}
}

func (p productRepository) SaveProduct(product entity.Product) (*entity.Product, error) {
	if err := p.Database.Connect(); err != nil {
		return nil, err
	}

	defer p.Database.Close()

	if err := p.Database.PrepareStmt(querySaveProduct); err != nil {
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
		return nil, err
	}

	return outProduct, nil
}

func (p productRepository) UpdateProductByID(id int64, product entity.Product) (*entity.Product, error) {
	if err := p.Database.Connect(); err != nil {
		return nil, err
	}

	defer p.Database.Close()

	if err := p.Database.PrepareStmt(queryUpdateProduct); err != nil {
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
		return nil, err
	}

	return outProduct, nil
}

func (p productRepository) GetProductByID(id int64) (*entity.Product, error) {
	if err := p.Database.Connect(); err != nil {
		return nil, err
	}

	defer p.Database.Close()

	var outProduct = new(entity.Product)

	if err := p.Database.Query(queryGetProductByID, id); err != nil {
		return nil, err
	}

	for p.Database.GetNextRows() {
		err := p.Database.Scan(
			&outProduct.ID,
			&outProduct.Name,
			&outProduct.Category,
			&outProduct.Image,
			&outProduct.Description,
			&outProduct.Price,
			&outProduct.CreatedAt.Value,
			&outProduct.DeactivatedAt.Value,
		)
		if err != nil {
			return nil, err
		}
	}

	if *outProduct == (entity.Product{}) {
		return nil, nil
	}

	return outProduct, nil
}
