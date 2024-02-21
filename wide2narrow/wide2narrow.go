/*
Package wide2narrow provides converter from Wide (Zenkaku) to Narrow (Hankaku) for Japanese.
*/
package wide2narrow

import (
	"bytes"
	_ "embed"
	"fmt"

	"github.com/koron-go/jaconv"
)

//go:embed wide2narrow.tsv
var data []byte

// Table is convert table from Wide (Zenkaku) to Narrow (Hankaku) for Japanese.
var Table *jaconv.Table

func init() {
	tbl, err := jaconv.Load(bytes.NewReader(data))
	if err != nil {
		panic(fmt.Sprintf("failed to load default table: %s", err))
	}
	Table = tbl
}

// Convert converts Wide (Zenkaku) to Narrow (Hankaku).
func Convert(s string) string {
	return Table.Convert(s)
}
