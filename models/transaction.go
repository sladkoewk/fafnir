package models

type Type string

const (
	Income  Type = "доход"
	Expense Type = "расход"
)

type Transaction struct {
	Comment    string
	Author     string
	Category   string
	Price      float64
	Сurrency   string
	Date       string
	CommitDate string
	Type       Type
}
