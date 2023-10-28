package client

import (
	"context"
	psqldb "fiap-hf-src/infrastructure/db/postgres"
	"fiap-hf-src/internal/core/domain/entity"
	"fiap-hf-src/internal/core/domain/valueObject"
)

var (
	queryGetClientByCPF = `SELECT * from client WHERE cpf = $1`
	querySaveClient     = `INSERT INTO client (id, name, cpf, email, created_at) VALUES (DEFAULT, $1, $2, $3, now()) RETURNING id, created_at`
)

type ClientRepository interface {
	GetClientByCPF(cpf string) (*entity.Client, error)
	SaveClient(client entity.Client) (*entity.Client, error)
}

type clientRepository struct {
	Ctx      context.Context
	Database psqldb.PostgresDB
}

func NewClientRepository(ctx context.Context, db psqldb.PostgresDB) ClientRepository {
	return clientRepository{Ctx: ctx, Database: db}
}

func (c clientRepository) GetClientByCPF(cpf string) (*entity.Client, error) {
	if err := c.Database.Connect(); err != nil {
		return nil, err
	}

	defer c.Database.Close()

	var outClient = new(entity.Client)

	if err := c.Database.Query(queryGetClientByCPF, cpf); err != nil {
		return nil, err
	}

	for c.Database.GetNextRows() {
		if err := c.Database.Scan(&outClient.ID, &outClient.Name, &outClient.CPF.Value, &outClient.Email, &outClient.CreatedAt.Value); err != nil {
			return nil, err
		}
	}

	if *outClient == (entity.Client{}) {
		return nil, nil
	}

	return outClient, nil
}

func (c clientRepository) SaveClient(client entity.Client) (*entity.Client, error) {

	if err := c.Database.Connect(); err != nil {
		return nil, err
	}

	defer c.Database.Close()

	if err := c.Database.PrepareStmt(querySaveClient); err != nil {
		return nil, err
	}

	defer c.Database.CloseStmt()

	var outClient = &entity.Client{
		Name: client.Name,
		CPF: valueObject.Cpf{
			Value: client.CPF.Value,
		},
		Email: client.Email,
	}

	c.Database.QueryRow(client.Name, client.CPF.Value, client.Email)

	if err := c.Database.ScanStmt(&outClient.ID, &outClient.CreatedAt.Value); err != nil {
		return nil, err
	}

	return outClient, nil
}
