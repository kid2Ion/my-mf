package query_expense

import (
	"net/http"

	"github.com/labstack/echo"
)

type (
	handler struct {
		usecase Usecase
	}
	Handler interface {
		GetAllExpensesInThisMonth() echo.HandlerFunc
	}
)

func NewHandler(usecase Usecase) Handler {
	return &handler{usecase: usecase}
}

func (t *handler) GetAllExpensesInThisMonth() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := t.usecase.GetAllExpensesInThisMonth()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, res)
	}
}
