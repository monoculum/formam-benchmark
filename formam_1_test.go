package formam_benchmark

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/ajg/form"
	formm "github.com/go-playground/form"
	"github.com/gorilla/schema"
	"github.com/monoculum/formam"
)

type Simple struct {
	Nest struct {
		Children []struct {
			ID   string
			Name string
		}
	}
	String string
	Slice  []int
	Bool   bool
}

var (
	valFormamT2 = url.Values{
		"Nest.Children[0].ID":   []string{"monoculum_id"},
		"Nest.Children[0].Name": []string{"Monoculum"},
		"String":                []string{"golang is very fun"},
		"Slice[0]":              []string{"1"},
		"Slice[1]":              []string{"2"},
		"Slice[2]":              []string{"3"},
		"Slice[3]":              []string{"4"},
		"Bool":                  []string{"true"},
	}
	valFormT2 = url.Values{
		"Nest.Children[0].ID":   []string{"monoculum_id"},
		"Nest.Children[0].Name": []string{"Monoculum"},
		"String":                []string{"golang is very fun"},
		"Slice[0]":              []string{"1"},
		"Slice[1]":              []string{"2"},
		"Slice[2]":              []string{"3"},
		"Slice[3]":              []string{"4"},
		"Bool":                  []string{"true"},
	}
	valSchemaT2 = url.Values{
		"Nest.Children.0.ID":   []string{"monoculum_id"},
		"Nest.Children.0.Name": []string{"Monoculum"},
		"String":               []string{"golang is very fun"},
		"Slice":                []string{"1", "2", "3", "4"},
		"Bool":                 []string{"true"},
	}
	valAJGT2 = url.Values{
		"Nest.Children.0.ID":   []string{"monoculum_id"},
		"Nest.Children.0.Name": []string{"Monoculum"},
		"String":               []string{"golang is very fun"},
		"Slice.0":              []string{"1"},
		"Slice.1":              []string{"2"},
		"Slice.2":              []string{"3"},
		"Slice.3":              []string{"4"},
		"Bool":                 []string{"true"},
	}
	valuesJSONT2 = `
	{
		"Nest":
			{
				"Children": [{"ID": "monoculum_id", "Name":"Monoculum"}]
			},
		"string": "golang is very fun",
		"Slice": [1, 2, 3, 4],
		"Bool": true
	}
	`
)

func BenchmarkAJGFormTestSIMPLE(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ne := new(Simple)
		if err := form.DecodeValues(ne, valAJGT2); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkSchemaTestSIMPLE(b *testing.B) {
	dec := schema.NewDecoder()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ne := new(Simple)
		if err := dec.Decode(ne, valSchemaT2); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkFormamTestSIMPLE(b *testing.B) {
	b.ReportAllocs()
	dec := formam.NewDecoder(nil)
	for i := 0; i < b.N; i++ {
		test := new(Simple)
		if err := dec.Decode(valFormamT2, test); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkFormTestSIMPLE(b *testing.B) {
	decoder := formm.NewDecoder()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		test := new(Simple)
		if err := decoder.Decode(test, valFormT2); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkJSONTestSIMPLE(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		test := new(Simple)
		if err := json.Unmarshal([]byte(valuesJSONT2), test); err != nil {
			b.Error(err)
		}
	}
}
