package model

import (
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
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
	softDelete bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
	Version    int
}

func NewModel() *Model {
	return &Model{
		softDelete: true,
	}
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
	Find() error
	SoftDelete() bool
}

func ID(s string) primitive.ObjectID {
	ids, err := primitive.ObjectIDFromHex(s)
	if err != nil {
		log.Println(err)
	}
	return ids
}

func UpdateOne(m Modeler, ops ...*options.UpdateOptions) error {
	m.BeforeUpdate()
	result, err := C(m._Name()).UpdateOne(context.TODO(), bson.M{
		"_id": m.GetID(),
	}, bson.M{
		"$set": m,
	}, ops...)
	if err == nil {
		log.Println(result.UpsertedID, result.MatchedCount, result.ModifiedCount)
	}

	return err
}

func InsertOne(m Modeler, ops ...*options.InsertOneOptions) error {
	m.BeforeInsert()
	result, err := C(m._Name()).InsertOne(context.TODO(), m, ops...)
	if err == nil {
		if v, b := result.InsertedID.(primitive.ObjectID); b {
			m.SetID(v)
		}
	}

	return err
}

func DeleteByID(m Modeler, ops ...*options.DeleteOptions) error {
	if m.SoftDelete() {
		err := m.Find()
		if err != nil {
			return err
		}
		m.BeforeDelete()
		return UpdateOne(m)
	}

	result, err := C(m._Name()).DeleteOne(context.TODO(), bson.M{
		"_id": m.GetID(),
	}, ops...)
	if err == nil {
		log.Println(result.DeletedCount)
	}

	return err
}

func FindByID(m Modeler, ops ...*options.FindOneOptions) error {
	if m.SoftDelete() {
		return C(m._Name()).FindOne(mgo.TimeOut(), bson.M{
			"_id":             m.GetID(),
			"model.deletedat": nil,
		}, ops...).Decode(m)
	}
	return C(m._Name()).FindOne(mgo.TimeOut(), bson.M{
		"_id": m.GetID(),
	}, ops...).Decode(m)
}

func (m *Model) SoftDelete() bool {
	return m.softDelete
}

func (m *Model) SetSoftDelete(b bool) {
	m.softDelete = b
}

func (m *Model) BeforeInsert() {
	m.CreatedAt = time.Now()
	m.UpdatedAt = m.CreatedAt
	m.Version = 1
}

func (m *Model) BeforeUpdate() {
	m.UpdatedAt = time.Now()
	m.Version += 1
}

func (m *Model) BeforeDelete() {
	t := time.Now()
	m.DeletedAt = &t
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
