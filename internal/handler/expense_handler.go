package handler

import (
	"my-mf/internal/usecase"
	"net/http"

	"golang.org/x/exp/slog"

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
		req := new(usecase.ExpenseCreateReq)
		if err := c.Bind(req); err != nil {
			slog.Error("failed to bind: %s", err.Error())
			return c.JSON(http.StatusBadRequest, err)
		}
		err := t.usecase.Create(req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, nil)
	}
}
