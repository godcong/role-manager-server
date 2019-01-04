package model

import (
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"log"
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

type Before interface {
	BeforeInsert()
	BeforeUpdate()
	BeforeDelete()
}

type After interface {
	AfterInsert()
	AfterUpdate()
	AfterDelete()
}

type Modeler interface {
	_Name() string
	Before
	After
	GetID() primitive.ObjectID
	SetID(id primitive.ObjectID)
	Create() error
	Update() error
	Delete() error
	FindByID(id string) error
}

func ID(s string) primitive.ObjectID {
	ids, err := primitive.ObjectIDFromHex(s)
	if err != nil {
		log.Println(err)
	}
	return ids
}

func UpdateOne(m Modeler, id primitive.ObjectID, v interface{}, ops ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	m.BeforeUpdate()
	return C(m._Name()).UpdateOne(context.TODO(), bson.M{
		"_id": id,
	}, v, ops...)
}

func InsertOne(m Modeler, v interface{}, ops ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	m.BeforeInsert()
	result, err := C(m._Name()).InsertOne(context.TODO(), v, ops...)
	if err == nil {
		if v, b := result.InsertedID.(primitive.ObjectID); b {
			m.SetID(v)
		}
	}

	return result, err
}

func DeleteByID(m Modeler, id primitive.ObjectID, ops ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return C(m._Name()).DeleteOne(context.TODO(), bson.M{
		"_id": id,
	}, ops...)
}

func FindByID(m Modeler, id string, v interface{}, ops ...*options.FindOneOptions) error {
	return C(m._Name()).FindOne(mgo.TimeOut(), bson.M{
		"_id": ID(id),
	}, ops...).Decode(v)
}

func (m *Model) BeforeInsert() {
	m.CreatedAt = time.Now()
	m.UpdatedAt = m.CreatedAt
	m.Version = 1
}

func (m *Model) BeforeUpdate() {
	m.Version += 1
}

func (m *Model) BeforeDelete() {
	return
}

func (m *Model) AfterInsert() {
	return
}

func (m *Model) AfterUpdate() {
	return
}

func (m *Model) AfterDelete() {
	return
}
