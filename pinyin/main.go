package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)
import "github.com/mozillazg/go-pinyin"

func main() {
	heteronym := flag.Bool("e", false, "启用多音字模式")
	flag.Parse()
	hans := flag.Args()
	args := pinyin.NewArgs()
	args.Style = pinyin.TONE

	if len(hans) == 0 {
		fmt.Println("请至少输入一个汉字: pinyin hans [hans ...]")
		os.Exit(1)
	}
	if *heteronym {
		args.Heteronym = true
	}
	pys := pinyin.Pinyin(strings.Join(hans, ""), args)
	for _, s := range pys {
		fmt.Print(strings.Join(s, ","), " ")
	}
	if len(pys) > 0 {
		fmt.Println()
	}
}
