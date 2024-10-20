package tutorial

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"tutorial.sqlc.dev/app/db/util"
)

func CreateRandomAccount(t *testing.T) Account{
	arg := CreateAccountParams{
		Owner: util.RandomOwner(),
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQuries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T){
	CreateRandomAccount(t)
}

func TestGetAccount(t *testing.T){

	account1 := CreateRandomAccount(t)
	account2, err := testQuries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt.Time, account2.CreatedAt.Time, time.Second)
}


func TestUpdateAccount(t *testing.T){

	// Create a random account
    account1 := CreateRandomAccount(t)

    // Prepare the update parameters
    arg := UpdateAccountParams{
        ID:      account1.ID,
        Balance: util.RandomMoney(),
    }

    err := testQuries.UpdateAccount(context.Background(), arg)
    require.NoError(t, err)

    // Fetch the updated account to verify the changes
    account2, err := testQuries.GetAccount(context.Background(), account1.ID)
    require.NoError(t, err)
    require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt.Time, account2.CreatedAt.Time, time.Second)
}


func TestDeleteAccount(t *testing.T) {
    // Create a random account
    account1 := CreateRandomAccount(t)

    // Attempt to delete the account
    err := testQuries.DeleteAccount(context.Background(), account1.ID)
    require.NoError(t, err, "Failed to delete account")

    // Attempt to fetch the account after deletion
    account2, err := testQuries.GetAccount(context.Background(), account1.ID)

    // Ensure the error returned is of type sql.ErrNoRows
    require.Error(t, err, "Expected an error when fetching a deleted account")
    require.True(t, errors.Is(err, sql.ErrNoRows), "Expected sql.ErrNoRows error")

    // Ensure no account is returned
    require.Empty(t, account2, "Expected empty account after deletion")
}


func TestListsAccount(t *testing.T) {
    // Create a random account

	for i := 0; i < 10; i++ {
		CreateRandomAccount(t)
	}

    arg := ListAccountsParams{
		Limit: 5,
		Offset: 5,
	} 

    accounts, err := testQuries.ListAccounts(context.Background(), arg)
 	require.NoError(t, err)
    require.Len(t, accounts, 5)
	
	for _, account := range accounts {
		require.NotEmpty(t, account)
	}

}

