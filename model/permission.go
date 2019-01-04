package model

import "github.com/mongodb/mongo-go-driver/bson/primitive"

// Permission ...
type Permission struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Name            string
	Slug            string
	Description     string
	PermissionModel string
	*Model
}

// CreateIfNotExist ...
func (p *Permission) CreateIfNotExist() error {
	return CreateIfNotExist(p)
}

// NewPermission ...
func NewPermission() *Permission {
	return &Permission{
		Model: NewModel(),
	}
}

// GetID ...
func (p *Permission) GetID() primitive.ObjectID {
	return p.ID
}

// SetID ...
func (p *Permission) SetID(id primitive.ObjectID) {
	p.ID = id
}

// Create ...
func (p *Permission) Create() error {
	return InsertOne(p)
}

// Update ...
func (p *Permission) Update() error {
	return UpdateOne(p)
}

// Delete ...
func (p *Permission) Delete() error {
	return DeleteByID(p)
}

// Find ...
func (p *Permission) Find() error {
	return FindByID(p)
}

func (p *Permission) _Name() string {
	return "permission"
}
