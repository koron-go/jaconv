/*
Package wide2narrow provides converter from Wide (Zenkaku) to Narrow (Hankaku) for Japanese.
*/
package wide2narrow

import (
	_ "embed"

	"github.com/koron-go/jaconv"
	"github.com/koron-go/jaconv/internal/loader"
)

//go:embed wide2narrow.tsv
var data []byte

// Table is convert table from Wide (Zenkaku) to Narrow (Hankaku) for Japanese.
var Table jaconv.Table = loader.Load(data)

// Convert converts Wide (Zenkaku) to Narrow (Hankaku).
func Convert(s string) string {
	return Table.Convert(s)
}
