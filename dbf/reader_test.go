package dbf

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"
	"time"
)

type Test struct {
	Name   string
	Header Header
	Attrs  [][]string
}

var tests = []*Test{
	// test_files/null
	&Test{
		Name: "null",
		Header: Header{
			Day:             time.Date(2019, 8, 4, 0, 0, 0, 0, time.UTC),
			NumberOfRecords: 2,
			NumberOfFields:  6,
			RecordLength:    170,
			Fields: []*Field{
				&Field{
					Name:         "NAME",
					Type:         TypeCharacter,
					Length:       10,
					DecimalCount: 0,
				},
				&Field{
					Name:         "CHAR",
					Type:         TypeCharacter,
					Length:       50,
					DecimalCount: 0,
				},
				&Field{
					Name:         "DATE",
					Type:         TypeDate,
					Length:       8,
					DecimalCount: 0,
				},
				&Field{
					Name:         "VALF",
					Type:         TypeFloat,
					Length:       50,
					DecimalCount: 4,
				},
				&Field{
					Name:         "VALL",
					Type:         TypeLogical,
					Length:       1,
					DecimalCount: 0,
				},
				&Field{
					Name:         "VALN",
					Type:         TypeNumeric,
					Length:       50,
					DecimalCount: 0,
				},
			},
		},
		Attrs: [][]string{
			{"a", "foo", "20190112", "100.4500", "T", "1000"},
			{"b", "bar", "19750712", "17.4500", "F", "42"},
		},
	},

	// test_files/point
	&Test{
		Name: "point",
		Header: Header{
			Day:             time.Date(2019, 8, 4, 0, 0, 0, 0, time.UTC),
			NumberOfRecords: 2,
			NumberOfFields:  6,
			RecordLength:    170,
			Fields: []*Field{
				&Field{
					Name:         "NAME",
					Type:         TypeCharacter,
					Length:       10,
					DecimalCount: 0,
				},
				&Field{
					Name:         "CHAR",
					Type:         TypeCharacter,
					Length:       50,
					DecimalCount: 0,
				},
				&Field{
					Name:         "DATE",
					Type:         TypeDate,
					Length:       8,
					DecimalCount: 0,
				},
				&Field{
					Name:         "VALF",
					Type:         TypeFloat,
					Length:       50,
					DecimalCount: 4,
				},
				&Field{
					Name:         "VALL",
					Type:         TypeLogical,
					Length:       1,
					DecimalCount: 0,
				},
				&Field{
					Name:         "VALN",
					Type:         TypeNumeric,
					Length:       50,
					DecimalCount: 0,
				},
			},
		},
		Attrs: [][]string{
			{"a", "foo", "20190113", "100.4500", "T", "1000"},
			{"b", "bar", "19750709", "17.4500", "F", "42"},
		},
	},

	// test_files/multipoint
	&Test{
		Name: "multipoint",
		Header: Header{
			Day:             time.Date(2019, 8, 4, 0, 0, 0, 0, time.UTC),
			NumberOfRecords: 2,
			NumberOfFields:  1,
			RecordLength:    255,
			Fields: []*Field{
				&Field{
					Name:         "DATA",
					Type:         TypeCharacter,
					Length:       254,
					DecimalCount: 0,
				},
			},
		},
		Attrs: [][]string{
			{"a"},
			{"b"},
		},
	},
}

func fieldsAreSame(a, b *Field) bool {
	return a.Name == b.Name &&
		a.Type == b.Type &&
		a.Length == b.Length &&
		a.DecimalCount == b.DecimalCount
}

func allFieldsAreSame(a, b []*Field) bool {
	if len(a) != len(b) {
		return false
	}

	for i, n := 0, len(a); i < n; i++ {
		if !fieldsAreSame(a[i], b[i]) {
			return false
		}
	}

	return true
}

func headersAreSame(a, b *Header) bool {
	return a.Day.Equal(b.Day) &&
		a.NumberOfRecords == b.NumberOfRecords &&
		a.NumberOfFields == b.NumberOfFields &&
		a.RecordLength == b.RecordLength &&
		allFieldsAreSame(a.Fields, b.Fields)
}

func attrsAreSame(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, n := 0, len(a); i < n; i++ {
		if len(a[i]) != len(b[i]) {
			return false
		}

		for j, m := 0, len(a[i]); j < m; j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}

func toJSON(data interface{}) []byte {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}
	return b
}

func runTest(t *testing.T, test *Test) {
	r, err := os.Open(fmt.Sprintf("../test_files/%s.dbf", test.Name))
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()

	dbf, err := NewReader(r)
	if err != nil {
		t.Fatal(err)
	}

	if !headersAreSame(&dbf.Header, &test.Header) {
		t.Fatalf("headers differ for test %s expected %s got %s",
			test.Name,
			toJSON(test.Header),
			toJSON(dbf.Header))
	}

	var attrs [][]string
	for {
		r, err := dbf.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			t.Fatal(err)
		}
		attrs = append(attrs, r.Attrs())
	}

	if !attrsAreSame(attrs, test.Attrs) {
		t.Fatalf("attrs differ for test %s expected %s got %s",
			test.Name,
			toJSON(test.Attrs),
			toJSON(attrs))
	}
}

func TestAll(t *testing.T) {
	for _, test := range tests {
		runTest(t, test)
	}
}
