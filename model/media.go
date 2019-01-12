package model

import (
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
)

// Media ...
type Media struct {
	Model        `bson:",inline"`
	Block        bool   `bson:"block"`    //禁止访问
	VIPFree      string `bson:"vip_free"` //Vip免费
	Photo        string `bson:"photo"`
	Name         string `bson:"name"`
	Type         string `bson:"type"`
	Language     string `bson:"language"`
	Output3D     string `bson:"output_3_d"`
	VR           string `bson:"vr"`
	Thumb        string `bson:"thumb"`
	Introduction string `bson:"introduction"`
	Starring     string `bson:"starring"`
	Director     string `bson:"director"`
	Episode      string `bson:"episode"`
	TotalNumber  string `bson:"total_number"`
	IPNSAddress  string `bson:"ipns_address"`
	IPFSAddress  string `bson:"ipfs_address"`
	KEYAddress   string `bson:"key_address"`
	Price        int    `bson:"price"`
	PlayType     string `bson:"play_type"`
	ExpireDate   int    `bson:"expire_date"`
}

// NewMedia ...
func NewMedia() *Media {
	return &Media{
		Model: model(),
	}
}

// CreateIfNotExist ...
func (m *Media) CreateIfNotExist() error {
	return CreateIfNotExist(m)
}

// Create ...
func (m *Media) Create() error {
	return InsertOne(m)
}

// Update ...
func (m *Media) Update() error {
	return UpdateOne(m)
}

// Delete ...
func (m *Media) Delete() error {
	return DeleteByID(m)
}

// ALL ...
func (m *Media) ALL() ([]*Media, error) {
	var medias []*Media
	b := bson.M{}
	err := Find(m, b, func(cursor mongo.Cursor) error {
		log.Println(cursor.DecodeBytes())
		var m Media
		err := cursor.Decode(&m)
		if err != nil {
			return err
		}
		medias = append(medias, &m)
		return nil
	})
	return medias, err
}

// Find ...
func (m *Media) Find() error {
	return FindByID(m)
}

func (m *Media) _Name() string {
	return "media"
}
