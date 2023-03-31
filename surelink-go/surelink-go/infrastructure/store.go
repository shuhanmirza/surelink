package infrastructure

import (
	"context"
	"database/sql"
	"fmt"
	"surelink-go/infrastructure/sqlc"
)

// Store provides all functions to execute Db Queries and transactions
type Store struct {
	Queries *sqlc.Queries
	Db      *sql.DB
}

// NewStore creates a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		Queries: sqlc.New(db),
		Db:      db,
	}
}

// execTx executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(queries *sqlc.Queries) error) error {
	tx, err := store.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := sqlc.New(tx)
	err = fn(q)
	if err != nil {
		rbErr := tx.Rollback()
		if rbErr != nil {
			return fmt.Errorf("tx err: %v ---- rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
