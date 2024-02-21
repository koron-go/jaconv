package jaconv

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"unicode/utf8"

	"github.com/koron-go/trietree"
)

type Table struct {
	tree    *trietree.STree
	entries []edgeEntry
}

type edgeEntry struct {
	emit   string
	remain string
}

func Load(r io.Reader) (*Table, error) {
	rr := csv.NewReader(r)
	rr.Comma = '\t'
	rr.Comment = '#'
	rr.FieldsPerRecord = -1
	rr.TrimLeadingSpace = true
	rr.ReuseRecord = true

	var tree trietree.DTree
	entries := make([]edgeEntry, 1) // an empty/dummy entries at 0

	for {
		records, err := rr.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}
		if len(records) < 2 {
			return nil, fmt.Errorf("requires at least 2 records per line")
		}
		key := records[0]
		entry := edgeEntry{emit: records[1]}
		if len(records) >= 3 {
			entry.remain = records[2]
		}
		id := tree.Put(key)
		if id != len(entries) {
			return nil, fmt.Errorf("duplicated entries for key %q", key)
		}
		entries = append(entries, entry)
	}
	return &Table{
		tree:    trietree.Freeze(&tree),
		entries: entries,
	}, nil
}

func (d *Table) Convert(s string) string {
	var buf bytes.Buffer
	for s != "" {
		prefix, id := d.tree.LongestPrefix(s)
		if id == 0 {
			r, n := utf8.DecodeRuneInString(s)
			if r == utf8.RuneError {
				// NOTE: n=0 is never happen
				buf.WriteString(s[:n])
				s = s[n:]
				continue
			}
			buf.WriteRune(r)
			s = s[n:]
			continue
		}
		entry := d.entries[id]
		buf.WriteString(entry.emit)
		s = s[len(prefix):]
		if entry.remain != "" {
			s = entry.remain + s
		}
	}
	return buf.String()
}
