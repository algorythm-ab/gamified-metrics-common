package methods

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"
	"strings"
)

// ReturnObject - Model for ReturnObject
type ReturnObject struct {
	Records any
	Count   Count
	Links   Links
}

// Count - Model for Count
type Count struct {
	Total           int    `json:"total"`
	Page            int    `json:"page"`
	Limit           int    `json:"limit"`
	SelfPagingToken string `json:"selfpagingtoken"`
	PrevPagingToken string `json:"prevpagingtoken"`
	NextPagingToken string `json:"nextpagingtoken"`
}

// Links - Model for Links
type Links struct {
	Self  string `json:"self"`
	First string `json:"first"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
	Last  string `json:"last"`
}

// RespondWithError - creating JSON response with error
func RespondWithError(w http.ResponseWriter, code int, err error) {
	RespondWithJSON(w, code, map[string]string{"error": err.Error()})
}

// RespondWithJSON - creating JSON response
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.MarshalIndent(payload, "", "  ")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// CreateLinks - Creating links for return object
func CreateLinks(r *http.Request, count Count) Links {
	out := Links{}

	page := count.Page
	limit := count.Limit
	total := count.Total
	min := 1
	max := int(math.Ceil(float64(total) / float64(limit)))

	//TODO: check that page= must be higher than 1

	out.Self = r.URL.RequestURI()
	out.First = strings.Replace(out.Self, "page="+strconv.Itoa(count.Page), "page="+strconv.Itoa(min), 1)
	if page <= 1 {
		out.Prev = strings.Replace(out.Self, "page="+strconv.Itoa(count.Page), "page="+strconv.Itoa(min), 1)
	} else {
		out.Prev = strings.Replace(out.Self, "page="+strconv.Itoa(count.Page), "page="+strconv.Itoa(page-1), 1)
	}
	if page >= max {
		out.Next = strings.Replace(out.Self, "page="+strconv.Itoa(count.Page), "page="+strconv.Itoa(max), 1)
	} else {
		out.Next = strings.Replace(out.Self, "page="+strconv.Itoa(count.Page), "page="+strconv.Itoa(page+1), 1)
	}
	out.Last = strings.Replace(out.Self, "page="+strconv.Itoa(count.Page), "page="+strconv.Itoa(max), 1)

	return out
}
