package service

import (
	"context"
	"github.com/saleh-ghazimoradi/BankHubGo/internal/repository"
	"github.com/saleh-ghazimoradi/BankHubGo/internal/service/service_model"
)

type Account interface {
	GetAccount(ctx context.Context, id int64) (*service_model.Account, error)
	GetAccounts(ctx context.Context) ([]*service_model.Account, error)
	CreateAccount(ctx context.Context, account *service_model.Account) error
	UpdateAccount(ctx context.Context, account *service_model.Account) error
	DeleteAccount(ctx context.Context, id int64) error
}

type accountService struct {
	accountRepo repository.Account
}

func (a *accountService) GetAccount(ctx context.Context, id int64) (*service_model.Account, error) {
	return a.accountRepo.GetAccount(ctx, id)
}

func (a *accountService) GetAccounts(ctx context.Context) ([]*service_model.Account, error) {
	return a.accountRepo.GetAccounts(ctx)
}

func (a *accountService) CreateAccount(ctx context.Context, account *service_model.Account) error {
	return a.accountRepo.CreateAccount(ctx, account)
}

func (a *accountService) UpdateAccount(ctx context.Context, account *service_model.Account) error {
	return a.accountRepo.UpdateAccount(ctx, account)
}

func (a *accountService) DeleteAccount(ctx context.Context, id int64) error {
	return a.accountRepo.DeleteAccount(ctx, id)
}

func NewAccountService(accountRepo repository.Account) Account {
	return &accountService{
		accountRepo: accountRepo,
	}
}
