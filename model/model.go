package model

import (
	"github.com/mongodb/mongo-go-driver/x/mongo/driver/uuid"
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
	Id        uuid.UUID
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
