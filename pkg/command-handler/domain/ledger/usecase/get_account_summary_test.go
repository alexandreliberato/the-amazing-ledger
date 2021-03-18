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

		accountName, err := entities.NewAccountName("liability")
		accountNameStr := "liability"
		paths := []entities.Path{{accountNameStr, 10, 20}}

		accountSummary, err := entities.NewAccountSummary(totalCredit, totalDebit, paths)
		assert.Nil(t, err)

		date := time.Now()

		useCase := newFakeGetAccountSummary(accountSummary, date, nil)
		a, err := useCase.GetAccountSummary(context.Background(), accountName, date, date)

		assert.Nil(t, err)
		assert.Equal(t, accountSummary.TotalDebit, a.TotalDebit)
	})
}
