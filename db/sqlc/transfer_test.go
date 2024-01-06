package db

import (
	"context"
	"database/sql"
	"github.com/Federico191/Simple_Bank/util"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func createRandomTransfer(t *testing.T) Transfer {
	accountFrom := createRandomAccount(t)
	accountTo := createRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: accountFrom.ID,
		ToAccountID:   accountTo.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	assert.NoError(t, err)
	assert.NotEmpty(t, transfer)

	assert.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	assert.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	assert.Equal(t, arg.Amount, transfer.Amount)

	return transfer
}

func TestQueries_CreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestQueries_GetTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, transfer2)

	assert.Equal(t, transfer1.ID, transfer2.ID)
	assert.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	assert.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	assert.Equal(t, transfer1.Amount, transfer2.Amount)
	assert.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestQueries_ListTransfer(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomTransfer(t)
	}

	arg := ListTransferParams{
		Limit:  5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransfer(context.Background(), arg)
	assert.NoError(t, err)

	for _, transfer := range transfers {
		assert.NotEmpty(t, transfer)
	}
}

func TestQueries_UpdateTransfers(t *testing.T) {
	transfer1 := createRandomTransfer(t)

	arg := UpdateTransfersParams{
		ID:     transfer1.ID,
		Amount: util.RandomMoney(),
	}

	transfer2, err := testQueries.UpdateTransfers(context.Background(), arg)
	assert.NoError(t, err)
	assert.NotEmpty(t, transfer2)

	assert.Equal(t, arg.ID, transfer2.ID)
	assert.Equal(t, arg.Amount, transfer2.Amount)
}

func TestQueries_DeleteTransfers(t *testing.T) {
	transfer1 := createRandomTransfer(t)

	err := testQueries.DeleteTransfers(context.Background(), transfer1.ID)
	assert.NoError(t, err)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	assert.Error(t, err)
	assert.EqualError(t, err, sql.ErrNoRows.Error())
	assert.Empty(t, transfer2)
}
