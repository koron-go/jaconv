/*
Package roma2hira provides Roma-ji to Hiragana converter for Japanese.
*/
package roma2hira

import (
	_ "embed"

	"github.com/koron-go/jaconv"
	"github.com/koron-go/jaconv/internal/loader"
)

//go:embed roma2hira.tsv
var data []byte

// Table is convert table for Roma-ji to Hiragana
var Table jaconv.Table = loader.Load(data)

// Convert converts Roma-ji in s to Hiragana.
func Convert(s string) string {
	return Table.Convert(s)
}
