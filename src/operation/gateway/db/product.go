package db

import (
	"fiap-hf-src/src/base/dto"
	"fiap-hf-src/src/base/interfaces"
)

var _ interfaces.ProductGateway = (*productDB)(nil)

type productDB struct {
	db interfaces.ProductDB
}

func NewProductGateway(db interfaces.ProductDB) *productDB {
	return &productDB{db: db}
}

func (p *productDB) SaveProduct(reqProduct dto.RequestProduct) (*dto.OutputProduct, error) {

	product := reqProduct.Product()

	outDB, err := p.db.SaveProduct(product)

	if err != nil {
		return nil, err
	}

	if outDB == nil {
		return nil, nil
	}

	out := &dto.OutputProduct{
		ID:            outDB.ID,
		Name:          outDB.Name,
		Category:      outDB.Category.Value,
		Image:         outDB.Image,
		Description:   outDB.Description,
		Price:         outDB.Price,
		CreatedAt:     outDB.CreatedAt.Format(),
		DeactivatedAt: outDB.DeactivatedAt.Format(),
	}

	return out, nil
}

func (p *productDB) UpdateProductByID(id int64, reqProduct dto.RequestProduct) (*dto.OutputProduct, error) {
	product := reqProduct.Product()

	outDB, err := p.db.UpdateProductByID(id, product)

	if err != nil {
		return nil, err
	}

	if outDB == nil {
		return nil, nil
	}

	out := &dto.OutputProduct{
		ID:            outDB.ID,
		Name:          outDB.Name,
		Category:      outDB.Category.Value,
		Image:         outDB.Image,
		Description:   outDB.Description,
		Price:         outDB.Price,
		CreatedAt:     outDB.CreatedAt.Format(),
		DeactivatedAt: outDB.DeactivatedAt.Format(),
	}

	return out, nil
}

func (p *productDB) GetProductByCategory(category string) ([]dto.OutputProduct, error) {

	outDB, err := p.db.GetProductByCategory(category)

	if err != nil {
		return nil, err
	}

	if outDB == nil {
		return nil, nil
	}

	var out = make([]dto.OutputProduct, 0)

	for i := 0; i < len(outDB); i++ {
		o := dto.OutputProduct{
			ID:            outDB[i].ID,
			Name:          outDB[i].Name,
			Category:      outDB[i].Category.Value,
			Image:         outDB[i].Image,
			Description:   outDB[i].Description,
			Price:         outDB[i].Price,
			CreatedAt:     outDB[i].CreatedAt.Format(),
			DeactivatedAt: outDB[i].DeactivatedAt.Format(),
		}

		out = append(out, o)
	}

	return out, nil
}

func (p *productDB) GetProductByID(id int64) (*dto.OutputProduct, error) {

	outDB, err := p.db.GetProductByID(id)

	if err != nil {
		return nil, err
	}

	if outDB == nil {
		return nil, nil
	}

	out := &dto.OutputProduct{
		ID:            outDB.ID,
		Name:          outDB.Name,
		Category:      outDB.Category.Value,
		Image:         outDB.Image,
		Description:   outDB.Description,
		Price:         outDB.Price,
		CreatedAt:     outDB.CreatedAt.Format(),
		DeactivatedAt: outDB.DeactivatedAt.Format(),
	}

	return out, nil

}

func (p *productDB) DeleteProductByID(id int64) error {

	if err := p.db.DeleteProductByID(id); err != nil {
		return err
	}

	return nil
}
