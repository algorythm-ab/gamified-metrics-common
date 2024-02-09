package models

// Account model
type Account struct {
	Status      string `bson:"status,omitempty" json:"status,omitempty"`
	AccountType string `bson:"accounttype,omitempty" json:"accounttype,omitempty"`
	Name        string `validate:"required" bson:"name,omitempty" json:"name,omitempty"`
	Public      string `valid:"required,min=56,max=56" bson:"public" json:"public,omitempty"`
	Secret      string `valid:"required,min=56,max=56" bson:"secret" json:"secret,omitempty"`
}

type AccountTypeId int

const (
	Primary AccountTypeId = iota
	Secondary
	Issuer
	Distributor
	OtherAccount
)

type AccountType struct {
	AccountTypeId `json:"id"`
	Name          string `json:"name"`
	Value         string `json:"value"`
}

var AccountTypes = []AccountType{
	{Primary, "primary", "primary"},
	{Secondary, "secondary", "secondary"},
	{Issuer, "issuer", "issuer"},
	{Distributor, "distributor", "distributor"},
	{OtherAccount, "otheraccount", "otheraccount"},
}

func (e *AccountType) GetAccountType() string {
	return e.Name
}

func (c AccountTypeId) GetAccountType() string {
	return AccountTypes[int(c)].Name
}
