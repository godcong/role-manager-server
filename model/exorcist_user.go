package model

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"time"
)

// ExorcistUser ...
type ExorcistUser struct {
	ID        primitive.ObjectID `bson:"_id"`
	Sn        []string           `bson:"sn"`
	IpfsID    []string           `bson:"ipfsid"`
	QuestList []struct {
		ID   string `bson:"id"`
		Code string `bson:"code"`
	} `bson:"questList"`
	Name                  string    `bson:"name"`
	Phone                 string    `bson:"phone"`
	Password              string    `bson:"password"`
	Nickname              string    `bson:"nickname"`
	PictureURL            string    `bson:"pictureUrl"`
	Level                 int       `bson:"level"`
	CreatedAt             time.Time `bson:"createdAt"`
	Binded                bool      `bson:"binded"`
	QueryApply            bool      `bson:"queryApply"`
	Order                 bool      `bson:"order"`
	WhaleCard             string    `bson:"whaleCard"`
	WhaleOrder            string    `bson:"whaleOrder"`
	SlotNum               int       `bson:"slot_num"`
	Approved              bool      `bson:"approved"`
	ParentID              string    `bson:"parentID"`
	Dvc                   string    `bson:"dvc"`
	WhaleDvc              string    `bson:"whaleDvc"`
	DragonBall            string    `bson:"dragonBall"`
	Master                string    `bson:"master"`
	WeChatUnionid         string    `bson:"weChatUnionid"`
	WeChatAppOpenid       string    `bson:"weChatAppOpenid"`
	WeChatAppToken        string    `bson:"weChatAppToken"`
	WeChatAppRefreshToken string    `bson:"weChatAppRefreshToken"`
	V                     int       `bson:"__v"`
}
