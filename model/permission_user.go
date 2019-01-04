package model

import "github.com/mongodb/mongo-go-driver/bson/primitive"

// PermissionUser ...
type PermissionUser struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	PermissionID primitive.ObjectID
	UserID       primitive.ObjectID
	permission   *Permission
	user         *User
	*Model
}

// User ...
func (r *PermissionUser) User() *User {
	return r.user
}

// SetUser ...
func (r *PermissionUser) SetUser(user *User) {
	r.user = user
	r.UserID = user.ID
}

// Permission ...
func (r *PermissionUser) Permission() *Permission {
	return r.permission
}

// SetPermission ...
func (r *PermissionUser) SetPermission(permission *Permission) {
	r.permission = permission
	r.PermissionID = permission.ID
}

// GetID ...
func (r *PermissionUser) GetID() primitive.ObjectID {
	return r.ID
}

// SetID ...
func (r *PermissionUser) SetID(id primitive.ObjectID) {
	r.ID = id
}

// Create ...
func (r *PermissionUser) Create() error {
	return InsertOne(r)
}

// Update ...
func (r *PermissionUser) Update() error {
	return UpdateOne(r)
}

// Delete ...
func (r *PermissionUser) Delete() error {
	return DeleteByID(r)
}

// Find ...
func (r *PermissionUser) Find() error {
	return FindByID(r)
}

// NewPermissionUser ...
func NewPermissionUser() *PermissionUser {
	return &PermissionUser{
		Model: NewModel(),
	}
}

func (r *PermissionUser) _Name() string {
	return "user_role"
}
