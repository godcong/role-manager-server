package model

import (
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	log "github.com/sirupsen/logrus"
)

// User ...
type User struct {
	Model          `bson:",inline"`
	Block          bool               `bson:"block"`           //禁止访问
	Name           string             `bson:"name"`            //名称
	Username       string             `bson:"username"`        //用户名
	Email          string             `bson:"email"`           //邮件
	Mobile         string             `bson:"mobile"`          //移动电话
	IDCardFacade   string             `bson:"id_card_facade"`  //身份证(正)
	IDCardObverse  string             `bson:"id_card_obverse"` //身份证(反)
	OrganizationID primitive.ObjectID `bson:"organization_id"` //组织ID
	Password       string             `bson:"password"`        //密码
	Certificate    string             `bson:"certificate"`     //证书
	PrivateKey     string             `bson:"private_key"`     //私钥
	LoginIP        string             `bson:"login_ip"`        //本次登录IP
	Token          string             `bson:"token"`
	role           *Role              `bson:"-"`
	organization   *Organization      `bson:"-"`
	permissions    []*Permission      `bson:"-"`
	menus          []*Menu            `bson:"menus"`
}

// IsExist ...
func (u *User) IsExist() bool {
	if u.ID != primitive.NilObjectID {
		return IDExist(u)
	}

	return IsExist(u, bson.M{
		"name": u.Name,
	})

}

// CreateIfNotExist ...
func (u *User) CreateIfNotExist() error {
	return CreateIfNotExist(u)
}

func (u *User) _Name() string {
	return "user"
}

// Update ...
func (u *User) Update() error {
	return UpdateOne(u)
}

// Delete ...
func (u *User) Delete() error {
	return DeleteByID(u)

}

// Create ...
func (u *User) Create() error {
	return InsertOne(u)

}

// Find ...
func (u *User) Find() error {
	return FindByID(u)
}

// SetPassword ...
func (u *User) SetPassword(pwd string) {
	u.Password = pwd
}

// ValidatePassword ...
func (u *User) ValidatePassword(pwd string) bool {
	return u.Password == pwd
}

// Role ...
func (u *User) Role() (*Role, error) {
	if u.role != nil {
		return u.role, nil
	}
	ru := NewRoleUser()
	ru.UserID = u.ID
	err := ru.Find()
	if err != nil {
		return nil, err
	}
	log.Println(*ru)
	role, err := ru.Role()
	if err != nil {
		return nil, err
	}
	u.role = role
	return role, nil
}

// Organization ...
func (u *User) Organization() (*Organization, error) {
	if u.organization != nil {
		return u.organization, nil
	}
	org := NewOrganization()
	org.ID = u.OrganizationID
	err := org.Find()
	if err != nil {
		return nil, err
	}
	u.organization = org
	return org, nil
}

func (u *User) Menus() ([]*Menu, error) {
	if u.menus != nil {
		return u.menus, nil
	}
	var ms []*Menu
	err := Find(NewPermissionUser(), bson.M{
		"user_id": u.ID,
	}, func(cursor mongo.Cursor) error {
		pu := NewPermissionUser()
		err := cursor.Decode(pu)
		if err != nil {
			return err
		}
		p, err := pu.Permission()
		if err != nil {
			return err
		}
		menu, err := p.Menu()
		if err != nil {
			return err
		}
		ms = append(ms, menu)
		return nil
	})
	return ms, err
}

// Permissions ...
func (u *User) Permissions() ([]*Permission, error) {
	if u.permissions != nil {
		return u.permissions, nil
	}
	var ps []*Permission
	err := Find(NewPermissionUser(), bson.M{
		"user_id": u.ID,
	}, func(cursor mongo.Cursor) error {
		pu := NewPermissionUser()
		err := cursor.Decode(pu)
		if err != nil {
			return err
		}
		p, err := pu.Permission()
		if err != nil {
			return err
		}
		ps = append(ps, p)
		return nil
	})
	u.permissions = ps
	return ps, err
}

// CheckPermission ...
func (u *User) CheckPermission(permission *Permission) error {
	pu := NewPermissionUser()
	err := FindOne(pu, bson.M{
		"user_id":       u.ID,
		"permission_id": permission.ID,
	})
	if err != nil {
		return err
	}
	return nil
}

// ALL ...
func (u *User) ALL() ([]*User, error) {
	var users []*User
	m := bson.M{}
	err := Find(u, m, func(cursor mongo.Cursor) error {
		log.Println(cursor.DecodeBytes())
		var u User
		err := cursor.Decode(&u)
		if err != nil {
			return err
		}
		users = append(users, &u)
		return nil
	})
	return users, err
}

// NewUser ...
func NewUser() *User {
	return &User{
		Model: model(),
	}
}
