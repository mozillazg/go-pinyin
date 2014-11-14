package pinyin

import (
	"strings"
)

const (
	NORMAL       = 0 // 普通风格，不带声调。如： pin yin
	TONE         = 1 // 声调风格1，拼音声调在韵母第一个字母上（默认风格）。如： pīn yīn
	TONE2        = 2 // 声调风格2，即拼音声调在各个拼音之后，用数字 [0-4] 进行表示。如： pi1n yi1n
	INITIALS     = 3 // 声母风格，只返回各个拼音的声母部分。如： 中国 的拼音 zh g
	FIRST_LETTER = 4 // 首字母风格，只返回拼音的首字母部分。如： p y
	FINALS       = 5 // 韵母风格1，只返回各个拼音的韵母部分，不带声调。如： ong uo
	FINALS_TONE  = 6 // 韵母风格2，带声调，声调在韵母第一个字母上。如： ōng uó
	FINALS_TONE2 = 7 // 韵母风格2，带声调，声调在各个拼音之后，用数字 [0-4] 进行表示。如： o1ng uo2
)

type Args struct {
	Style     int
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
