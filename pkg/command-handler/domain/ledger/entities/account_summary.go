package entities

type Path struct {
	Account string
	Credit  int
	Debit   int
}

type AccountSummary struct {
	TotalCredit    int
	TotalDebit     int
	Paths          []Path
	CurrentVersion Version
}

func NewAccountSummary(totalCredit, totalDebit int, paths []Path, version Version) (*AccountSummary, error) {
	if paths == nil || len(paths) < 1 {
		return nil, ErrInvalidAccountSummaryStructure
	}

	return &AccountSummary{
		TotalCredit:    totalCredit,
		TotalDebit:     totalDebit,
		Paths:          paths,
		CurrentVersion: version,
	}, nil
}
