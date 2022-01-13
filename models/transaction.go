package models

import "time"

type Type string

const (
	Income  Type = "Расход"
	Expense Type = "Доход"
)

type Transaction struct {
	Type       Type
	Comment    string
	Author     string
	Category   string
	Price      float64
	Date       time.Time
	CommitDate time.Time
}
