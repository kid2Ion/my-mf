package usecase

import "my-mf/internal/domain/repository"

type (
	expenseUsecase struct {
		repo repository.ExpenseRepository
	}
	ExpenseUsecase interface {
		Create() error
	}
	ExpenseCreateReq struct {
		Title  string `json:"title"`
		Amount string `json:"amount"`
	}
)

func NewExpenseUsecase(expenseRepository repository.ExpenseRepository) ExpenseUsecase {
	return &expenseUsecase{repo: expenseRepository}
}

func (t *expenseUsecase) Create() error {
	return nil
}
