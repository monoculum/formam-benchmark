package formam_benchmark

import (
	"encoding/json"
	"github.com/ajg/form"
	"github.com/monoculum/formam"
	"net/url"
	"testing"
)

type Anonymous struct {
	Int int `formam:"int" json:"int"`
}

type Bench struct {
	Nest struct {
		Children []struct {
			Id   string
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

type BenchAJG struct {
	Nest struct {
		Children []struct {
			Id   string
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
	valuesFormam = url.Values{
		"Nest.Children[0].Id":   []string{"monoculum_id"},
		"Nest.Children[0].Name": []string{"Monoculum"},
		"Map.es_Es[0]":          []string{"javier"},
		"Map.es_Es[1]":          []string{"javier"},
		"Map.es_Es[2]":          []string{"javier"},
		"Map.es_Es[3]":          []string{"javier"},
		"Map.es_Es[4]":          []string{"javier"},
		"Map.es_Es[5]":          []string{"javier"},
		"string":                []string{"golang is very fun"},
		"Slice[0]":              []string{"1"},
		"Slice[1]":              []string{"2"},
		"int":                   []string{"1"},
		"Bool":                  []string{"true"},
	}
	valuesAJGForm = url.Values{
		"Nest.Children.0.Id":   []string{"monoculum_id"},
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
	valuesJSON = `
	{
		"Nest":
			{
				"Children": [{"Id": "monoculum_id", "Name":"Monoculum"}]
			},
		"string": "golang is very fun",
		"Map": {"es_Es": ["javier", "javier", "javier", "javier", "javier", "javier"]},
		"Slice": [1, 2],
		"int": 20,
		"Bool": true
	}
	`
)

func BenchmarkAJGForm(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ne := new(BenchAJG)
		if err := form.DecodeValues(ne, valuesAJGForm); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkFormam(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test := new(Bench)
		if err := formam.Decode(valuesFormam, test); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test := new(Bench)
		if err := json.Unmarshal([]byte(valuesJSON), test); err != nil {
			b.Error(err)
		}
	}
}

/*
func BenchmarkSchema(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ne := new(Test)
		dec := schema.NewDecoder()
		if err := dec.Decode(ne, values); err != nil {
			b.Error(err)
		}
	}
}
*/
