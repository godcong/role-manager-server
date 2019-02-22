package model

import (
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
		log "github.com/sirupsen/logrus"
)

// Log ...
type Log struct {
	Model      `bson:",inline"`
	UserID     primitive.ObjectID `bson:"user_id"`
	Method     string             `bson:"method"`
	URL        string             `bson:"url"`
	Permission string             `bson:"permission"`
	Err        string             `bson:"err"`
	Detail     string             `bson:"detail"`
	VisitIP    string             `bson:"visit_ip"`
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

// ALL ...
func (l *Log) ALL() ([]*Log, error) {
	var logs []*Log
	m := bson.M{}
	err := Find(l, m, func(cursor mongo.Cursor) error {
		log.Println(cursor.DecodeBytes())
		var l Log
		err := cursor.Decode(&l)
		if err != nil {
			return err
		}
		logs = append(logs, &l)
		return nil
	})
	return logs, err
}

// Pages ...
func (l *Log) Pages(order, limit, current int64) ([]*Log, int64) {
	var logs []*Log
	m := bson.M{}

	i, _ := Count(l, m)

	err := Pages(l, m, order, limit, current, func(cursor mongo.Cursor) error {
		log.Println(cursor.DecodeBytes())
		var l Log
		err := cursor.Decode(&l)
		if err != nil {
			return err
		}
		logs = append(logs, &l)
		return nil
	})
	if err != nil {
		return nil, 0
	}
	return logs, i
}

func (l *Log) _Name() string {
	return 	log "github.com/sirupsen/logrus"
}
