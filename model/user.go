package model

import (
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
)

// User ...
type User struct {
	Model          `bson:",inline"`
	Name           string             `bson:"name"`
	Username       string             `bson:"username"`
	Email          string             `bson:"email"`
	Mobile         string             `bson:"mobile"`
	IDCardFacade   string             `bson:"id_card_facade"`
	IDCardObverse  string             `bson:"id_card_obverse"`
	OrganizationID primitive.ObjectID `bson:"organization_id"`
	Password       string             `bson:"password"`
	Certificate    string             `bson:"certificate"`
	PrivateKey     string             `bson:"private_key"`

	Token string `bson:"token"`
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
	ru := NewRoleUser()
	ru.UserID = u.ID
	err := ru.Find()
	if err != nil {
		return nil, err
	}
	log.Println(*ru)
	return ru.Role()
}

// Permissions ...
func (u *User) Permissions() ([]*Permission, error) {
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
		var p User
		err := cursor.Decode(&p)
		if err != nil {
			return err
		}
		users = append(users, &p)
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
