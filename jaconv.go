/*
Package jaconv provides conversion table for Japanese.
*/
package jaconv

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"unicode/utf8"

	"github.com/koron-go/trietree/trie2"
)

// Table is conversion table for Japanese.
type Table struct {
	tree *trie2.STrie[edgeEntry]
}

type edgeEntry struct {
	emit   string
	remain string
}

// Load loads conversion table data in TSV format.
func Load(r io.Reader) (Table, error) {
	rr := csv.NewReader(r)
	rr.Comma = '\t'
	rr.Comment = '#'
	rr.FieldsPerRecord = -1
	rr.ReuseRecord = true

	var tree trie2.DTrie[edgeEntry]

	for {
		records, err := rr.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return Table{}, err
		}
		if len(records) < 2 {
			return Table{}, fmt.Errorf("requires at least 2 records per line")
		}
		key := records[0]
		entry := edgeEntry{emit: records[1]}
		if len(records) >= 3 {
			entry.remain = records[2]
		}
		tree.Put(key, entry)
	}
	return Table{tree: tree.Freeze(false)}, nil
}

// Convert converts s with this table.
func (tbl Table) Convert(s string) string {
	var buf bytes.Buffer
	for s != "" {
		entry, prefix, ok := tbl.tree.LongestPrefix(s)
		if !ok {
			r, n := utf8.DecodeRuneInString(s)
			if r == utf8.RuneError {
				// n=0 will never be happen
				buf.WriteString(s[:n])
				s = s[n:]
				continue
			}
			buf.WriteRune(r)
			s = s[n:]
			continue
		}
		buf.WriteString(entry.emit)
		s = s[len(prefix):]
		if entry.remain != "" {
			s = entry.remain + s
		}
	}
	return buf.String()
}
