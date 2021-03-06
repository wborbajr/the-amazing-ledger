package postgres

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/stone-co/the-amazing-ledger/pkg/command-handler/domain/ledger/entities"
)

func (r *LedgerRepository) CreateTransaction(ctx context.Context, transaction *entities.Transaction) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		_ = tx.Rollback(ctx)
	}()

	query := `
		INSERT INTO
			entries (
				id,
				account_id,
	  			operation,
				amount,
				version,
				transaction_id
			) VALUES ($1, $2, $3, $4, $5, $6)
	`

	var batch pgx.Batch

	for _, entry := range transaction.Entries {
		batch.Queue(
			query,
			entry.ID,
			entry.Account.Name(),
			entry.Operation.String(),
			entry.Amount,
			entry.Version,
			transaction.ID,
		)
	}

	br := tx.SendBatch(ctx, &batch)
	err = br.Close()
	if err != nil {
		// TODO: assuming that is duplicate key.
		return entities.ErrIdempotencyKey
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}
