package infra

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() *SqlHandler {
	conn, err := sql.Open("sqlite3", "db/my-mf.db")
	if err != nil {
		panic(err)
	}
	tName := "expenses"
	cmd := fmt.Sprintf(`
		create table if not exists %s (
			uuid uuid not null,
			title string not null,
			created_at timestamp not null,
			updated_at timestamp not null
		)`, tName)
	_, err = conn.Exec(cmd)
	if err != nil {
		panic(err)
	}
	return &SqlHandler{Conn: conn}
}
