package pinyin

import "fmt"

func ExamplePinyin_default() {
	hans := "中国人"
	a := NewArgs()
	fmt.Println("default:", Pinyin(hans, a))
	// Output: default: [[zhong] [guo] [ren]]
}

func ExamplePinyin_normal() {
	hans := "中国人"
	a := NewArgs()
	a.Style = Normal
	fmt.Println("Normal:", Pinyin(hans, a))
	// Output: Normal: [[zhong] [guo] [ren]]
}

func ExamplePinyin_tone() {
	hans := "中国人"
	a := NewArgs()
	a.Style = Tone
	fmt.Println("Tone:", Pinyin(hans, a))
	// Output: Tone: [[zhōng] [guó] [rén]]
}

func ExamplePinyin_tone2() {
	hans := "中国人"
	a := NewArgs()
	a.Style = Tone2
	fmt.Println("Tone2:", Pinyin(hans, a))
	// Output: Tone2: [[zho1ng] [guo2] [re2n]]
}

func ExamplePinyin_initials() {
	hans := "中国人"
	a := NewArgs()
	a.Style = Initials
	fmt.Println("Initials:", Pinyin(hans, a))
	// Output: Initials: [[zh] [g] [r]]
}

func ExamplePinyin_firstLetter() {
	hans := "中国人"
	a := NewArgs()
	a.Style = FirstLetter
	fmt.Println("FirstLetter:", Pinyin(hans, a))
	// Output: FirstLetter: [[z] [g] [r]]
}

func ExamplePinyin_finals() {
	hans := "中国人"
	a := NewArgs()
	a.Style = Finals
	fmt.Println("Finals:", Pinyin(hans, a))
	// Output: Finals: [[ong] [uo] [en]]
}

func ExamplePinyin_finalsTone() {
	hans := "中国人"
	a := NewArgs()
	a.Style = FinalsTone
	fmt.Println("FinalsTone:", Pinyin(hans, a))
	// Output: FinalsTone: [[ōng] [uó] [én]]
}

func ExamplePinyin_finalsTone2() {
	hans := "中国人"
	a := NewArgs()
	a.Style = FinalsTone2
	fmt.Println("FinalsTone2:", Pinyin(hans, a))
	// Output: FinalsTone2: [[o1ng] [uo2] [e2n]]
}

func ExamplePinyin_heteronym() {
	hans := "中国人"
	a := NewArgs()
	a.Heteronym = true
	a.Style = Tone2
	fmt.Println("Heteronym true Tone2:", Pinyin(hans, a))
	// Output: Heteronym true Tone2: [[zho1ng zho4ng] [guo2] [re2n]]
}

func ExampleLazyPinyin() {
	hans := "中国人"
	a := NewArgs()
	fmt.Println("LazyPinyin default:", LazyPinyin(hans, a))
	// Output: LazyPinyin default: [zhong guo ren]
}

func ExampleSlug() {
	hans := "中国人"
	a := NewArgs()
	fmt.Println("Slug default:", Slug(hans, a))
	// Output: Slug default: zhong-guo-ren
}
