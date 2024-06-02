package repository

import "my-mf/internal/infra"

type (
	expenseRepository struct {
		infra infra.ExpenseInfra
	}
	ExpenseRepository interface {
		Create() error
	}
)

func NewExpenseRepository(expenseInfra infra.ExpenseInfra) ExpenseRepository {
	return &expenseRepository{infra: expenseInfra}
}

func (t *expenseRepository) Create() error {
	return nil
}
