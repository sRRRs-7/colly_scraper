package mecab

import (
	"fmt"
	"strings"
)

var st = "I'm a Mr.Smith BOS/EOS (end of sentence) sub title "

func NewNltk() {
	str := strings.TrimRight(st, " ")
	str = strings.Replace(st, "'", " ' ", -1)	// single quote replace
	str = strings.Replace(str, ".", " . ", -1)	// period replace
	str = strings.Replace(str, ",", " , ", -1)	// comma replace
	str = strings.Replace(str, "-", " - ", -1)	// hyphen replace
	str = strings.Replace(str, "/", " / ", -1)	// slash replace
	str = strings.Replace(str, ":", " : ", -1)	// colon replace
	str = strings.Replace(str, ";", " ; ", -1)	// semicolon replace
	str = strings.Replace(str, "&", " & ", -1)	// and replace
	str = strings.Replace(str, "%", " % ", -1)	// percent replace
	str = strings.Replace(str, "(", " ( ", -1)	// left parenthesis replace
	str = strings.Replace(str, ")", " ) ", -1)	// right parenthesis replace
	str = strings.Replace(str, "_", " _ ", -1)	// under bar replace

	s := strings.Split(str, " ")
	a := []string{}
	for i, x := range s {
		if x == "" {
			a = append(s[:i], s[i+1:]...)
		}
	}

	fmt.Println(a)
	fmt.Println(len(a))
}