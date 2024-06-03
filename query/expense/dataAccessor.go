package query_expense

import (
	"log/slog"
	"my-mf/adapter"
	"time"

	"github.com/google/uuid"
)

type (
	dataAccessor struct {
		adapter.SqlHandler
	}
	DataAccessor interface {
		GetAllExpensesInThisMonth() ([]getAllExpensesInThisMonthRes, error)
	}
)

func NewDataAccessor(sh adapter.SqlHandler) DataAccessor {
	return &dataAccessor{SqlHandler: sh}
}

func (t *dataAccessor) GetAllExpensesInThisMonth() ([]getAllExpensesInThisMonthRes, error) {
	now := time.Now()
	sDay := time.Date(now.Year(), time.Month(now.Month()), 1, 0, 0, 0, 0, time.Local)
	eDay := sDay.AddDate(0, 1, 0)
	q := `select e.uuid, e.amount, rec.category_uuid, c.name
	from expenses e
	inner join rel_expenses_categories rec
	on rec.expense_uuid = e.uuid
	inner join categories c
	on c.uuid = rec.category_uuid
	where e.created_at >= $1 and e.created_at < $2
	order by e.created_at`
	rows, err := t.SqlHandler.Conn.Query(q, sDay, eDay)
	if err != nil {
		slog.Error("failed to get all in this month", err.Error())
		return nil, err
	}
	res := make([]getAllExpensesInThisMonthRes, 0)
	defer rows.Close()
	for rows.Next() {
		var rslt struct {
			Uuid   uuid.UUID `db:"uuid"`
			Amount int       `db:"amount"`
			CUuid  uuid.UUID `db:"category_uuid"`
			cName  string    `db:"name"`
		}
		err := rows.Scan(&rslt.Uuid, &rslt.Amount, &rslt.CUuid, &rslt.cName)
		if err != nil {
			slog.Error("failed to sqan get all in this month")
			return nil, err
		}
		res = append(res, getAllExpensesInThisMonthRes{
			Uuid:   rslt.Uuid.String(),
			Amount: rslt.Amount,
			Category: CategoryRes{
				Uuid: rslt.CUuid.String(),
				Name: rslt.cName,
			},
		})
	}
	return res, nil
}
