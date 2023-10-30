package valueObject

import (
	"errors"
	"strings"
)

type Category struct {
	Value string `json:"value,omitempty"`
}

const (
	LancheCategory         = "LANCHE"
	AcompanhamentoCategory = "ACOMPANHAMENTO"
	BebidaCategory         = "BEBIDA"
)

var categoryMap = map[string]string{
	"lanche":         LancheCategory,
	"acompanhamento": AcompanhamentoCategory,
	"bebida":         BebidaCategory,
}

func (v *Category) Validate() error {

	status, ok := categoryMap[strings.ToLower(v.Value)]

	if !ok {
		return errors.New("category is not valid")
	}

	v.Value = status

	return nil
}
