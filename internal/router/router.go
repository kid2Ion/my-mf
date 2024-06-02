package router

import (
	"my-mf/internal/di"
	"net/http"

	"github.com/labstack/echo"
)

func InitTemplateRouter(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", nil)
	})
	e.GET("/category", func(c echo.Context) error {
		return c.Render(http.StatusOK, "category.html", nil)
	})
}

// expense
func InitExpenseRouter(e *echo.Echo) {
	handler := di.InjectExpenseHandler()
	e.POST("expense/create", handler.Create())
}

// category
func InitCategoryRouter(e *echo.Echo) {
	handler := di.InjectCategoryHandler()
	e.POST("category/create", handler.Create())
	e.GET("category/get-all", handler.GetAll())
}
