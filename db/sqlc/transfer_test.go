package db

import (
	"context"
	"github.com/ZeinabAshjaei/go-master/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func getRandomTransfer(account1, account2 Account, t *testing.T) Transfer {
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        utils.RandomMoney(),
	}
	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, transfer.Amount, arg.Amount)
	require.Equal(t, transfer.Amount, arg.Amount)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestQueries_CreateTransfer(t *testing.T) {
	account1 := getARandomAccount(t)
	account2 := getARandomAccount(t)
	getRandomTransfer(account1, account2, t)
}

func TestQueries_GetTransfer(t *testing.T) {
	account1 := getARandomAccount(t)
	account2 := getARandomAccount(t)
	transfer := getRandomTransfer(account1, account2, t)

	retrievedTransfer, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, retrievedTransfer)

	require.Equal(t, retrievedTransfer.Amount, transfer.Amount)
	require.Equal(t, retrievedTransfer.Amount, transfer.Amount)

	require.Equal(t, transfer.FromAccountID, retrievedTransfer.FromAccountID)
	require.Equal(t, transfer.ToAccountID, retrievedTransfer.ToAccountID)
	require.Equal(t, transfer.Amount, retrievedTransfer.Amount)

	require.NotZero(t, retrievedTransfer.CreatedAt)
}

func TestQueries_ListTransfers(t *testing.T) {
	account1 := getARandomAccount(t)
	account2 := getARandomAccount(t)

	for i := 0; i < 10; i++ {
		getRandomTransfer(account1, account2, t)
	}

	arg := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
