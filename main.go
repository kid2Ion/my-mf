package main

import (
	"html/template"
	"io"
	"log"
	"my-mf/router"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET_KEY")))

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
	// session認証
	// session無限なっている（謎）
	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		envUsername := os.Getenv("BASIC_AUTH_USERNAME")
		envPassword := os.Getenv("BASIC_AUTH_PASSWORD")
		if username == envUsername && password == envPassword {
			session, _ := store.Get(c.Request(), os.Getenv("SESSION_SECRET_NAME"))
			session.Values["authenticated"] = true
			session.Options = &sessions.Options{
				MaxAge: 5,
			}
			session.Save(c.Request(), c.Response())
			return true, nil
		}
		return false, nil
	}))
	// session検証
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			session, _ := store.Get(c.Request(), os.Getenv("SESSION_SECRET_NAME"))
			if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "unauthorized",
				})
			}
			return next(c)
		}
	})

	// template
	router.InitTemplateRouter(e)

	// domain
	router.InitExpenseRouter(e)
	router.InitCategoryRouter(e)
	// query
	router.InitQueryExpense(e)
	e.Start(":8080")
}
