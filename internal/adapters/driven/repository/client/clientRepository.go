package client

import (
	"context"
	psqldb "fiap-hf-src/internal/adapters/driven/repository/db"
	"fiap-hf-src/internal/core/domain/entity"
)

var (
	queryGetClientByID = `SELECT * from client WHERE id = ?`
)

type ClientRepository interface {
	GetClientByID(clientID string) (*entity.Client, error)
}

type clientRepository struct {
	Ctx      context.Context
	Database psqldb.PostgresDB
}

func NewClientRepository(ctx context.Context, db psqldb.PostgresDB) ClientRepository {
	return clientRepository{Ctx: ctx, Database: db}
}

func (c clientRepository) GetClientByID(clientID string) (*entity.Client, error) {
	if err := c.Database.Connect(); err != nil {
		return nil, err
	}

	defer c.Database.Close()

	if err := c.Database.PrepareStmt(queryGetClientByID); err != nil {
		return nil, err
	}

	defer c.Database.CloseStmt()

	var outClient = new(entity.Client)

	c.Database.QueryRow(clientID)

	if err := c.Database.Scan(&outClient.ID, &outClient.Name, &outClient.CPF.Value, &outClient.Email); err != nil {
		return nil, err
	}

	return outClient, nil
}
