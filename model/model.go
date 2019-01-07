package model

import (
	"errors"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"log"
	"time"
)

// SyncAble ...
type SyncAble interface {
	Sync() error
}

// CountAble ...
type CountAble interface {
	Count() (int64, error)
}

// CreateAble ...
type CreateAble interface {
	Create() (int64, error)
}

// GetAble ...
type GetAble interface {
	Get() (bool, error)
	List(v interface{}) error
}

// UpdateAble ...
type UpdateAble interface {
	Update() (int64, error)
	UpdateOnly(cols ...string) (int64, error)
}

// BaseAble ...
type BaseAble interface {
	SyncAble
	CountAble
	CreateAble
	GetAble
	UpdateAble
}

// Model ...
type Model struct {
	softDelete bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
	Version    int
}

// NewModel ...
func NewModel() *Model {
	return &Model{
		softDelete: true,
	}
}

// Before ...
type Before interface {
	BeforeInsert()
	BeforeUpdate()
	BeforeDelete()
}

// After ...
type After interface {
	AfterInsert()
	AfterUpdate()
	AfterDelete()
}

// Modeler ...
type Modeler interface {
	_Name() string
	Before
	After

	IsExist() bool
	GetID() primitive.ObjectID
	SetID(id primitive.ObjectID)
	CreateIfNotExist() error
	Create() error
	Update() error
	Delete() error
	Find() error
	SoftDelete() bool
}

// ID ...
func ID(s string) primitive.ObjectID {
	ids, err := primitive.ObjectIDFromHex(s)
	if err != nil {
		log.Println(err)
	}
	return ids
}

// UpdateOne ...
func UpdateOne(m Modeler, ops ...*options.UpdateOptions) error {
	m.BeforeUpdate()
	result, err := C(m._Name()).UpdateOne(mgo.TimeOut(), bson.M{
		"_id": m.GetID(),
	}, bson.M{
		"$set": m,
	}, ops...)
	if err == nil {
		log.Println(result.UpsertedID, result.MatchedCount, result.ModifiedCount)
	}

	return err
}

// InsertOne ...
func InsertOne(m Modeler, ops ...*options.InsertOneOptions) error {
	m.BeforeInsert()
	result, err := C(m._Name()).InsertOne(mgo.TimeOut(), m, ops...)
	if err == nil {
		if v, b := result.InsertedID.(primitive.ObjectID); b {
			m.SetID(v)
		}
	}

	return err
}

// DeleteByID ...
func DeleteByID(m Modeler, ops ...*options.DeleteOptions) error {
	if m.SoftDelete() {
		err := FindByID(m)
		if err != nil {
			return err
		}
		m.BeforeDelete()
		return UpdateOne(m)
	}

	result, err := C(m._Name()).DeleteOne(mgo.TimeOut(), bson.M{
		"_id": m.GetID(),
	}, ops...)
	if err == nil {
		log.Println(result.DeletedCount)
	}

	return err
}

// FindDecode ...
type FindDecode func(cursor mongo.Cursor) error

// Find ...
func Find(m Modeler, v bson.M, dec FindDecode, ops ...*options.FindOptions) error {
	if m.SoftDelete() {
		v["model.deletedat"] = nil
	}
	find, err := C(m._Name()).Find(mgo.TimeOut(), v, ops...)
	if err != nil {
		return err
	}
	for find.Next(mgo.TimeOut()) {
		err := dec(find)
		if err != nil {
			return err
		}
	}
	return nil
}

// FindOne ...
func FindOne(m Modeler, v bson.M, ops ...*options.FindOneOptions) error {
	if m.SoftDelete() {
		v["model.deletedat"] = nil
	}
	return C(m._Name()).FindOne(mgo.TimeOut(), v, ops...).Decode(m)
}

// FindByID ...
func FindByID(m Modeler, ops ...*options.FindOneOptions) error {
	v := bson.M{
		"_id": m.GetID(),
	}
	if m.SoftDelete() {
		v["model.deletedat"] = nil
	}
	return C(m._Name()).FindOne(mgo.TimeOut(), v, ops...).Decode(m)
}

// Count ...
func Count(m Modeler, v bson.M) (int64, error) {
	if m.SoftDelete() {
		v["model.deletedat"] = nil
	}
	//result := C(m._Name()).FindOne(mgo.TimeOut(), v)
	return C(m._Name()).Count(mgo.TimeOut(), v)

}

// IsExist ...
func IsExist(m Modeler, v bson.M) bool {
	i, err := Count(m, v)
	log.Println(i, err)
	if err == nil && i != 0 {
		return true
	}
	return false
}

// CreateIfNotExist ...
func CreateIfNotExist(m Modeler) error {
	if !m.IsExist() {
		return m.Create()
	}
	return errors.New(m._Name() + " is exist")
}

// IsExist ...
func (m *Model) IsExist() bool {
	return false
}

// SoftDelete ...
func (m *Model) SoftDelete() bool {
	return m.softDelete
}

// SetSoftDelete ...
func (m *Model) SetSoftDelete(b bool) {
	m.softDelete = b
}

// BeforeInsert ...
func (m *Model) BeforeInsert() {
	m.CreatedAt = time.Now()
	m.UpdatedAt = m.CreatedAt
	m.Version = 1
}

// BeforeUpdate ...
func (m *Model) BeforeUpdate() {
	m.UpdatedAt = time.Now()
	m.Version++
}

// BeforeDelete ...
func (m *Model) BeforeDelete() {
	t := time.Now()
	m.DeletedAt = &t
	return
}

// AfterInsert ...
func (m *Model) AfterInsert() {
	return
}

// AfterUpdate ...
func (m *Model) AfterUpdate() {
	return
}

// AfterDelete ...
func (m *Model) AfterDelete() {
	return
}

// TransactionDo ...
type TransactionDo func() error

// Transaction 事物
func Transaction(fn TransactionDo) error {
	session, err := DB().StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(mgo.TimeOut())
	err = session.StartTransaction()
	if err != nil {
		return err
	}
	err = fn()
	if err != nil {
		_ = session.AbortTransaction(mgo.TimeOut())
		return err
	}
	err = session.CommitTransaction(mgo.TimeOut())
	if err != nil {
		_ = session.AbortTransaction(mgo.TimeOut())
	}
	return nil
}
