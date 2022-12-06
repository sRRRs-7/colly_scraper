package mecab

import (
	"fmt"
	"strings"
)

var st = "I'm a Mr. Smith"

func NewNltk() {
	s := []string{}
	s = strings.Split(st, "")

	for i, x := range s {
		if x == `'` {
			s[:i]
		}
	}

	fmt.Println(len(s))
}