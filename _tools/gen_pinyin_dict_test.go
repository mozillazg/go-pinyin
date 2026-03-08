package main

import (
	"strings"
	"testing"
)

func TestGenCode(t *testing.T) {
	input := `# Comment line
U+4E2D: zhōng,zhòng  # 中
U+56FD: guó  # 国
`
	expectedOutput := `package pinyin

// PinyinDict is data map
// Warning: Auto-generated file, don't edit.
var PinyinDict = map[int]string{
	0x4E2D: "zhōng,zhòng",
	0x56FD: "guó",
}
`
	r := strings.NewReader(input)
	var w strings.Builder
	genCode(r, &w)

	if w.String() != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, w.String())
	}
}

func TestGenCode_Empty(t *testing.T) {
	input := ""
	expectedOutput := `package pinyin

// PinyinDict is data map
// Warning: Auto-generated file, don't edit.
var PinyinDict = map[int]string{
}
`
	r := strings.NewReader(input)
	var w strings.Builder
	genCode(r, &w)

	if w.String() != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, w.String())
	}
}

func TestGenCode_OnlyComments(t *testing.T) {
	input := "# only comments\n# another comment"
	expectedOutput := `package pinyin

// PinyinDict is data map
// Warning: Auto-generated file, don't edit.
var PinyinDict = map[int]string{
}
`
	r := strings.NewReader(input)
	var w strings.Builder
	genCode(r, &w)

	if w.String() != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, w.String())
	}
}
