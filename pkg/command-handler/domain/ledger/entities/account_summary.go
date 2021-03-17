package entities

type Path struct {
	Account string
	Credit  int
	Debit   int
}

type AccountSummary struct {
	TotalCredit int
	TotalDebit  int
	Paths       []Path
}

func NewAccountSummary(totalCredit, totalDebit int, paths []Path) (*AccountSummary, error) {

	if paths == nil || len(paths) < 1 {
		return nil, ErrInvalidAccountSummaryStructure
	}

	// TODO check empty path account

	return &AccountSummary{
		TotalCredit: totalCredit,
		TotalDebit:  totalDebit,
		Paths:       paths,
	}, nil
}
