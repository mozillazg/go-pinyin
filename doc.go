// Package pinyin : 汉语拼音转换工具
//
// Usage
//
//	package main
//
// 	import (
// 		"fmt"
// 		"github.com/mozillazg/go-pinyin"
// 	)
//
// 	func main() {
// 		hans := "中国人"
// 		a := pinyin.NewArgs()
// 		// 默认输出 [[zhong] [guo] [ren]]
// 		fmt.Println(pinyin.Pinyin(hans, a))
//
// 		// 包含声调 [[zhōng] [guó] [rén]]
// 		a.Style = pinyin.Tone
// 		fmt.Println(pinyin.Pinyin(hans, a))
//
// 		// 声调用数字表示 [[zho1ng] [guo2] [re2n]]
// 		a.Style = pinyin.Tone2
// 		fmt.Println(pinyin.Pinyin(hans, a))
//
// 		// 开启多音字模式 [[zhong zhong] [guo] [ren]]
// 		a = pinyin.NewArgs()
// 		a.Heteronym = true
// 		fmt.Println(pinyin.Pinyin(hans, a))
// 		// [[zho1ng zho4ng] [guo2] [re2n]]
// 		a.Style = pinyin.Tone2
// 		fmt.Println(pinyin.Pinyin(hans, a))
// 	}
//
package pinyin
