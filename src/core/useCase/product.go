package useCase

import (
	"errors"
	"fiap-hf-src/src/base/dto"
	"fiap-hf-src/src/base/interfaces"
	c "fiap-hf-src/src/operation/presenter/common"
	"strings"
)

var _ interfaces.ProductUseCase = (*productUseCase)(nil)

type productUseCase struct {
	gateway interfaces.ProductGateway
}

func NewProductUseCase(gateway interfaces.ProductGateway) *productUseCase {
	return &productUseCase{gateway: gateway}
}

func (p *productUseCase) SaveProduct(reqProduct dto.RequestProduct) (*dto.OutputProduct, error) {
	product := reqProduct.Product()

	if err := product.Category.Validate(); err != nil {
		return nil, err
	}

	reqProduct.Category = product.Category.Value

	return p.gateway.SaveProduct(reqProduct)
}

func (p *productUseCase) UpdateProductByID(id int64, reqProduct dto.RequestProduct) (*dto.OutputProduct, error) {
	if id < 1 {
		return nil, errors.New("the id is not valid for consult")
	}

	product := reqProduct.Product()

	if err := product.Category.Validate(); err != nil {
		return nil, err
	}

	reqProduct.Category = product.Category.Value

	return p.gateway.UpdateProductByID(id, reqProduct)
}

func (p *productUseCase) GetProductByID(id int64) (*dto.OutputProduct, error) {
	if id < 1 {
		return nil, errors.New("the id is not valid for consult")
	}
	return p.gateway.GetProductByID(id)
}

func (p *productUseCase) GetProductByCategory(category string) ([]dto.OutputProduct, error) {
	if len(category) < 1 {
		return nil, errors.New("the category is not valid for consult")
	}

	if _, ok := c.CategoryMap[strings.ToLower(category)]; !ok {
		return nil, errors.New("category is not valid")
	}

	return p.gateway.GetProductByCategory(category)
}

func (p *productUseCase) DeleteProductByID(id int64) error {
	if id < 1 {
		return errors.New("the id is not valid for consult")
	}
	return p.gateway.DeleteProductByID(id)
}
