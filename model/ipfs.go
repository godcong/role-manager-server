package model

import (
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// IPFS ...
type IPFS struct {
	Model       `bson:",inline"`
	MediaID     primitive.ObjectID `bson:"media_id"`
	FileID      string             `bson:"file_id"`
	IPFSAddress string             `bson:"ipfs_address"`
	IPNSAddress string             `bson:"ipns_address"`
	IpnsKey     string             `bson:"ipns_key"`
	Status      string             `bson:"status"`
}

// NewIPFS ...
func NewIPFS() *IPFS {
	return &IPFS{
		Model: model(),
	}
}

// CreateIfNotExist ...
func (i *IPFS) CreateIfNotExist() error {
	return CreateIfNotExist(i)
}

// Create ...
func (i *IPFS) Create() error {
	return InsertOne(i)
}

// Update ...
func (i *IPFS) Update() error {
	return UpdateOne(i)
}

// Delete ...
func (i *IPFS) Delete() error {
	return DeleteByID(i)
}

// FindByFileID ...
func (i *IPFS) FindByFileID() error {
	return FindOne(i, bson.M{
		"file_id": i.FileID,
	})
}

// Find ...
func (i *IPFS) Find() error {
	return FindByID(i)
}

func (i *IPFS) _Name() string {
	return "ipfs"
}
