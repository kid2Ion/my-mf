package model

import "github.com/google/uuid"

type Expense struct {
	Uuid   uuid.UUID
	Amount int
}

func NewExpense(amount int) *Expense {
	e := new(Expense)
	e.Uuid = uuid.New()
	e.Amount = amount
	return e
}
