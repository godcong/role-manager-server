package model

import "github.com/mongodb/mongo-go-driver/bson/primitive"

// Report ...
type Report struct {
	Model         `bson:",inline"`
	MediaID       primitive.ObjectID `bson:"media_id"`       //举报视频ID
	ExoID         primitive.ObjectID `bson:"exo_id"`         //用户ID
	Types         string             `bson:"types"`          //举报类型
	Detail        string             `bson:"detail"`         //举报详情
	ProcessResult string             `bson:"process_result"` //处理结果
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

func (r *Report) _Name() string {
	return "report"
}

// NewReport ...
func NewReport() *Report {
	return &Report{
		Model: model(),
	}
}
