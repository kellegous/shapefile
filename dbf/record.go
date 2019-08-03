package dbf

import "bytes"

// Record ...
type Record struct {
	data []byte
	r    *Reader
}

// Attr ...
func (r *Record) Attr(idx int) string {
	f := r.r.Fields[idx]
	return string(bytes.TrimSpace(r.data[f.off : f.off+int(f.Length)]))
}

// Attrs ...
func (r *Record) Attrs() []string {
	attrs := make([]string, 0, len(r.r.Fields))
	for i, n := 0, len(r.r.Fields); i < n; i++ {
		attrs = append(attrs, r.Attr(i))
	}
	return attrs
}
