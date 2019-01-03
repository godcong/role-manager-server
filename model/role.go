package model

type Role struct {
	*Model
	Name        string
	Slug        string
	Description string
	Level       int
}
