package nihon_test

import (
	"testing"

	"github.com/kinbiko/nihon-go"
)

func TestRomaji2Kana(t *testing.T) {
	t.Run("hiragana", func(st *testing.T) {
		for _, tc := range []struct{ exp, in string }{
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
		} {
			st.Run(tc.in, func(sst *testing.T) {
				got, err := nihon.RomajiToKana(tc.in)
				if err != nil {
					sst.Fatalf("Unexpected error: %s", err.Error())
				}
				if tc.exp != got {
					sst.Fatalf("expected '%s' but got '%s'", tc.exp, got)
				}
			})
		}
	})

	t.Run("katakana", func(st *testing.T) {
		for _, tc := range []struct{ exp, in string }{
			{in: "A", exp: "ア"},
			{in: "AKA", exp: "アカ"},
			{in: "BAAGENSEERU", exp: "バーゲンセール"},
			{in: "GURUDOBURANDOSEN", exp: "グルドブランドセン"},
			{in: "PYUU", exp: "ピュー"},
			{in: "TOUKYOU", exp: "トーキョー"},
		} {
			st.Run(tc.in, func(sst *testing.T) {
				got, err := nihon.RomajiToKana(tc.in)
				if err != nil {
					sst.Fatalf("Unexpected error: %s", err.Error())
				}
				if tc.exp != got {
					sst.Fatalf("expected '%s' but got '%s'", tc.exp, got)
				}
			})
		}
	})

	t.Run("negative cases", func(st *testing.T) {
		for _, tc := range []struct{ exp, in string }{
			{in: "fishu", exp: "parse error: parsed up until 'しゅい' when getting 'could not find right-most kana for 'f''"},
			{in: "-", exp: "parse error: parsed up until '' when getting 'could not find right-most kana for '-''"},
		} {
			st.Run(tc.in, func(sst *testing.T) {
				got, err := nihon.RomajiToKana(tc.in)
				if err == nil {
					sst.Fatalf("Expected error but successfully parsed '%s' as '%s'", tc.in, got)
				}
				if err.Error() != tc.exp {
					sst.Fatalf("Expected error '%s' but got '%s'", tc.exp, err.Error())
				}
			})
		}
	})
}
