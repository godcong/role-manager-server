package model

import (
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
)

// Media ...
type Media struct {
	Model             `bson:",inline"`
	OrganizationID    primitive.ObjectID `bson:"organization_id"`     //组织id
	CensorID          primitive.ObjectID `bson:"censor_id"`           //AI检查状态
	CensorResult      string             `bson:"censor_result"`       //鉴定结果
	Block             bool               `bson:"block"`               //禁止访问
	VIPFree           string             `bson:"vip_free"`            //Vip免费
	Photo             string             `bson:"photo"`               //照片
	Name              string             `bson:"name"`                //名称
	Type              string             `bson:"type"`                //类别
	Language          string             `bson:"language"`            //语言
	Output3D          string             `bson:"output_3d"`           //3D
	VR                string             `bson:"vr"`                  //VR
	Thumb             string             `bson:"thumb"`               //缩略图
	Introduction      string             `bson:"introduction"`        //简介
	Starring          string             `bson:"starring"`            //主演
	Director          string             `bson:"director"`            //导演
	Episode           string             `bson:"episode"`             //集数
	TotalNumber       string             `bson:"total_number"`        //总集数
	IPNSAddress       string             `bson:"ipns_address"`        //ipns地址
	IPFSAddress       string             `bson:"ipfs_address"`        //ipfs地址
	VideoOSSAddress   string             `bson:"video_oss_address"`   //视频oss地址
	PictureOSSAddress []string           `bson:"picture_oss_address"` //图片oss地址
	KEYAddress        string             `bson:"key_address"`         //key地址
	Price             string             `bson:"price"`               //价格
	PlayType          string             `bson:"play_type"`           //播放类型(单次,多次)
	ExpireDate        string             `bson:"expire_date"`         //过期时间(48H,24H,0H)
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
	if m.CensorID != primitive.NilObjectID {
		return m.FindByCensor()
	}
	return FindByID(m)
}

// FindByOrg ...
func (m *Media) FindByOrg() (medias []*Media, err error) {
	err = Find(m, bson.M{
		"organization_id": m.OrganizationID,
	}, func(cursor mongo.Cursor) error {
		var media Media
		err := cursor.Decode(&media)
		if err != nil {
			return err
		}
		medias = append(medias, &media)
		return nil
	})
	return
}

// FindByCensor ...
func (m *Media) FindByCensor() (err error) {
	return FindOne(m, bson.M{
		"censor_id": m.CensorID,
	})
	return
}

// Censors ...
func (m *Media) Censors() (mcs []*MediaCensor, err error) {
	mc := NewMediaCensor()
	err = Find(mc, bson.M{
		"media_id": m.ID,
		//"verify":   "verification",
	}, func(cursor mongo.Cursor) error {
		mc := NewMediaCensor()
		err := cursor.Decode(mc)
		if err != nil {
			return err
		}

		mcs = append(mcs, mc)
		return nil
	})
	return
}

func (m *Media) _Name() string {
	return "medias"
}
