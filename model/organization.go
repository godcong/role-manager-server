package model

import (
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
)

// VerifyApplication 申请中
const VerifyApplication = "application"

// VerifyPass 通过
const VerifyPass = "pass"

// VerifyReturn 打回
const VerifyReturn = "return"

// VerifyClosed 关闭
const VerifyClosed = "closed"

// Organization ...
type Organization struct {
	Model                  `bson:",inline"`
	IsDefault              bool   `bson:"is_default"`                //是否为默认
	Verify                 string `bson:"verify"`                    //验证状态
	Corporate              string `bson:"corporate"`                 //企业法人
	CorporateIDCardFacade  string `bson:"corporate_id_card_facade"`  //法人身份证(正)
	CorporateIDCardObverse string `bson:"corporate_id_card_obverse"` //法人身份证(反)
	BusinessLicense        string `bson:"business_license"`          //营业执照
	Name                   string `bson:"name"`                      //商户名称
	Code                   string `bson:"code"`                      //社会统一信用代码
	Contact                string `bson:"contact"`                   //商户联系人
	Position               string `bson:"position"`                  //联系人职位
	Phone                  string `bson:"phone"`                     //联系人手机号
	Mailbox                string `bson:"mailbox"`                   //联系人邮箱
	IDCardFacade           string `bson:"id_card_facade"`            //联系人身份证(正)
	IDCardObverse          string `bson:"id_card_obverse"`           //联系人身份证(反)
	Description            string `bson:"description"`               //描述
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
		return IDExist(o)
	}
	return IsExist(o, bson.M{
		"name": o.Name,
	})
}

// ALL ...
func (o *Organization) ALL() ([]*Organization, error) {
	var orgs []*Organization
	m := bson.M{}
	err := Find(o, m, func(cursor mongo.Cursor) error {
		log.Println(cursor.DecodeBytes())
		var o Organization
		err := cursor.Decode(&o)
		if err != nil {
			return err
		}
		orgs = append(orgs, &o)
		return nil
	})
	return orgs, err
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

// Users ...
func (o *Organization) Users() ([]*User, error) {
	var users []*User
	err := Find(NewUser(), bson.M{
		"organization_id": o.ID,
	}, func(cursor mongo.Cursor) error {
		user := NewUser()
		err := cursor.Decode(user)
		if err != nil {
			return err
		}
		users = append(users, user)
		return nil
	})
	return users, err
}

func (o *Organization) _Name() string {
	return "organization"
}
