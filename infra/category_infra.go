package infra

import (
	"my-mf/adapter"
	"my-mf/domain/model"
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/slog"
)

type (
	CategoryInfra interface {
		Create(req *model.Category) error
		GetAll() ([]model.Category, error)
	}
	categoryInfra struct {
		adapter.SqlHandler
	}
)

func NewCategoryInfra(sh adapter.SqlHandler) CategoryInfra {
	return &categoryInfra{SqlHandler: sh}
}

func (t *categoryInfra) Create(req *model.Category) error {
	now := time.Now()
	cmd := `insert into categories (uuid, name, created_at, updated_at) values ($1, $2, $3, $4);`
	_, err := t.SqlHandler.Conn.Exec(cmd, req.Uuid, req.Name, now, now)
	if err != nil {
		slog.Error("failed to insert to categories", err.Error())
		return err
	}
	return nil
}

func (t *categoryInfra) GetAll() ([]model.Category, error) {
	q := `select c.uuid, c.name from categories c order by c.created_at`
	rows, err := t.SqlHandler.Conn.Query(q)
	if err != nil {
		slog.Error("failed to get categories", err.Error())
		return nil, err
	}
	res := make([]model.Category, 0)
	defer rows.Close()
	for rows.Next() {
		var rslt struct {
			Uuid uuid.UUID `db:"uuid"`
			Name string    `db:"name"`
		}

		err := rows.Scan(&rslt.Uuid, &rslt.Name)
		if err != nil {
			slog.Error("failed to sqan categories")
			return nil, err
		}
		res = append(res, model.Category{rslt.Uuid, rslt.Name})
	}
	return res, nil
}
