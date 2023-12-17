package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	WithTransaction(ctx context.Context, fn func(pgx.Tx) error) error
}

type StoreImpl struct {
	Database *pgxpool.Pool
}

func NewUserStore(database *pgxpool.Pool) Store {
	return &StoreImpl{
		Database: database,
	}
}

func (r *StoreImpl) WithTransaction(ctx context.Context, fn func(pgx.Tx) error) error {
	tx, err := r.Database.Begin(ctx)
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		_ = tx.Rollback(ctx)
		return err
	}

	return tx.Commit(ctx)
}
