package db

import "github.com/jackc/pgx/v5/pgxpool"

type Store interface {
	Querier
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	connPool *pgxpool.Pool
	*Queries
}

// NewStore creates a new store
func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}

//type Store struct {
//	*Queries
//	db *sql.DB
//}
//
//func NewStore(db *sql.DB) *Store {
//	return &Store{
//		db:      db,
//		Queries: New(db),
//	}
//}
//
//func (store *Store) execTx(ctx context.Context, fn func(queries *Queries) error) error {
//	tx, err := store.db.BeginTx(ctx, nil)
//	if err != nil {
//		return err
//	}
//	q := New(tx)
//	err = fn(q)
//	if err != nil {
//		if rbErr := tx.Rollback(); rbErr != nil {
//			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
//		}
//		return err
//	}
//	return tx.Commit()
//}
