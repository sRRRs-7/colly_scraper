package morphosis

import (
	"unicode/utf8"

	"golang.org/x/exp/utf8string"
)

func HebonConveert(kana string, index int) {
	var hebon string
	var char string

	utfStr := utf8string.NewString(kana)
	if index + 1 < utf8.RuneCountInString(kana) {
		char = utfStr.Slice(index, index+2)
		hebon = hebonMap[char]
	}
	if hebon == "" && index < utfStr.RuneCount() {
		char = utfStr.Slice(index, index+1)
		hebon = hebonMap[char]
	}
}