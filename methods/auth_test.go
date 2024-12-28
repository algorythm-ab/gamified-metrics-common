package methods

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestExtractBearerToken(t *testing.T) {
	tests := []struct {
		name     string
		keycloak Keycloak
		token    string
		want     string
	}{
		{
			name:     "test case 1",
			keycloak: Keycloak{},
			token:    "Bearer token",
			want:     "token",
		},
		{
			name:     "test case 2",
			keycloak: Keycloak{},
			token:    "token",
			want:     "token",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.keycloak.extractBearerToken(tt.token); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bearer = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
func TestBasicAuth(t *testing.T) {
	type args struct {
		h    http.HandlerFunc
		code int
		err  error
	}
	tests := []struct {
		keycloak Keycloak
		name     string
		args     args
	}{
		{
			name:     "test case 1",
			keycloak: Keycloak{},
			args:     args{httptest.HandlerFunc, 502, errors.New("My Error")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.keycloak.VerifyToken(tt.args.h)
		})
	}
}
*/

func TestBasicAuthA(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, clienta")
	}))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	greeting, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", greeting)
}

func TestBasicAuthB(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	greeting, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", greeting)
}
