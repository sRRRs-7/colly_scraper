package utils

import "strings"

func PriceReplace(s string) string {
	strings.Replace(s, "¥", "", -1)
}