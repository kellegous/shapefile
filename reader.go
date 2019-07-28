package shp

import (
	"encoding/binary"
	"errors"
	"io"
)

// Reader ...
type Reader struct {
	shp io.Reader
	dbf io.Reader
	Header
}

// Header ...
type Header struct {
	FileLength int32
	ShapeType  int32
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
	hdr.ShapeType, err = readInteger(r, binary.LittleEndian)
	if err != nil {
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
	if _, err := readInteger(r.shp, binary.BigEndian); err != nil {
		return nil, err
	}

	cl, err := readInteger(r.shp, binary.BigEndian)
	if err == io.EOF {
		return nil, io.ErrUnexpectedEOF
	} else if err != nil {
		return nil, err
	}

	var st ShapeType
	if err := binary.Read(r.shp, binary.LittleEndian, &st); err != nil {
		return nil, err
	}

	switch st {
	case TypeNull:
		return readNull(r.shp, cl)
	case TypePoint:
		return readPoint(r.shp, cl)
	case TypeMultiPoint:
		return readMultiPoint(r.shp, cl)
	case TypePolyline:
		return readPolyline(r.shp, cl)
	case TypePolygon:
		return readPolygon(r.shp, cl)
	case TypePointM:
		return readPointM(r.shp, cl)
	case TypeMultiPointM:
		return readMultiPointM(r.shp, cl)
	case TypePolylineM:
		return readPolylineM(r.shp, cl)
	case TypePolygonM:
		return readPolygonM(r.shp, cl)
	case TypePointZ:
		return readPointZ(r.shp, cl)
	}

	// TODO(knorton): just discard the bytes for now.
	buf := make([]byte, cl*2-4)
	if _, err := io.ReadAtLeast(r.shp, buf, len(buf)); err != nil {
		return nil, err
	}

	return nil, nil
}

// NewReader ...
func NewReader(shp io.Reader, opts ...Option) (*Reader, error) {
	r := &Reader{
		shp: shp,
	}

	for _, opt := range opts {
		opt(r)
	}

	if err := readHeader(shp, &r.Header); err != nil {
		return nil, err
	}

	return r, nil
}
