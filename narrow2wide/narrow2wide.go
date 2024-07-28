/*
Package narrow2wide provides converter from Narrow (Hankaku) to Wide (Zenkaku) for Japanese.
*/
package narrow2wide

import (
	"bytes"
	_ "embed"
	"fmt"

	"github.com/koron-go/jaconv"
)

//go:embed narrow2wide.tsv
var data []byte

// Table is convert table from Narrow (Hankaku) to Wide (Zenkaku) for Japanese.
var Table jaconv.Table

func init() {
	tbl, err := jaconv.Load(bytes.NewReader(data))
	if err != nil {
		panic(fmt.Sprintf("failed to load default table: %s", err))
	}
	Table = tbl
}

// Convert converts Narrow (Hankaku) to Wide (Zenkaku).
func Convert(s string) string {
	return Table.Convert(s)
}
