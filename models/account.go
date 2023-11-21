package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Account model
type Account struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	User_ID     primitive.ObjectID `bson:"user_id,omitempty" json:"user_id"`
	Created     time.Time          `bson:"created,omitempty" json:"created"`
	Modified    time.Time          `bson:"modified,omitempty" json:"modified"`
	Status      Status             `bson:"status,omitempty" json:"status"`
	AccountType AccountType        `bson:"accounttype,omitempty" json:"accounttype"`
	Name        string             `validate:"required" bson:"name,omitempty" json:"name"`
	Public      string             `valid:"required,min=56,max=56" bson:"public" json:"public"`
	Secret      string             `valid:"required,min=56,max=56" bson:"secret" json:"secret"`
}

type AccountTypeId int

const (
	Primary AccountTypeId = iota
	Secondary
	Issuer
	Distributor
	OtherAccount
)

type AccountType struct {
	AccountTypeId `json:"id"`
	Name          string `json:"name"`
	Value         string `json:"value"`
}

var AccountTypes = []AccountType{
	{Primary, "primary", "primary"},
	{Secondary, "secondary", "secondary"},
	{Issuer, "issuer", "issuer"},
	{Distributor, "distributor", "distributor"},
	{OtherAccount, "otheraccount", "otheraccount"},
}

func (e *AccountType) GetAccountType() string {
	return e.Name
}

func (c AccountTypeId) GetAccountType() string {
	return AccountTypes[int(c)].Name
}