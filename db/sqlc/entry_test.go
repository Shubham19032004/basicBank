package db

import (
	"bank/utils"
	"context"
	"database/sql"

	"testing"

	"github.com/stretchr/testify/require"
)

func CreateRendomEntry(t *testing.T)(int64,Entry){
	account:=randomAccount(t);
	args:=CreateEntryParams{
		AccountID:account.ID,
		Amount:utils.RandomMoney(),
	}
	
	entry,err:=testQueries.CreateEntry(context.Background(),args);
	require.NoError(t,err)
	require.NotEmpty(t,entry.Amount)
	require.Equal(t,account.ID,entry.AccountID)
	return entry.ID,entry
}
func TestCreateEnity(t *testing.T) {
	CreateRendomEntry(t)

}
func TestGetEntry(t *testing.T){
	entryId,entry:=CreateRendomEntry(t)

	entry2,err:=testQueries.GetEntry(context.Background(),entryId);
	require.NoError(t,err)
	require.Equal(t, entry.ID, entry2.ID)
	require.Equal(t, entry.AccountID, entry2.AccountID)
	require.Equal(t, entry.Amount, entry2.Amount)
}

func TestUpdateEntry(t *testing.T){
	entryId,entry1:=CreateRendomEntry(t)
	arr:=UpdateEntryParams{
		Amount:utils.RandomMoney(),
		ID: entryId,
	}
	entry2,err:=testQueries.UpdateEntry(context.Background(),arr)
	require.NoError(t, err)
	require.NotEmpty(t,entry1)
	require.NotEmpty(t,entry2)
	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.NotEqual(t, entry1.Amount, entry2.Amount)
}
func TestDeleteEntry(t *testing.T){
	entryId,_:=CreateRendomEntry(t)
	err:=testQueries.DeleteEntry(context.Background(),entryId);
	require.NoError(t,err)
	entry2,err:=testQueries.GetEntry(context.Background(),entryId)
	require.Error(t,err)
	require.EqualError(t,err,sql.ErrNoRows.Error())
	require.Empty(t,entry2)
}
func TestListEntry(t *testing.T){
	for i:=0;i<10;i++{
		CreateRendomEntry(t)
	}
	arg:=ListEntryParams{
		Limit :5,
		Offset: 5,
	}
	entry,err:=testQueries.ListEntry(context.Background(),arg)
	require.NoError(t,err)
	require.Len(t,entry,5)


	for _,e:=range entry{
		require.NotEmpty(t,e)
	}
}