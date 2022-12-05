package morphosis

import (
	"unicode/utf8"

	"golang.org/x/exp/utf8string"
)

func HebonConveert(kana string) {
	str := utf8string.NewString(kana)
	utf8.RuneCountInString()
}