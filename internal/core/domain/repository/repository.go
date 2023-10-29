package repository

import "fiap-hf-src/internal/core/domain/entity"

type ClientRepository interface {
	GetClientByID(id int64) (*entity.Client, error)
	GetClientByCPF(cpf string) (*entity.Client, error)
	SaveClient(client entity.Client) (*entity.Client, error)
}

type OrderRepository interface {
	SaveOrder(order entity.Order) (*entity.Order, error)
}
