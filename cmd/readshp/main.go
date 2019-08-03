package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/kellegous/shapefile/shp"
)

// Emit ...
func Emit(w io.Writer, src string) error {
	r, err := os.Open(src)
	if err != nil {
		return err
	}
	defer r.Close()

	sr, err := shp.NewReader(r)
	if err != nil {
		return err
	}

	if _, err := fmt.Fprintf(w, "%#v\n", sr.Header); err != nil {
		return err
	}

	for {
		s, err := sr.Next()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}

		if _, err := fmt.Fprintf(w, "%#v\n", s); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	flag.Parse()

	for _, arg := range flag.Args() {
		if err := Emit(os.Stdout, arg); err != nil {
			log.Panic(err)
		}
	}
}
