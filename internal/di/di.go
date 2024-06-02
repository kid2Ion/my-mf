package di

import (
	"my-mf/internal/domain/repository"
	"my-mf/internal/handler"
	"my-mf/internal/infra"
	"my-mf/internal/usecase"
)

// sqlHansler
func injectDB() infra.SqlHandler {
	sh := infra.NewSqlHandler()
	return *sh
}

// infra
func injectExpenseInfra() infra.ExpenseInfra {
	sh := injectDB()
	return infra.NewExpenseInfra(sh)
}

// repository
func injectExpenseRepository() repository.ExpenseRepository {
	ei := injectExpenseInfra()
	return repository.NewExpenseRepository(ei)
}

// usecase
func injectExpenseUsecase() usecase.ExpenseUsecase {
	er := injectExpenseRepository()
	return usecase.NewExpenseUsecase(er)
}

// handler
func InjectExpenseHandler() handler.ExpenseHandler {
	eu := injectExpenseUsecase()
	return handler.NewExpenseHandler(eu)
}
