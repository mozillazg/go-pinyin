package pinyin

import (
	"strings"
)

type Args struct {
	Style     string
	Heteronym bool
}

func SinglePinyin(r rune, a Args) []string {
	value, ok := PinyinDict[int(r)]
	pys := []string{}
	if ok {
		if len(value) < 1 || a.Heteronym {
			pys = strings.Split(value, ",")
		} else {
			pys = strings.Split(value, ",")[:1]
		}
	}
	return pys
}

func Pinyin(s string, a Args) [][]string {
	hans := []rune(s)
	pys := [][]string{}
	for _, r := range hans {
		pys = append(pys, SinglePinyin(r, a))
	}
	return pys
}
