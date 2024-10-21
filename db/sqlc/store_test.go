package tutorial

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTrasnsferTx(t *testing.T){
	store := NewStore(testDB)

	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)

	// run a cuncurrent transfer transaction
	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
    go func() {
        result, err := store.TransferTx(context.Background(), TransferTxParams{
            FromAccountID: account1.ID,
            ToAccountID:   account2.ID,
            Amount:        amount,
        })
        errs <- err
        if err == nil {
            results <- result
        } else {
            // Send an empty result to avoid blocking
            results <- TransferTxResult{}
        }
    }()
	}

	for i := 0; i < n; i++ {
		err := <-errs
        require.NoError(t, err)

		result := <- results
		require.NotEmpty(t, result)

		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		// check acccount entries
		toEntry := result.FromEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, account2.ID, toEntry.AccountID)
		require.Equal(t, -amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		
		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		//TODO check acccount balance

	}
}