package model

import (
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
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt  time.Time          `bson:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at"`
	DeletedAt  *time.Time         `bson:"deleted_at"`
	Version    int                `bson:"version"`
	softDelete bool
}

func model() Model {
	return Model{
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		DeletedAt:  nil,
		Version:    1,
		softDelete: true,
	}
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

// FindDecodeLoop ...
type FindDecodeLoop func(cursor mongo.Cursor) error

func find(m Modeler, v bson.M, dec FindDecodeLoop, prefix bool, ops ...*options.FindOptions) error {
	SoftDelete(m, &v)

	col := C(m._Name())
	if !prefix {
		log.Println(prefix)
		col = C(m._Name(), Prefix(true))
	}
	log.Println(col.Name())
	find, err := col.Find(mgo.TimeOut(), v, ops...)
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

// FindWithPrefix ...
func FindWithPrefix(m Modeler, v bson.M, dec FindDecodeLoop, prefix bool, ops ...*options.FindOptions) error {
	return find(m, v, dec, prefix, ops...)
}

// Find ...
func Find(m Modeler, v bson.M, dec FindDecodeLoop, ops ...*options.FindOptions) error {
	return find(m, v, dec, true, ops...)
}

// Pages ...
func Pages(m Modeler, v bson.M, order, limit, current int64, dec FindDecodeLoop) error {
	skip := current * limit
	return Find(m, v, dec, &options.FindOptions{
		AllowPartialResults: nil,
		BatchSize:           nil,
		Collation:           nil,
		Comment:             nil,
		CursorType:          nil,
		Hint:                nil,
		Limit:               &limit,
		Max:                 nil,
		MaxAwaitTime:        nil,
		MaxTime:             nil,
		Min:                 nil,
		NoCursorTimeout:     nil,
		OplogReplay:         nil,
		Projection:          nil,
		ReturnKey:           nil,
		ShowRecordID:        nil,
		Skip:                &skip,
		Snapshot:            nil,
		Sort: bson.M{
			"created_at": order,
		},
	})
}

// FindOne ...
func FindOne(m Modeler, v bson.M, ops ...*options.FindOneOptions) error {
	SoftDelete(m, &v)
	return C(m._Name()).FindOne(mgo.TimeOut(), v, ops...).Decode(m)
}

// FindByID ...
func FindByID(m Modeler, ops ...*options.FindOneOptions) error {
	v := bson.M{
		"_id": m.GetID(),
	}
	SoftDelete(m, &v)
	return C(m._Name()).FindOne(mgo.TimeOut(), v, ops...).Decode(m)
}

// Count ...
func Count(m Modeler, v bson.M) (int64, error) {
	SoftDelete(m, &v)
	//result := C(m._Name()).FindOne(mgo.TimeOut(), v)
	return C(m._Name()).Count(mgo.TimeOut(), v)

}

// IDExist ...
func IDExist(m Modeler) bool {
	return IsExist(m, bson.M{
		"_id": m.GetID(),
	})
}

// IsExist ...
func IsExist(m Modeler, v bson.M) bool {
	i, err := Count(m, v)
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
	return m.Find()
	//return errors.New(m._Name() + " is exist")
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

// GetID ...
func (m *Model) GetID() primitive.ObjectID {
	return m.ID
}

// SetID ...
func (m *Model) SetID(id primitive.ObjectID) {
	m.ID = id
}

// SoftDelete ...
func SoftDelete(modeler Modeler, v *bson.M) bool {
	if modeler.SoftDelete() {
		(*v)["deleted_at"] = nil
		return true
	}
	return false
}
