package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type cmdArgs struct {
	inputFile  string
	outputFile string
}

func genCode(inFile *os.File, outFile *os.File) {
	rd := bufio.NewReader(inFile)
	output := `package pinyin

// phraseDict is data map
//
// Generate from:
// https://github.com/hotoo/pinyin/blob/master/data/phrases-dict.js
//
// Warning: Auto-generated file, don't edit.
// If you want add more words, use phrase_dict_addition.go
var phraseDict = map[string]string{
`
	lines := []string{}

	for {
		line, err := rd.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		// Remove prefix space
		line = strings.TrimSpace(line)

		// `"后来居上": [["hòu"], ["lái"], ["jū"], ["shàng"]],` to `"后来居上": "hòu lái jū shàng",`
		if !strings.HasPrefix(line, `"`) {
			continue
		}

		line = strings.ReplaceAll(line, `[`, "")
		line = strings.ReplaceAll(line, `]`, "")
		line = strings.ReplaceAll(line, `", "`, " ")

		lines = append(lines, line)
	}

	output += strings.Join(lines, "\n")
	output += "\n}\n"
	outFile.WriteString(output)
	return
}

func parseCmdArgs() cmdArgs {
	flag.Parse()
	inputFile := flag.Arg(0)
	outputFile := flag.Arg(1)
	return cmdArgs{inputFile, outputFile}
}

func main() {
	args := parseCmdArgs()
	usage := "gen_phrase_dict INPUT OUTPUT"
	inputFile := args.inputFile
	outputFile := args.outputFile
	if inputFile == "" || outputFile == "" {
		fmt.Println(usage)
		os.Exit(1)
	}

	inFp, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("open file %s error", inputFile)
		panic(err)
	}
	outFp, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("open file %s error", outputFile)
		panic(err)
	}
	defer inFp.Close()
	defer outFp.Close()

	genCode(inFp, outFp)
}
