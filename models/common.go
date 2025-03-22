package models

// Count - Model for Count
type Count struct {
	Total int `json:"total"`
	Page  int `json:"page"`
	Limit int `json:"limit"`
	//CurrentPage int `json:"currentpage"`
	//PageSize    int `json:"pagesize"`
	SelfPagingToken string `json:"selfpagingtoken"`
	PrevPagingToken string `json:"prevpagingtoken"`
	NextPagingToken string `json:"nextpagingtoken"`
}

// Links - Model for Links
type Links struct {
	Self  string `json:"self,omitempty"`
	First string `json:"first,omitempty"`
	Prev  string `json:"prev,omitempty"`
	Next  string `json:"next,omitempty"`
	Last  string `json:"last,omitempty"`
}

// StatusObject - Model for handling status
type StatusObject struct {
	Description string `json:"description,omitempty"`
	Success     bool   `json:"success"`
	Error       string `json:"error,omitempty"`
}

// ReturnObject - Model for ReturnObject
type ReturnObject struct {
	Records any          `json:"records,omitempty"`
	Record  any          `json:"record,omitempty"`
	Count   Count        `json:"count,omitzero"`
	Links   Links        `json:"links,omitzero"`
	Status  StatusObject `json:"status,omitempty"`
}
