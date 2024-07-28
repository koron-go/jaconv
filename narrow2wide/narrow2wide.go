/*
Package narrow2wide provides converter from Narrow (Hankaku) to Wide (Zenkaku) for Japanese.
*/
package narrow2wide

import (
	_ "embed"

	"github.com/koron-go/jaconv"
	"github.com/koron-go/jaconv/internal/loader"
)

//go:embed narrow2wide.tsv
var data []byte

// Table is convert table from Narrow (Hankaku) to Wide (Zenkaku) for Japanese.
var Table jaconv.Table = loader.Load(data)

// Convert converts Narrow (Hankaku) to Wide (Zenkaku).
func Convert(s string) string {
	return Table.Convert(s)
}
