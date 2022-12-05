package utils

import "regexp"

func regexRetriever(reg, s string) string {
	return regexp.MustCompile(reg).FindString(s)
}