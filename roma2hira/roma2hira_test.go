package roma2hira

import "testing"

func TestConvert(t *testing.T) {
	for i, c := range []struct{ src, want string }{
		{"aiueo", "あいうえお"},
		{"konnnichiha", "こんにちは"},
		{"kon'nichiha", "こんにちは"},
		{"konbini", "こんびに"},
		{"combini", "こんびに"},
		{"tottemo", "とっても"},
		{"kstnhmyrw", "kstんhmyrw"},
		{"\xff", "\xff"},
	} {
		got := Convert(c.src)
		if got != c.want {
			t.Errorf("unexpected result: #%d case=%+v got=%s", i, c, got)
		}
	}
}
