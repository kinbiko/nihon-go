package nihon_test

import (
	"fmt"
	"testing"

	"github.com/kinbiko/nihon-go"
)

func TestFindsUniqueKanji(t *testing.T) {
	for _, tc := range []struct {
		in  string
		exp []string
	}{
		{in: ""},
		{in: "just english characters"},
		{in: "ひらがな"},
		{in: "カタカナ"},
		{in: "々ヶ。、？！ー"},
		{in: "一", exp: []string{"一"}},
		{in: "人人", exp: []string{"人"}},
		{
			in:  "たった二つの真実見抜く、見た目は子供、頭脳は大人。その名は、名探偵困難",
			exp: []string{"二", "真", "実", "見", "抜", "目", "子", "供", "頭", "脳", "大", "人", "名", "探", "偵", "困", "難"},
		},
	} {
		t.Run(fmt.Sprintf("should extract %v from %s", tc.exp, tc.in), func(st *testing.T) {
			if got := nihon.UniqueKanji(tc.in); !same(got, tc.exp) {
				st.Errorf("expected to pull out \n%v\nfrom the string '%s' but got \n%v", tc.exp, tc.in, got)
			}
		})
	}
}

func same(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
