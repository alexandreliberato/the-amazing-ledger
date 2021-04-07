package postgres

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/stone-co/the-amazing-ledger/pkg/command-handler/domain/ledger/entities"
)

func (r *LedgerRepository) GetSyntheticReport(ctx context.Context, accountName string, startTime time.Time, endTime time.Time) (*entities.SyntheticReport, error) {
	var index uint16 = 2
	var params = []string{}

	var columns string = `
		SELECT
			account_class
	`

	var query string = `
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

	var paths []string

	if accountName != "" {
		paths = strings.Split(accountName, ":")
	}

	var class string
	var group string
	var subgroup string
	var id string

	if len(paths) >= 1 {
		class = paths[0]

		index++

		columns += ",account_group"
		query += fmt.Sprint("AND account_class = $", index)
		query += "AND account_group is not null and account_group != '' "
		groupBy += ",2"

		params = append(params, class)

		if len(paths) >= 1 {
			group = paths[1]

			index++

			columns += ",account_subgroup"
			query += fmt.Sprint(`AND account_group = $`, index)
			query += "AND account_subgroup is not null and account_group != '' "
			groupBy += ",3"

			params = append(params, group)

			if len(paths) >= 1 {
				subgroup = paths[2]

				index++

				columns += ",account_id"
				query += fmt.Sprint(`AND account_subgroup = $`, index)
				query += "AND account_id is not null and account_group != '' "
				groupBy += ",4"

				params = append(params, subgroup)

				if len(paths) >= 1 {
					id = paths[3]
					index++
					query += fmt.Sprint(`AND account_id = $`, index)
					params = append(params, id)
				}
			}
		}
	}

	var dates string

	if !startTime.IsZero() {
		dates += "AND date_part('year', created_at) >= $7 and date_part('year', created_at)  <= coalesce($8, date_part('year', created_at)::integer)"
		dates += "AND date_part('month', created_at) >= $9 and date_part('month', created_at)  <= coalesce($10, date_part('month', created_at)::integer)"
		dates += "AND date_part('day', created_at) >= $11 and date_part('day', created_at)  <= coalesce($12, date_part('day', created_at)::integer)"

		startYear := startTime.Year()
		startMonth := int(startTime.Month())
		startDay := startTime.Day()

		var endYear string
		var endMonth string
		var endDay string

		params = append(params, strconv.Itoa(startYear))
		params = append(params, strconv.Itoa(startMonth))
		params = append(params, strconv.Itoa(startDay))

		if !endTime.IsZero() {
			endYear = strconv.Itoa(endTime.Year())
			endMonth = strconv.Itoa(int(endTime.Month()))
			endYear = strconv.Itoa(endTime.Day())

			params = append(params, endYear)
			params = append(params, endMonth)
			params = append(params, endDay)
		}
	}

	finalQuery := columns + query + groupBy

	creditOperation := entities.CreditOperation.String()
	debitOperation := entities.DebitOperation.String()

	rows, errQuery := r.db.Query(
		context.Background(),
		finalQuery,
		creditOperation, // obrigatorio
		debitOperation,  // obrigatorio
		params,
	)

	defer rows.Close()

	if errQuery != nil {
		return nil, errQuery
	}

	pathsReport := []entities.Path{}
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

		pathsReport = append(pathsReport, path)

		totalCredit = totalCredit + credit
		totalDebit = totalDebit + debit
	}

	syntheticReport, errEntity := entities.NewSyntheticReport(totalCredit, totalDebit, pathsReport, entities.Version(currentVersion))
	if errEntity != nil {
		return nil, errEntity
	}

	return syntheticReport, nil

}
