package repository

import (
	"context"
	"database/sql"
)

type Transfer interface {
	BeginTx(ctx context.Context) (*sql.Tx, error)
}

type transferRepository struct {
	db *sql.DB
}

func (t *transferRepository) BeginTx(ctx context.Context) (*sql.Tx, error) {
	return t.db.BeginTx(ctx, nil)
}

func NewTransferRepository(db *sql.DB) Transfer {
	return &transferRepository{db: db}
}
