package adapter

import (
	"database/sql"

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
