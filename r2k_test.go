package nihon_test

import (
	"github.com/kinbiko/nihon-go"
	"testing"
)

func TestHiragana(t *testing.T) {
	tt := []struct{ exp, in string }{
		{in: "", exp: ""},
		{in: "a", exp: "あ"},
		{in: "ka", exp: "か"},
		{in: "tsu", exp: "つ"},
		{in: "shi", exp: "し"},
		{in: "nani", exp: "なに"},
		{
			// Kana pangram/isogram. Only uses basic forms of hiragana (i.e. no 'ga' and no 'gya', 'ttsu', 'nma')
			in:  "irohanihohetochirinuruwowakayotaresotsunenaramuuwinookuyamakefukoeteasakiyumemishiwehimosesu",
			exp: "いろはにほへとちりぬるをわかよたれそつねならむうゐのおくやまけふこえてあさきゆめみしゑひもせす",
		},
		{in: "nandeyanen", exp: "なんでやねん"},
		{in: "chuushajou", exp: "ちゅうしゃじょう"},
		{in: "chottomatte", exp: "ちょっとまって"},
		{in: "sakki", exp: "さっき"},
		{in: "massugu", exp: "まっすぐ"},
		{in: "mettani", exp: "めったに"},
		{in: "happyou", exp: "はっぴょう"},
		{in: "joji", exp: "じょじ"},
	}
	for _, tc := range tt {
		t.Run(tc.in, func(st *testing.T) {
			got, err := nihon.ToKana(tc.in)
			if err != nil {
				st.Fatalf("Unexpected error: %s", err.Error())
			}
			if tc.exp != got {
				st.Fatalf("expected '%s' but got '%s'", tc.exp, got)
			}
		})
	}
}

func TestKatakana(t *testing.T) {
	tt := []struct{ exp, in string }{
		{in: "A", exp: "ア"},
		{in: "AKA", exp: "アカ"},
		{in: "BAAGENSEERU", exp: "バーゲンセール"},
		{in: "GURUDOBURANDOSEN", exp: "グルドブランドセン"},
		{in: "PYUU", exp: "ピュー"},
		{in: "TOUKYOU", exp: "トーキョー"},
	}
	for _, tc := range tt {
		t.Run(tc.in, func(st *testing.T) {
			got, err := nihon.ToKana(tc.in)
			if err != nil {
				st.Fatalf("Unexpected error: %s", err.Error())
			}
			if tc.exp != got {
				st.Fatalf("expected '%s' but got '%s'", tc.exp, got)
			}
		})
	}
}

func TestNegativeCases(t *testing.T) {
	tt := []struct{ exp, in string }{
		{in: "fishu", exp: "parse error: parsed up until 'しゅい' when getting 'could not find right-most kana for 'f''"},
		{in: "-", exp: "parse error: parsed up until '' when getting 'could not find right-most kana for '-''"},
	}
	for _, tc := range tt {
		t.Run(tc.in, func(st *testing.T) {
			got, err := nihon.ToKana(tc.in)
			if err == nil {
				st.Fatalf("Expected error but successfully parsed '%s' as '%s'", tc.in, got)
			}
			if err.Error() != tc.exp {
				st.Fatalf("Expected error '%s' but got '%s'", tc.exp, err.Error())
			}
		})
	}
}
