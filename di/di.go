package di

import (
	"my-mf/adapter"
	"my-mf/domain/repository"
	"my-mf/handler"
	"my-mf/infra"
	query_expense "my-mf/query/expense"
	"my-mf/usecase"
)

// sqlHansler
func injectDB() adapter.SqlHandler {
	sh := adapter.NewSqlHandler()
	return *sh
}

// infra
func injectExpenseInfra() infra.ExpenseInfra {
	sh := injectDB()
	return infra.NewExpenseInfra(sh)
}
func injectCategoryInfra() infra.CategoryInfra {
	sh := injectDB()
	return infra.NewCategoryInfra(sh)
}

// repository
func injectExpenseRepository() repository.ExpenseRepository {
	ei := injectExpenseInfra()
	return repository.NewExpenseRepository(ei)
}
func injectCategoryRepository() repository.CategoryRepository {
	ci := injectCategoryInfra()
	return repository.NewCategoryRepository(ci)
}

// usecase
func injectExpenseUsecase() usecase.ExpenseUsecase {
	er := injectExpenseRepository()
	return usecase.NewExpenseUsecase(er)
}
func injectCategoryUsecase() usecase.CategoryUsecase {
	cu := injectCategoryRepository()
	return usecase.NewCategoryUsecase(cu)
}

// handler
func InjectExpenseHandler() handler.ExpenseHandler {
	eu := injectExpenseUsecase()
	return handler.NewExpenseHandler(eu)
}
func InjectCategoryHandler() handler.CategoryHandler {
	ch := injectCategoryUsecase()
	return handler.NewCategoryHandler(ch)
}

// query
// expense
func injectQueryExpenseDataAccessor() query_expense.DataAccessor {
	qh := injectDB()
	return query_expense.NewDataAccessor(qh)
}
func injectQueryExpenseUsecase() query_expense.Usecase {
	qeda := injectQueryExpenseDataAccessor()
	return query_expense.NewUsecase(qeda)
}
func InjectQueryExpenseHandler() query_expense.Handler {
	qeu := injectQueryExpenseUsecase()
	return query_expense.NewHandler(qeu)
}
