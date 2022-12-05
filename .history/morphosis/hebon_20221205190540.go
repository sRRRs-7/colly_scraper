package morphosis

import (
	"unicode/utf8"

	"golang.org/x/exp/utf8string"
)

func hebonConvert(kana string, index int) charHebon {
	var hebon string
	var char string

	utfStr := utf8string.NewString(kana)
	// 2 string hit
	if index + 1 < utf8.RuneCountInString(kana) {
		char = utfStr.Slice(index, index+2)
		hebon = hebonMap[char]
	}
	// 1 string hit
	if hebon == "" && index < utfStr.RuneCount() {
		char = utfStr.Slice(index, index+1)
		hebon = hebonMap[char]
	}
	return charHebon{
		Char: char,
		Hebon: hebon,
	}
}

func ToHebon(kana string) string {
	var hebon string
	var lastHebon string

	i := 0
	for {
		ch := hebonConvert(kana, i)

		// example: mitchel
		if ch.Char == "っ" {
			nextCh := hebonConvert(kana, i+1)
			if nextCh.Hebon != "" {
				ch.Hebon = "t"
			}
		} else if ch.Char == "-" {
			// "-" is ignore
			ch.Hebon = ""
		} else if ch.Char == "ん" {
			nextCh := hebonConvert(kana, i)
			if nextCh.Hebon != "" {}
		}
	}
}