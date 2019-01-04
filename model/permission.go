package model

import "github.com/mongodb/mongo-go-driver/bson/primitive"

type Permission struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Name            string
	Slug            string
	Description     string
	PermissionModel string
	*Model
}

func NewPermission() *Permission {
	return &Permission{
		Model: NewModel(),
	}
}

func (p *Permission) GetID() primitive.ObjectID {
	return p.ID
}

func (p *Permission) SetID(id primitive.ObjectID) {
	p.ID = id
}

func (p *Permission) Create() error {
	return InsertOne(p)
}

func (p *Permission) Update() error {
	return UpdateOne(p)
}

func (p *Permission) Delete() error {
	return DeleteByID(p)
}

func (p *Permission) Find() error {
	return FindByID(p)
}

func (p *Permission) _Name() string {
	return "permission"
}
