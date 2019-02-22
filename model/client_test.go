package model

import (
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	log "github.com/sirupsen/logrus"
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
	//	db.user.aggregate(
	//		[
	//		{
	//			"$project" : {
	//				"_id" : NumberInt(0),
	//				"user" : "$$ROOT"
	//			}
	//		},
	//	{
	//		"$lookup" : {
	//		"localField" : "user._id",
	//			"from" : "role_user",
	//			"foreignField" : "userid",
	//			"as" : "role_user"
	//	}
	//	},
	//	{
	//		"$unwind" : {
	//		"path" : "$role_user",
	//			"preserveNullAndEmptyArrays" : true
	//	}
	//	},
	//	{
	//		"$lookup" : {
	//		"localField" : "role_user.roleid",
	//			"from" : "role",
	//			"foreignField" : "_id",
	//			"as" : "role"
	//	}
	//	},
	//	{
	//		"$unwind" : {
	//		"path" : "$role",
	//			"preserveNullAndEmptyArrays" : true
	//	}
	//	},
	//	{
	//		"$match" : {
	//		"user._id" : ObjectId("5c33711e06b5362b5f8dccbf")
	//	}
	//	}
	//],
	//	{
	//	"allowDiskUse" : true
	//	}
	//	);
	cursor, e := C("user").Aggregate(context.Background(),
		[]primitive.E{
			primitive.E{
				Key: "$project",
				Value: bson.M{
					"_id":  0,
					"user": "$$ROOT",
				},
			},
			primitive.E{
				Key: "$lookup",
				Value: bson.M{
					"localField":   "user._id",
					"from":         "role_user",
					"foreignField": "userid",
					"as":           "role_user",
				},
			},
			primitive.E{
				Key: "$unwind",
				Value: bson.M{
					"path":                       "$role_user",
					"preserveNullAndEmptyArrays": true,
				},
			},
			primitive.E{
				Key: "$lookup",
				Value: bson.M{
					"localField":   "role_user.roleid",
					"from":         "role",
					"foreignField": "_id",
					"as":           "role",
				},
			},
			primitive.E{
				Key: "$unwind",
				Value: bson.M{"path": "$role",
					"preserveNullAndEmptyArrays": true,
				},
			},
			primitive.E{
				Key: "$match",
				Value: bson.M{
					"user._id": ID("5c384909078d4d5bd20177be"),
				},
			},
		},
		&options.AggregateOptions{
			AllowDiskUse: &b,
		})

	log.Println(e)
	log.Println(cursor)
}
