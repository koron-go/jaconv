/*
Package hira2kata provides Hiragana to Katakana convert talbe for Japanese.
*/
package hira2kata

import (
	_ "embed"

	"github.com/koron-go/jaconv"
	"github.com/koron-go/jaconv/internal/loader"
)

//go:embed hira2kata.tsv
var data []byte

// Table is convert table from Hiragana to Katakana
var Table jaconv.Table = loader.Load(data)

// Convert converts Hiragana to Katakana.
func Convert(s string) string {
	return Table.Convert(s)
}
