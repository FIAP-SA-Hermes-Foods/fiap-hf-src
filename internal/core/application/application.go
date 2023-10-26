package application

import "fiap-hf-src/internal/core/domain/entity"

func whatever() {
	var c entity.Client

	c.CPF.Validate()
}
