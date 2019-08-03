package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/kellegous/shapefile/dbf"
)

// Field ...
type Field struct {
	Name        [11]byte
	Type        byte
	ReservedA   [4]byte
	FieldLength byte
	FieldCount  byte
	WorkAreaID  uint16
	Example     byte
	ReservedB   [10]byte
	Index       byte
}

func read(src string) error {
	r, err := os.Open(src)
	if err != nil {
		return err
	}
	defer r.Close()

	var h [4]byte
	if _, err := io.ReadAtLeast(r, h[:], len(h)); err != nil {
		return err
	}

	for _, b := range h[:] {
		fmt.Printf("%x\n", b)
	}

	var numberOfRecords int32
	if err := binary.Read(r, binary.LittleEndian, &numberOfRecords); err != nil {
		return err
	}
	fmt.Printf("NumberOfRecords: %d\n", numberOfRecords)

	var numberOfBytesInHeader int16
	if err := binary.Read(r, binary.LittleEndian, &numberOfBytesInHeader); err != nil {
		return err
	}
	fmt.Printf("NumberOfBytesInHeader: %d\n", numberOfBytesInHeader)

	var numberOfBytesInRecord int16
	if err := binary.Read(r, binary.LittleEndian, &numberOfBytesInRecord); err != nil {
		return err
	}
	fmt.Printf("NumberOfBytesInRecord: %d\n", numberOfBytesInRecord)

	var headerJunk [20]byte
	if _, err := io.ReadAtLeast(r, headerJunk[:], len(headerJunk)); err != nil {
		return err
	}
	fmt.Printf("Junk: %#v\n", headerJunk)

	numberOfFields := (numberOfBytesInHeader - 33) / 32
	fmt.Printf("NumberOfFields: %d\n", numberOfFields)

	fields := make([]Field, numberOfFields)
	if err := binary.Read(r, binary.LittleEndian, &fields); err != nil {
		return err
	}

	for i, field := range fields {
		fmt.Printf("%d name = %s, type = %c, data = %v\n", i, field.Name, field.Type, field)
	}

	var endOfFields [1]byte
	if _, err := r.Read(endOfFields[:]); err != nil {
		return err
	}
	fmt.Printf("%x\n", endOfFields[0])

	buf := make([]byte, numberOfBytesInRecord)
	for i := int32(0); i < numberOfRecords; i++ {
		if _, err := io.ReadAtLeast(r, buf, len(buf)); err != nil {
			return err
		}
		fmt.Printf("%v\n", buf)
	}

	var end [1]byte
	if _, err := r.Read(end[:]); err != nil {
		return err
	}

	fmt.Printf("%x\n", end[0])
	return nil
}

func readWithLib(src string) error {
	r, err := os.Open(src)
	if err != nil {
		return err
	}
	defer r.Close()

	dbf, err := dbf.NewReader(r)
	if err != nil {
		return err
	}

	fmt.Printf("Day = %s\n", dbf.Day)
	fmt.Printf("NumberOfRecords = %d\n", dbf.NumberOfRecords)
	fmt.Printf("NumberOfFields = %d\n", dbf.NumberOfFields)
	fmt.Printf("RecordLength = %d\n", dbf.RecordLength)

	for _, field := range dbf.Fields {
		fmt.Printf("  %#v\n", field)
	}

	for i := 0; i < int(dbf.NumberOfRecords); i++ {
		r, err := dbf.Next()
		if err != nil {
			return err
		}
		fmt.Printf("%#v\n", r.Attrs())
	}

	return nil
}

func main() {
	flag.Parse()

	for _, arg := range flag.Args() {
		if err := readWithLib(arg); err != nil {
			log.Panic(err)
		}
	}
}
