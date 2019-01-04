package model

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
	"time"
)

// DefaultInterval ...
const DefaultInterval = 5 * time.Second

// MongoDB ...
type MongoDB struct {
	ctx  context.Context
	host string
	*mongo.Client
	Interval time.Duration
	database string
}

var mgo *MongoDB

func init() {
	mgo = defaultDB()
}

func newMongoDB() *MongoDB {
	//ctx, _ := context.WithCancel(context.Background())

	return &MongoDB{
		ctx:      context.Background(),
		host:     "mongodb://localhost:27017",
		database: "database",
		Interval: DefaultInterval,
	}
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
func C(name string) *mongo.Collection {
	return DB().D().Collection(name)
}
