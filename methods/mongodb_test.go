package methods

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestNewMongoPaginate(t *testing.T) {
	type args struct {
		page  int
		limit int
	}
	tests := []struct {
		name string
		args args
		want *MongoPaginate
	}{
		{
			name: "should return a pointer to a MongoPaginate",
			args: args{1, 1},
			want: &MongoPaginate{page: 1, limit: 1},
		},
		{
			name: "should return a pointer to a MongoPaginate",
			args: args{1, 2},
			want: &MongoPaginate{page: 1, limit: 2},
		},
		{
			name: "should return a pointer to a MongoPaginate",
			args: args{5, 5},
			want: &MongoPaginate{page: 5, limit: 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMongoPaginate(tt.args.page, tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMongoPaginate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMongoPaginate_GetPaginatedOpts(t *testing.T) {
	page := 1
	limit := 1
	mpa := NewMongoPaginate(page, limit)
	tests := []struct {
		name string
		mp   *MongoPaginate
		want *options.FindOptions
	}{
		{
			name: "should return a pointer to FindOptions",
			//mp:   NewMongoPaginate(page, limit),
			mp: mpa,
			//mp:   &MongoPaginate{page: page, limit: limit},
			//mp:   &MongoPaginate{1, 1},
			//want: &options.FindOptions{Limit: &page, Skip: &limit},
			want: &options.FindOptions{Limit: &mpa.limit, Skip: &mpa.page},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mp.GetPaginatedOpts(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MongoPaginate.GetPaginatedOpts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToBSON_M(t *testing.T) {
	type args struct {
		fields string
		values string
	}
	tests := []struct {
		name string
		args args
		want bson.M
	}{
		{
			name: "a",
			args: args{"Test", "1"},
			want: bson.M{"Test": "1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToBSON_M(tt.args.fields, tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToBSON_M() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToBSON_D(t *testing.T) {
	type args struct {
		fields string
		values string
	}
	tests := []struct {
		name string
		args args
		want bson.D
	}{
		{
			name: "a",
			args: args{"Test", "asc"},
			want: bson.D{{Key: "Test", Value: 1}},
		},
		{
			name: "a",
			args: args{"Test", "desc"},
			want: bson.D{{Key: "Test", Value: -1}},
		},
		{
			name: "a",
			args: args{"Test1,Test2", "asc,desc"},
			want: bson.D{{Key: "Test1", Value: 1}, {Key: "Test2", Value: -1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToBSON_D(tt.args.fields, tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToBSON_D() = %v, want %v", got, tt.want)
			}
		})
	}
}
