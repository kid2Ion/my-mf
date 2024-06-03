package query_expense

type (
	usecase struct {
		da DataAccessor
	}
	Usecase interface {
		GetAllExpensesInThisMonth() ([]getAllExpensesInThisMonthRes, error)
	}
	getAllExpensesInThisMonthRes struct {
		Uuid     string      `json:"uuid"`
		Amount   int         `json:"amount"`
		Category CategoryRes `json:"category"`
	}
	CategoryRes struct {
		Uuid string `json:"uuid"`
		Name string `json:"name"`
	}
)

func NewUsecase(da DataAccessor) Usecase {
	return &usecase{da: da}
}

func (t *usecase) GetAllExpensesInThisMonth() ([]getAllExpensesInThisMonthRes, error) {
	return t.da.GetAllExpensesInThisMonth()
}
