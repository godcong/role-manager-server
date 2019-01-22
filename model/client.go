package model

import (
	"context"
	"fmt"
	"github.com/godcong/role-manager-server/config"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
	"log"
	"time"
)

// DefaultInterval ...
const DefaultInterval = 30 * time.Second

// NoPrefix ...
type NoPrefix bool

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

// InitDB ...
func InitDB(cfg *config.Configure) {
	mgo = defaultDB(cfg)

	log.Printf("%+v", mgo)
	err := Ping()
	if err != nil {
		panic(err)
	}
}

func newMongoDB(cfg *config.Configure) *MongoDB {
	host := fmt.Sprintf("mongodb://%s:%s@%s%s/%s",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Addr,
		cfg.Database.Port,
		cfg.Database.DB,
	)
	return &MongoDB{
		ctx:      context.Background(),
		host:     host,
		prefix:   cfg.Database.Prefix,
		Client:   nil,
		Interval: DefaultInterval,
		database: cfg.Database.DB,
	}
}

func defaultDB(cfg *config.Configure) *MongoDB {
	db := newMongoDB(cfg)
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
	return defaultDB(config.Config())
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
	if !Prefix(values...) {
		return DB().D().Collection(name)
	}
	return DB().D().Collection(mgo.prefix + "_" + name)
}

// Prefix ...
func Prefix(values ...interface{}) bool {
	for _, value := range values {
		if v, b := value.(NoPrefix); b {
			return !(bool)(v)
		}
	}
	return true
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
