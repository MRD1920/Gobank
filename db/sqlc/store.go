package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
}

// Store provides all functions to execute db queries and transactions.
// The by defaul Queries struct doesn't support tranasctions, so we need to create a new Sturct named store to handle transactions.
type SQLStore struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction.
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {

	//Using default isolation level
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// Create a new Queries struct with the transaction
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()

}

// TransferTxParams contians the input parameters of the transfer transaction
type TransferTxParams struct {
	FromAccountID int64 `json:"from_acount_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// TransferTxResult is the result of the transfer transaction
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

var txKey = struct{}{}

// TransferTx performs a money transfer from one accout to the other.
// It creates a transfer record, add record entry, and update accounts' balance within a single transaction.
func (store *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		txName := ctx.Value(txKey)
		fmt.Println(txName, "create transfer")
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})

		if err != nil {
			return err
		}

		fmt.Println(txName, "create entry 1")
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})

		if err != nil {
			return err
		}

		fmt.Println(txName, "create entry 2")

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    arg.Amount,
		})

		if err != nil {
			return err
		}

		fmt.Println(txName, "get account 1")

		account1, err := q.GetAccountForUpdate(ctx, arg.FromAccountID)
		if err != nil {
			return err
		}

		fmt.Println(txName, "update account 1")
		result.FromAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
			ID:      arg.FromAccountID,
			Balance: account1.Balance - arg.Amount,
		})
		if err != nil {
			return err
		}
		fmt.Println(txName, "get account 2")

		account2, err := q.GetAccountForUpdate(ctx, arg.FromAccountID)
		if err != nil {
			return err
		}

		fmt.Println(txName, "update account 2")

		result.FromAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
			ID:      arg.FromAccountID,
			Balance: account2.Balance + arg.Amount,
		})
		if err != nil {
			return err
		}
		return nil
	})

	return result, err
}
