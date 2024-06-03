package model

import "github.com/google/uuid"

type Category struct {
	Uuid uuid.UUID
	Name string
}

func NewCategory(name string) *Category {
	c := new(Category)
	c.Uuid = uuid.New()
	c.Name = name
	return c
}
