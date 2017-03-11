package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/mattn/go-isatty"
	"github.com/mozillazg/go-pinyin"
)

func main() {
	heteronym := flag.Bool("e", false, "启用多音字模式")
	style := flag.String("s", "zh4ao", "指定拼音风格。可选值：zhao, zh4ao, zha4o, zhao4, zh, z, ao, 4ao, a4o, ao4")
	flag.Parse()
	hans := flag.Args()
	stdin := []byte{}
	if !isatty.IsTerminal(os.Stdin.Fd()) {
		stdin, _ = ioutil.ReadAll(os.Stdin)
	}
	if len(stdin) > 0 {
		hans = append(hans, string(stdin))
	}

	if len(hans) == 0 {
		fmt.Println("请至少输入一个汉字: pinyin [-e] [-s STYLE] HANS [HANS ...]")
		os.Exit(1)
	}

	args := pinyin.NewArgs()
	if *heteronym {
		args.Heteronym = true
	}
	switch *style {
	case "zhao":
		args.Style = pinyin.Normal
	case "zha4o":
		args.Style = pinyin.Tone2
	case "zhao4":
		args.Style = pinyin.Tone3
	case "zh":
		args.Style = pinyin.Initials
	case "z":
		args.Style = pinyin.FirstLetter
	case "ao":
		args.Style = pinyin.Finals
	case "4ao":
		args.Style = pinyin.FinalsTone
	case "a4o":
		args.Style = pinyin.FinalsTone2
	case "ao4":
		args.Style = pinyin.FinalsTone3
	default:
		args.Style = pinyin.Tone
	}

	pys := pinyin.Pinyin(strings.Join(hans, ""), args)
	for _, s := range pys {
		fmt.Print(strings.Join(s, ","), " ")
	}
	if len(pys) > 0 {
		fmt.Println()
	}
}
