package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Workflow model
type Workflow struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Created  time.Time          `bson:"created,omitempty" json:"created"`
	Modified time.Time          `bson:"modified,omitempty" json:"modified"`
	Status   string             `bson:"status,omitempty" json:"status"`
	User     User               `bson:"user,omitempty" json:"user"`
	Actions  []Action           `bson:"actions,omitempty" json:"actions"`
	Name     string             `validate:"required" bson:"name,omitempty" json:"name"`
	Note     string             `bson:"note,omitempty" json:"note"`
}

//User_Id  primitive.ObjectID `bson:"user_id,omitempty" json:"user_id"`
