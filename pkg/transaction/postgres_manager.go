package transaction

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type PostgresTxManager struct {
	conn *pgx.Conn
}

func NewPostgresTxManager(conn *pgx.Conn) *PostgresTxManager {
	return &PostgresTxManager{conn: conn}
}

func (manager *PostgresTxManager) RunTransaction(ctx context.Context, fn func(context.Context) error) error {
	return manager.RunTransactionWithOptions(ctx, fn, nil)
}

func (manager *PostgresTxManager) RunTransactionWithOptions(ctx context.Context, fn func(context.Context) error, options *pgx.TxOptions) error {
	if options == nil {
		options = &pgx.TxOptions{}
	}

	tx, err := manager.conn.BeginTx(ctx, *options)
	if err != nil {
		return err
	}

	contextWithTx := Inject(ctx, &tx)
	if err = fn(contextWithTx); err != nil {
		if errRollback := tx.Rollback(contextWithTx); errRollback != nil {
			return errRollback
		}
		return err
	}

	if err = tx.Commit(contextWithTx); err != nil {
		if errRollback := tx.Rollback(contextWithTx); errRollback != nil {
			return errRollback
		}
		return err
	}

	return nil
}
