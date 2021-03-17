package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/stone-co/the-amazing-ledger/pkg/command-handler/domain/ledger/entities"
)

func TestLedgerUseCase_GetAccountSummary(t *testing.T) {
	t.Run("GetAccountSummary with startDate and endDate being equals must return the summary for the year", func(t *testing.T) {
		totalCredit := 150
		totalDebit := 130

		accountName := "liability"
		paths := []entities.Path{{accountName, 10, 20}}

		accountBalance, err := entities.NewAccountSummary(totalCredit, totalDebit, paths)
		assert.Nil(t, err)

		//TODO add date

		date := time.Now()

		useCase := newFakeGetAccountSummary(accountSummary, date, nil)
		a, err := useCase.GetAccountSummary(context.Background(), accountBalance.AccountName, date, date)

		assert.Nil(t, err)
		assert.Equal(t, accountBalance.TotalCredit, a.TotalCredit)
		assert.Equal(t, accountBalance.TotalDebit, a.TotalDebit)
		assert.Equal(t, expectedBalance, a.Balance())
	})
	t.Run("GetAccountSummary with no endDate should return the summary since the startDate until now", func(t *testing.T) {
		totalCredit := 150
		totalDebit := 130
		expectedBalance := totalCredit - totalDebit

		accountName, err := entities.NewAccountName("liability:stone:clients:user-1")
		assert.Nil(t, err)

		accountBalance := entities.NewAccountBalance(*accountName, 3, totalCredit, totalDebit)
		//TODO add date

		date := time.Now()

		useCase := newFakeGetAccountSummary(accountBalance, date, nil)
		a, err := useCase.GetAccountSummary(context.Background(), accountBalance.AccountName, date, time.Time{})
		assert.Nil(t, err)
		assert.Equal(t, accountBalance.TotalCredit, a.TotalCredit)
		assert.Equal(t, accountBalance.TotalDebit, a.TotalDebit)
		assert.Equal(t, expectedBalance, a.Balance())
	})

	t.Run("The max version for account path must be version in account balance", func(t *testing.T) {
		expectedVersion := entities.Version(5)

		accountName, err := entities.NewAccountName("liability:stone:clients:user-1")
		assert.Nil(t, err)

		accountBalance := entities.NewAccountBalance(*accountName, expectedVersion, 0, 0)

		useCase := newFakeGetAccountBalance(accountBalance, nil)
		a, err := useCase.GetAccountBalance(context.Background(), accountBalance.AccountName)

		assert.Nil(t, err)
		assert.Equal(t, expectedVersion, a.CurrentVersion)
	})
}
