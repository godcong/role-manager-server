package model

import "github.com/mongodb/mongo-go-driver/bson/primitive"

// Report ...
type Report struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	MediaID       primitive.ObjectID `bson:"media_id"`       //举报视频ID
	ExoID         primitive.ObjectID `bson:"exo_id"`         //用户ID
	Types         string             `bson:"types"`          //举报类型
	Detail        string             `bson:"detail"`         //举报详情
	ProcessResult string             `bson:"process_result"` //处理结果
}
