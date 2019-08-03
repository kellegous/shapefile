package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/kellegous/shp"
)

var sources = []string{
	"multipointm",
	"polylinem",
	"polygonm",
	"multipointz",
	"polylinez",
	"polygonz",
	"multipatch",
}

func sizeOfMData(st shp.ShapeType, rec []byte) int32 {
	switch st {
	case shp.TypeMultiPointM, shp.TypeMultiPointZ:
		numPoints := int32(binary.LittleEndian.Uint32(rec[36:40]))
		return 8 + numPoints*4
	case shp.TypePolylineM, shp.TypePolylineZ, shp.TypePolygonM, shp.TypePolygonZ, shp.TypeMultiPatch:
		numPoints := int32(binary.LittleEndian.Uint32(rec[40:44]))
		return 8 + numPoints*4
	}
	return 0
}

func writeHeader(w io.Writer, h *shp.Header) error {
	// File Code
	if err := binary.Write(w, binary.BigEndian, int32(9994)); err != nil {
		return err
	}

	var zeros [32]byte

	// Unused (20 bytes)
	if _, err := w.Write(zeros[:20]); err != nil {
		return err
	}

	// File Length
	if err := binary.Write(w, binary.BigEndian, h.FileLength); err != nil {
		return err
	}

	// Version
	if err := binary.Write(w, binary.LittleEndian, int32(1000)); err != nil {
		return err
	}

	// ShapeType
	if err := binary.Write(w, binary.LittleEndian, h.ShapeType); err != nil {
		return err
	}

	// BBox
	if err := binary.Write(w, binary.LittleEndian, &h.BBox); err != nil {
		return err
	}

	// Unused (32 bytes)
	if _, err := w.Write(zeros[:32]); err != nil {
		return err
	}

	return nil
}

func rewriteSHP(src, dst string) error {
	r, err := os.Open(src)
	if err != nil {
		return err
	}
	defer r.Close()

	w, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer w.Close()

	// FileLength will need to be updated
	sr, err := shp.NewReader(r)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	for {
		var num int32
		if err := binary.Read(r, binary.BigEndian, &num); err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		var cl int32
		if err := binary.Read(r, binary.BigEndian, &cl); err != nil {
			return err
		}

		rec := make([]byte, cl*2)
		if _, err := io.ReadAtLeast(r, rec, len(rec)); err != nil {
			return err
		}

		st := shp.ShapeType(int32(binary.LittleEndian.Uint32(rec[:4])))

		ml := sizeOfMData(st, rec)
		sr.Header.FileLength -= ml

		if err := binary.Write(&buf, binary.BigEndian, num); err != nil {
			return err
		}

		if err := binary.Write(&buf, binary.BigEndian, cl-ml); err != nil {
			return err
		}

		if _, err := buf.Write(rec[:(cl-ml)*2]); err != nil {
			return err
		}
	}

	if err := writeHeader(w, &sr.Header); err != nil {
		return err
	}

	if _, err := w.Write(buf.Bytes()); err != nil {
		return err
	}

	return nil
}

func copyFile(src, dst string) error {
	r, err := os.Open(src)
	if err != nil {
		return err
	}
	defer r.Close()

	w, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer w.Close()

	if _, err := io.Copy(w, r); err != nil {
		return err
	}

	return nil
}

func main() {
	flagRoot := flag.String("root",
		".",
		"root directory")
	flag.Parse()

	for _, source := range sources {
		if err := rewriteSHP(
			filepath.Join(*flagRoot, fmt.Sprintf("test_files/%s.shp", source)),
			filepath.Join(*flagRoot, fmt.Sprintf("test_files/%s_no_m.shp", source))); err != nil {
			log.Panic(err)
		}

		if err := copyFile(
			filepath.Join(*flagRoot, fmt.Sprintf("test_files/%s.dbf", source)),
			filepath.Join(*flagRoot, fmt.Sprintf("test_files/%s_no_m.dbf", source))); err != nil {
			log.Panic(err)
		}
	}
}
