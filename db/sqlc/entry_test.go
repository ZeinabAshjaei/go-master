package db

import (
	"context"
	"github.com/ZeinabAshjaei/go-master/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func getRandomEntry(account Account, t *testing.T) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    utils.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	return entry
}

func TestQueries_CreateEntry(t *testing.T) {
	account := getARandomAccount(t)
	getRandomEntry(account, t)
}

func TestQueries_GetEntry(t *testing.T) {
	account := getARandomAccount(t)
	entry := getRandomEntry(account, t)

	retrievedEntry, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, retrievedEntry)
	require.Equal(t, retrievedEntry.AccountID, entry.AccountID)
	require.Equal(t, retrievedEntry.Amount, entry.Amount)
	require.NotZero(t, retrievedEntry.CreatedAt)

}

func TestQueries_ListEntries(t *testing.T) {
	account := getARandomAccount(t)
	for i := 0; i < 10; i++ {
		getRandomEntry(account, t)
	}

	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	}
	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)

	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
