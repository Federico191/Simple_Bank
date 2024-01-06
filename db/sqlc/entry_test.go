package db

import (
	"context"
	"database/sql"
	"github.com/Federico191/Simple_Bank/util"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func createRandomEntry(t *testing.T) Entry {
	account := createRandomAccount(t)

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	assert.NoError(t, err)
	assert.NotEmpty(t, entry)

	assert.Equal(t, arg.AccountID, entry.AccountID)
	assert.Equal(t, arg.Amount, entry.Amount)

	return entry
}

func TestQueries_CreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestQueries_GetEntry(t *testing.T) {
	entry1 := createRandomEntry(t)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, entry2)

	assert.Equal(t, entry1.ID, entry2.ID)
	assert.Equal(t, entry1.AccountID, entry2.AccountID)
	assert.Equal(t, entry1.Amount, entry2.Amount)
	assert.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestQueries_ListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}

	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	assert.NoError(t, err)
	assert.NotEmpty(t, entries)

	for _, entry := range entries {
		assert.NotEmpty(t, entry)
	}
}

func TestQueries_UpdateEntries(t *testing.T) {
	entry1 := createRandomEntry(t)

	arg := UpdateEntriesParams{
		ID:     entry1.ID,
		Amount: util.RandomMoney(),
	}

	entry2, err := testQueries.UpdateEntries(context.Background(), arg)
	assert.NoError(t, err)
	assert.NotEmpty(t, entry2)

	assert.Equal(t, arg.ID, entry2.ID)
	assert.Equal(t, arg.Amount, entry2.Amount)
}

func TestQueries_DeleteEntries(t *testing.T) {
	entry1 := createRandomEntry(t)

	err := testQueries.DeleteEntries(context.Background(), entry1.ID)
	assert.NoError(t, err)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	assert.Error(t, err)
	assert.EqualError(t, err, sql.ErrNoRows.Error())
	assert.Empty(t, entry2)
}
