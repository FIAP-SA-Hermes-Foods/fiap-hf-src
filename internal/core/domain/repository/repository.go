package repository

import "fiap-hf-src/internal/core/domain/entity"

type ClientRepository interface {
	GetClientByID(clientID string) (*entity.Client, error)
}
