package mecab

import (
	"fmt"
	"strings"
)

var st = "I'm a Mr. Smith"

func NewNltk() {
	s := []string{}
	s = strings.Split(st, "")

	for _, x := range s {
		if x == "'" {}
	}

	fmt.Println(len(s))
}