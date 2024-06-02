package usecase

import "my-mf/internal/domain/repository"

type (
	categoryUsecase struct {
		repo repository.CategoryRepository
	}
	CategoryUsecase interface {
		Create() error
	}
	CategoryCreateReq struct {
		Name string `json:"name"`
	}
)

func NewCategoryUsecase(categoryRepository repository.CategoryRepository) CategoryUsecase {
	return &categoryUsecase{repo: categoryRepository}
}

func (t *categoryUsecase) Create() error {
	return nil
}
