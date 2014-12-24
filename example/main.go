package main

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
)

func main() {
	hans := "中国人"
	a := pinyin.NewArgs()
	fmt.Println("default: ", a)

	a.Style = pinyin.NORMAL
	fmt.Println("NORMAL: ", a)

	a.Style = pinyin.TONE
	fmt.Println("TONE: ", a)

	a.Style = pinyin.TONE2
	fmt.Println("TONE2: ", a)

	a.Style = pinyin.INITIALS
	fmt.Println("INITIALS: ", a)

	a.Style = pinyin.FIRST_LETTER
	fmt.Println("FIRST_LETTER: ", a)

	a.Style = pinyin.FINALS
	fmt.Println("FINALS: ", a)

	a.Style = pinyin.FINALS_TONE
	fmt.Println("FINALS_TONE: ", a)

	a.Style = pinyin.FINALS_TONE2
	fmt.Println("FINALS_TONE2: ", a)

	a = NewArgs()
	a.Heteronym = true
	fmt.Println("Heteronym true: ", a)

	a.Style = pinyin.INITIALS
	fmt.Println("Heteronym true INITIALS: ", a)

	a.Style = pinyin.TONE2
	fmt.Println("Heteronym true TONE2: ",
		pinyin.Pinyin(hans, a),
	)

	a = NewArgs()
	fmt.Println("LazyPinyin default: ",
		pinyin.LazyPinyin(hans, a),
	)
	fmt.Println("Slug default: ",
		pinyin.Slug(hans, a),
	)
}
