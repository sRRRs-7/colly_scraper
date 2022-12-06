package utils

import "strings"

func PriceReplace(s string) {
	strings.Replace(s, "¥", "", -1)
}