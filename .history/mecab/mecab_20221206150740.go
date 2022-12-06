package mecab

import (
	"fmt"

	"github.com/shogo82148/go-mecab"
)

type element struct {
	Str string
	Part string
}

func NewMecab() {
	tagger, err := mecab.New(map[string]string{"output-format-type": "wakati"})
	if err != nil {
		panic(err)
	}
	defer tagger.Destroy()

	s := "私はカツ丼が好き"
	node := parseMecab(tagger, s)

	list := mecabDetail(node)

	strs := mecabRetrieve(list, "名刺")
}

func parseMecab(tagger mecab.MeCab, s string) mecab.Node {
	node, err := tagger.ParseToNode(s)
	if err != nil {
		panic(err)
	}
	return node
}

func mecabDetail(node mecab.Node) []element {
	el := element{}
	list := make([]element, 0)
	for ; node != (mecab.Node{}); node = node.Next() {
		if node.Surface() == "" { continue }
		el.Str = node.Surface()
		el.Part = node.Feature()
		list = append(list, el)
		fmt.Printf("%s\t%s\n", node.Surface(), node.Feature())
	}
	return list
}

func mecabRetrieve(list []element, p string) []string {
	strs := make([]string, 0)
	for _, e := range list {
		if e.Part == p {
			strs = append(strs, e.Str)
		}
	}
	return strs
}

