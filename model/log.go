package model

import "github.com/mongodb/mongo-go-driver/bson/primitive"

// Log ...
type Log struct {
	Model
	UserID     primitive.ObjectID
	Method     primitive.ObjectID
	Permission string
	Detail     string
}

// CreateIfNotExist ...
func (l *Log) CreateIfNotExist() error {
	return CreateIfNotExist(l)
}

// Create ...
func (l *Log) Create() error {
	return InsertOne(l)
}

// Update ...
func (l *Log) Update() error {
	return UpdateOne(l)
}

// Delete ...
func (l *Log) Delete() error {
	return DeleteByID(l)
}

// Find ...
func (l *Log) Find() error {
	return FindByID(l)
}

func (l *Log) _Name() string {
	return "log"
}
