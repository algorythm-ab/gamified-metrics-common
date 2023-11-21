package models

type StatusId int

const (
	Active StatusId = iota
	Inactive
	Delete
)

type Status struct {
	StatusId `json:"statusid"`
	Name     string `json:"name"`
	Value    string `json:"value"`
}

var Statuses = []Status{
	{Active, "active", "active"},
	{Inactive, "inactive", "inactive"},
	{Delete, "delete", "delete"},
}

func (e *Status) GetStatus() string {
	return e.Name
}

func (c StatusId) GetStatus() string {
	return Statuses[int(c)].Name
}
