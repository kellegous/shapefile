package shp

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

// Reader ...
type Reader struct {
	Header
	r io.Reader
}

// Header ...
type Header struct {
	FileLength int32
	ShapeType  ShapeType
	BBox       BBox
}

func readInteger(r io.Reader, o binary.ByteOrder) (int32, error) {
	var v int32
	if err := binary.Read(r, o, &v); err != nil {
		return 0, err
	}
	return v, nil
}

func readHeader(r io.Reader, hdr *Header) error {
	// File Code
	if code, err := readInteger(r, binary.BigEndian); err != nil {
		return err
	} else if code != 9994 {
		return errors.New("invalid file code")
	}

	var ignore [32]byte

	// Unused (20 bytes)
	if _, err := io.ReadAtLeast(r, ignore[:20], 20); err != nil {
		return err
	}

	// File Length
	var err error
	hdr.FileLength, err = readInteger(r, binary.BigEndian)
	if err != nil {
		return err
	}

	// Version
	if ver, err := readInteger(r, binary.LittleEndian); err != nil {
		return err
	} else if ver != 1000 {
		return errors.New("invalid version")
	}

	// Shape Type
	if err := binary.Read(r, binary.LittleEndian, &hdr.ShapeType); err != nil {
		return err
	}

	// BBox
	if err := binary.Read(r, binary.LittleEndian, &hdr.BBox); err != nil {
		return err
	}

	// Unused (32 bytes)
	if _, err := io.ReadAtLeast(r, ignore[:32], 32); err != nil {
		return err
	}

	return nil
}

// Next ...
func (r *Reader) Next() (Shape, error) {
	if _, err := readInteger(r.r, binary.BigEndian); err != nil {
		return nil, err
	}

	cl, err := readInteger(r.r, binary.BigEndian)
	if err == io.EOF {
		return nil, io.ErrUnexpectedEOF
	} else if err != nil {
		return nil, err
	}

	var st ShapeType
	if err := binary.Read(r.r, binary.LittleEndian, &st); err != nil {
		return nil, err
	}

	switch st {
	case TypeNull:
		return readNull(r.r, cl)
	case TypePoint:
		return readPoint(r.r, cl)
	case TypeMultiPoint:
		return readMultiPoint(r.r, cl)
	case TypePolyline:
		return readPolyline(r.r, cl)
	case TypePolygon:
		return readPolygon(r.r, cl)
	case TypePointM:
		return readPointM(r.r, cl)
	case TypeMultiPointM:
		return readMultiPointM(r.r, cl)
	case TypePolylineM:
		return readPolylineM(r.r, cl)
	case TypePolygonM:
		return readPolygonM(r.r, cl)
	case TypePointZ:
		return readPointZ(r.r, cl)
	case TypeMultiPointZ:
		return readMultiPointZ(r.r, cl)
	case TypePolylineZ:
		return readPolylineZ(r.r, cl)
	case TypePolygonZ:
		return readPolygonZ(r.r, cl)
	case TypeMultiPatch:
		return readMultiPatch(r.r, cl)
	}

	return nil, fmt.Errorf("unknown ShapeType %d", st)
}

// NewReader ...
func NewReader(r io.Reader) (*Reader, error) {
	sr := &Reader{
		r: r,
	}

	if err := readHeader(r, &sr.Header); err != nil {
		return nil, err
	}

	return sr, nil
}
