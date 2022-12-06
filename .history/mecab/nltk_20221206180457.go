package mecab

import (
	"fmt"
	"strings"
)

var st = "I'm a Mr. Smith BOS/EOS "

func NewNltk() {
	str := strings.Replace(st, "'", " ' ", -1)	// single quote replace
	str = strings.Replace(str, ".", " . ", -1)	// period replace
	str = strings.Replace(str, ",", " , ", -1)	// comma replace
	str = strings.Replace(str, "-", " - ", -1)	// hyphen replace
	str = strings.Replace(str, "/", " / ", -1)	// slash replace
	str = strings.Replace(str, ":", " : ", -1)	// colon replace
	str = strings.Replace(str, ";", " ; ", -1)	// semicolon replace
	str = strings.Replace(str, "&", " & ", -1)	// and replace
	str = strings.Replace(str, "%", " % ", -1)	// percent replace
	str = strings.Replace(str, "(", " ( ", -1)	// percent replace
	str = strings.Replace(str, ")", " ) ", -1)	// percent replace

	s := strings.Split(str, " ")

	fmt.Println(str)
	fmt.Println(s)
}