package roma2hira_test

import (
	"fmt"

	"github.com/koron-go/jaconv/roma2hira"
)

func ExampleConvert() {
	fmt.Println(roma2hira.Convert("ohayougozaimasu"))
	// Output:
	// おはようございます
}
