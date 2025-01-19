package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Measurement model
type Measurement struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Created   time.Time          `bson:"created,omitempty" json:"created,omitempty"`
	Modified  time.Time          `bson:"modified,omitempty" json:"modified,omitempty"`
	Status    string             `bson:"status,omitempty" json:"status,omitempty"`
	Owner     User               `bson:"owner,omitempty" json:"owner,omitempty"`
	Asset     Asset              `bson:"asset,omitempty" json:"asset,omitempty"`
	Amount    int                `bson:"amount,omitempty" json:"amount,omitempty"`
	Interval  string             `validate:"required" bson:"interval,omitempty" json:"interval,omitempty"`
	StartDate time.Time          `bson:"startdate,omitempty" json:"startdate,omitempty"`
	EndDate   time.Time          `bson:"enddate,omitempty" json:"enddate,omitempty"`
	MeasurementType
	Executions  []Execution  `bson:"executions,omitempty" json:"executions,omitempty"`
	Identifiers []Identifier `bson:"identifiers" json:"identifiers"`
	Type        string       `bson:"type,omitempty" json:"type,omitempty"`
	Name        string       `validate:"required" bson:"name,omitempty" json:"name,omitempty"`
	Note        string       `bson:"note,omitempty" json:"note,omitempty"`
}

type MeasurementType struct {
	Bonus      bool `bson:"bonus" json:"bonus"`
	GroupShare bool `bson:"groupshare" json:"groupshare"`

	MeasurementOnTime
	MeasurementOnTarget
	MeasurementConfidence
}

// Handling OnTime type of Measurement
type MeasurementOnTime struct {
	CheckDay int `bson:"checkday" json:"checkday"`
}

// Handling OnTarget type of Measurement
type MeasurementOnTarget struct {
	TargetReached bool `bson:"targetreached" json:"targetreached"`
	FixedAmount   bool `bson:"fixedamount" json:"fixedamount"`
}

// Handling Confidence type of Measurement
type MeasurementConfidence struct {
	ConfidenceStart  int  `bson:"confidencestart" json:"confidencestart"`
	ConfidenceEnd    int  `bson:"confidenceend" json:"confidenceend"`
	ConfidenceStep   int  `bson:"confidencestep" json:"confidencestep"`
	ConfidenceSigned bool `bson:"confidencesigned" json:"confidencesigned"`
}

// MeasurementSimplified model
type MeasurementSimplified struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Status string             `bson:"status,omitempty" json:"status,omitempty"`
	Name   string             `validate:"required" bson:"name,omitempty" json:"name,omitempty"`
}

// Execution model
type Execution struct {
	ExecutionInterval string    `bson:"executioninterval,omitempty" json:"executioninterval,omitempty"`
	ExecutionDate     time.Time `bson:"executiondate,omitempty" json:"executiondate,omitempty"`
	OkDate            time.Time `bson:"okdate,omitempty" json:"okdate,omitempty"`
	ExecutedDate      time.Time `bson:"executeddate,omitempty" json:"executeddate,omitempty"`
	Executed          bool      `bson:"executed,omitempty" json:"executed,omitempty"`
}

// Identifier model
type Identifier struct {
	Name   string `validate:"required" bson:"name,omitempty" json:"name,omitempty"`
	Value  string `validate:"required" bson:"value,omitempty" json:"value,omitempty"`
	Target int    `bson:"target" json:"target"` //or read target from erp
}
