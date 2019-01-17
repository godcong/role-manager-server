package model

import (
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"log"
	"testing"
	"time"
)

// TestInitClient ...
func TestInitClient(t *testing.T) {
	client, _ := InitClient(context.Background(), "")
	time.Sleep(10 * time.Second)
	client.Database("db1").Collection("numbers").InsertOne(context.Background(), bson.M{
		"value": "hello world",
		"name":  "test",
	})
}

// TestRelate ...
func TestRelate(t *testing.T) {
	b := true
	cursor, e := C("user").Aggregate(context.Background(), bson.A{
		bson.M{"$project": bson.M{
			"_id":  0,
			"user": "$$ROOT",
		}},
		bson.M{"$lookup": bson.M{
			"$lookup": bson.M{
				"localField":   "user._id",
				"from":         "role_user",
				"foreignField": "userid",
				"as":           "role_user",
			},
			"$unwind": bson.M{
				"path":                       "$role_user",
				"preserveNullAndEmptyArrays": true,
			},
		}},
		bson.M{"$lookup": bson.M{
			"localField":   "role_user.roleid",
			"from":         "role",
			"foreignField": "_id",
			"as":           "role",
		}},
		bson.M{"$unwind": bson.M{
			"path":                       "$role",
			"preserveNullAndEmptyArrays": true,
		}},
		bson.M{"$match": bson.M{
			"user._id": ID("5c384909078d4d5bd20177be"),
		}},
	}, &options.AggregateOptions{
		AllowDiskUse: &b,
	})

	log.Println(e)
	log.Println(cursor.Next(context.Background()))

}
