package model

import (
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// Organization ...
type Organization struct {
	Model       `bson:",inline"`
	IsDefault   bool   `bson:"is_default"`
	Verify      string `bson:"verify"`   //验证状态
	Name        string `bson:"name"`     //商户名称
	Code        string `bson:"code"`     //社会统一信用代码
	Contact     string `bson:"contact"`  //商户联系人
	Position    string `bson:"position"` //联系人职位
	Phone       string `bson:"phone"`    //联系人手机号
	Mailbox     string `bson:"mailbox"`  //联系人邮箱
	Description string `bson:"description"`
}

// NewOrganization ...
func NewOrganization() *Organization {
	return &Organization{
		Model: model(),
	}
}

// IsExist ...
func (o *Organization) IsExist() bool {
	if o.ID != primitive.NilObjectID {
		return IsExist(o, bson.M{
			"_id": o.ID,
		})
	}
	return IsExist(o, bson.M{
		"name": o.Name,
	})
}

// CreateIfNotExist ...
func (o *Organization) CreateIfNotExist() error {
	return CreateIfNotExist(o)
}

// Create ...
func (o *Organization) Create() error {
	return InsertOne(o)
}

// Update ...
func (o *Organization) Update() error {
	return UpdateOne(o)
}

// Delete ...
func (o *Organization) Delete() error {
	return DeleteByID(o)
}

// Find ...
func (o *Organization) Find() error {
	return FindByID(o)
}

func (o *Organization) _Name() string {
	return "organization"
}
