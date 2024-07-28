package loader

import (
	"bytes"

	"github.com/koron-go/jaconv"
)

func Load(data []byte) jaconv.Table {
	tbl, err := jaconv.Load(bytes.NewReader(data))
	if err != nil {
		panic("failed to load default table: " + err.Error())
	}
	return tbl
}
