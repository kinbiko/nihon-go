package nihon_test

import (
	"fmt"

	"github.com/kinbiko/nihon-go"
)

func ExampleRomajiToKana_hiragana() {
	s, _ := nihon.RomajiToKana("yurushiteyattaradouya")
	fmt.Println(s)
	// output: ゆるしてやったらどうや
}

func ExampleRomajiToKana_katakana() {
	s, _ := nihon.RomajiToKana("HAPPIIBAASUDEI")
	fmt.Println(s)
	// output: ハッピーバースデー
}

func ExampleUniqueKanji() {
	kanji := nihon.UniqueKanji("有名な名探偵")
	for _, k := range kanji {
		fmt.Print(k)
	}
	// output: 有名探偵
}
