package service

import "github.com/saleh-ghazimoradi/BankHubGo/internal/repository"

type Transfer interface {
}

type transferService struct {
	transferRepo repository.Transfer
}

func NewTransferService(transferRepo repository.Transfer) Transfer {
	return &transferService{transferRepo: transferRepo}
}
