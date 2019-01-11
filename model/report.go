package model

import "github.com/mongodb/mongo-go-driver/bson/primitive"

// Report ...
type Report struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	MediaID       primitive.ObjectID //举报视频ID
	ExoID         primitive.ObjectID //用户ID
	Types         string             //举报类型
	Detail        string             //举报详情
	ProcessResult string             //处理结果
}
