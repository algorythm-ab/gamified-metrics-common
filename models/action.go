package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Action model
type Action struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	//Owner_Id   primitive.ObjectID `bson:"owner_id,omitempty" json:"owner_id,omitempty"`
	//User_Id    primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
	Created    time.Time `bson:"created,omitempty" json:"created,omitempty"`
	Modified   time.Time `bson:"modified,omitempty" json:"modified,omitempty"`
	Deadline   time.Time `bson:"deadline,omitempty" json:"deadline,omitempty"`
	Done       time.Time `bson:"done,omitempty" json:"done,omitempty"`
	Owner      User      `bson:"owner,omitempty" json:"owner,omitempty"`
	User       User      `bson:"user,omitempty" json:"user,omitempty"`
	Asset      Asset     `bson:"asset,omitempty" json:"asset,omitempty"`
	Status     string    `bson:"status,omitempty" json:"status,omitempty"`
	ActionType string    `bson:"actiontype,omitempty" json:"actiontype,omitempty"`
	Amount     int       `bson:"amount,omitempty" json:"amount,omitempty"`
	Name       string    `validate:"required" bson:"name,omitempty" json:"name,omitempty"`
}

type ActionTypeId int

const (
	ActionPayment ActionTypeId = iota
	ActionMessage
	ActionEmail
)

type ActionType struct {
	ActionTypeId `json:"id"`
	Name         string `json:"name"`
	Value        string `json:"value"`
}

var ActionTypes = []ActionType{
	{ActionPayment, "payment", "payment"},
	{ActionMessage, "message", "message"},
	{ActionEmail, "email", "email"},
}

func (e *ActionType) GetActionType() string {
	return e.Name
}

func (c ActionTypeId) GetActionType() string {
	return AccountTypes[int(c)].Name
}
