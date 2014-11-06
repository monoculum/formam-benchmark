package formam_benchmark

import (
	"github.com/gorilla/schema"
	"testing"
	"net/url"
	"github.com/monoculum/formam"
)

type BenchFormamSchema struct {
	Nest struct {
		Children []struct {
			Id   string
			Name string
		}
	}
	String string
	Slice  []int
	Bool   bool
}

var (
	valFormam = url.Values{
		"Nest.Children[0].Id":   []string{"monoculum_id"},
		"Nest.Children[0].Name": []string{"Monoculum"},
		"String":                []string{"golang is very fun"},
		"Slice[0]":              []string{"1"},
		"Slice[1]":              []string{"2"},
		"Slice[2]":              []string{"3"},
		"Slice[3]":              []string{"4"},
		"Bool":                  []string{"true"},
	}
	valSchema = url.Values{
		"Nest.Children.0.Id":   []string{"monoculum_id"},
		"Nest.Children.0.Name": []string{"Monoculum"},
		"String":               []string{"golang is very fun"},
		"Slice":              []string{"1", "2", "3", "4"},
		"Bool":                 []string{"true"},
	}
)

func Benchmark2Formam(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test := new(BenchFormamSchema)
		if err := formam.Decode(valFormam, test); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkSchema(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ne := new(BenchFormamSchema)
		dec := schema.NewDecoder()
		if err := dec.Decode(ne, valSchema); err != nil {
			b.Error(err)
		}
	}
}
