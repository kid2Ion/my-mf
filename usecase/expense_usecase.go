package usecase

import (
	"log/slog"
	"my-mf/domain/model"
	"my-mf/domain/repository"
	"strconv"

	"github.com/google/uuid"
)

type (
	expenseUsecase struct {
		repo repository.ExpenseRepository
	}
	ExpenseUsecase interface {
		Create(req *ExpenseCreateReq) error
	}
	ExpenseCreateReq struct {
		Amount     string `json:"amount"`
		CategoryID string `json:"category_id"`
	}
)

func NewExpenseUsecase(expenseRepository repository.ExpenseRepository) ExpenseUsecase {
	return &expenseUsecase{repo: expenseRepository}
}

func (t *expenseUsecase) Create(req *ExpenseCreateReq) error {
	intA, err := strconv.Atoi(req.Amount)
	if err != nil {
		slog.Error("failed to convert amount", err)
		return err
	}
	cUUID, err := uuid.Parse(req.CategoryID)
	if err != nil {
		slog.Error("failed to parse category uuid")
		return err
	}
	e := model.NewExpense(intA, cUUID)
	err = t.repo.Create(e)
	if err != nil {
		return err
	}
	return nil
}
