package interfaces

import "fiap-hf-src/src/core/entity"

type ClientRepository interface {
	GetClientByID(id int64) (*entity.Client, error)
	GetClientByCPF(cpf string) (*entity.Client, error)
	SaveClient(client entity.Client) (*entity.Client, error)
}

type ClientService interface {
	GetClientByID(id int64) error
	GetClientByCPF(cpf string) error
	SaveClient(client entity.Client) (*entity.Client, error)
}
