package narrow2wide

import "testing"

func TestConvert(t *testing.T) {
	for i, c := range []struct{ src, want string }{
		{"ｱｲｳｴｵ", "アイウエオ"},
		{"ｺﾝﾆﾁﾊ", "コンニチハ"},
		{"ｺﾝﾆﾁﾜ", "コンニチワ"},
		{"ｺﾝﾋﾞﾆ", "コンビニ"},
		{"ﾄｯﾃﾓ", "トッテモ"},
		{"kstﾝhmyrw", "ｋｓｔンｈｍｙｒｗ"},
		{"\xff", "\xff"},
		{"\"", "”"},
		{"ABC123xyz", "ＡＢＣ１２３ｘｙｚ"},
		{" ", "　"},
	} {
		got := Convert(c.src)
		if got != c.want {
			t.Errorf("unexpected result: #%d case=%+v got=%s", i, c, got)
		}
	}
}
