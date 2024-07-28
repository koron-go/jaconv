package hira2kata_test

import (
	"fmt"

	"github.com/koron-go/jaconv/hira2kata"
)

func ExampleConvert() {
	fmt.Println(hira2kata.Convert("おはようございます"))
	// Output:
	// オハヨウゴザイマス
}
