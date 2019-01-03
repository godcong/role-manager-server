package model

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
	"time"
)

const DefaultInterval = 5

type MongoDB struct {
	//ctx    context.Context
	host   string
	client *mongo.Client
	*mongo.Database
	Interval time.Duration
	//Database string
	//*mongo.Client
}

var mgo *MongoDB

func init() {
	mgo = defaultDB()
}

func newMongoDB() *MongoDB {
	//ctx, _ := context.WithCancel(context.Background())

	return &MongoDB{
		//ctx:      ctx,
		host:     "",
		Interval: DefaultInterval,
	}
}

func defaultDB() *MongoDB {
	db := newMongoDB()
	client, err := InitClient(db.TimeOut(), "mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	db.client = client
	db.Database = client.Database("db1")
	return db
}

func (m *MongoDB) TimeOut() context.Context {
	ctx, _ := context.WithTimeout(context.TODO(), m.Interval)
	return ctx
}

func DB() *MongoDB {
	if mgo != nil {
		return mgo
	}
	return defaultDB()
}

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

func Ping() error {
	return mgo.client.Ping(mgo.TimeOut(), readpref.Primary())
}

func Reconnect() error {
	if err := Ping(); err != nil {
		return mgo.client.Connect(mgo.TimeOut())
	}
	return nil
}

func Collection(name string) *mongo.Collection {
	return DB().Collection(name)
}

func InsertOne(name string, v interface{}, ops ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return Collection(name).InsertOne(context.TODO(), v, ops...)
}
