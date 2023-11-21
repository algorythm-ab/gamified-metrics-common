package methods

import (
	"net/http"
	"reflect"
	"testing"
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
		count Count
	}
	tests := []struct {
		name string
		args args
		want Links
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateLinks(tt.args.r, tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateLinks() = %v, want %v", got, tt.want)
			}
		})
	}
}
