package infra

import "fmt"

type (
	CategoryInfra interface {
		Create() error
	}
	categoryInfra struct {
		SqlHandler
	}
)

func NewCategoryInfra(sh SqlHandler) CategoryInfra {
	return &categoryInfra{SqlHandler: sh}
}

func (t *categoryInfra) Create() error {
	fmt.Println("bun bun hello")
	return nil
}
