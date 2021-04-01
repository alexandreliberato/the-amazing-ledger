package ledger

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/stone-co/the-amazing-ledger/pkg/command-handler/domain/ledger/entities"
)

type TransactionsMock struct {
	OnCreateTransaction    func(ctx context.Context, id uuid.UUID, entries []entities.Entry) error
	OnLoadObjectsIntoCache func(ctx context.Context) error
	OnGetAccountBalance    func(ctx context.Context, accountName entities.AccountName) (*entities.AccountBalance, error)
	OnGetSyntheticReport   func(ctx context.Context, accountName entities.AccountName, startTime time.Time, endTime time.Time) (*entities.SyntheticReport, error)
}

func (m TransactionsMock) CreateTransaction(ctx context.Context, id uuid.UUID, entries []entities.Entry) error {
	return m.OnCreateTransaction(ctx, id, entries)
}

func (m TransactionsMock) LoadObjectsIntoCache(ctx context.Context) error {
	return m.OnLoadObjectsIntoCache(ctx)
}

func (m TransactionsMock) GetAccountBalance(ctx context.Context, accountName entities.AccountName) (*entities.AccountBalance, error) {
	return m.OnGetAccountBalance(ctx, accountName)
}

func SuccessfulTransactionMock() TransactionsMock {
	return TransactionsMock{
		OnCreateTransaction: func(ctx context.Context, id uuid.UUID, entries []entities.Entry) error {
			return nil
		},
		OnLoadObjectsIntoCache: func(ctx context.Context) error {
			return nil
		},
		OnGetAccountBalance: func(ctx context.Context, accountName entities.AccountName) (*entities.AccountBalance, error) {
			return &entities.AccountBalance{
				AccountName:    entities.AccountName{},
				CurrentVersion: 0,
				TotalCredit:    0,
				TotalDebit:     0,
			}, nil
		},
		OnGetSyntheticReport: func(ctx context.Context, accountName entities.AccountName, startTime time.Time, endTime time.Time) (*entities.SyntheticReport, error) {
			return &entities.SyntheticReport{
				CurrentVersion: 0,
				TotalCredit:    0,
				TotalDebit:     0,
				Paths: []entities.Path{{
					Account: "",
					Debit:   0,
					Credit:  0,
				}},
			}, nil
		},
	}
}
