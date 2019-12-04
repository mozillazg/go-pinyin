package pinyin

import (
	"regexp"
	"strings"
	"unicode"
)

var (
	splacesRegexp    = regexp.MustCompile("[ ]+")
	allowCharsRegexp = regexp.MustCompile("[a-zA-Z0-9\\.,\\?\\!;\\(\\)\\[\\]\\&\\=\\-_@ ]")
	hansMarks        = map[string]string{
		"？": "?",
		"！": "!",
		"：": ":",
		"。": ".",
		"，": ",",
		"；": ";",
		"（": "(",
		"）": ")",
		"【": "[",
		"】": "]",
	}
	paragraphOption = Args{
		Style:     NORMAL,
		Heteronym: true,
	}
)

// ParagraphPinyin convert a Chinese paragraph into pinyin, including letters, numbers, symbols
func ParagraphPinyin(p string) (s string) {
	p = pinyinPhrase(p)

	for _, r := range p {
		if unicode.Is(unicode.Han, r) {
			s += " " + string(Pinyin(string(r), paragraphOption)[0][0]) + " "
		} else {
			char := string(r)

			if allowCharsRegexp.MatchString(char) {
				s += char
			} else {
				if hansMarks[char] != "" {
					s += hansMarks[char]
				}
			}
		}
	}

	// 去连续两个空格
	s = splacesRegexp.ReplaceAllString(s, " ")
	// 去掉 , . ? 前面的空格
	s = strings.Replace(s, " ,", ",", -1)
	s = strings.Replace(s, " .", ".", -1)
	s = strings.Replace(s, " ?", "?", -1)
	s = strings.Replace(s, " ;", ";", -1)
	s = strings.Replace(s, " !", "!", -1)
	s = strings.Replace(s, "( ", "(", -1)
	s = strings.Replace(s, " )", ")", -1)
	s = strings.Replace(s, "[ ", "[", -1)
	s = strings.Replace(s, " ]", "]", -1)
	s = strings.Replace(s, " :", ":", -1)
	s = strings.TrimSpace(s)
	return
}
