package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Item model
type Item struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Created      time.Time          `bson:"created,omitempty" json:"created,omitzero"`
	Modified     time.Time          `bson:"modified,omitempty" json:"modified,omitzero"`
	Status       string             `bson:"status,omitempty" json:"status,omitempty"`
	Owner        User               `bson:"owner,omitempty" json:"owner,omitzero"`
	SerialNumber string             `bson:"serialnumber,omitempty" json:"serialnumber,omitempty"`
	PhysicalType string             `bson:"physicaltype,omitempty" json:"physicaltype,omitempty"`
	ItemType     string             `bson:"itemtype,omitempty" json:"itemtype,omitempty"`
	Actions      []Action           `bson:"actions,omitempty" json:"actions,omitempty"`
	Network      string             `bson:"network,omitempty" json:"network,omitempty"`
	Repeat       bool               `bson:"repeat,omitempty" json:"repeat,omitempty"`
	Name         string             `validate:"required" bson:"name,omitempty" json:"name"`
	Location     string             `bson:"location,omitempty" json:"location,omitempty"`
	Note         string             `bson:"note,omitempty" json:"note,omitempty"`
}

// PhysicalType
type PhysicalTypeId int

const (
	Sticker PhysicalTypeId = iota
	Tag
	Card
	Key
	OtherPhysical
)

type PhysicalType struct {
	PhysicalTypeId `json:"id"`
	Name           string `json:"name"`
	Value          string `json:"value"`
}

var PhysicalTypes = []PhysicalType{
	{Sticker, "sticker", "sticker"},
	{Tag, "tag", "tag"},
	{Card, "card", "card"},
	{Key, "key", "key"},
	{OtherPhysical, "other", "other"},
}

func (e *PhysicalType) GetPhysicalType() string {
	return e.Name
}

func (c PhysicalTypeId) GetPhysicalType() string {
	return PhysicalTypes[int(c)].Name
}

// ItemType
type ItemTypeId int

const (
	ItemTask ItemTypeId = iota
	ItemAction
	ItemNFC
	ItemOther
)

type ItemType struct {
	ItemTypeId `json:"id"`
	Name       string `json:"name"`
	Value      string `json:"value"`
}

var ItemTypes = []ItemType{
	{ItemTask, "task", "task"},
	{ItemAction, "action", "action"},
	{ItemNFC, "nfc", "nfc"},
	{ItemOther, "other", "other"},
}

func (e *ItemType) GetItemType() string {
	return e.Name
}

func (c ItemTypeId) GetItemType() string {
	return ItemTypes[int(c)].Name
}
