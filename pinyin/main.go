package main

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
)

func main() {
	hans := "中国人"
	fmt.Println(pinyin.Pinyin(hans, pinyin.Args{}))
	fmt.Println(pinyin.Pinyin(hans, pinyin.Args{Heteronym: true}))
}
