package utils

import "regexp"

func RegexRetriever(reg, s string) string {
	return regexp.MustCompile(reg).FindString(s)
}

func RegexCheck() {}