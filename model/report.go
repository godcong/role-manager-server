package model

import "github.com/mongodb/mongo-go-driver/bson/primitive"

// Report ...
type Report struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	VideoID string
	Types   string
	Detail  string
}
