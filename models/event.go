package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Event model
type Event struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Created  time.Time          `bson:"created,omitempty" json:"created"`
	Modified time.Time          `bson:"modified,omitempty" json:"modified"`
	Status   string             `bson:"status,omitempty" json:"status"`
}

//Tag_Id    primitive.ObjectID `bson:"tag_id,omitempty" json:"tag_id"`
//Action_Id primitive.ObjectID `bson:"action_id,omitempty" json:"action_id"`
