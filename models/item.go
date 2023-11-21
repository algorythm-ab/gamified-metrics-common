package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Item model
type Item struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Created      time.Time          `bson:"created,omitempty" json:"created"`
	Modified     time.Time          `bson:"modified,omitempty" json:"modified"`
	Status       string             `bson:"status,omitempty" json:"status"`
	SerialNumber string             `bson:"serialnumber,omitempty" json:"serialnumber"`
	PhysicalType string             `bson:"physicaltype,omitempty" json:"physicaltype"`
	TagType      string             `bson:"tagtype,omitempty" json:"tagtype"`
	Network      string             `bson:"network,omitempty" json:"network"`
	Name         string             `validate:"required" bson:"name,omitempty" json:"name"`
	Location     string             `bson:"location,omitempty" json:"location"`
	Note         string             `bson:"note,omitempty" json:"note"`
}

//TODO: PhysicalType
//TODO: TagType
