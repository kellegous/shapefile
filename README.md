## Go Shapefile Reader

[![Build Status](https://travis-ci.org/kellegous/shapefile.svg)](https://travis-ci.org/kellegous/shapefile)  [![GoDoc](https://godoc.org/github.com/kellegous/shapefile?status.svg)](http://godoc.org/github.com/kellegous/shapefile)

Pure Go implementation for reading of ESRI Shapefiles as specified by the [ESRI Shapfile Technical Description](https://www.esri.com/library/whitepapers/pdfs/shapefile.pdf). I created this as part of a project that ingests GIS data provided by the National Hurricane Center as I was unable to find similar libraries that didn't need signficant modification.

### Use

#### Importing
```go
import "github.com/kellegous/shapefile"
```

There are also subpackages for reading the SHP and DBF formats, the imports for those are `github.com/kellegous/shapefile/shp` and `github.com/kellegous/shapefile/dbf` respectively.

### Examples

#### Reading the records in a shapfile (shp and dbf)
```go
sr, err := os.Open("data.shp")
if err != nil {
    panic(err)
}
defer sr.Close()

dr, err := os.Open("data.dbf")
if err != nil {
    panic(err)
}
defer dr.Close()

r, err := shapefile.NewReader(sr, shapefile.WithDBF(dr))
if err != nil {
    panic(err)
}

for {
    rec, err := r.Next()
    if err == io.EOF {
        break
    } else if err != nil {
        panic(err)
    }

    switch s := rec.Shape.(type) {
    case *shp.Point:
        fmt.Printf("{X:%0.2f, Y:%.2f}", s.X, s.Y)
        // ...
    }

    for i, field := range r.Fields() {
        fmt.Printf("%s = %s\n", field.Name, rec.Attr(i))
    }
}
```

#### Reading shapfile contained in zip archive
```go
	s, err := os.Stat("al062018_5day_035.zip")
	if err != nil {
        panic(err)
	}

	r, err := os.Open("al062018_5day_035.zip")
	if err != nil {
        panic(err)
	}
	defer r.Close()

	zr, err := zip.NewReader(r, s.Size())
	if err != nil {
        panic(err)
    }

	sr, err := NewReaderFromZip(zr, "al062018-035_5day_pts")
	if err != nil {
        panic(err)
	}
    defer r.Close()
    
    for {
        rec, err := sr.Next()
        if err == io.EOF {
            break
        } else if err != nil {
            panic(err)
        }

        switch s := rec.Shape.(type) {
        case *shp.Point:
            fmt.Printf("{X:%0.2f, Y:%.2f}", s.X, s.Y)
            // ...
        }

        for i, field := range r.Fields() {
            fmt.Printf("%s = %s\n", field.Name, rec.Attr(i))
        }
    }
```