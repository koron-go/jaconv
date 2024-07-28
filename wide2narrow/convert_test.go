package wide2narrow_test

import (
	"fmt"

	"github.com/koron-go/jaconv/wide2narrow"
)

func ExampleConvert() {
	fmt.Println(wide2narrow.Convert("オハヨウゴザイマス"))
	// Output:
	// ｵﾊﾖｳｺﾞｻﾞｲﾏｽ
}
