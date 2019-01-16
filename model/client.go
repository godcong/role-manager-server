package model

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
	"time"
)

// DefaultInterval ...
const DefaultInterval = 30 * time.Second

// Prefix ...
type Prefix bool

// MongoDB ...
type MongoDB struct {
	ctx      context.Context
	host     string
	prefix   string
	database string
	limit    int64
	*mongo.Client
	Interval time.Duration
}

// Limit ...
func (m *MongoDB) Limit() *int64 {
	return &m.limit
}

// SetLimit ...
func (m *MongoDB) SetLimit(limit int64) {
	m.limit = limit
}

var mgo *MongoDB

func init() {
	mgo = defaultDB()
}

func newMongoDB() *MongoDB {
	//ctx, _ := context.WithCancel(context.Background())

	return &MongoDB{}
}

func defaultDB() *MongoDB {
	db := newMongoDB()
	client, err := InitClient(db.ctx, db.host)
	if err != nil {
		panic(err)
	}

	db.Client = client
	return db
}

// TimeOut ...
func (m *MongoDB) TimeOut() context.Context {
	ctx, _ := context.WithTimeout(m.ctx, m.Interval)
	return ctx
}

// DB ...
func DB() *MongoDB {
	if mgo != nil {
		return mgo
	}
	return defaultDB()
}

// D ...
func (m *MongoDB) D() *mongo.Database {
	return m.Database(m.database)
}

// InitClient ...
func InitClient(ctx context.Context, ip string) (*mongo.Client, error) {
	client, err := mongo.NewClient(ip)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// Ping ...
func Ping() error {
	return mgo.Ping(mgo.TimeOut(), readpref.Primary())
}

// Reconnect ...
func Reconnect() error {
	if err := Ping(); err != nil {
		return mgo.Connect(mgo.TimeOut())
	}
	return nil
}

// C return a collection
func C(name string, values ...interface{}) *mongo.Collection {
	if NoPrefix(values...) {
		return DB().D().Collection(name)
	}
	return DB().D().Collection(mgo.prefix + "_" + name)
}

// NoPrefix ...
func NoPrefix(values ...interface{}) bool {
	for _, value := range values {
		switch val := value.(type) {
		case Prefix:
			return (bool)(val)
		}
	}
	return false
}

// RelateInfo ...
type RelateInfo struct {
	From         string `bson:"from"`
	LocalField   string `bson:"localField"`
	ForeignField string `json:"foreignField"`
	As           string `json:"as"`
}

// RelateFunc ...
type RelateFunc func() (Modeler, error)

// RelateMakeFunc ...
type RelateMakeFunc func(a, b Modeler) error

// RelateMaker ...
func RelateMaker(fa, fb RelateFunc, f RelateMakeFunc) error {
	return Transaction(func() error {
		a, err := fa()
		if err != nil {
			return err
		}
		b, err := fb()
		if err != nil {
			return err
		}
		return f(a, b)
	})
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
