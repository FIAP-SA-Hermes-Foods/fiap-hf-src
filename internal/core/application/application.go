package application

import "hermes-foods/internal/core/domain/entity"

func whatever() {
	var c entity.Client

	c.CPF.Validate()
}