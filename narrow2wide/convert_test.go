package narrow2wide_test

import (
	"fmt"

	"github.com/koron-go/jaconv/narrow2wide"
)

func ExampleConvert() {
	fmt.Println(narrow2wide.Convert("ｵﾊﾖｳｺﾞｻﾞｲﾏｽ"))
	// Output:
	// オハヨウゴザイマス
}
