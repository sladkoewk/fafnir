package models

import "time"

type Transaction struct {
	Type       string
	Comment    string
	Author     string
	Category   string
	Price      uint64
	CommitDate time.Time
}

const (
	INCOME    = "Income"
	COST      = "Cost"
	DIVIDENDS = "Dividends"
)
