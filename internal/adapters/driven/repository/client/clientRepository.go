package client

import (
	"context"
	l "fiap-hf-src/infrastructure/logger"
	"fiap-hf-src/internal/core/db"
	"fiap-hf-src/internal/core/domain/entity"
	"fiap-hf-src/internal/core/domain/valueObject"
)

var (
	queryGetClientByCPF = `SELECT * FROM client WHERE cpf = $1`
	queryGetClientByID  = `SELECT * FROM client WHERE id = $1`
	querySaveClient     = `INSERT INTO client (id, name, cpf, email, created_at) VALUES (DEFAULT, $1, $2, $3, now()) RETURNING id, created_at`
)

type ClientRepository interface {
	GetClientByID(id int64) (*entity.Client, error)
	GetClientByCPF(cpf string) (*entity.Client, error)
	SaveClient(client entity.Client) (*entity.Client, error)
}

type clientRepository struct {
	Ctx      context.Context
	Database db.SQLDatabase
}

func NewClientRepository(ctx context.Context, db db.SQLDatabase) ClientRepository {
	return clientRepository{Ctx: ctx, Database: db}
}

func (c clientRepository) GetClientByID(id int64) (*entity.Client, error) {
	l.Infof("GetClientByID received input: ", " | ", id)
	if err := c.Database.Connect(); err != nil {
		l.Errorf("GetClientByID connect error:", " | ", err)
		return nil, err
	}

	defer c.Database.Close()

	var outClient = new(entity.Client)

	if err := c.Database.Query(queryGetClientByID, id); err != nil {
		l.Errorf("GetClientByID error to connect database: ", " | ", err)
		return nil, err
	}

	for c.Database.GetNextRows() {
		if err := c.Database.Scan(&outClient.ID, &outClient.Name, &outClient.CPF.Value, &outClient.Email, &outClient.CreatedAt.Value); err != nil {
			l.Errorf("GetClientByID error to scan database: ", " | ", err)
			return nil, err
		}
	}

	if *outClient == (entity.Client{}) {
		l.Infof("GetClientByID output: ", " | ", "nil")
		return nil, nil
	}

	l.Infof("GetClientByID output: ", " | ", outClient.MarshalString())
	return outClient, nil
}

func (c clientRepository) GetClientByCPF(cpf string) (*entity.Client, error) {
	if err := c.Database.Connect(); err != nil {
		l.Errorf("GetClientByCPF connect error: ", " | ", err)
		return nil, err
	}

	defer c.Database.Close()

	var outClient = new(entity.Client)

	if err := c.Database.Query(queryGetClientByCPF, cpf); err != nil {
		l.Errorf("GetClientByCPF error to connect database: ", " | ", err)
		return nil, err
	}

	for c.Database.GetNextRows() {
		if err := c.Database.Scan(&outClient.ID, &outClient.Name, &outClient.CPF.Value, &outClient.Email, &outClient.CreatedAt.Value); err != nil {
			l.Errorf("GetClientByCPF error to scan database: ", " | ", err)
			return nil, err
		}
	}

	if *outClient == (entity.Client{}) {
		l.Infof("GetClientByCPF output: ", " | ", "nil")
		return nil, nil
	}
	l.Infof("GetClientByCPF output: ", " | ", outClient.MarshalString())
	return outClient, nil
}

func (c clientRepository) SaveClient(client entity.Client) (*entity.Client, error) {

	err := c.Database.Connect()
	if err != nil {
		l.Errorf("SaveClient connect error:", " | ", err)
		return nil, err
	}

	defer c.Database.Close()

	if err := c.Database.PrepareStmt(querySaveClient); err != nil {
		l.Errorf("SaveClient error to connect database: ", " | ", err)
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
		l.Errorf("SaveClient error to scan database: ", " | ", err)
		return nil, err
	}

	return outClient, nil
}
