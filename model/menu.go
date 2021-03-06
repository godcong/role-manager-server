package model

import (
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
)

// Menu ...
type Menu struct {
	Model       `bson:",inline"`
	PID         primitive.ObjectID `bson:"pid"`         //菜单关系
	Name        string             `bson:"name"`        //菜单名称
	Icon        string             `bson:"icon"`        //图标
	Slug        string             `bson:"slug"`        //菜单对应的权限
	URL         string             `bson:"url"`         //菜单链接地址
	Active      string             `bson:"active"`      //菜单高亮地址
	Description string             `bson:"description"` //描述
	Sort        int                `bson:"sort"`        //排序
}

// Find ...
func (m *Menu) Find() error {
	return FindByID(m)
}

// NewMenu ...
func NewMenu() *Menu {
	return &Menu{
		Model: model(),
	}
}
func (m *Menu) _Name() string {
	return "menus"
}

// CreateIfNotExist ...
func (m *Menu) CreateIfNotExist() error {
	return CreateIfNotExist(m)
}

// Create ...
func (m *Menu) Create() error {
	return InsertOne(m)
}

// Update ...
func (m *Menu) Update() error {
	return UpdateOne(m)
}

// Delete ...
func (m *Menu) Delete() error {
	return DeleteByID(m)
}

// All ...
func (m *Menu) All() ([]*Menu, error) {
	var Menus []*Menu
	b := bson.M{}
	err := Find(m, b, func(cursor mongo.Cursor) error {
		log.Println(cursor.DecodeBytes())
		var m Menu
		err := cursor.Decode(&m)
		if err != nil {
			return err
		}
		Menus = append(Menus, &m)
		return nil
	})
	return Menus, err
}
