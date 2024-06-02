package repository

import "my-mf/internal/infra"

type (
	categoryRepository struct {
		infra infra.CategoryInfra
	}
	CategoryRepository interface {
		Create() error
	}
)

func NewCategoryRepository(categoryInfra infra.CategoryInfra) CategoryRepository {
	return &categoryRepository{infra: categoryInfra}
}

func (t *categoryRepository) Create() error {
	return nil
}
