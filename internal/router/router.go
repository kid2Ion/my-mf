package router

import (
	"my-mf/internal/handler"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func InitTemplateRouter(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		id := uuid.New()
		data := struct {
			UUID string
		}{
			UUID: id.String(),
		}
		return c.Render(http.StatusOK, "index.html", data)
	})
	e.GET("/about", func(c echo.Context) error {
		return c.Render(http.StatusOK, "about.html", nil)
	})
}

// まずは支出(expenses)をcreate, get
// まずは動くようにしよう。その後設計調べて修正していこう
func InitExpenseRouter(e *echo.Echo, handler handler.ExpenseHandler) {
	e.POST("expense/create", handler.Create())
}
