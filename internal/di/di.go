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
func injectCategoryInfra() infra.CategoryInfra {
	sh := injectDB()
	return infra.NewCategoryInfra(sh)
}
func injectRelExpenseCategoryInfra() infra.RelExpensesCategoriesInfra {
	sh := injectDB()
	return infra.NewRelExpensesCategoriesInfra(sh)
}

// repository
func injectExpenseRepository() repository.ExpenseRepository {
	ei := injectExpenseInfra()
	rr := injectRelExpenseCategoryInfra()
	return repository.NewExpenseRepository(ei, rr)
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
