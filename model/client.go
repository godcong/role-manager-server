package model

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
	"time"
)

const DefaultInterval = 5

type MongoDB struct {
	ctx    context.Context
	host   string
	client *mongo.Client
	*mongo.Database
	Interval time.Duration
	//Database string
	//*mongo.Client
}

var mgo *MongoDB

func init() {
	mgo = newMongoDB()
	client, err := InitClient(mgo.ctx, "mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	mgo.client = client
	mgo.Database = client.Database("db1")
}

func newMongoDB() *MongoDB {
	ctx, _ := context.WithTimeout(context.Background(), DefaultInterval)

	return &MongoDB{
		ctx:      ctx,
		host:     "",
		Interval: DefaultInterval,
	}
}

func DB() *MongoDB {
	return mgo
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
	return mgo.client.Ping(mgo.ctx, readpref.Primary())
}

func Reconnect() error {
	if err := Ping(); err != nil {
		return mgo.client.Connect(mgo.ctx)
	}
	return nil
}
