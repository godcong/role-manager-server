package model

import (
	"errors"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// MediaUser ...
type MediaUser struct {
	Model   `bson:",inline"`
	UserID  primitive.ObjectID `bson:"user_id"`
	MediaID primitive.ObjectID `bson:"media_id"`
	user    *User
	media   *Media
}

// NewMediaUser ...
func NewMediaUser() *MediaUser {
	return &MediaUser{}
}

func (u *MediaUser) _Name() string {
	return "media_user"
}

// CreateIfNotExist ...
func (u *MediaUser) CreateIfNotExist() error {
	return CreateIfNotExist(u)
}

// IsExist ...
func (u *MediaUser) IsExist() bool {
	return IsExist(u, bson.M{
		"media_id": u.MediaID,
		"user_id":  u.UserID,
	})
}

// Media ...
func (u *MediaUser) Media() (*Media, error) {
	if u.ID == primitive.NilObjectID {
		return nil, errors.New("id is null")
	}
	if u.MediaID != primitive.NilObjectID {
		md := NewMedia()
		md.ID = u.MediaID
		err := md.Find()
		if err != nil {
			return nil, err
		}
		u.media = md
		return md, nil
	}
	return nil, errors.New("role not found")
}

// SetMedia ...
func (u *MediaUser) SetMedia(media *Media) {
	u.media = media
	u.MediaID = media.ID
}

// User ...
func (u *MediaUser) User() (*User, error) {
	if u.ID == primitive.NilObjectID {
		return nil, errors.New("id is null")
	}
	if u.UserID != primitive.NilObjectID {
		user := NewUser()
		user.ID = u.UserID
		err := user.Find()
		if err != nil {
			return nil, err
		}
		u.user = user
		return user, nil
	}
	return nil, errors.New("permission not found")
}

// SetUser ...
func (u *MediaUser) SetUser(user *User) {
	u.user = user
	u.UserID = user.ID
}

// GetID ...
func (u *MediaUser) GetID() primitive.ObjectID {
	return u.ID
}

// SetID ...
func (u *MediaUser) SetID(id primitive.ObjectID) {
	u.ID = id
}

// Create ...
func (u *MediaUser) Create() error {
	return InsertOne(u)
}

// Update ...
func (u *MediaUser) Update() error {
	return UpdateOne(u)
}

// Delete ...
func (u *MediaUser) Delete() error {
	return DeleteByID(u)
}

// Find ...
func (u *MediaUser) Find() error {
	return FindByID(u)
}
