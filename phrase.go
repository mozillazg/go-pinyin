package pinyin

import (
	"strings"

	"github.com/yanyiwu/gojieba"
)

var (
	jieba = gojieba.NewJieba()
)

func cutWords(s string) []string {
	return jieba.CutAll(s)
}

func pinyinPhrase(s string) string {
	words := cutWords(s)
	for _, word := range words {
		match := phraseDict[word]
		match = toFixed(match, paragraphOption)
		if match != "" {
			s = strings.ReplaceAll(s, word, " "+match+" ")
		}
	}

	splacesRegexp.ReplaceAllString(s, " ")

	return s
}
