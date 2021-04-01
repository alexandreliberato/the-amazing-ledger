package ledger

import (
	"context"

	"time"

	"github.com/google/uuid"
	"github.com/stone-co/the-amazing-ledger/pkg/command-handler/domain/ledger/entities"
)

type UseCase interface {
	CreateTransaction(ctx context.Context, id uuid.UUID, entries []entities.Entry) error
	LoadObjectsIntoCache(ctx context.Context) error
	GetAccountBalance(ctx context.Context, accountName entities.AccountName) (*entities.AccountBalance, error)
	GetSyntheticReport(ctx context.Context, accountName string, startTime time.Time, endTime time.Time) (*entities.SyntheticReport, error)
}
