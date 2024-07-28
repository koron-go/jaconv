package hira2kata

import "testing"

func TestConvert(t *testing.T) {
	for i, c := range []struct{ src, want string }{
		{"あいうえお", "アイウエオ"},
		{"こんにちは", "コンニチハ"},
		{"こんにちわ", "コンニチワ"},
		{"こんびに", "コンビニ"},
		{"とっても", "トッテモ"},
		{"kstんhmyrw", "kstンhmyrw"},
		{"\xff", "\xff"},
	} {
		got := Convert(c.src)
		if got != c.want {
			t.Errorf("unexpected result: #%d case=%+v got=%s", i, c, got)
		}
	}
}
