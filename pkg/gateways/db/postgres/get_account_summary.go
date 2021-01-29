package postgres

import (
	"context"

	"time"

	"github.com/stone-co/the-amazing-ledger/pkg/command-handler/domain/ledger/entities"
)

func (r *LedgerRepository) GetAccountSummary(ctx context.Context, accountName entities.AccountName, startTime time.Time, endTime time.Time) (*entities.AccountBalance, error) {
	query := `
		SELECT
			account_class,
			account_group,
			account_subgroup,
			account_id,
			MAX(version) as current_version,
			SUM(CASE operation
				WHEN $1 THEN amount
				ELSE 0
				END) AS total_credit,
			SUM(CASE operation
				WHEN $2 THEN amount
				ELSE 0
				END) AS total_debit
		FROM entries
		WHERE 1=1

	`

	if (accountName != entities.AccountName{}) {
		query += `AND account_class = ` + accountName.Class.String()
	}

	if accountName.Group != "" {
		query += `AND account_group = ` + accountName.Group
	}

	if accountName.Subgroup != "" {
		query += `AND account_subgroup = ` + accountName.Subgroup
	}

	if accountName.ID != "" {
		query += `AND account_id = ` + accountName.ID
	}

	creditOperation := entities.CreditOperation.String()
	debitOperation := entities.DebitOperation.String()

	row := r.db.QueryRow(
		context.Background(),
		query,
		creditOperation,
		debitOperation,
	)

	var currentVersion uint64
	var totalCredit int
	var totalDebit int

	err := row.Scan(
		nil,
		nil,
		nil,
		nil,
		&currentVersion,
		&totalCredit,
		&totalDebit,
	)

	if err != nil {
		return nil, err
	}

	accountBalance := entities.NewAccountBalance(accountName, entities.Version(currentVersion), totalCredit, totalDebit)
	return accountBalance, nil

}
