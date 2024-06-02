package infra

import (
	"my-mf/internal/domain/model"
	"time"

	"golang.org/x/exp/slog"
)

type (
	ExpenseInfra interface {
		Create(req *model.Expense) error
	}
	expenseInfra struct {
		SqlHandler
	}
)

func NewExpenseInfra(sh SqlHandler) ExpenseInfra {
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
	return nil
}
