package nihon

// UniqueKanji takes a string of arbitrary characters and returns each kanji in
// the string exactly once, regardless of how many existed in the input, in the
// same order as they appear in the input.
// 々 ,ヶ, full-width punctuation and kana does not count as being kanji.
func UniqueKanji(in string) []string {
	withDupes := []string{}
	for _, r := range in {
		if r == '？' || r == '！' {
			continue
		}
		// Values greater than the beginning of the the 'common' han ideographs.
		// Definitely not bullet proof, but will do the trick for what I need
		if r >= '\u4e00' {
			withDupes = append(withDupes, string(r))
		}
	}
	vals := []string{}
	previouslySeen := map[string]bool{}
	for _, k := range withDupes {
		if !previouslySeen[k] {
			vals = append(vals, k)
			previouslySeen[k] = true
		}
	}
	return vals
}
