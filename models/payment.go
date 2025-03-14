package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
	SignerAccount Account          `json:"signeraccount,omitzero"`
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
	Transaction   Transaction `json:"transaction,omitzero"`
	SignerAccount Account     `json:"signeraccount,omitzero"`
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

// OfferObject - used to handle json input
type OfferObject struct {
	StatusObject
	Seller  Account `json:"seller"`
	Selling Asset   `json:"selling"`
	Buying  Asset   `json:"buying"`
	Amount  string  `json:"amount"`
	Price   string  `json:"price"`
}
