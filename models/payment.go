package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StatusObject struct {
	Description string `json:"description,omitempty"`
	Success     bool   `json:"success,omitempty"`
	Error       error  `json:"error,omitempty"`
}

// SetTrustObject - used to handle json input
type SetTrustObject struct {
	StatusObject
	Account Account `json:"account"`
	Asset   Asset   `json:"asset"`
	Limit   string  `json:"limit,omitempty"`
}

// SetOptionsObject - used to handle json input
type SetOptionsObject struct {
	StatusObject
	MainAccount   Account          `json:"mainaccount"`
	SignerAccount Account          `json:"signeraccount,omitempty"`
	Thresholds    *ThresholdObject `json:"thresholds,omitempty"`
	Weight        int              `json:"weight,omitempty"`
}

type ThresholdObject struct {
	LowThreshold    int `json:"lowthreshold"`
	MediumThreshold int `json:"mediumthreshold"`
	HighThreshold   int `json:"highthreshold"`
}

// PaymentObject - used to handle json input
type PaymentObject struct {
	StatusObject
	FromAccount Account `json:"fromaccount"`
	ToAccount   Account `json:"toaccount"`
	Asset       Asset   `json:"asset"`
	Amount      string  `json:"amount"`
}

// SignObject - used to handle json input
type SignObject struct {
	StatusObject
	Transaction   string  `json:"transaction,omitempty"`
	SignerAccount Account `json:"signeraccount,omitempty"`
}

// TransObject - used to handle json input
type TransObject struct {
	TransactionId     primitive.ObjectID `json:"transactionid"`
	Created           time.Time          `json:"created"`
	Modified          time.Time          `json:"modified"`
	Status            string             `json:"status"`
	SourceAccount     Account            `json:"sourceaccount"`
	MaxTime           string             `json:"maxtime"`
	Envelope          string             `json:"envelope"`
	TransactionWeight int                `json:"transactionweight"`
	CurrentWeight     int                `json:"currentweight"`
}
