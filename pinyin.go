package pinyin

import (
	"strings"
)

// 拼音风格
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

// 声母表
var _INITIALS = strings.Split(
	"zh,ch,sh,b,p,m,f,d,t,n,l,g,k,h,j,q,x,r,z,c,s,yu,y,w",
	",",
)

type Args struct {
	Style     int  // 拼音风格
	Heteronym bool // 是否启用多音字模式
}

// 获取单个拼音中的声母
func initial(p string) string {
	s := ""
	for _, v := range _INITIALS {
		if strings.HasPrefix(p, v) {
			s = v
		}
	}
	return s
}

// 获取单个拼音中的韵母
func final(p string) string {
	i := initial(p)
	if i == "" {
		return p
	} else {
		return strings.Join(strings.SplitN(p, i, 1), "")
	}
}

func applyStyle(p []string, a Args) []string {
	newP := []string{}
	for _, v := range p {
		if a.Style == INITIALS {
			newP = append(newP, initial(v))
		} else {
			newP = append(newP, v)
		}
	}
	return newP
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
	return applyStyle(pys, a)
}

func Pinyin(s string, a Args) [][]string {
	hans := []rune(s)
	pys := [][]string{}
	for _, r := range hans {
		pys = append(pys, SinglePinyin(r, a))
	}
	return pys
}
