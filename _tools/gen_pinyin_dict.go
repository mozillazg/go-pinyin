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

// PinyinDict is data map
// Warning: Auto-generated file, don't edit.
var PinyinDict = map[int]string{
`
	lines := []string{}

	for {
		line, err := rd.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		if strings.HasPrefix(line, "#") {
			continue
		}

		// line: `U+4E2D: zhōng,zhòng  # 中`
		dataSlice := strings.Split(line, "  #")
		dataSlice = strings.Split(dataSlice[0], ": ")
		// 0x4E2D
		hexCode := strings.Replace(dataSlice[0], "U+", "0x", 1)
		// zhōng,zhòng
		pinyin := dataSlice[1]
		lines = append(lines, fmt.Sprintf("\t%s: \"%s\",", hexCode, pinyin))
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
	usage := "gen_pinyin_dict INPUT OUTPUT"
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
