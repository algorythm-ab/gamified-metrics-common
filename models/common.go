package models

type Count struct {
	CurrentPage int `json:"currentpage"`
	PageSize    int `json:"pagesize"`
	Total       int `json:"total"`
}

type ReturnObject struct {
	Records []any `json:"Records"`
	Count   Count `json:"Count"`
}
