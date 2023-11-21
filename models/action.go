package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Action model
type Action struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Created  time.Time          `bson:"created,omitempty" json:"created"`
	Modified time.Time          `bson:"modified,omitempty" json:"modified"`
	Deadline time.Time          `bson:"deadline,omitempty" json:"deadline"`
	Status   string             `bson:"status,omitempty" json:"status"`
	Asset    Asset              `bson:"asset,omitempty" json:"asset"`
	User     User               `bson:"user,omitempty" json:"user"`
	Events   []Event            `bson:"events,omitempty" json:"events"`
	Payment  string             `bson:"payment,omitempty" json:"payment"`
	Name     string             `validate:"required" bson:"name,omitempty" json:"name"`
	Amount   int                `bson:"amount,omitempty" json:"amount"`
	Note     string             `bson:"note,omitempty" json:"note"`
}

//Workflow_Id primitive.ObjectID `bson:"workflow_id,omitempty" json:"workflow_id"`
//Asset_Id primitive.ObjectID `bson:"asset_id,omitempty" json:"asset_id"`
//User_Id  primitive.ObjectID `bson:"user_id,omitempty" json:"user_id"`
