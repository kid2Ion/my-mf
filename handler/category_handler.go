package handler

import (
	"my-mf/usecase"
	"net/http"

	"golang.org/x/exp/slog"

	"github.com/labstack/echo/v4"
)

type (
	categoryHandler struct {
		usecase usecase.CategoryUsecase
	}
	CategoryHandler interface {
		Create() echo.HandlerFunc
		GetAll() echo.HandlerFunc
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
		err := t.usecase.Create(req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, nil)
	}
}

func (t *categoryHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := t.usecase.GetAll()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, res)
	}
}
