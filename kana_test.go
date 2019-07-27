package nihon

import "testing"

// This ensures that the skip values are correct for all elements in the katakana and hiragana maps, but this should really be solved
// by redesigning instead of testing.
func TestKanaSkipValues(t *testing.T) {
	for k, v := range hiragana {
		t.Run(k, func(st *testing.T) {
			if len(k) != v.skip {
				st.Errorf("Expected '%s' to have skip %d but was %d", k, len(k), v.skip)
			}
		})
	}
	for k, v := range katakana {
		t.Run(k, func(st *testing.T) {
			if len(k) != v.skip {
				st.Errorf("Expected '%s' to have skip %d but was %d", k, len(k), v.skip)
			}
		})
	}
}
