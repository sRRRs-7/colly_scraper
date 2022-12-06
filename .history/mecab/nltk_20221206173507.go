package mecab

import "strings"

var st = "I'm a Mr. Smith"

func NewNltk() {
	s := []string{}
	s = strings.Split(st, " ")
	s = strings.Split(st, "'")
}