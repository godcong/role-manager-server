package model

import (
	"github.com/json-iterator/go"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	log "github.com/sirupsen/logrus"
)

// ReportResultObtained ...
const ReportResultObtained = "obtained"

// Report ...
type Report struct {
	Model         `bson:",inline"`
	MediaID       primitive.ObjectID `json:"media_id" bson:"media_id"`             //举报视频ID
	ExoID         primitive.ObjectID `json:"exo_id" bson:"exo_id"`                 //用户ID
	Types         string             `json:"types" bson:"types"`                   //举报类型
	Detail        string             `json:"detail" bson:"detail"`                 //举报详情
	ProcessResult string             `json:"process_result" bson:"process_result"` //处理结果
}

// CreateIfNotExist ...
func (r *Report) CreateIfNotExist() error {
	return CreateIfNotExist(r)
}

// Create ...
func (r *Report) Create() error {
	return InsertOne(r)
}

// Update ...
func (r *Report) Update() error {
	return UpdateOne(r)
}

// Delete ...
func (r *Report) Delete() error {
	return DeleteByID(r)
}

// Find ...
func (r *Report) Find() error {
	return FindByID(r)
}

// UnmarshalJSON ...
func (r *Report) UnmarshalJSON(b []byte) error {
	m := map[string]string{}
	err := jsoniter.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	for key, val := range m {
		switch key {
		case "media_id":
			r.MediaID = ID(val)
		case "exo_id":
			r.ExoID = ID(val)
		case "types":
			r.Types = val
		case "detail":
			r.Detail = val
		}
	}

	return nil
}

func (r *Report) _Name() string {
	return "report"
}

// ALL ...
func (r *Report) ALL() ([]*Report, error) {
	var orgs []*Report
	m := bson.M{}
	err := Find(r, m, func(cursor mongo.Cursor) error {
		log.Println(cursor.DecodeBytes())
		var r Report
		err := cursor.Decode(&r)
		if err != nil {
			return err
		}
		orgs = append(orgs, &r)
		return nil
	})
	return orgs, err
}

// Media ...
func (r *Report) Media() (*Media, error) {
	media := NewMedia()
	media.ID = r.MediaID
	err := media.Find()
	return media, err
}

// NewReport ...
func NewReport() *Report {
	return &Report{
		Model: model(),
	}
}
