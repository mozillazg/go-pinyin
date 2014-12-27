package main

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
)

func main() {
	hans := "中国人"
	a := pinyin.NewArgs()
	fmt.Println("default: ", a)

	a.Style = pinyin.Normal
	fmt.Println("Normal: ", a)

	a.Style = pinyin.Tone
	fmt.Println("Tone: ", a)

	a.Style = pinyin.Tone2
	fmt.Println("Tone2: ", a)

	a.Style = pinyin.Initials
	fmt.Println("Initials: ", a)

	a.Style = pinyin.FirstLetter
	fmt.Println("FirstLetter: ", a)

	a.Style = pinyin.Finals
	fmt.Println("Finals: ", a)

	a.Style = pinyin.FinalsTone
	fmt.Println("FinalsTone: ", a)

	a.Style = pinyin.FinalsTone2
	fmt.Println("FinalsTone2: ", a)

	a = NewArgs()
	a.Heteronym = true
	fmt.Println("Heteronym true: ", a)

	a.Style = pinyin.Initials
	fmt.Println("Heteronym true Initials: ", a)

	a.Style = pinyin.Tone2
	fmt.Println("Heteronym true Tone2: ",
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
