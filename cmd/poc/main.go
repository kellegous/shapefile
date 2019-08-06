package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/kellegous/shapefile"
)

// Filename ...
const Filename = "al062018_5day_030.zip"

func main() {
	s, err := os.Stat(Filename)
	if err != nil {
		log.Panic(err)
	}

	r, err := os.Open(Filename)
	if err != nil {
		log.Panic(err)
	}
	defer r.Close()

	zr, err := zip.NewReader(r, s.Size())
	if err != nil {
		log.Panic(err)
	}

	sr, err := shapefile.NewReaderFromZip(
		zr,
		"al062018-030_5day_pts")
	if err != nil {
		log.Panic(err)
	}
	defer sr.Close()

	for {
		rec, err := sr.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Panic(err)
		}
		fmt.Printf("%#v\n%#v\n", rec.Shape, rec.Attrs())
	}
}
