package usecase

import (
	"context"

	"time"

	"github.com/jackc/pgx/v4"

	"github.com/stone-co/the-amazing-ledger/pkg/command-handler/domain/ledger/entities"
)

func (l *LedgerUseCase) GetSyntheticReport(ctx context.Context, accountName string, startTime time.Time, endTime time.Time) (*entities.SyntheticReport, error) {

	syntheticReport, err := l.repository.GetSyntheticReport(ctx, accountName, startTime, endTime)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, entities.ErrAccountNotFound
		}
		return nil, err
	}

	return syntheticReport, nil
}
