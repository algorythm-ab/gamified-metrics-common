package methods

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/algorythm-ab/gamified-metrics-common/models"
	"github.com/pkg/errors"
)

func TestRespondWithError(t *testing.T) {
	type args struct {
		w    http.ResponseWriter
		code int
		err  error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test case 1",
			args: args{httptest.NewRecorder(), 502, errors.New("My Error")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RespondWithError(tt.args.w, tt.args.code, tt.args.err)
		})
	}
}

func TestRespondWithJSON(t *testing.T) {
	type args struct {
		w       http.ResponseWriter
		code    int
		payload interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test case 1",
			args: args{httptest.NewRecorder(), 200, "{'a': 'b'}"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RespondWithJSON(tt.args.w, tt.args.code, tt.args.payload)
		})
	}
}

func TestCreateLinks(t *testing.T) {
	type args struct {
		r     *http.Request
		count models.Count
	}
	tests := []struct {
		name string
		args args
		want models.Links
	}{
		{
			name: "test case 1",
			args: args{httptest.NewRequest(http.MethodGet, "/url?page=1", nil), models.Count{Total: 100, Page: 1, Limit: 1}},
			want: models.Links{Self: "/url?page=1", First: "/url?page=1", Prev: "/url?page=1", Next: "/url?page=2", Last: "/url?page=100"},
		},
		{
			name: "test case 2",
			args: args{httptest.NewRequest(http.MethodGet, "/url?page=1", nil), models.Count{Total: 100, Page: 1, Limit: 10}},
			want: models.Links{Self: "/url?page=1", First: "/url?page=1", Prev: "/url?page=1", Next: "/url?page=2", Last: "/url?page=10"},
		},
		{
			name: "test case 3",
			args: args{httptest.NewRequest(http.MethodGet, "/url?page=2&limit=10", nil), models.Count{Total: 100, Page: 2, Limit: 10}},
			want: models.Links{Self: "/url?page=2&limit=10", First: "/url?page=1&limit=10", Prev: "/url?page=1&limit=10", Next: "/url?page=3&limit=10", Last: "/url?page=10&limit=10"},
		},
		{
			name: "test case 4",
			args: args{httptest.NewRequest(http.MethodGet, "/url?page=3&limit=10", nil), models.Count{Total: 100, Page: 3, Limit: 10}},
			want: models.Links{Self: "/url?page=3&limit=10", First: "/url?page=1&limit=10", Prev: "/url?page=2&limit=10", Next: "/url?page=4&limit=10", Last: "/url?page=10&limit=10"},
		},
		{
			name: "test case 5",
			args: args{httptest.NewRequest(http.MethodGet, "/url?page=10&limit=10", nil), models.Count{Total: 100, Page: 10, Limit: 10}},
			want: models.Links{Self: "/url?page=10&limit=10", First: "/url?page=1&limit=10", Prev: "/url?page=9&limit=10", Next: "/url?page=10&limit=10", Last: "/url?page=10&limit=10"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateLinks(tt.args.r, tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateLinks() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
func TestRequestHandler(t *testing.T) {

	expected := "Hello john"

	req := httptest.NewRequest(http.MethodGet, "/greet?name=john", nil)

	w := httptest.NewRecorder()

	RequestHandler(w, req)

	res := w.Result()

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {

		t.Errorf("Error: %v", err)

	}

	if string(data) != expected {

		t.Errorf("Expected Hello john but got %v", string(data))

	}

}
*/
