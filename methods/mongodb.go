package methods

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoPaginate struct {
	page  int64
	limit int64
}

func NewMongoPaginate(page int, limit int) *MongoPaginate {
	return &MongoPaginate{
		page:  int64(page),
		limit: int64(limit),
	}
}

func (mp *MongoPaginate) GetPaginatedOpts() *options.FindOptions {
	l := mp.limit
	skip := mp.page*mp.limit - mp.limit
	fOpt := options.FindOptions{Limit: &l, Skip: &skip}

	return &fOpt
}

func ToBSON_M(fields string, values string) bson.M {
	fieldsA := strings.Split(fields, ",")
	valuesA := strings.Split(values, ",")
	out := bson.M{}

	i := 0
	if len(fields) > 0 {
		for _, field := range fieldsA {
			if field[len(field)-3:] == "_id" {
				p, _ := primitive.ObjectIDFromHex(valuesA[i])
				out[field] = p
			} else {
				out[field] = valuesA[i]
			}
			i++
		}
	}

	return out
}

func ToBSON_D(fields string, values string) bson.D {
	fieldsA := strings.Split(fields, ",")
	valuesA := strings.Split(values, ",")
	out := bson.D{}

	i := 0
	if len(fields) > 0 {
		for _, field := range fieldsA {
			if valuesA[i] == "asc" {
				out = append(out, bson.E{Key: field, Value: 1})
			} else {
				out = append(out, bson.E{Key: field, Value: -1})
			}
			i++
		}
	}

	return out
}
