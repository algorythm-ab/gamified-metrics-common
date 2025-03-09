package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Action model
type Action struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Created    time.Time          `bson:"created,omitempty" json:"created,omitzero"`
	Modified   time.Time          `bson:"modified,omitempty" json:"modified,omitzero"`
	Deadline   time.Time          `bson:"deadline,omitempty" json:"deadline,omitzero"`
	Done       time.Time          `bson:"done,omitempty" json:"done,omitzero"`
	Owner      User               `bson:"owner,omitempty" json:"owner,omitzero"`
	User       User               `bson:"user,omitempty" json:"user,omitzero"`
	Asset      Asset              `bson:"asset,omitempty" json:"asset,omitzero"`
	Status     string             `bson:"status,omitempty" json:"status,omitempty"`
	Logic      string             `bson:"logic,omitempty" json:"logic,omitempty"`
	Repeat     bool               `bson:"repeat,omitempty" json:"repeat,omitempty"`
	ActionType string             `bson:"actiontype,omitempty" json:"actiontype,omitempty"`
	Amount     int                `bson:"amount,omitempty" json:"amount,omitempty"`
	Name       string             `validate:"required" bson:"name,omitempty" json:"name,omitempty"`
}

type ActionTypeId int

const (
	ActionPayment ActionTypeId = iota
	ActionSign
	ActionMessage
	ActionEmail
	ActionCollection
)

type ActionType struct {
	ActionTypeId `json:"id"`
	Name         string `json:"name"`
	Value        string `json:"value"`
}

var ActionTypes = []ActionType{
	{ActionPayment, "payment", "payment"},
	{ActionSign, "sign", "sign"},
	{ActionMessage, "message", "message"},
	{ActionEmail, "email", "email"},
	{ActionCollection, "collection", "collection"},
}

func (e *ActionType) GetActionType() string {
	return e.Name
}

func (c ActionTypeId) GetActionType() string {
	return AccountTypes[int(c)].Name
}
