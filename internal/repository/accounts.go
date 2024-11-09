package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/saleh-ghazimoradi/BankHubGo/internal/service/service_model"
)

type Account interface {
	GetAccount(ctx context.Context, id int64) (*service_model.Account, error)
	GetAccounts(ctx context.Context, p service_model.Pagination) ([]*service_model.Account, error)
	CreateAccount(ctx context.Context, account *service_model.Account) error
	UpdateAccount(ctx context.Context, account *service_model.Account) error
	DeleteAccount(ctx context.Context, id int64) error
}

type accountRepository struct {
	db *sql.DB
}

func (a *accountRepository) GetAccount(ctx context.Context, id int64) (*service_model.Account, error) {
	var account service_model.Account

	query := `SELECT id, owner, balance, currency, created_at FROM accounts WHERE id = $1;`
	err := a.db.QueryRowContext(ctx, query, id).Scan(
		&account.ID,
		&account.Owner,
		&account.Balance,
		&account.Currency,
		&account.CreatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}
	return &account, nil
}

func (a *accountRepository) GetAccounts(ctx context.Context, p service_model.Pagination) ([]*service_model.Account, error) {

	var accounts []*service_model.Account

	query := `SELECT id, owner, balance, currency, created_at FROM accounts ORDER BY created_at ` + p.Sort + `
	   LIMIT $1 OFFSET $2`

	rows, err := a.db.QueryContext(ctx, query, p.Limit, p.Offset)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var account service_model.Account
		err = rows.Scan(
			&account.ID,
			&account.Owner,
			&account.Balance,
			&account.Currency,
			&account.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, &account)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}

func (a *accountRepository) CreateAccount(ctx context.Context, account *service_model.Account) error {
	query := `INSERT INTO accounts (owner, balance, currency) VALUES ($1, $2, $3) RETURNING id, created_at;`

	err := a.db.QueryRowContext(ctx, query, account.Owner, account.Balance, account.Currency).Scan(&account.ID, &account.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (a *accountRepository) UpdateAccount(ctx context.Context, account *service_model.Account) error {
	query := `UPDATE accounts SET balance = $1 WHERE id = $2;`
	res, err := a.db.ExecContext(ctx, query, account.Balance, account.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

func (a *accountRepository) DeleteAccount(ctx context.Context, id int64) error {
	query := `DELETE FROM accounts WHERE id = $1;`
	res, err := a.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

func NewAccountRepository(db *sql.DB) Account {
	return &accountRepository{
		db: db,
	}
}
