package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Asset model
type Asset struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Created            time.Time          `bson:"created,omitempty" json:"created,omitempty"`
	Modified           time.Time          `bson:"modified,omitempty" json:"modified,omitempty"`
	Status             string             `bson:"status,omitempty" json:"status,omitempty"`
	AssetType          string             `bson:"assettype,omitempty" json:"assettype,omitempty"`
	IssuerAccount      Account            `bson:"issueraccount,omitempty" json:"issueraccount,omitempty"`
	DistributorAccount Account            `bson:"distributoraccount,omitempty" json:"distributoraccount,omitempty"`
	Name               string             `validate:"required" bson:"name,omitempty" json:"name,omitempty"`
	Abbr               string             `valid:"required,min=3,max=12" bson:"abbr,omitempty" json:"abbr,omitempty"`
	Public             string             `valid:"required,min=56,max=56" bson:"public,omitempty" json:"public,omitempty"`
	Note               string             `bson:"note,omitempty" json:"note,omitempty"`
}

type AssetTypeId int

const (
	KPI AssetTypeId = iota
	Unit
	Commercial
	NFT
	OtherAsset
)

type AssetType struct {
	AssetTypeId `json:"assettypeid"`
	Name        string `json:"name"`
	Value       string `json:"value"`
}

var AssetTypes = []AssetType{
	{KPI, "kpi", "kpi"},
	{Unit, "unit", "unit"},
	{Commercial, "commercial", "commercial"},
	{NFT, "nft", "nft"},
	{OtherAsset, "otherasset", "otherasset"},
}

func (e *AssetType) GetAssetType() string {
	return e.Name
}

func (c AssetTypeId) GetAssetType() string {
	return AssetTypes[int(c)].Name
}
