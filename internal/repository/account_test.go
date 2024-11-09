package repository

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/saleh-ghazimoradi/BankHubGo/internal/service/service_model"
	utils "github.com/saleh-ghazimoradi/BankHubGo/utils/connections"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func setupTestDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock, *accountRepository) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error initializing mock DB: %v", err)
	}
	repo := NewAccountRepository(db)
	return db, mock, repo.(*accountRepository)
}

func TestGetAccount(t *testing.T) {
	db, mock, repo := setupTestDB(t)
	defer db.Close()

	ctx := context.Background()
	accountID := utils.RandomInt(1, 1000)
	expectedAccount := &service_model.Account{
		ID:        accountID,
		Owner:     utils.RandomOwner(),
		Balance:   utils.RandomMoney(),
		Currency:  utils.RandomCurrency(),
		CreatedAt: time.Now(),
	}

	mock.ExpectQuery(`SELECT id, owner, balance, currency, created_at FROM accounts WHERE id = \$1`).
		WithArgs(accountID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "owner", "balance", "currency", "created_at"}).
			AddRow(expectedAccount.ID, expectedAccount.Owner, expectedAccount.Balance, expectedAccount.Currency, expectedAccount.CreatedAt))

	account, err := repo.GetAccount(ctx, accountID)
	assert.NoError(t, err, "unexpected error fetching account")
	assert.Equal(t, expectedAccount, account)

	assert.NoError(t, mock.ExpectationsWereMet(), "there were unfulfilled expectations")
}

func TestGetAccounts(t *testing.T) {
	db, mock, repo := setupTestDB(t)
	defer db.Close()

	ctx := context.Background()
	pagination := service_model.Pagination{
		Limit:  2,
		Offset: 0,
		Sort:   "ASC",
	}

	expectedAccounts := []*service_model.Account{
		{ID: utils.RandomInt(1, 1000), Owner: utils.RandomOwner(), Balance: utils.RandomMoney(), Currency: utils.RandomCurrency(), CreatedAt: time.Now()},
		{ID: utils.RandomInt(1, 1000), Owner: utils.RandomOwner(), Balance: utils.RandomMoney(), Currency: utils.RandomCurrency(), CreatedAt: time.Now()},
	}

	mock.ExpectQuery(`SELECT id, owner, balance, currency, created_at FROM accounts ORDER BY created_at ASC LIMIT \$1 OFFSET \$2`).
		WithArgs(pagination.Limit, pagination.Offset).
		WillReturnRows(sqlmock.NewRows([]string{"id", "owner", "balance", "currency", "created_at"}).
			AddRow(expectedAccounts[0].ID, expectedAccounts[0].Owner, expectedAccounts[0].Balance, expectedAccounts[0].Currency, expectedAccounts[0].CreatedAt).
			AddRow(expectedAccounts[1].ID, expectedAccounts[1].Owner, expectedAccounts[1].Balance, expectedAccounts[1].Currency, expectedAccounts[1].CreatedAt))

	accounts, err := repo.GetAccounts(ctx, pagination)
	assert.NoError(t, err, "unexpected error fetching accounts")
	assert.Equal(t, expectedAccounts, accounts)
	assert.NoError(t, mock.ExpectationsWereMet(), "there were unfulfilled expectations")
}

func TestCreateAccount(t *testing.T) {
	db, mock, repo := setupTestDB(t)
	defer db.Close()

	ctx := context.Background()

	newAccount := &service_model.Account{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	mock.ExpectQuery(`INSERT INTO accounts \(owner, balance, currency\) VALUES \(\$1, \$2, \$3\) RETURNING id, created_at`).
		WithArgs(newAccount.Owner, newAccount.Balance, newAccount.Currency).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at"}).AddRow(1, time.Now()))

	err := repo.CreateAccount(ctx, newAccount)
	assert.NoError(t, err)
	assert.NotZero(t, newAccount.ID)
	assert.NotZero(t, newAccount.CreatedAt)

	assert.NoError(t, mock.ExpectationsWereMet(), "there were unfulfilled expectations")
}

func TestUpdateAccount(t *testing.T) {
	db, mock, repo := setupTestDB(t)
	defer db.Close()

	ctx := context.Background()
	accountID := utils.RandomInt(1, 1000)
	newBalance := utils.RandomMoney()

	mock.ExpectExec(`UPDATE accounts SET balance = \$1 WHERE id = \$2`).
		WithArgs(newBalance, accountID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.UpdateAccount(ctx, &service_model.Account{ID: accountID, Balance: newBalance})
	assert.NoError(t, err, "unexpected error updating account")

	assert.NoError(t, mock.ExpectationsWereMet(), "there were unfulfilled expectations")
}

func TestDeleteAccount(t *testing.T) {
	db, mock, repo := setupTestDB(t)
	defer db.Close()

	ctx := context.Background()
	accountID := utils.RandomInt(1, 1000)

	mock.ExpectExec(`DELETE FROM accounts WHERE id = \$1`).
		WithArgs(accountID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.DeleteAccount(ctx, accountID)
	assert.NoError(t, err, "unexpected error deleting account")
	assert.NoError(t, mock.ExpectationsWereMet(), "there were unfulfilled expectations")
}
