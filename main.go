package main

import (
	"html/template"
	"io"
	"log"
	"log/slog"
	"my-mf/router"
	"net/http"
	"os"

	"github.com/joho/godotenv"
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

	if os.Getenv("GO_EXEC_ENV") == "local" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	allowedIPs := map[string]bool{
		os.Getenv(("IP1")): true,
		os.Getenv(("IP2")): true,
		os.Getenv(("IP3")): true,
	}
	ipRestrictionMiddleware := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ip := c.RealIP()
			if _, allowed := allowedIPs[ip]; !allowed {
				slog.Error("Access denied: %s", ip)
				return c.JSON(http.StatusForbidden, map[string]string{
					"message": "Access denied",
				})
			}
			return next(c)
		}
	}
	e.Use(ipRestrictionMiddleware)

	// template
	router.InitTemplateRouter(e)

	// domain
	router.InitExpenseRouter(e)
	router.InitCategoryRouter(e)
	// query
	router.InitQueryExpense(e)
	e.Start(":8080")
}
