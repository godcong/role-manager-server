package model

import "github.com/mongodb/mongo-go-driver/bson/primitive"

const (
	LevelGenesis = 0
	LevelAdmin   = iota
	LevelOrg     = iota
	Levelmonitor = iota
)

type Role struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string
	Slug        string
	Description string
	Level       int
	*Model
}

func NewRole() *Role {
	return &Role{
		Model: NewModel(),
	}
}

func (r *Role) SetID(id primitive.ObjectID) {
	r.ID = id
}

func (r *Role) Create() error {
	return InsertOne(r)
}

func (r *Role) Update() error {
	return UpdateOne(r)
}

func (r *Role) Delete() error {
	return DeleteByID(r)
}

func (r *Role) Find() error {
	return FindByID(r)
}

func (r *Role) _Name() string {
	return "role"
}

func (r *Role) GetID() primitive.ObjectID {
	return r.ID
}
