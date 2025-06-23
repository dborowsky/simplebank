package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/dborowsky/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)
	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	return account
}

func TestCreateAccount(t *testing.T) {
	user := createRandomUser(t)
	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, account.Owner, arg.Owner)
	require.Equal(t, account.Balance, arg.Balance)
	require.Equal(t, account.Currency, arg.Currency)
	require.NotZero(t, account.ID) // ID should be auto-assigned by the database
	require.NotEmpty(t, account.CreatedAt)
}

func TestGetAccountByID(t *testing.T) {
	account := createRandomAccount(t)
	got, err := testQueries.GetAccount(context.Background(), account.ID)

	require.NoError(t, err)
	require.Equal(t, got.ID, account.ID)
	require.Equal(t, got.Owner, account.Owner)
	require.Equal(t, got.Balance, account.Balance)
	require.Equal(t, got.Currency, account.Currency)
	require.WithinDuration(t, got.CreatedAt, account.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account := createRandomAccount(t)
	newBalance := util.RandomMoney()

	got, err := testQueries.UpdateAccount(context.Background(), UpdateAccountParams{
		ID:      account.ID,
		Balance: newBalance,
	})

	require.NoError(t, err)
	require.Equal(t, got.Balance, newBalance)
}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)
	got, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, got)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 3; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountParams{Limit: 2, Offset: 2}

	accounts, err := testQueries.ListAccount(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 2)
	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
