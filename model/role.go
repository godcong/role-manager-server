package model

import "github.com/mongodb/mongo-go-driver/bson/primitive"

type Role struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string
	Slug        string
	Description string
	Level       int
	Model
}

func (Role) CollectionName() string {
	panic("implement me")
}

func (Role) Create() error {
	panic("implement me")
}

func (Role) Update() error {
	panic("implement me")
}

func (Role) Delete() error {
	panic("implement me")
}

func (Role) FindByID(id string) error {
	panic("implement me")
}

func (Role) _Name() string {
	panic("implement me")
}
