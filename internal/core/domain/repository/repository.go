package repository

import "fiap-hf-src/internal/core/domain/entity"

type ClientRepository interface {
	GetClientByCPF(cpf string) (*entity.Client, error)
	SaveClient(client entity.Client) (*entity.Client, error)
}
