package main

import (
	"fmt"
	"io"
	"os"

	"github.com/kellegous/shapefile"
	"github.com/kellegous/shapefile/shp"
)

func main() {
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
}
