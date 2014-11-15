package pinyin

import (
	"regexp"
	"strings"
)

// Meta
const (
	Version   = "0.1.0"
	Author    = "mozillazg, 闲耘"
	License   = "MIT"
	Copyright = "Copyright (c) 2014 mozillazg, 闲耘"
)

// 拼音风格
const (
	NORMAL       = 0 // 普通风格，不带声调（默认风格）。如： pin yin
	TONE         = 1 // 声调风格1，拼音声调在韵母第一个字母上。如： pīn yīn
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

// 所有带声调的字符
var rePhoneticSymbolSource = func(m map[string]string) string {
	s := ""
	for k := range m {
		s = s + k
	}
	return s
}(phoneticSymbol)

// 匹配带声调字符的正则表达式
var re_PHONETIC_SYMBOL = regexp.MustCompile("[" + rePhoneticSymbolSource + "]")

// 匹配使用数字标识声调的字符的正则表达式
var re_TONE2 = regexp.MustCompile("([aeoiuvnm])([0-4])$")

type Args struct {
	Style     int    // 拼音风格（默认： NORMAL)
	Heteronym bool   // 是否启用多音字模式（默认：禁用）
	Separator string // Slug 中使用的分隔符（默认：-)
}

// 获取单个拼音中的声母
func initial(p string) string {
	s := ""
	for _, v := range _INITIALS {
		if strings.HasPrefix(p, v) {
			s = v
			break
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
		return strings.Join(strings.SplitN(p, i, 2), "")
	}
}

func toFixed(p string, a Args) string {
	if a.Style == INITIALS {
		return initial(p)
	}

	// 替换拼音中的带声调字符
	py := re_PHONETIC_SYMBOL.ReplaceAllStringFunc(p, func(m string) string {
		symbol, _ := phoneticSymbol[m]
		switch a.Style {
		// 不包含声调
		case NORMAL, FIRST_LETTER, FINALS:
			// 去掉声调: a1 -> a
			m = re_TONE2.ReplaceAllString(symbol, "$1")
		case TONE2, FINALS_TONE2:
			// 返回使用数字标识声调的字符
			m = symbol
		default:
			// 	// 声调在头上
		}
		return m
	})

	switch a.Style {
	// 首字母
	case FIRST_LETTER:
		py = string([]byte(py)[0])
	// 韵母
	case FINALS, FINALS_TONE, FINALS_TONE2:
		py = final(py)
	}
	return py
}

func applyStyle(p []string, a Args) []string {
	newP := []string{}
	for _, v := range p {
		newP = append(newP, toFixed(v, a))
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

func LazyPinyin(s string, a Args) []string {
	a.Heteronym = false
	pys := []string{}
	for _, v := range Pinyin(s, a) {
		pys = append(pys, v[0])
	}
	return pys
}

func Slug(s string, a Args) string {
	separator := "-"
	if a.Separator != "" {
		separator = a.Separator
	}
	return strings.Join(LazyPinyin(s, a), separator)
}
