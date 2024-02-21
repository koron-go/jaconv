/*
Package roma2hira provides Roma-ji to Hiragana converter for Japanese.
*/
package roma2hira

import (
	"bytes"
	_ "embed"
	"fmt"

	"github.com/koron-go/jaconv"
)

//go:embed roma2hira.tsv
var data []byte

// Table is convert table for Roma-ji to Hiragana
var Table *jaconv.Table

func init() {
	tbl, err := jaconv.Load(bytes.NewReader(data))
	if err != nil {
		panic(fmt.Sprintf("failed to load default table: %s", err))
	}
	Table = tbl
}

// Convert converts Roma-ji in s to Hiragana.
func Convert(s string) string {
	return Table.Convert(s)
}
