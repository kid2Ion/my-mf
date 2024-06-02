package usecase

import (
	"my-mf/internal/domain/model"
	"my-mf/internal/domain/repository"
)

type (
	categoryUsecase struct {
		repo repository.CategoryRepository
	}
	CategoryUsecase interface {
		Create(req *CategoryCreateReq) error
		GetAll() ([]CategoryGetAllRes, error)
	}
	CategoryCreateReq struct {
		Name string `json:"name"`
	}
	CategoryGetAllRes struct {
		Uuid string `json:"uuid"`
		Name string `json:"name"`
	}
)

func NewCategoryUsecase(categoryRepository repository.CategoryRepository) CategoryUsecase {
	return &categoryUsecase{repo: categoryRepository}
}

func (t *categoryUsecase) Create(req *CategoryCreateReq) error {
	c := model.NewCategory(req.Name)
	return t.repo.Create(c)
}

func (t *categoryUsecase) GetAll() ([]CategoryGetAllRes, error) {
	r, err := t.repo.GetAll()
	if err != nil {
		return nil, err
	}
	res := make([]CategoryGetAllRes, 0)
	for _, v := range r {
		res = append(res, CategoryGetAllRes{v.Uuid.String(), v.Name})
	}
	return res, nil
}
