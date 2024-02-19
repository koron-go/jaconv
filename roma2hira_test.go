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
		got, err := Convert(c.src)
		if err != nil {
			t.Errorf("failed to convert: #%d case=%+v", i, c)
		}
		if got != c.want {
			t.Errorf("unexpected result: #%d case=%+v got=%s", i, c, got)
		}
	}
}
