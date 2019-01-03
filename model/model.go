package model

import (
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"time"
)

type SyncAble interface {
	Sync() error
}

type CountAble interface {
	Count() (int64, error)
}

type CreateAble interface {
	Create() (int64, error)
}

type GetAble interface {
	Get() (bool, error)
	List(v interface{}) error
}

type UpdateAble interface {
	Update() (int64, error)
	UpdateOnly(cols ...string) (int64, error)
}

type BaseAble interface {
	SyncAble
	CountAble
	CreateAble
	GetAble
	UpdateAble
}

type Model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Version   int
}

type Modeler interface {
	CollectionName() string
	Create() error
	Update() error
	Delete() error
	Find() error
}

func UpdateOne(collection *mongo.Collection, id primitive.ObjectID, v interface{}, ops ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return collection.UpdateOne(context.TODO(), bson.M{
		"_id": id,
	}, v, ops...)
}

func InsertOne(collection *mongo.Collection, name string, v interface{}, ops ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return collection.InsertOne(context.TODO(), v, ops...)
}

func DeleteByID(collection *mongo.Collection, id primitive.ObjectID, ops ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return collection.DeleteOne(context.TODO(), bson.M{
		"_id": id,
	}, ops...)
}

func FindByID(collection *mongo.Collection, id string, v interface{}, ops ...*options.FindOneOptions) error {
	ids, _ := primitive.ObjectIDFromHex(id)
	return collection.FindOne(context.TODO(), bson.M{
		"_id": ids,
	}, ops...).Decode(v)
}

func (m *Model) BeforeInsert() {
	m.CreatedAt = time.Now()
	m.UpdatedAt = m.CreatedAt
	m.Version = 1
}
