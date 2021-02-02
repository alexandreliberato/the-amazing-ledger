package usecase

import (
	"context"

	"time"

	"github.com/jackc/pgx/v4"

	"github.com/stone-co/the-amazing-ledger/pkg/command-handler/domain/ledger/entities"
)

func (l *LedgerUseCase) GetAccountSummary(ctx context.Context, accountName entities.AccountName, startTime time.Time, endTime time.Time) (*entities.AccountBalance, error) {

	accountBalance, err := l.repository.GetAccountSummary(ctx, accountName, startTime, endTime)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, entities.ErrAccountNotFound
		}
		return nil, err
	}

	return accountBalance, nil
}
