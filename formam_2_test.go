package formam_benchmark

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/ajg/form"
	formm "github.com/go-playground/form"
	"github.com/monoculum/formam"
)

type Anonymous struct {
	Int int `formam:"int" json:"int"`
}

type Medium struct {
	Nest struct {
		Children []struct {
			ID   string
			Name string
		}
	}
	String string `formam:"string" json:"string"`
	Slice  []int
	Map    map[string][]string
	Bool   bool
	Ptr    *string
	Anonymous
}

type Medium2 struct {
	Nest struct {
		Children []struct {
			ID   string
			Name string
		}
	}
	String string `form:"string"`
	Slice  []int
	Map    map[string][]string
	Int    int `form:"int"`
	Bool   bool
	Anonymous
}

var (
	valuesFormamT1 = url.Values{
		"Nest.Children[0].ID":   []string{"monoculum_id"},
		"Nest.Children[0].Name": []string{"Monoculum"},
		"Map[es_Es][0]":          []string{"javier"},
		"Map[es_Es][1]":          []string{"javier"},
		"Map[es_Es][2]":          []string{"javier"},
		"Map[es_Es][3]":          []string{"javier"},
		"Map[es_Es][4]":          []string{"javier"},
		"Map[es_Es][5]":          []string{"javier"},
		"string":                []string{"golang is very fun"},
		"Slice[0]":              []string{"1"},
		"Slice[1]":              []string{"2"},
		"int":                   []string{"1"},
		"Bool":                  []string{"true"},
	}
	valuesFormT1 = url.Values{
		"Nest.Children[0].ID":   []string{"monoculum_id"},
		"Nest.Children[0].Name": []string{"Monoculum"},
		"Map[es_Es][0]":         []string{"javier"},
		"Map[es_Es][1]":         []string{"javier"},
		"Map[es_Es][2]":         []string{"javier"},
		"Map[es_Es][3]":         []string{"javier"},
		"Map[es_Es][4]":         []string{"javier"},
		"Map[es_Es][5]":         []string{"javier"},
		"string":                []string{"golang is very fun"},
		"Slice[0]":              []string{"1"},
		"Slice[1]":              []string{"2"},
		"int":                   []string{"1"},
		"Bool":                  []string{"true"},
	}
	valuesAJGFormT1 = url.Values{
		"Nest.Children.0.ID":   []string{"monoculum_id"},
		"Nest.Children.0.Name": []string{"Monoculum"},
		"Map.es_Es.0":          []string{"javier"},
		"Map.es_Es.1":          []string{"javier"},
		"Map.es_Es.2":          []string{"javier"},
		"Map.es_Es.3":          []string{"javier"},
		"Map.es_Es.4":          []string{"javier"},
		"Map.es_Es.5":          []string{"javier"},
		"string":               []string{"golang is very fun"},
		"Slice.0":              []string{"1"},
		"Slice.1":              []string{"2"},
		"int":                  []string{"1"},
		"Bool":                 []string{"true"},
	}
	valuesJSONT1 = `
	{
		"Nest":
			{
				"Children": [{"ID": "monoculum_id", "Name":"Monoculum"}]
			},
		"string": "golang is very fun",
		"Map": {"es_Es": ["javier", "javier", "javier", "javier", "javier", "javier"]},
		"Slice": [1, 2],
		"int": 20,
		"Bool": true
	}
	`
)

func BenchmarkAJGFormTestMEDIUM(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ne := new(Medium2)
		if err := form.DecodeValues(ne, valuesAJGFormT1); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkFormamTestMEDIUM(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		test := new(Medium)
		if err := formam.Decode(valuesFormamT1, test); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkFormTestMEDIUM(b *testing.B) {
	decoder := formm.NewDecoder()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		test := new(Medium)
		if err := decoder.Decode(test, valuesFormT1); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkJSONTestMEDIUM(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		test := new(Medium)
		if err := json.Unmarshal([]byte(valuesJSONT1), test); err != nil {
			b.Error(err)
		}
	}
}
