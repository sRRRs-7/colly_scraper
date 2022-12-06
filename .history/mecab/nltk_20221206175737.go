package mecab

import (
	"fmt"
	"strings"
)

var st = "I'm a Mr. Smith"

func NewNltk() {
	s := []string{}
	str := strings.Replace(st, "'", " ' ", -1)
	s = strings.Split(st, " ")


	fmt.Println(s)
}