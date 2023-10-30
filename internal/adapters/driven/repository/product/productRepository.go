package product

import (
	"context"
	psqldb "fiap-hf-src/infrastructure/db/postgres"
	"fiap-hf-src/internal/core/domain/entity"
	"fiap-hf-src/internal/core/domain/valueObject"
)

var (
	querySaveProduct = `INSERT INTO product (id, name, category, image, description, price, created_at, deactivated_at) VALUES (DEFAULT, $1, $2, $3, $4, $5, now(), NULL) RETURNING id, created_at`
)

type ProductRepository interface {
	SaveProduct(product entity.Product) (*entity.Product, error)
}

type productRepository struct {
	Ctx      context.Context
	Database psqldb.PostgresDB
}

func NewProductRepository(ctx context.Context, db psqldb.PostgresDB) ProductRepository {
	return productRepository{Ctx: ctx, Database: db}
}

func (o productRepository) SaveProduct(product entity.Product) (*entity.Product, error) {
	if err := o.Database.Connect(); err != nil {
		return nil, err
	}

	defer o.Database.Close()

	if err := o.Database.PrepareStmt(querySaveProduct); err != nil {
		return nil, err
	}

	defer o.Database.CloseStmt()

	var outProduct = &entity.Product{
		Name: product.Name,
		Category: valueObject.Category{
			Value: product.Category.Value,
		},
		Image:       product.Image,
		Description: product.Description,
		Price:       product.Price,
	}

	o.Database.QueryRow(product.Name, product.Category.Value, product.Image, product.Description, product.Price)

	if err := o.Database.ScanStmt(&outProduct.ID, &outProduct.CreatedAt.Value); err != nil {
		return nil, err
	}

	return outProduct, nil
}
