package repository

import "hermes-foods/internal/core/domain/entity"

type ClientRepository interface {
	GetClientByID(clientID string) (*entity.Client, error)
}
