package models

type IntervalId int

const (
	Daily IntervalId = iota
	Weekly
	Biweekly
	Monthly
	Quarterly
	Halftearly
	Yearly
)

type Interval struct {
	IntervalId `json:"intervalid"`
	Name       string `json:"name"`
	Value      string `json:"value"`
}

var Intervals = []Interval{
	{Daily, "daily", "daily"},
	{Weekly, "weekly", "weekly"},
	{Biweekly, "biweekly", "biweekly"},
	{Monthly, "monthly", "monthly"},
	{Quarterly, "quarterly", "quarterly"},
	{Halftearly, "halftearly", "halftearly"},
	{Yearly, "yearly", "yearly"},
}

func (e *Interval) GetInterval() string {
	return e.Name
}

func (c IntervalId) GetInterval() string {
	return Intervals[int(c)].Name
}
