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
		fmt.Fprintln(os.Stderr, "请至少输入一个汉字: pinyin [-e] [-s STYLE] HANS [HANS ...]")
		os.Exit(1)
	}

	args := pinyin.NewArgs()
	if *heteronym {
		args.Heteronym = true
	}

	styleValues := map[string]int{
		"zhao":  pinyin.Normal,
		"zh4ao": pinyin.Tone,
		"zha4o": pinyin.Tone2,
		"zhao4": pinyin.Tone3,
		"zh":    pinyin.Initials,
		"z":     pinyin.FirstLetter,
		"ao":    pinyin.Finals,
		"4ao":   pinyin.FinalsTone,
		"a4o":   pinyin.FinalsTone2,
		"ao4":   pinyin.FinalsTone3,
	}
	if value, ok := styleValues[*style]; !ok {
		fmt.Fprintf(os.Stderr, "无效的拼音风格：%s\n", *style)
		os.Exit(1)
	} else {
		args.Style = value
	}

	pys := pinyin.Pinyin(strings.Join(hans, ""), args)
	for _, s := range pys {
		fmt.Print(strings.Join(s, ","), " ")
	}
	if len(pys) > 0 {
		fmt.Println()
	}
}
