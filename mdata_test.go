package shp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"testing"
)

var srcs = []string{
	"multipointm",
	"polylinem",
	"polygonm",
	"multipointz",
	"polylinez",
	"polygonz",
	"multipatch",
}

func sizeOfMData(st ShapeType, rec []byte) int32 {
	switch st {
	case TypeMultiPointM, TypeMultiPointZ:
		numPoints := int32(binary.LittleEndian.Uint32(rec[36:40]))
		return 8 + numPoints*4
	case TypePolylineM, TypePolylineZ, TypePolygonM, TypePolygonZ, TypeMultiPatch:
		numPoints := int32(binary.LittleEndian.Uint32(rec[40:44]))
		return 8 + numPoints*4
	}
	return 0
}

func writeHeader(w io.Writer, h *Header) error {
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

func rewrite(src, dst string) error {
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
	var hdr Header
	if err := readHeader(r, &hdr); err != nil {
		return err
	}

	var buf bytes.Buffer
	for {
		num, err := readInteger(r, binary.BigEndian)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		cl, err := readInteger(r, binary.BigEndian)
		if err != nil {
			return err
		}

		rec := make([]byte, cl*2)
		if _, err := io.ReadAtLeast(r, rec, len(rec)); err != nil {
			return err
		}

		fmt.Printf("rec: %d\n", len(rec))

		st := ShapeType(int32(binary.LittleEndian.Uint32(rec[:4])))
		fmt.Printf("Shapetype: %d\n", st)

		ml := sizeOfMData(st, rec)
		hdr.FileLength -= ml

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

	if err := writeHeader(w, &hdr); err != nil {
		return err
	}

	if _, err := w.Write(buf.Bytes()); err != nil {
		return err
	}

	return nil
}

func TestCreateTests(t *testing.T) {
	for _, src := range srcs {
		fmt.Printf("file: %s\n", src)
		if err := rewrite(fmt.Sprintf("test_files/%s.shp", src),
			fmt.Sprintf("test_files/%s_no_m.shp", src)); err != nil {
			t.Fatal(err)
		}
	}
}
