package tutorial

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

//  Store provides all funcctions to execute db quries and transactions
type Store struct {
	*Queries
	db *pgxpool.Pool
}

// New store creates a new store
func NewStore(db *pgxpool.Pool) *Store {
    return &Store{db: db, Queries: New(db)}
}

// ExecTx creates a function within a databased transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries)error) error {
	tx, err := store.db.BeginTx(ctx, pgx.TxOptions{}) 
    if err!= nil {
        return err
    }
    defer tx.Rollback(ctx)

    queries := New(tx)
    err = fn(queries)
    if err!= nil {
        return err
    }

    return tx.Commit(ctx)
}


// TransferTxParams contains the input parameters of the transfer transaction
type TransferTxParams struct{
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID int64 `json:"to_account_id"`
	Amount int64 `json:"amount"`
}

// TransferTxResult contains the result of the transfer transaction
type TransferTxResult struct{
	Transfer Transfer `json:"tranfer"`
	FromAccountID Account `json:"from_account_id"`
	ToAccountID Account `json:"to_account_id"`
	FromEntry Entry `json:"from_entry"`
	ToEntry Entry `json:"to_entry"`
}

func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error){

	var result TransferTxResult

	err := store.execTx(ctx, func (q *Queries) error  {
		var err error
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
            ToAccountID: arg.ToAccountID,
            Amount: arg.Amount,
		})
		if(err != nil){
			return err
		}
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount: -arg.Amount,
		})
		if(err!= nil){
            return err
        }
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount: arg.Amount,
		})
		if(err!= nil){
            return err
        }

		// TODO: Update account balance
		return nil
	})

	return result, err

}