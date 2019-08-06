package shapefile

import (
	"archive/zip"
	"fmt"
	"io"
)

// ReadCloser provides a Reader that also has associated
// resources that must be closed by the caller.
type ReadCloser struct {
	*Reader
	io.Closer
}

type closeAll []io.Closer

func (c closeAll) Close() error {
	var err error
	for _, r := range c {
		err = r.Close()
	}
	return err
}

func findFile(files []*zip.File, name string) *zip.File {
	for _, file := range files {
		if name == file.Name {
			return file
		}
	}
	return nil
}

// NewReaderFromZip is a convenience utility for the common task of opening
// a shapefile and its associated attribute database from within a zip archive.
// The name provided should not include a file extension. For instance, if you
// with to open the archived file pts.shp (which also implies opening pts.dbf
// if it is present), name should be provided as "pts".
func NewReaderFromZip(r *zip.Reader, name string) (*ReadCloser, error) {
	var closer closeAll

	s := findFile(r.File, name+".shp")
	if s == nil {
		return nil, fmt.Errorf("file not found: %s", name)
	}

	sr, err := s.Open()
	if err != nil {
		return nil, err
	}

	closer = append(closer, sr)

	var opts []Option
	d := findFile(r.File, name+".dbf")
	if d != nil {
		dr, err := d.Open()
		if err != nil {
			closer.Close()
			return nil, err
		}

		closer = append(closer, dr)
		opts = append(opts, WithDBF(dr))
	}

	rr, err := NewReader(sr, opts...)
	if err != nil {
		closer.Close()
		return nil, err
	}

	return &ReadCloser{
		Reader: rr,
		Closer: closer,
	}, nil
}
