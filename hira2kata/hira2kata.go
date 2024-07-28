/*
Package hira2kata provides Hiragana to Katakana convert talbe for Japanese.
*/
package hira2kata

import (
	"bytes"
	_ "embed"
	"fmt"

	"github.com/koron-go/jaconv"
)

//go:embed hira2kata.tsv
var data []byte

// Table is convert table from Hiragana to Katakana
var Table jaconv.Table

func init() {
	tbl, err := jaconv.Load(bytes.NewReader(data))
	if err != nil {
		panic(fmt.Sprintf("failed to load default table: %s", err))
	}
	Table = tbl
}

// Convert converts Hiragana to Katakana.
func Convert(s string) string {
	return Table.Convert(s)
}
