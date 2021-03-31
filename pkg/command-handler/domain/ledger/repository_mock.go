package ledger

import (
	"context"
	"time"

	"github.com/stone-co/the-amazing-ledger/pkg/command-handler/domain/ledger/entities"
)

type RepositoryMock struct {
	OnCreateTransaction    func(context.Context, *entities.Transaction) error
	OnLoadObjectsIntoCache func(ctx context.Context, cachedAccounts *entities.CachedAccounts) (entities.Version, error)
	OnGetAccountBalance    func(ctx context.Context, accountName entities.AccountName) (*entities.AccountBalance, error)
	OnGetSyntheticReport   func(ctx context.Context, accountName string, startTime time.Time, endTime time.Time) (*entities.SyntheticReport, error)
}

func (s RepositoryMock) CreateTransaction(ctx context.Context, transaction *entities.Transaction) error {
	return s.OnCreateTransaction(ctx, transaction)
}

func (s RepositoryMock) LoadObjectsIntoCache(ctx context.Context, cachedAccounts *entities.CachedAccounts) (entities.Version, error) {
	return s.OnLoadObjectsIntoCache(ctx, cachedAccounts)
}

func (s RepositoryMock) GetAccountBalance(ctx context.Context, accountName entities.AccountName) (*entities.AccountBalance, error) {
	return s.OnGetAccountBalance(ctx, accountName)
}

func (s RepositoryMock) GetSyntheticReport(ctx context.Context, accountName string, startTime time.Time, endTime time.Time) (*entities.SyntheticReport, error) {
	return s.OnGetSyntheticReport(ctx, accountName, startTime, endTime)
}
