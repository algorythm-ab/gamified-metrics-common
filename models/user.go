package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User model
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Created  time.Time          `bson:"created,omitempty" json:"created"`
	Modified time.Time          `bson:"modified,omitempty" json:"modified"`
	Status   Status             `bson:"status,omitempty" json:"status"`
	UserType UserType           `bson:"usertype,omitempty" json:"usertype"`
	Accounts []Account          `bson:"accounts,omitempty" json:"accounts,omitempty"`
	Name     string             `validate:"required" bson:"name,omitempty" json:"name"`
	Email    string             `validate:"required,email" bson:"email,omitempty" json:"email"`
	Note     string             `bson:"note,omitempty" json:"note"`
}

type UserTypeId int

const (
	Default UserTypeId = iota
	AssetUser
	Trade
	Test
	OtherUser
)

type UserType struct {
	UserTypeId `json:"usertypeid"`
	Name       string `json:"name"`
	Value      string `json:"value"`
}

var UserTypes = []UserType{
	{Default, "user", "user"},
	{AssetUser, "assetuser", "assetuser"},
	{Trade, "trade", "trade"},
	{Test, "test", "test"},
	{OtherUser, "otheruser", "otheruser"},
}

func (e *UserType) GetUserType() string {
	return e.Name
}

func (c UserTypeId) GetUserType() string {
	return UserTypes[int(c)].Name
}
