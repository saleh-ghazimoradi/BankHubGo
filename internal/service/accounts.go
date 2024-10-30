package service

import (
	"context"
	"github.com/saleh-ghazimoradi/BankHubGo/internal/repository"
	"github.com/saleh-ghazimoradi/BankHubGo/internal/service/service_model"
)

type Account interface {
	GetAccount(ctx context.Context, id string) (*service_model.Account, error)
	GetAccounts(ctx context.Context) ([]*service_model.Account, error)
	CreateAccount(ctx context.Context, account *service_model.Account) error
	UpdateAccount(ctx context.Context, account *service_model.Account) error
	DeleteAccount(ctx context.Context, id string) error
}

type accountService struct {
	accountRepo repository.Account
}

func (a *accountService) GetAccount(ctx context.Context, id string) (*service_model.Account, error) {
	return nil, nil
}

func (a *accountService) GetAccounts(ctx context.Context) ([]*service_model.Account, error) {
	return nil, nil
}

func (a *accountService) CreateAccount(ctx context.Context, account *service_model.Account) error {
	return nil
}

func (a *accountService) UpdateAccount(ctx context.Context, account *service_model.Account) error {
	return nil
}

func (a *accountService) DeleteAccount(ctx context.Context, id string) error {
	return nil
}

func NewAccountService(accountRepo repository.Account) Account {
	return &accountService{
		accountRepo: accountRepo,
	}
}
