package pinyin

import (
	"regexp"
	"strings"
	"unicode"
)

var (
	splacesRegexp    = regexp.MustCompile(`[\s]+`)
	allowCharsRegexp = regexp.MustCompile(`[a-zA-Z0-9\.,\?\!;\(\)\[\]\&\=\-_@\s]`)
	hansSymbols      = map[string]string{
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

// Paragraph convert a Chinese paragraph into pinyin, including letters, numbers, symbols
func Paragraph(p string) (s string) {
	p = pinyinPhrase(p)

	for _, r := range p {
		if unicode.Is(unicode.Han, r) {
			// Han chars
			result := Pinyin(string(r), paragraphOption)
			if len(result) == 0 {
				continue
			}
			if len(result[0]) == 0 {
				continue
			}

			s += " " + string(result[0][0]) + " "
		} else {
			// Other chars
			char := string(r)

			if allowCharsRegexp.MatchString(char) {
				s += char
			} else {
				if hansSymbols[char] != "" {
					s += hansSymbols[char]
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
