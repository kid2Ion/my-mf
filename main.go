package main

import (
	"html/template"
	"io"
	"my-mf/internal/di"
	"my-mf/internal/router"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "static")

	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = t

	router.InitTemplateRouter(e)
	expenseHandler := di.InjectExpenseHandler()
	router.InitExpenseRouter(e, expenseHandler)
	e.Start(":8080")
}
