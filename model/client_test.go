package model

import (
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
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
	Relate("user", &RelateInfo{
		From:         "",
		LocalField:   "",
		ForeignField: "",
		As:           "",
	})
}
