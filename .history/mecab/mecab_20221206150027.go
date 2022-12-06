package mecab

import (
	"fmt"

	"github.com/shogo82148/go-mecab"
)



func NewMecab() {
	tagger, err := mecab.New(map[string]string{"output-format-type": "wakati"})
	if err != nil {
		panic(err)
	}
	defer tagger.Destroy()

	s := "私はカツ丼が好き"
	node := parseMecab(tagger, s)

	mecabDetail(node)
}

func parseMecab(tagger mecab.MeCab, s string) mecab.Node {
	node, err := tagger.ParseToNode(s)
	if err != nil {
		panic(err)
	}
	return node
}

type element struct {
	Str string
	Part string
}

func mecabDetail(node mecab.Node) {
	el := element{}
	list := make([]element, 100)
	for ; node != (mecab.Node{}); node = node.Next() {
		if node.Surface() == "" { continue }
		el.Str = node.Surface()
		el.Part = node.Feature()
		list = append(list, el)
		fmt.Printf("%s\t%s\n", node.Surface(), node.Feature())
	}
}

