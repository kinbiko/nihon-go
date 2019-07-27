package nihon

import "fmt"

type kana struct {
	val  string
	skip int
}

type kanas []kana

func (ks kanas) String() string {
	res := ""
	for _, k := range ks {
		res += k.val
	}
	return res
}

// ToKana returns the kana equivalent of the given romaji string.
// The input is assumed to have no punctuation, whitespace, and consist strictly of English characters.
// Any unexpected characters found in the input string will result in an error.
// Capitalised letters will be converted into katakana, while lower-cased letters are treated as hiragana.
// E.g. "a" -> "あ", and "A" -> "ア"
func ToKana(in string) (string, error) {
	var (
		rev kanas
		top = len(in)
	)
	for top > 0 {
		k, err := sub(in[0:top])
		if err != nil {
			return "", fmt.Errorf("parse error: parsed up until '%s' when getting '%s'", rev.String(), err.Error())
		}
		rev = append(rev, k)
		top -= k.skip
	}
	return reverse(rev).String(), nil
}

func sub(in string) (kana, error) {
	var k kana
	for l := 5; k.val == "" && l > 0; l-- {
		if len(in) >= l {
			k = hiragana[in[len(in)-l:]]
			if k.val == "" {
				k = katakana[in[len(in)-l:]]
			}
		}
	}
	if k.val == "" {
		return kana{}, fmt.Errorf("could not find right-most kana for '%s'", in)
	}
	return k, nil
}

func reverse(ks kanas) kanas {
	for i := len(ks)/2 - 1; i >= 0; i-- {
		opp := len(ks) - 1 - i
		ks[i], ks[opp] = ks[opp], ks[i]
	}
	return ks
}
