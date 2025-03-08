package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TemplateWorkflow model
type TemplateWorkflow struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Owner_Id      primitive.ObjectID `bson:"owner_id,omitempty" json:"owner_id,omitempty"`
	Created       time.Time          `bson:"created,omitempty" json:"created,omitzero"`
	Modified      time.Time          `bson:"modified,omitempty" json:"modified,omitzero"`
	Status        string             `bson:"status,omitempty" json:"status,omitempty"`
	Version       string             `bson:"version,omitempty" json:"version,omitempty"`
	TemplateTasks *[]TemplateTask    `bson:"templatetasks,omitempty" json:"templatetasks,omitempty"`
	Name          string             `validate:"required" bson:"name,omitempty" json:"name,omitempty"`
	Note          string             `bson:"note,omitempty" json:"note,omitempty"`
}

//TODO:
//Add Owner (Owner_Id) as User?
