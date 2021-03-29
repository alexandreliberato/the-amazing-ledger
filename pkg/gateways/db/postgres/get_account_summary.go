package postgres

import (
	"context"

	"time"

	"github.com/stone-co/the-amazing-ledger/pkg/command-handler/domain/ledger/entities"
)

func (r *LedgerRepository) GetAccountSummary(ctx context.Context, accountName entities.AccountName, startTime time.Time, endTime time.Time) (*entities.AccountSummary, error) {
	var columns string = `
		SELECT
			account_class
	`

	query := `
			MAX(version) as current_version,
			SUM(CASE operation
				WHEN $1 THEN amount
				ELSE 0
				END) AS credit,
			SUM(CASE operation
				WHEN $2 THEN amount
				ELSE 0
				END) AS debit
		FROM entries
		WHERE 1=1

	`

	var groupBy string = " GROUP BY 1"

	// TODO dates
	if (accountName != entities.AccountName{}) {
		columns += ",account_group"
		query += "AND account_class = $3"
		query += "AND account_group is not null and account_group != '' "
		groupBy += ",2"

		if accountName.Group != "" {
			columns += ",account_subgroup"
			query += `AND account_group = $4`
			query += "AND account_subgroup is not null and account_group != '' "
			groupBy += ",3"

			if accountName.Subgroup != "" {
				columns += ",account_id"
				query += `AND account_subgroup = $5`
				query += "AND account_id is not null and account_group != '' "
				groupBy += ",4"

				if accountName.ID != "" {
					query += `AND account_id = $6`
				}
			}
		}
	}

	finalQuery := columns + query + groupBy

	creditOperation := entities.CreditOperation.String()
	debitOperation := entities.DebitOperation.String()

	rows, errQuery := r.db.Query(
		context.Background(),
		finalQuery,
		creditOperation,
		debitOperation,
		accountName.Class.String(),
		accountName.Group,
		accountName.Subgroup,
		accountName.ID,
	)

	defer rows.Close()

	if errQuery != nil {
		return nil, errQuery
	}

	paths := []entities.Path{}
	var currentVersion uint64
	var totalCredit int
	var totalDebit int

	for rows.Next() {
		var class string
		var group string
		var subgroup string
		var id string

		var credit int
		var debit int

		err := rows.Scan(
			&class,
			&group,
			&subgroup,
			&id,
			&currentVersion,
			&credit,
			&debit,
		)

		if err != nil {
			return nil, err
		}

		acc := class + ":" + group + ":" + subgroup + ":" + id

		path := entities.Path{
			Account: acc,
			Credit:  credit,
			Debit:   debit,
		}

		paths = append(paths, path)

		totalCredit = totalCredit + credit
		totalDebit = totalDebit + debit
	}

	accountSummary, errEntity := entities.NewAccountSummary(totalCredit, totalDebit, paths, entities.Version(currentVersion))
	if errEntity != nil {
		return nil, errEntity
	}

	return accountSummary, nil

}
