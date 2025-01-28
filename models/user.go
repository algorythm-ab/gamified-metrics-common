package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User model
type User struct {
	ID        *primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Created   *time.Time          `bson:"created,omitempty" json:"created,omitempty"`
	Modified  *time.Time          `bson:"modified,omitempty" json:"modified,omitempty"`
	Status    string              `bson:"status,omitempty" json:"status,omitempty"`
	UserType  string              `bson:"usertype,omitempty" json:"usertype,omitempty"`
	Accounts  []Account           `bson:"accounts,omitempty" json:"accounts,omitempty"`
	FirstName string              `validate:"required" bson:"firstname,omitempty" json:"firstname,omitempty"`
	LastName  string              `validate:"required" bson:"lastname,omitempty" json:"lastname,omitempty"`
	Name      string              `validate:"required" bson:"name,omitempty" json:"name,omitempty"`
	Email     string              `validate:"required,email" bson:"email,omitempty" json:"email,omitempty"`
	Note      string              `bson:"note,omitempty" json:"note,omitempty"`
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
