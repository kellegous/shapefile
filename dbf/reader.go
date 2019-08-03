package dbf

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"time"
)

// Reader ...
type Reader struct {
	Header
	rec Record
	r   io.Reader
	idx uint32
}

// Header ...
type Header struct {
	Day             time.Time
	NumberOfRecords uint32
	NumberOfFields  uint8
	RecordLength    uint16
	Fields          []*Field
}

// Next ...
func (r *Reader) Next() (*Record, error) {
	if r.idx == r.NumberOfRecords {
		return nil, io.EOF
	}

	if _, err := io.ReadAtLeast(r.r, r.rec.data, len(r.rec.data)); err != nil {
		return nil, err
	}

	r.idx++

	return &r.rec, nil
}

func readHeader(r io.Reader, h *Header) error {
	var buf [32]byte
	if _, err := io.ReadAtLeast(r, buf[:], len(buf)); err != nil {
		return err
	}

	h.Day = time.Date(1900+int(buf[1]),
		time.Month(buf[2]),
		int(buf[3]),
		0, 0, 0, 0, time.UTC)

	h.NumberOfRecords = binary.LittleEndian.Uint32(buf[4:8])

	hdrLen := binary.LittleEndian.Uint16(buf[8:10])
	h.NumberOfFields = uint8((hdrLen - 33) / 32)
	h.RecordLength = binary.LittleEndian.Uint16(buf[10:12])

	off := 1
	h.Fields = make([]*Field, 0, h.NumberOfFields)
	for i := uint8(0); i < h.NumberOfFields; i++ {
		if _, err := io.ReadAtLeast(r, buf[:], len(buf)); err != nil {
			return err
		}

		h.Fields = append(h.Fields, &Field{
			Name:         string(bytes.TrimRight(buf[:11], "\x00")),
			Type:         FieldType(buf[11]),
			Length:       buf[16],
			DecimalCount: buf[17],
			off:          off,
		})

		off += int(buf[16])
	}

	if _, err := r.Read(buf[:1]); err != nil {
		return err
	} else if buf[0] != 0xd {
		return fmt.Errorf("invalid dbf file, invalid field descriptor terminator: %x", buf[0])
	}

	return nil
}

// NewReader ...
func NewReader(r io.Reader) (*Reader, error) {
	rr := &Reader{
		r: r,
	}

	// read the field header
	if err := readHeader(r, &rr.Header); err != nil {
		return nil, err
	}

	// initialize the record buffer
	rr.rec = Record{
		data: make([]byte, rr.RecordLength),
		r:    rr,
	}

	return rr, nil
}
