package usecase

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/stone-co/the-amazing-ledger/pkg/command-handler/domain/ledger"
	"github.com/stone-co/the-amazing-ledger/pkg/command-handler/domain/ledger/entities"
)

func newFakeCreateTransactionUseCase(result error) *LedgerUseCase {
	log := logrus.New()

	mockRepository := &ledger.RepositoryMock{}
	mockRepository.OnCreateTransaction = func(context.Context, *entities.Transaction) error {
		return result
	}

	return NewLedgerUseCase(log, mockRepository)
}

func newFakeLoadObjectsIntoCacheUseCase(maxVersion entities.Version, result error) *LedgerUseCase {
	log := logrus.New()

	mockRepository := &ledger.RepositoryMock{
		OnLoadObjectsIntoCache: func(ctx context.Context, cachedAccounts *entities.CachedAccounts) (entities.Version, error) {
			return maxVersion, result
		},
	}

	return NewLedgerUseCase(log, mockRepository)
}

func newFakeGetAccountBalance(accountBalance *entities.AccountBalance, result error) *LedgerUseCase {
	log := logrus.New()

	mockRepository := &ledger.RepositoryMock{
		OnGetAccountBalance: func(ctx context.Context, accountName entities.AccountName) (*entities.AccountBalance, error) {
			return accountBalance, result
		},
	}

	return NewLedgerUseCase(log, mockRepository)
}

func newFakeGetSyntheticReport(syntheticReport *entities.SyntheticReport, date time.Time, result error) *LedgerUseCase {
	log := logrus.New()

	mockRepository := &ledger.RepositoryMock{
		OnGetSyntheticReport: func(ctx context.Context, accountName string, startTime time.Time, endTime time.Time) (*entities.SyntheticReport, error) {
			return syntheticReport, result
		},
	}

	return NewLedgerUseCase(log, mockRepository)
}
