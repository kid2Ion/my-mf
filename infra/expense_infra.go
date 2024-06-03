package infra

import (
	"my-mf/adapter"
	"my-mf/domain/model"
	"time"

	"golang.org/x/exp/slog"
)

type (
	ExpenseInfra interface {
		Create(req *model.Expense) error
	}
	expenseInfra struct {
		adapter.SqlHandler
	}
)

func NewExpenseInfra(sh adapter.SqlHandler) ExpenseInfra {
	return &expenseInfra{SqlHandler: sh}
}

func (t *expenseInfra) Create(req *model.Expense) error {
	now := time.Now()
	cmd := `insert into expenses (uuid, amount, created_at, updated_at) values ($1, $2, $3, $4);`
	_, err := t.SqlHandler.Conn.Exec(cmd, req.Uuid, req.Amount, now, now)
	if err != nil {
		slog.Error("failed to insert to expenses", err.Error())
		return err
	}
	cmd = `insert into rel_expenses_categories (expense_uuid, category_uuid) values ($1, $2);`
	_, err = t.SqlHandler.Conn.Exec(cmd, req.Uuid, req.CategoryUuid)
	if err != nil {
		slog.Error("failed to insert to rel_expenses_categories", err.Error())
		return err
	}
	return nil
}
