package wide2narrow

import "testing"

func TestConvert(t *testing.T) {
	for i, c := range []struct{ src, want string }{
		{"アイウエオ", "ｱｲｳｴｵ"},
		{"コンニチハ", "ｺﾝﾆﾁﾊ"},
		{"コンニチワ", "ｺﾝﾆﾁﾜ"},
		{"コンビニ", "ｺﾝﾋﾞﾆ"},
		{"トッテモ", "ﾄｯﾃﾓ"},
		{"ｋｓｔンｈｍｙｒｗ", "kstﾝhmyrw"},
		{"\xff", "\xff"},
		{"”", "\""},
		{"ＡＢＣ１２３ｘｙｚ", "ABC123xyz"},
		{"　", " "},
	} {
		got := Convert(c.src)
		if got != c.want {
			t.Errorf("unexpected result: #%d case=%+v got=%s", i, c, got)
		}
	}
}
