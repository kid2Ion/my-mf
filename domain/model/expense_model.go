package model

import "github.com/google/uuid"

// relテーブルはcommandsの場合domainとしてはここに所属させる
type Expense struct {
	Uuid         uuid.UUID
	Amount       int
	CategoryUuid uuid.UUID
}

func NewExpense(amount int, cUUID uuid.UUID) *Expense {
	e := new(Expense)
	e.Uuid = uuid.New()
	e.Amount = amount
	e.CategoryUuid = cUUID
	return e
}
