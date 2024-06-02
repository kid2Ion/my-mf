package handler

import (
	"fmt"
	"my-mf/internal/usecase"
	"net/http"

	"golang.org/x/exp/slog"

	"github.com/labstack/echo"
)

type (
	categoryHandler struct {
		usecase usecase.CategoryUsecase
	}
	CategoryHandler interface {
		Create() echo.HandlerFunc
	}
)

func NewCategoryHandler(categoryUsecase usecase.CategoryUsecase) CategoryHandler {
	return &categoryHandler{usecase: categoryUsecase}
}

func (t *categoryHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(usecase.CategoryCreateReq)
		if err := c.Bind(req); err != nil {
			slog.Error("failed to bind: %s", err.Error())
			return c.JSON(http.StatusBadRequest, err)
		}
		fmt.Println("---")
		fmt.Println(req)
		fmt.Println("---")
		return c.JSON(http.StatusOK, nil)
	}
}
