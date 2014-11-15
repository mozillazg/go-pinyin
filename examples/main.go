package main

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
)

func main() {
	hans := "中国人"
	fmt.Println("default: ",
		pinyin.Pinyin(hans, pinyin.Args{}),
	)
	fmt.Println("NORMAL: ",
		pinyin.Pinyin(hans, pinyin.Args{Style: pinyin.NORMAL}),
	)
	fmt.Println("TONE: ",
		pinyin.Pinyin(hans, pinyin.Args{Style: pinyin.TONE}),
	)
	fmt.Println("TONE2: ",
		pinyin.Pinyin(hans, pinyin.Args{Style: pinyin.TONE2}),
	)
	fmt.Println("INITIALS: ",
		pinyin.Pinyin(hans, pinyin.Args{Style: pinyin.INITIALS}),
	)
	fmt.Println("FIRST_LETTER: ",
		pinyin.Pinyin(hans, pinyin.Args{Style: pinyin.FIRST_LETTER}),
	)
	fmt.Println("FINALS: ",
		pinyin.Pinyin(hans, pinyin.Args{Style: pinyin.FINALS}),
	)
	fmt.Println("FINALS_TONE: ",
		pinyin.Pinyin(hans, pinyin.Args{Style: pinyin.FINALS_TONE}),
	)
	fmt.Println("FINALS_TONE2: ",
		pinyin.Pinyin(hans, pinyin.Args{Style: pinyin.FINALS_TONE2}),
	)
	fmt.Println("Heteronym true: ",
		pinyin.Pinyin(hans, pinyin.Args{Heteronym: true}),
	)
	fmt.Println("Heteronym true INITIALS: ",
		pinyin.Pinyin(hans, pinyin.Args{Heteronym: true,
			Style: pinyin.INITIALS},
		),
	)
	fmt.Println("Heteronym true TONE2: ",
		pinyin.Pinyin(hans, pinyin.Args{Heteronym: true,
			Style: pinyin.TONE2},
		),
	)
	fmt.Println("LazyPinyin default: ",
		pinyin.LazyPinyin(hans, pinyin.Args{}),
	)
}
