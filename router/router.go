package router

import (
	"my-mf/di"
	"net/http"

	"github.com/labstack/echo/v4"
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

// CQRSを採用。集約を跨ぐqueryが必要な際にはこちら
// query
// expense
func InitQueryExpense(e *echo.Echo) {
	handler := di.InjectQueryExpenseHandler()
	e.GET("expense/get-all-in-this-month", handler.GetAllExpensesInThisMonth())
}
