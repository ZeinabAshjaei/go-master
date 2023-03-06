package db

import (
	"context"
	"database/sql"
	"github.com/ZeinabAshjaei/go-master/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func getARandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomInt(0, 1000),
		Currency: utils.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	require.Equal(t, account.Owner, arg.Owner)
	require.Equal(t, account.Balance, arg.Balance)
	require.Equal(t, account.Currency, arg.Currency)

	return account
}

func TestQueries_CreateAccount(t *testing.T) {
	getARandomAccount(t)
}

func TestQueries_GetAccount(t *testing.T) {
	insertedAccount := getARandomAccount(t)
	retrievedAccount, err := testQueries.GetAccount(context.Background(), insertedAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, retrievedAccount)

	require.Equal(t, insertedAccount.ID, retrievedAccount.ID)
	require.Equal(t, insertedAccount.Balance, retrievedAccount.Balance)
	require.Equal(t, insertedAccount.Owner, retrievedAccount.Owner)
	require.Equal(t, insertedAccount.Currency, retrievedAccount.Currency)
	require.WithinDuration(t, insertedAccount.CreatedAt, retrievedAccount.CreatedAt, time.Second)

}

func TestQueries_UpdateAccount(t *testing.T) {
	insertedAccount := getARandomAccount(t)

	arg := UpdateAccountParams{
		ID:      insertedAccount.ID,
		Balance: utils.RandomInt(0, 1000),
	}

	retrievedAccount, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, retrievedAccount)

	require.Equal(t, insertedAccount.ID, retrievedAccount.ID)
	require.Equal(t, arg.Balance, retrievedAccount.Balance)
	require.Equal(t, insertedAccount.Owner, retrievedAccount.Owner)
	require.Equal(t, insertedAccount.Currency, retrievedAccount.Currency)
	require.WithinDuration(t, insertedAccount.CreatedAt, retrievedAccount.CreatedAt, time.Second)

}

func TestQueries_DeleteAccount(t *testing.T) {
	insertedAccount := getARandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), insertedAccount.ID)
	require.NoError(t, err)

	retrievedAccount, err := testQueries.GetAccount(context.Background(), insertedAccount.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, retrievedAccount)

}

func TestQueries_ListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		getARandomAccount(t)
	}

	arg := ListAccountParams{Limit: 5, Offset: 5}
	accounts, err := testQueries.ListAccount(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
