package model

import "github.com/mongodb/mongo-go-driver/bson/primitive"

type Permission struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Name            string
	Slug            string
	Description     string
	PermissionModel string
	Model
}
