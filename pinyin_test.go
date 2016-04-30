package pinyin

import (
	"reflect"
	"testing"
)

type pinyinFunc func(string, Args) [][]string
type testCase struct {
	args   Args
	result [][]string
}

func testPinyin(t *testing.T, s string, d []testCase, f pinyinFunc) {
	for _, tc := range d {
		v := f(s, tc.args)
		if !reflect.DeepEqual(v, tc.result) {
			t.Errorf("Expected %s, got %s", tc.result, v)
		}
	}
}

func TestPinyin(t *testing.T) {
	hans := "中国人"
	testData := []testCase{
		// default
		testCase{
			Args{Style: Normal},
			[][]string{
				[]string{"zhong"},
				[]string{"guo"},
				[]string{"ren"},
			},
		},
		// default
		testCase{
			NewArgs(),
			[][]string{
				[]string{"zhong"},
				[]string{"guo"},
				[]string{"ren"},
			},
		},
		// Normal
		testCase{
			Args{Style: Normal},
			[][]string{
				[]string{"zhong"},
				[]string{"guo"},
				[]string{"ren"},
			},
		},
		// Tone
		testCase{
			Args{Style: Tone},
			[][]string{
				[]string{"zhōng"},
				[]string{"guó"},
				[]string{"rén"},
			},
		},
		// Tone2
		testCase{
			Args{Style: Tone2},
			[][]string{
				[]string{"zho1ng"},
				[]string{"guo2"},
				[]string{"re2n"},
			},
		},
		// Initials
		testCase{
			Args{Style: Initials},
			[][]string{
				[]string{"zh"},
				[]string{"g"},
				[]string{"r"},
			},
		},
		// FirstLetter
		testCase{
			Args{Style: FirstLetter},
			[][]string{
				[]string{"z"},
				[]string{"g"},
				[]string{"r"},
			},
		},
		// Finals
		testCase{
			Args{Style: Finals},
			[][]string{
				[]string{"ong"},
				[]string{"uo"},
				[]string{"en"},
			},
		},
		// FinalsTone
		testCase{
			Args{Style: FinalsTone},
			[][]string{
				[]string{"ōng"},
				[]string{"uó"},
				[]string{"én"},
			},
		},
		// FinalsTone2
		testCase{
			Args{Style: FinalsTone2},
			[][]string{
				[]string{"o1ng"},
				[]string{"uo2"},
				[]string{"e2n"},
			},
		},
		// Heteronym
		testCase{
			Args{Heteronym: true},
			[][]string{
				[]string{"zhong", "zhong"},
				[]string{"guo"},
				[]string{"ren"},
			},
		},
	}

	testPinyin(t, hans, testData, Pinyin)

	// 测试不是多音字的 Heteronym
	hans = "你"
	testData = []testCase{
		testCase{
			Args{},
			[][]string{
				[]string{"ni"},
			},
		},
		testCase{
			Args{Heteronym: true},
			[][]string{
				[]string{"ni"},
			},
		},
	}
	testPinyin(t, hans, testData, Pinyin)
}

func TestNoneHans(t *testing.T) {
	s := "abc"
	v := Pinyin(s, NewArgs())
	value := [][]string{}
	if !reflect.DeepEqual(v, value) {
		t.Errorf("Expected %s, got %s", value, v)
	}
}

func TestLazyPinyin(t *testing.T) {
	s := "中国人"
	v := LazyPinyin(s, Args{})
	value := []string{"zhong", "guo", "ren"}
	if !reflect.DeepEqual(v, value) {
		t.Errorf("Expected %s, got %s", value, v)
	}

	s = "中国人abc"
	v = LazyPinyin(s, Args{})
	value = []string{"zhong", "guo", "ren"}
	if !reflect.DeepEqual(v, value) {
		t.Errorf("Expected %s, got %s", value, v)
	}
}

func TestSlug(t *testing.T) {
	s := "中国人"
	v := Slug(s, Args{})
	value := "zhongguoren"
	if v != value {
		t.Errorf("Expected %s, got %s", value, v)
	}

	v = Slug(s, Args{Separator: ","})
	value = "zhong,guo,ren"
	if v != value {
		t.Errorf("Expected %s, got %s", value, v)
	}

	a := NewArgs()
	v = Slug(s, a)
	value = "zhong-guo-ren"
	if v != value {
		t.Errorf("Expected %s, got %s", value, v)
	}

	s = "中国人abc，,中"
	v = Slug(s, a)
	value = "zhong-guo-ren-zhong"
	if v != value {
		t.Errorf("Expected %s, got %s", value, v)
	}
}

func TestFinal(t *testing.T) {
	value := "an"
	v := final("an")
	if v != value {
		t.Errorf("Expected %s, got %s", value, v)
	}
}

func TestFallback(t *testing.T) {
	hans := "中国人abc"
	testData := []testCase{
		// default
		testCase{
			NewArgs(),
			[][]string{
				[]string{"zhong"},
				[]string{"guo"},
				[]string{"ren"},
			},
		},
		// custom
		testCase{
			Args{
				Fallback: func(r rune, a Args) []string {
					return []string{"la"}
				},
			},
			[][]string{
				[]string{"zhong"},
				[]string{"guo"},
				[]string{"ren"},
				[]string{"la"},
				[]string{"la"},
				[]string{"la"},
			},
		},
		// custom
		testCase{
			Args{
				Heteronym: true,
				Fallback: func(r rune, a Args) []string {
					return []string{"la", "wo"}
				},
			},
			[][]string{
				[]string{"zhong", "zhong"},
				[]string{"guo"},
				[]string{"ren"},
				[]string{"la", "wo"},
				[]string{"la", "wo"},
				[]string{"la", "wo"},
			},
		},
	}
	testPinyin(t, hans, testData, Pinyin)
}

type testItem struct {
	hans   string
	args   Args
	result [][]string
}

func testPinyinUpdate(t *testing.T, d []testItem, f pinyinFunc) {
	for _, tc := range d {
		v := f(tc.hans, tc.args)
		if !reflect.DeepEqual(v, tc.result) {
			t.Errorf("Expected %s, got %s", tc.result, v)
		}
	}
}

func TestUpdated(t *testing.T) {
	testData := []testItem{
		// 误把 yu 放到声母列表了
		testItem{"鱼", Args{Style: Tone2}, [][]string{[]string{"yu2"}}},
		testItem{"鱼", Args{Style: Finals}, [][]string{[]string{"v"}}},
		testItem{"雨", Args{Style: Tone2}, [][]string{[]string{"yu3"}}},
		testItem{"雨", Args{Style: Finals}, [][]string{[]string{"v"}}},
		testItem{"元", Args{Style: Tone2}, [][]string{[]string{"yua2n"}}},
		testItem{"元", Args{Style: Finals}, [][]string{[]string{"van"}}},
		// y, w 也不是拼音, yu的韵母是v, yi的韵母是i, wu的韵母是u
		testItem{"呀", Args{Style: Initials}, [][]string{[]string{""}}},
		testItem{"呀", Args{Style: Tone2}, [][]string{[]string{"ya"}}},
		testItem{"呀", Args{Style: Finals}, [][]string{[]string{"ia"}}},
		testItem{"无", Args{Style: Initials}, [][]string{[]string{""}}},
		testItem{"无", Args{Style: Tone2}, [][]string{[]string{"wu2"}}},
		testItem{"无", Args{Style: Finals}, [][]string{[]string{"u"}}},
		testItem{"衣", Args{Style: Tone2}, [][]string{[]string{"yi1"}}},
		testItem{"衣", Args{Style: Finals}, [][]string{[]string{"i"}}},
		testItem{"万", Args{Style: Tone2}, [][]string{[]string{"wa4n"}}},
		testItem{"万", Args{Style: Finals}, [][]string{[]string{"uan"}}},
		// ju, qu, xu 的韵母应该是 v
		testItem{"具", Args{Style: FinalsTone}, [][]string{[]string{"ǜ"}}},
		testItem{"具", Args{Style: FinalsTone2}, [][]string{[]string{"v4"}}},
		testItem{"具", Args{Style: Finals}, [][]string{[]string{"v"}}},
		testItem{"取", Args{Style: FinalsTone}, [][]string{[]string{"ǚ"}}},
		testItem{"取", Args{Style: FinalsTone2}, [][]string{[]string{"v3"}}},
		testItem{"取", Args{Style: Finals}, [][]string{[]string{"v"}}},
		testItem{"徐", Args{Style: FinalsTone}, [][]string{[]string{"ǘ"}}},
		testItem{"徐", Args{Style: FinalsTone2}, [][]string{[]string{"v2"}}},
		testItem{"徐", Args{Style: Finals}, [][]string{[]string{"v"}}},
	}
	testPinyinUpdate(t, testData, Pinyin)
}
