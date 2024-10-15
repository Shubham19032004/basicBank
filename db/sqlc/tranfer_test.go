package db

import (
	"bank/utils"
	"context"

	"database/sql"

	"testing"

	"github.com/stretchr/testify/require"
)

func CreateRndomTransfer(t *testing.T) Transfer {
	account1 := randomAccount(t)
	account2 := randomAccount(t)
	args := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        utils.RandomMoney(),
	}
	transfer, err := testQueries.CreateTransfer(context.Background(), args)
	require.NoError(t, err)
	require.Equal(t, account1.ID, transfer.FromAccountID)
	require.Equal(t, account2.ID, transfer.ToAccountID)
	require.NotEmpty(t, transfer.Amount)
	return transfer

}

func TestCreateTranfer(t *testing.T) {
	CreateRndomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer1 := CreateRndomTransfer(t)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)

}
func TestDeleteTransfer(t *testing.T) {
	transfer1 := CreateRndomTransfer(t)
	err := testQueries.DeleteTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t,transfer2)
}

func TestUpadateTransfer(t *testing.T) {
	transfer1 := CreateRndomTransfer(t)
	arr:=UpdateTransferParams{
		Amount:utils.RandomMoney(),
		ID: transfer1.ID,
	}
	transfer2, err := testQueries.UpdateTransfer(context.Background(), arr)
	require.NoError(t, err)
	require.NotEmpty(t,transfer1)
	require.NotEmpty(t,transfer2)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ID, transfer2.ID)
	require.NotEqual(t, transfer1.Amount, transfer2.Amount)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
}
func TestListTransfer(t *testing.T){
	for i:=0;i<10;i++{
		CreateRndomTransfer(t)
	}
	arg:=ListTransferParams{
		Limit :5,
		Offset: 5,
	}
	transfers,err:=testQueries.ListTransfer(context.Background(),arg)
	require.NoError(t,err)
	require.Len(t,transfers,5)


	for _,transfer:=range transfers{
		require.NotEmpty(t,transfer)
	}
}