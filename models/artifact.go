package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Artifact model
type Artifact struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Item_Id  primitive.ObjectID `bson:"item_id,omitempty" json:"item_id,omitempty"`
	Created  time.Time          `bson:"created,omitempty" json:"created,omitempty"`
	Modified time.Time          `bson:"modified,omitempty" json:"modified,omitempty"`
	Status   string             `bson:"status,omitempty" json:"status,omitempty"`
	Version  int                `bson:"version,omitempty" json:"version,omitempty"`
	Name     string             `validate:"required" bson:"name,omitempty" json:"name,omitempty"`
	Url      string             `bson:"url,omitempty" json:"url,omitempty"`
	Code     string             `bson:"code,omitempty" json:"code,omitempty"`
}

//TODO:
//Add Item (Item_Id) as Item?
