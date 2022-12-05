package morphosis

import (
	"unicode/utf8"

	"golang.org/x/exp/utf8string"
)

func HebonConveert(kana string, index int) {
	var hebon string
	var char string

	str := utf8string.NewString(kana)
	if index + 1 < utf8.RuneCountInString(kana) {
		char = str.Slice()
		hebon = hebonMap[char]
	}
}