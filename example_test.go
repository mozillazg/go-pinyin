package pinyin_test

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
)

func ExamplePinyin_default() {
	hans := "中国人"
	a := pinyin.NewArgs()
	fmt.Println("default:", pinyin.Pinyin(hans, a))
	// Output: default: [[zhong] [guo] [ren]]
}

func ExamplePinyin_normal() {
	hans := "中国人"
	a := pinyin.NewArgs()
	a.Style = pinyin.Normal
	fmt.Println("Normal:", pinyin.Pinyin(hans, a))
	// Output: Normal: [[zhong] [guo] [ren]]
}

func ExamplePinyin_tone() {
	hans := "中国人"
	a := pinyin.NewArgs()
	a.Style = pinyin.Tone
	fmt.Println("Tone:", pinyin.Pinyin(hans, a))
	// Output: Tone: [[zhōng] [guó] [rén]]
}

func ExamplePinyin_tone2() {
	hans := "中国人"
	a := pinyin.NewArgs()
	a.Style = pinyin.Tone2
	fmt.Println("Tone2:", pinyin.Pinyin(hans, a))
	// Output: Tone2: [[zho1ng] [guo2] [re2n]]
}

func ExamplePinyin_initials() {
	hans := "中国人"
	a := pinyin.NewArgs()
	a.Style = pinyin.Initials
	fmt.Println("Initials:", pinyin.Pinyin(hans, a))
	// Output: Initials: [[zh] [g] [r]]
}

func ExamplePinyin_firstLetter() {
	hans := "中国人"
	a := pinyin.NewArgs()
	a.Style = pinyin.FirstLetter
	fmt.Println(pinyin.Pinyin(hans, a))
	// Output: [[z] [g] [r]]
}

func ExamplePinyin_finals() {
	hans := "中国人"
	a := pinyin.NewArgs()
	a.Style = pinyin.Finals
	fmt.Println(pinyin.Pinyin(hans, a))
	// Output: [[ong] [uo] [en]]
}

func ExamplePinyin_finalsTone() {
	hans := "中国人"
	a := pinyin.NewArgs()
	a.Style = pinyin.FinalsTone
	fmt.Println(pinyin.Pinyin(hans, a))
	// Output: [[ōng] [uó] [én]]
}

func ExamplePinyin_finalsTone2() {
	hans := "中国人"
	a := pinyin.NewArgs()
	a.Style = pinyin.FinalsTone2
	fmt.Println(pinyin.Pinyin(hans, a))
	// Output: [[o1ng] [uo2] [e2n]]
}

func ExamplePinyin_heteronym() {
	hans := "中国人"
	a := pinyin.NewArgs()
	a.Heteronym = true
	a.Style = pinyin.Tone2
	fmt.Println(pinyin.Pinyin(hans, a))
	// Output: [[zho1ng zho4ng] [guo2] [re2n]]
}

func ExampleLazyPinyin() {
	hans := "中国人"
	a := pinyin.NewArgs()
	fmt.Println(pinyin.LazyPinyin(hans, a))
	// Output: [zhong guo ren]
}

func ExampleSlug() {
	hans := "中国人"
	a := pinyin.NewArgs()
	fmt.Println(pinyin.Slug(hans, a))
	// Output: zhong-guo-ren
}
