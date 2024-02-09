package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Item model
type Item struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Created      time.Time          `bson:"created,omitempty" json:"created,omitempty"`
	Modified     time.Time          `bson:"modified,omitempty" json:"modified,omitempty"`
	Status       string             `bson:"status,omitempty" json:"status,omitempty"`
	SerialNumber string             `bson:"serialnumber,omitempty" json:"serialnumber,omitempty"`
	PhysicalType string             `bson:"physicaltype,omitempty" json:"physicaltype,omitempty"`
	TagType      string             `bson:"tagtype,omitempty" json:"tagtype,omitempty"`
	Network      string             `bson:"network,omitempty" json:"network,omitempty"`
	Name         string             `validate:"required" bson:"name,omitempty" json:"name"`
	Location     string             `bson:"location,omitempty" json:"location,omitempty"`
	Note         string             `bson:"note,omitempty" json:"note,omitempty"`
}

//TODO: PhysicalType
//TODO: TagType
