package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Workflow model
type Workflow struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Owner_Id            primitive.ObjectID `bson:"owner_id,omitempty" json:"owner_id,omitempty"`
	User_Id             primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
	TemplateWorkflow_Id primitive.ObjectID `bson:"templateworkflow_id,omitempty" json:"templateworkflow_id,omitempty"`
	Created             time.Time          `bson:"created,omitempty" json:"created,omitzero"`
	Modified            time.Time          `bson:"modified,omitempty" json:"modified,omitzero"`
	Done                time.Time          `bson:"done,omitempty" json:"done,omitzero"`
	Status              string             `bson:"status,omitempty" json:"status,omitempty"`
	Tasks               *[]Task            `bson:"tasks,omitempty" json:"tasks,omitempty"`
	Name                string             `validate:"required" bson:"name,omitempty" json:"name,omitempty"`
	Note                string             `bson:"note,omitempty" json:"note,omitempty"`
}

//TODO:
//Add Owner (Owner_Id) as User?
//Add User (User_Id) as User?
