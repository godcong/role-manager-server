package model

import "github.com/mongodb/mongo-go-driver/bson/primitive"

// Log ...
type Log struct {
	Model      `bson:",inline"`
	UserID     primitive.ObjectID `bson:"user_id"`
	Method     string             `bson:"method"`
	URL        string             `bson:"url"`
	Permission string             `bson:"permission"`
	Err        string             `json:"err"`
	Detail     string             `bson:"detail"`
}

// NewLog ...
func NewLog() *Log {
	return &Log{
		Model: model(),
	}
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
