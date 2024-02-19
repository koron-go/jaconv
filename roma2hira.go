package roma2hira

import (
	"bytes"
	_ "embed"
	"fmt"
)

//go:embed roma2hira.txt
var defaultData []byte

var defaultDict *Dict

func init() {
	d, err := Load(bytes.NewReader(defaultData))
	if err != nil {
		panic(fmt.Sprintf("failed to load default dict: %s", err))
	}
	defaultDict = d
}

func Convert(s string) (string, error) {
	return defaultDict.Convert(s)
}
