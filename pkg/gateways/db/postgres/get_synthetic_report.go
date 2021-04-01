package postgres

import (
	"context"
	"strings"
	"time"

	"github.com/stone-co/the-amazing-ledger/pkg/command-handler/domain/ledger/entities"
)

func (r *LedgerRepository) GetSyntheticReport(ctx context.Context, accountName string, startTime time.Time, endTime time.Time) (*entities.SyntheticReport, error) {
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
	paths := strings.Split(accountName, ":")

	// account.Class
	if len(paths) >= 1 {
		columns += ",account_group"
		query += "AND account_class = $3"
		query += "AND account_group is not null and account_group != '' "
		groupBy += ",2"

		// accountName.Group
		if len(paths) >= 1 {
			columns += ",account_subgroup"
			query += `AND account_group = $4`
			query += "AND account_subgroup is not null and account_group != '' "
			groupBy += ",3"

			// accountName.Subgroup
			if len(paths) >= 1 {
				columns += ",account_id"
				query += `AND account_subgroup = $5`
				query += "AND account_id is not null and account_group != '' "
				groupBy += ",4"

				// accountName.ID
				if len(paths) >= 1 {
					query += `AND account_id = $6`
				}
			}
		}
	}

	var dates string

	if !startTime.IsZero() {
		dates += "AND date_part('year', created_at) >= $7 and date_part('year', created_at)  <= coalesce($8, date_part('year', created_at)::integer)"
		dates += "AND date_part('month', created_at) >= $9 and date_part('month', created_at)  <= coalesce($10, date_part('month', created_at)::integer)"
		dates += "AND date_part('day', created_at) >= $11 and date_part('day', created_at)  <= coalesce($12, date_part('day', created_at)::integer)"
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
		startTime.Year,
		endTime.Year,
		startTime.Month,
		endTime.Month,
		startTime.Day,
		endTime.Day,
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

	syntheticReport, errEntity := entities.NewSyntheticReport(totalCredit, totalDebit, paths, entities.Version(currentVersion))
	if errEntity != nil {
		return nil, errEntity
	}

	return syntheticReport, nil

}
