package repository

import (
	"my-mf/domain/model"
	"my-mf/infra"
)

type (
	// expense集約(expense, rel_expense_categoryモデル)
	expenseRepository struct {
		expenseInfra infra.ExpenseInfra
	}
	ExpenseRepository interface {
		Create(req *model.Expense) error
	}
)

func NewExpenseRepository(expenseInfra infra.ExpenseInfra) ExpenseRepository {
	return &expenseRepository{expenseInfra: expenseInfra}
}

func (t *expenseRepository) Create(req *model.Expense) error {
	err := t.expenseInfra.Create(req)
	if err != nil {
		return err
	}
	return nil
}
