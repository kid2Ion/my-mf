package infra

import "fmt"

type (
	ExpenseInfra interface {
		Create() error
	}
	expenseInfra struct {
		SqlHandler
	}
)

func NewExpenseInfra(sh SqlHandler) ExpenseInfra {
	return &expenseInfra{SqlHandler: sh}
}

func (t *expenseInfra) Create() error {
	fmt.Println("bun bun hello")
	return nil
}
