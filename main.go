package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type UUIDreq struct {
	UUID string `json: "uuid"`
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

	e.GET("/", func(c echo.Context) error {
		id := uuid.New()
		data := struct {
			UUID string
		}{
			UUID: id.String(),
		}
		return c.Render(http.StatusOK, "index.html", data)
	})
	e.POST("/uuid", func(c echo.Context) error {
		req := new(UUIDreq)
		if err := c.Bind(req); err != nil {
			return err
		}
		fmt.Println("---")
		fmt.Println(req)
		fmt.Println("---")
		uuid := uuid.New()
		fmt.Println("---")
		fmt.Println(uuid)
		fmt.Println("---")
		return c.JSON(http.StatusOK, map[string]string{
			"uuid": uuid.String(),
		})
	})
	e.Start(":8080")
}
