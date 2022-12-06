package mecab

import (
	"fmt"

	"github.com/shogo82148/go-mecab"
)

func NewMecab() {
	tagger, err := mecab.New(map[string]string{})
if err != nil {
    panic(err)
}
defer tagger.Destroy()

result, err := tagger.Parse("こんにちは世界")
if err != nil {
    panic(err)
}
fmt.Println(result)
}