package repository

import (
	"context"
	"database/sql"
	"github.com/saleh-ghazimoradi/BankHubGo/internal/service/service_model"
)

type Account interface {
	GetAccount(ctx context.Context, id int64) (*service_model.Account, error)
	GetAccounts(ctx context.Context) ([]*service_model.Account, error)
	CreateAccount(ctx context.Context, account *service_model.Account) error
	UpdateAccount(ctx context.Context, account *service_model.Account) error
	DeleteAccount(ctx context.Context, id int64) error
}

type accountRepository struct {
	db *sql.DB
}

func (a *accountRepository) GetAccount(ctx context.Context, id int64) (*service_model.Account, error) {
	return nil, nil
}

func (a *accountRepository) GetAccounts(ctx context.Context) ([]*service_model.Account, error) {
	return nil, nil
}

func (a *accountRepository) CreateAccount(ctx context.Context, account *service_model.Account) error {
	return nil
}

func (a *accountRepository) UpdateAccount(ctx context.Context, account *service_model.Account) error {
	return nil
}

func (a *accountRepository) DeleteAccount(ctx context.Context, id int64) error {
	return nil
}

func NewAccountRepository(db *sql.DB) Account {
	return &accountRepository{
		db: db,
	}
}
