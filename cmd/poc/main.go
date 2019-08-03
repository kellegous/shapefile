package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/kellegous/shapefile/shp"
)

func main() {
	flagShp := flag.String("shp", "", "shp file")
	flag.Parse()

	rs, err := os.Open(*flagShp)
	if err != nil {
		log.Panic(err)
	}
	defer rs.Close()

	r, err := shp.NewReader(rs)
	if err != nil {
		log.Panic(err)
	}
	log.Println(r)

	for {
		s, err := r.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Panic(err)
		}

		fmt.Printf("%#v\n", s)
	}
}
