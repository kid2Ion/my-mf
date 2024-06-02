package repository

import (
	"my-mf/internal/domain/model"
	"my-mf/internal/infra"
)

type (
	categoryRepository struct {
		infra infra.CategoryInfra
	}
	CategoryRepository interface {
		Create(req *model.Category) error
		GetAll() ([]model.Category, error)
	}
)

func NewCategoryRepository(categoryInfra infra.CategoryInfra) CategoryRepository {
	return &categoryRepository{infra: categoryInfra}
}

func (t *categoryRepository) Create(req *model.Category) error {
	return t.infra.Create(req)
}

func (t *categoryRepository) GetAll() ([]model.Category, error) {
	res, err := t.infra.GetAll()
	if err != nil {
		return nil, err
	}
	return res, nil
}
