package adapter

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() *SqlHandler {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PW"),
		os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PORT"))
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	cmd := `
		create table if not exists expenses (
			uuid uuid primary key not null,
			amount integer not null,
			created_at timestamp not null,
			updated_at timestamp not null
		);
		create table if not exists categories (
			uuid uuid primary key not null,
			name text not null,
			created_at timestamp not null,
			updated_at timestamp not null
		);
		create table if not exists rel_expenses_categories (
			expense_uuid uuid,
			category_uuid uuid,
			primary key (expense_uuid, category_uuid),
			-- 支出が削除された場合はリレーションレコードも削除
			foreign key (expense_uuid) references expenses(uuid) on delete cascade,
			-- カテゴリーが削除された場合はリレーションレコードにnullをセット
			foreign key(category_uuid) references categories(uuid) on delete set null
		);
		`
	_, err = conn.Exec(cmd)
	if err != nil {
		panic(err)
	}
	return &SqlHandler{Conn: conn}
}
