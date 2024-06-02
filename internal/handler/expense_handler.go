package handler

import (
	"my-mf/internal/usecase"
	"net/http"

	"github.com/labstack/echo"
)

type (
	expenseHandler struct {
		usecase usecase.ExpenseUsecase
	}
	ExpenseHandler interface {
		Create() echo.HandlerFunc
	}
)

func NewExpenseHandler(expenseUsecase usecase.ExpenseUsecase) ExpenseHandler {
	return &expenseHandler{usecase: expenseUsecase}
}

func (t *expenseHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	}
}
