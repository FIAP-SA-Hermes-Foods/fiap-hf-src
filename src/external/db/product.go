package db

import (
	"context"
	"encoding/json"
	"fiap-hf-src/src/base/interfaces"
	l "fiap-hf-src/src/base/logger"
	"fiap-hf-src/src/core/entity"
	com "fiap-hf-src/src/operation/presenter/common"
	ps "fiap-hf-src/src/operation/presenter/strings"
)

var (
	queryGetProductByID       = `SELECT * FROM product where id = $1`
	queryGetProductByCategory = `SELECT * FROM product where category = $1`
	queryDeleteProduct        = `DELETE FROM product where id = $1 RETURNING id`
	querySaveProduct          = `INSERT INTO product (id, name, category, image, description, price, created_at, deactivated_at) VALUES (DEFAULT, $1, $2, $3, $4, $5, now(), NULL) RETURNING id, created_at`
	queryUpdateProduct        = `UPDATE product SET name = $1, category = $2, image = $3, description = $4, price = $5, deactivated_at = $6 WHERE id = $7 RETURNING id, created_at`
)

var _ interfaces.ProductDB = (*productDB)(nil)

type productDB struct {
	Ctx      context.Context
	Database interfaces.SQLDatabase
}

func NewProductDB(ctx context.Context, db interfaces.SQLDatabase) *productDB {
	return &productDB{Ctx: ctx, Database: db}
}

func (p *productDB) SaveProduct(product entity.Product) (*entity.Product, error) {
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
		Category: com.Category{
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
	l.Infof("SaveProduct output: ", " | ", ps.MarshalString(outProduct))
	return outProduct, nil
}

func (p *productDB) UpdateProductByID(id int64, product entity.Product) (*entity.Product, error) {
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
		Category: com.Category{
			Value: product.Category.Value,
		},
		Image:       product.Image,
		Description: product.Description,
		Price:       product.Price,
		DeactivatedAt: com.DeactivatedAt{
			Value: product.DeactivatedAt.Value,
		},
	}

	p.Database.QueryRow(product.Name, product.Category.Value, product.Image, product.Description, product.Price, product.DeactivatedAt.Value, id)

	if err := p.Database.ScanStmt(&outProduct.ID, &outProduct.CreatedAt.Value); err != nil {
		l.Errorf("UpdateProductByID scan error: ", " | ", err)
		return nil, err
	}

	l.Infof("UpdateProductByID output: ", " | ", ps.MarshalString(outProduct))
	return outProduct, nil
}

func (p *productDB) GetProductByID(id int64) (*entity.Product, error) {
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

	l.Infof("GetProductByID output: ", " | ", ps.MarshalString(outProduct))
	return outProduct, nil
}

func (p *productDB) GetProductByCategory(category string) ([]entity.Product, error) {
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

	plStr, err := json.Marshal(productList)

	if err != nil {
		l.Errorf("GetProductByCategory error to unmarshal: ", " | ", err)
		return nil, err
	}

	l.Infof("GetProductByCategory output: ", " | ", string(plStr))
	return productList, nil
}

func (p *productDB) DeleteProductByID(id int64) error {
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
