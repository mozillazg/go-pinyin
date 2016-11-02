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
	style := flag.String("s", "Tone", "指定拼音风格。可选值：Normal, Tone, Tone2, Tone3, Initials, FirstLetter, Finals, FinalsTone, FinalsTone2, FinalsTone3")
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
	case "Normal":
		args.Style = pinyin.Normal
	case "Tone2":
		args.Style = pinyin.Tone2
	case "Tone3":
		args.Style = pinyin.Tone3
	case "Initials":
		args.Style = pinyin.Initials
	case "FirstLetter":
		args.Style = pinyin.FirstLetter
	case "Finals":
		args.Style = pinyin.Finals
	case "FinalsTone":
		args.Style = pinyin.FinalsTone
	case "FinalsTone2":
		args.Style = pinyin.FinalsTone2
	case "FinalsTone3":
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
