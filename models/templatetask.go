package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TemplateTask model
type TemplateTask struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Asset_Id  primitive.ObjectID `bson:"asset_id,omitempty" json:"asset_id,omitempty"`
	Owner_Id  primitive.ObjectID `bson:"owner_id,omitempty" json:"owner_id,omitempty"`
	Created   time.Time          `bson:"created,omitempty" json:"created,omitempty"`
	Modified  time.Time          `bson:"modified,omitempty" json:"modified,omitempty"`
	Deadline  int                `bson:"deadline,omitempty" json:"deadline,omitempty"`
	Status    string             `bson:"status,omitempty" json:"status,omitempty"`
	Version   string             `bson:"version,omitempty" json:"version,omitempty"`
	Artifacts *[]Artifact        `bson:"artifacts,omitempty" json:"artifacts,omitempty"`
	Name      string             `validate:"required" bson:"name,omitempty" json:"name,omitempty"`
	Amount    int                `bson:"amount,omitempty" json:"amount,omitempty"`
	Note      string             `bson:"note,omitempty" json:"note,omitempty"`
}

//TODO:
//Add Asset (Asset_Id) as Asset?
//Add Owner (Owner_Id) as User?
