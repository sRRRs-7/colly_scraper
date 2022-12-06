package mecab

import (
	"fmt"
	"strings"
)

var st = "I'm a Mr. Smith"

func NewNltk() {
	s := []string{}
	s = strings.Split(st, "")
	str := strings.Join(s, " ")

	fmt.Println(len(s))
	fmt.Println(s)
}