package formam_benchmark

import (
	"encoding/json"
	"github.com/ajg/form"
	"github.com/gorilla/schema"
	"github.com/monoculum/formam"
	"net/url"
	"testing"
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
		"Slice":                []string{"1", "2", "3", "4"},
		"Bool":                 []string{"true"},
	}
	valAJG = url.Values{
		"Nest.Children.0.Id":   []string{"monoculum_id"},
		"Nest.Children.0.Name": []string{"Monoculum"},
		"String":               []string{"golang is very fun"},
		"Slice.0":              []string{"1"},
		"Slice.1":              []string{"2"},
		"Slice.2":              []string{"3"},
		"Slice.3":              []string{"4"},
		"Bool":                 []string{"true"},
	}
	valuesJSON = `
	{
		"Nest":
			{
				"Children": [{"Id": "monoculum_id", "Name":"Monoculum"}]
			},
		"string": "golang is very fun",
		"Slice": [1, 2, 3, 4],
		"Bool": true
	}
	`
)

func BenchmarkAJGForm(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ne := new(BenchFormamSchema)
		if err := form.DecodeValues(ne, valAJG); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkSchema(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ne := new(BenchFormamSchema)
		dec := schema.NewDecoder()
		if err := dec.Decode(ne, valSchema); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkFormam(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		test := new(BenchFormamSchema)
		if err := formam.Decode(valFormam, test); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkJSON(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		test := new(BenchFormamSchema)
		if err := json.Unmarshal([]byte(valuesJSON), test); err != nil {
			b.Error(err)
		}
	}
}
