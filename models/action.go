package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Action model
type Action struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Created  time.Time          `bson:"created,omitempty" json:"created,omitempty"`
	Modified time.Time          `bson:"modified,omitempty" json:"modified,omitempty"`
	Deadline time.Time          `bson:"deadline,omitempty" json:"deadline,omitempty"`
	Status   string             `bson:"status,omitempty" json:"status,omitempty"`
	Asset    *Asset             `bson:"asset,omitempty" json:"asset,omitempty"`
	User     *User              `bson:"user,omitempty" json:"user,omitempty"`
	Events   *[]Event           `bson:"events,omitempty" json:"events,omitempty"`
	Payment  string             `bson:"payment,omitempty" json:"payment,omitempty"`
	Name     string             `validate:"required" bson:"name,omitempty" json:"name,omitempty"`
	Amount   int                `bson:"amount,omitempty" json:"amount,omitempty"`
	Note     string             `bson:"note,omitempty" json:"note,omitempty"`
}

//Workflow_Id primitive.ObjectID `bson:"workflow_id,omitempty" json:"workflow_id"`
//Asset_Id primitive.ObjectID `bson:"asset_id,omitempty" json:"asset_id"`
//User_Id  primitive.ObjectID `bson:"user_id,omitempty" json:"user_id"`
