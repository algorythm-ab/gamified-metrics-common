package methods

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/algorythm-ab/gamified-metrics-common/models"
)

// RespondWithError - creating JSON response with error
func RespondWithError(w http.ResponseWriter, code int, err error) {
	var status models.StatusObject
	status.Description = err.Error()
	status.Success = false
	status.Error = err.Error()

	var returner models.ReturnObject
	returner.Status = status

	RespondWithJSON(w, code, returner)
}

// RespondWithJSON - creating JSON response
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.MarshalIndent(payload, "", "  ")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// CreateLinks - Creating links for return object
func CreateLinks(r *http.Request, count models.Count) models.Links {
	out := models.Links{}

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

// VersionNotFound - Return version not found
func VersionNotFound(w http.ResponseWriter, _ *http.Request) {
	var status models.StatusObject
	status.Description = "api version not found"
	status.Success = false
	status.Error = "api version not found"

	var returner models.ReturnObject
	returner.Status = status

	RespondWithJSON(w, http.StatusOK, returner)
}
