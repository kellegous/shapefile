package shapefile

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"testing"
)

type TestFile struct {
	Name            string
	NumberOfRecords int
	NumberOfAttrs   int
}

type Test struct {
	File  string
	Files []*TestFile
}

func runTestFile(t *testing.T, zr *zip.Reader, file *TestFile) {
	r, err := NewReaderFromZip(zr, file.Name)
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()

	numberOfRecords := 0
	for i := 0; ; i++ {
		rec, err := r.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			t.Fatal(err)
		}
		numberOfRecords++

		if len(rec.Attrs()) != file.NumberOfAttrs {
			t.Fatalf("in %s record %d has %d attrs instead of %d",
				file.Name, i, len(rec.Attrs()), file.NumberOfAttrs)
		}
	}

	if numberOfRecords != file.NumberOfRecords {
		t.Fatalf("%s had %d records instead of %d",
			file.Name,
			numberOfRecords,
			file.NumberOfRecords)
	}
}

func runTest(t *testing.T, test *Test) {
	path := filepath.Join("test_files", test.File)
	s, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}

	r, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()

	zr, err := zip.NewReader(r, s.Size())
	if err != nil {
		t.Fatal(err)
	}

	for _, file := range test.Files {
		runTestFile(t, zr, file)
	}
}

func TestReaderFromZip(t *testing.T) {
	tests := []*Test{
		&Test{
			File: "all.zip",
			Files: []*TestFile{
				&TestFile{
					Name:            "pts",
					NumberOfRecords: 8,
					NumberOfAttrs:   23,
				},
				&TestFile{
					Name:            "pgn",
					NumberOfRecords: 1,
					NumberOfAttrs:   7,
				},
			},
		},
		&Test{
			File: "shp.zip",
			Files: []*TestFile{
				&TestFile{
					Name:            "pts",
					NumberOfRecords: 8,
					NumberOfAttrs:   0,
				},
				&TestFile{
					Name:            "pgn",
					NumberOfRecords: 1,
					NumberOfAttrs:   0,
				},
			},
		},
	}

	for _, test := range tests {
		runTest(t, test)
	}
}
