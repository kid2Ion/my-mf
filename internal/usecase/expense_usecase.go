package usecase

import "my-mf/internal/domain/repository"

type (
	expenseUsecase struct {
		repo repository.ExpenseRepository
	}
	ExpenseUsecase interface {
		Create() error
	}
)

func NewExpenseUsecase(expenseRepository repository.ExpenseRepository) ExpenseUsecase {
	return &expenseUsecase{repo: expenseRepository}
}

func (t *expenseUsecase) Create() error {
	return nil
}
