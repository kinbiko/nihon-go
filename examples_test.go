package nihon_test

import (
	"fmt"

	"github.com/kinbiko/nihon-go"
)

func ExampleToKana_hiragana() {
	s, _ := nihon.ToKana("yurushiteyattaradouya")
	fmt.Println(s)
	// output: ゆるしてやったらどうや
}

func ExampleToKana_katakana() {
	s, _ := nihon.ToKana("HAPPIIBAASUDEI")
	fmt.Println(s)
	// output: ハッピーバースデー
}
