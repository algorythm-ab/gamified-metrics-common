package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Transaction model
type Transaction struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Created           time.Time          `bson:"created,omitempty" json:"created"`
	Modified          time.Time          `bson:"modified,omitempty" json:"modified"`
	Status            Status             `bson:"status,omitempty" json:"status"`
	TransactionType   TransactionType    `bson:"transactiontype,omitempty" json:"transactiontype"`
	SourceAccount     Account            `bson:"sourceaccount,omitempty" json:"sourceaccount"`
	SignerAccounts    []Account          `bson:"signeraccounts,omitempty" json:"signeraccounts,omitempty"`
	Maxtime           string             `bson:"maxtime,omitempty" json:"maxtime"`
	Envelope          string             `bson:"envelope,omitempty" json:"envelope"`
	TransactionWeight int                `bson:"transactionweight,omitempty" json:"transactionweight"`
	CurrentWeight     int                `bson:"currentweight,omitempty" json:"currentweight"`
}

type TransactionTypeId int

const (
	Payment TransactionTypeId = iota
	SetOptions
	SetTrust
)

type TransactionType struct {
	TransactionTypeId `json:"transactiontypeid"`
	Name              string `json:"name"`
	Value             string `json:"value"`
}

var TransactionTypes = []TransactionType{
	{Payment, "payment", "payment"},
	{SetOptions, "setoptions", "setoptions"},
	{SetTrust, "settrust", "settrust"},
}

func (e *UserType) GetTransactionType() string {
	return e.Name
}

func (c TransactionTypeId) GetTransactionType() string {
	return TransactionTypes[int(c)].Name
}
