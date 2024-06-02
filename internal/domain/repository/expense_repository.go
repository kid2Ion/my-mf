package repository

import (
	"my-mf/internal/domain/model"
	"my-mf/internal/infra"

	"github.com/google/uuid"
)

type (
	expenseRepository struct {
		expenseInfra            infra.ExpenseInfra
		relExpenseCategoryInfra infra.RelExpensesCategoriesInfra
	}
	ExpenseRepository interface {
		Create(req *model.Expense, categoryUUID uuid.UUID) error
	}
)

func NewExpenseRepository(expenseInfra infra.ExpenseInfra, relExpenseCategoryInfra infra.RelExpensesCategoriesInfra) ExpenseRepository {
	return &expenseRepository{expenseInfra: expenseInfra, relExpenseCategoryInfra: relExpenseCategoryInfra}
}

func (t *expenseRepository) Create(req *model.Expense, categoryUUID uuid.UUID) error {
	err := t.expenseInfra.Create(req)
	if err != nil {
		return err
	}
	err = t.relExpenseCategoryInfra.Create(req.Uuid, categoryUUID)
	if err != nil {
		return err
	}
	return nil
}
