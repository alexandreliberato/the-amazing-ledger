package ledger

import (
	"context"

	"time"

	"github.com/stone-co/the-amazing-ledger/pkg/command-handler/domain/ledger/entities"
)

type Repository interface {
	CreateTransaction(context.Context, *entities.Transaction) error
	LoadObjectsIntoCache(ctx context.Context, objects *entities.CachedAccounts) (entities.Version, error)
	GetAccountBalance(ctx context.Context, accountName entities.AccountName) (*entities.AccountBalance, error)
	GetSyntheticReport(ctx context.Context, accountName string, startTime time.Time, endTime time.Time) (*entities.SyntheticReport, error)
}
