package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Task model
type Task struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Asset_Id  primitive.ObjectID `bson:"asset_id,omitempty" json:"asset_id,omitempty"`
	User_Id   primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
	Created   time.Time          `bson:"created,omitempty" json:"created,omitzero"`
	Modified  time.Time          `bson:"modified,omitempty" json:"modified,omitzero"`
	Deadline  time.Time          `bson:"deadline,omitempty" json:"deadline,omitzero"`
	Done      time.Time          `bson:"done,omitempty" json:"done,omitzero"`
	Status    string             `bson:"status,omitempty" json:"status,omitempty"`
	Artifacts *[]Artifact        `bson:"artifacts,omitempty" json:"artifacts,omitempty"`
	Items     *[]Item            `bson:"items,omitempty" json:"items,omitempty"`
	Payment   string             `bson:"payment,omitempty" json:"payment,omitempty"`
	Name      string             `validate:"required" bson:"name,omitempty" json:"name,omitempty"`
	Amount    int                `bson:"amount,omitempty" json:"amount,omitempty"`
	Note      string             `bson:"note,omitempty" json:"note,omitempty"`
}

//TODO:
//Add Asset (Asset_Id) as Asset?
//Add User (User_Id) as User?
