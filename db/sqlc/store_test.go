package db

import (
	"context"
	"github.com/ZeinabAshjaei/go-master/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStore_TransferTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := getARandomAccount(t)
	account2 := getARandomAccount(t)

	numberOfConcurrentTx := 5
	amount := utils.RandomMoney()

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < numberOfConcurrentTx; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})
			errs <- err
			results <- result
		}()
	}

	for i := 0; i < numberOfConcurrentTx; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results

		transfer := result.Transfer
		require.NotEmpty(t, result)
		require.Equal(t, transfer.FromAccountID, account1.ID)
		require.Equal(t, transfer.ToAccountID, account2.ID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), result.Transfer.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.Equal(t, toEntry.AccountID, account2.ID)
		require.Equal(t, toEntry.Amount, amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		fromEntry := result.FromEntry
		require.Equal(t, fromEntry.AccountID, account1.ID)
		require.Equal(t, -fromEntry.Amount, amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)
	}
}
