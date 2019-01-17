package model

import (
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
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
	Block                 bool      `bson:"block"`
	IP                    string    `bson:"ip"`
	LastLogin             time.Time `bson:"lastLogin"`
	V                     int       `bson:"__v"`
}

// NoPrefix ...
func (u *ExorcistUser) NoPrefix() NoPrefix {
	return true
}

// NewExorcistUser ...
func NewExorcistUser() *ExorcistUser {
	return &ExorcistUser{}
}

// BeforeInsert ...
func (u *ExorcistUser) BeforeInsert() {
	u.CreatedAt = time.Now()

}

// BeforeUpdate ...
func (u *ExorcistUser) BeforeUpdate() {

}

// BeforeDelete ...
func (u *ExorcistUser) BeforeDelete() {

}

// AfterInsert ...
func (u *ExorcistUser) AfterInsert() {

}

// AfterUpdate ...
func (u *ExorcistUser) AfterUpdate() {

}

// AfterDelete ...
func (u *ExorcistUser) AfterDelete() {

}

// IsExist ...
func (u *ExorcistUser) IsExist() bool {
	return false
}

// GetID ...
func (u *ExorcistUser) GetID() primitive.ObjectID {
	return u.ID
}

// SetID ...
func (u *ExorcistUser) SetID(id primitive.ObjectID) {
	u.ID = id
}

// CreateIfNotExist ...
func (u *ExorcistUser) CreateIfNotExist() error {
	return CreateIfNotExist(u)
}

// Create ...
func (u *ExorcistUser) Create() error {
	return InsertOne(u)
}

// Update ...
func (u *ExorcistUser) Update() error {
	return UpdateOne(u)
}

// Delete ...
func (u *ExorcistUser) Delete() error {
	return DeleteByID(u)
}

// All ...
func (u *ExorcistUser) All() ([]*ExorcistUser, error) {

	var users []*ExorcistUser
	m := bson.M{}
	err := Find(u, m, func(cursor mongo.Cursor) error {
		log.Println(cursor.DecodeBytes())
		var u ExorcistUser
		err := cursor.Decode(&u)
		if err != nil {
			return err
		}
		users = append(users, &u)
		return nil
	})
	return users, err
}

// Find ...
func (u *ExorcistUser) Find() error {
	return FindByID(u)
}

// SoftDelete ...
func (u *ExorcistUser) SoftDelete() bool {
	return false
}

func (u *ExorcistUser) _Name() string {
	return "dbusers"
}
