package infra

import (
	"github.com/google/uuid"
	"golang.org/x/exp/slog"
)

type (
	RelExpensesCategoriesInfra interface {
		Create(eUUID uuid.UUID, cUUID uuid.UUID) error
	}
	relExpensesCategoriesInfra struct {
		SqlHandler
	}
)

func NewRelExpensesCategoriesInfra(sh SqlHandler) RelExpensesCategoriesInfra {
	return &relExpensesCategoriesInfra{SqlHandler: sh}
}

func (t *relExpensesCategoriesInfra) Create(eUUID uuid.UUID, cUUID uuid.UUID) error {
	cmd := `insert into rel_expenses_categories (expense_uuid, category_uuid) values ($1, $2);`
	_, err := t.SqlHandler.Conn.Exec(cmd, eUUID, cUUID)
	if err != nil {
		slog.Error("failed to insert to rel_expenses_categories", err.Error())
		return err
	}
	return nil
}
