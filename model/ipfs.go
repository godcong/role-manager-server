package model

import "github.com/mongodb/mongo-go-driver/bson/primitive"

// IPFS ...
type IPFS struct {
	Model   `bson:",inline"`
	MediaID primitive.ObjectID `json:"media_id"`
	FileID  string             `json:"file_id"`
	Status  string             `json:"status"`
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

// Find ...
func (i *IPFS) Find() error {
	return FindByID(i)
}

func (i *IPFS) _Name() string {
	return "ipfs"
}
