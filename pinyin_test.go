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
		{
			Args{Style: Normal},
			[][]string{
				{"zhong"},
				{"guo"},
				{"ren"},
			},
		},
		// default
		{
			NewArgs(),
			[][]string{
				{"zhong"},
				{"guo"},
				{"ren"},
			},
		},
		// Normal
		{
			Args{Style: Normal},
			[][]string{
				{"zhong"},
				{"guo"},
				{"ren"},
			},
		},
		// Tone
		{
			Args{Style: Tone},
			[][]string{
				{"zhōng"},
				{"guó"},
				{"rén"},
			},
		},
		// Tone2
		{
			Args{Style: Tone2},
			[][]string{
				{"zho1ng"},
				{"guo2"},
				{"re2n"},
			},
		},
		// Tone3
		{
			Args{Style: Tone3},
			[][]string{
				{"zhong1"},
				{"guo2"},
				{"ren2"},
			},
		},
		// Initials
		{
			Args{Style: Initials},
			[][]string{
				{"zh"},
				{"g"},
				{"r"},
			},
		},
		// FirstLetter
		{
			Args{Style: FirstLetter},
			[][]string{
				{"z"},
				{"g"},
				{"r"},
			},
		},
		// Finals
		{
			Args{Style: Finals},
			[][]string{
				{"ong"},
				{"uo"},
				{"en"},
			},
		},
		// FinalsTone
		{
			Args{Style: FinalsTone},
			[][]string{
				{"ōng"},
				{"uó"},
				{"én"},
			},
		},
		// FinalsTone2
		{
			Args{Style: FinalsTone2},
			[][]string{
				{"o1ng"},
				{"uo2"},
				{"e2n"},
			},
		},
		// FinalsTone3
		{
			Args{Style: FinalsTone3},
			[][]string{
				{"ong1"},
				{"uo2"},
				{"en2"},
			},
		},
		// Heteronym
		{
			Args{Heteronym: true},
			[][]string{
				{"zhong", "zhong"},
				{"guo"},
				{"ren"},
			},
		},
	}

	testPinyin(t, hans, testData, Pinyin)

	// 测试不是多音字的 Heteronym
	hans = "你"
	testData = []testCase{
		{
			Args{},
			[][]string{
				{"ni"},
			},
		},
		{
			Args{Heteronym: true},
			[][]string{
				{"ni"},
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
		{
			NewArgs(),
			[][]string{
				{"zhong"},
				{"guo"},
				{"ren"},
			},
		},
		// custom
		{
			Args{
				Fallback: func(r rune, a Args) []string {
					return []string{"la"}
				},
			},
			[][]string{
				{"zhong"},
				{"guo"},
				{"ren"},
				{"la"},
				{"la"},
				{"la"},
			},
		},
		// custom
		{
			Args{
				Heteronym: true,
				Fallback: func(r rune, a Args) []string {
					return []string{"la", "wo"}
				},
			},
			[][]string{
				{"zhong", "zhong"},
				{"guo"},
				{"ren"},
				{"la", "wo"},
				{"la", "wo"},
				{"la", "wo"},
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
		{"鱼", Args{Style: Tone2}, [][]string{{"yu2"}}},
		{"鱼", Args{Style: Tone3}, [][]string{{"yu2"}}},
		{"鱼", Args{Style: Finals}, [][]string{{"v"}}},
		{"雨", Args{Style: Tone2}, [][]string{{"yu3"}}},
		{"雨", Args{Style: Tone3}, [][]string{{"yu3"}}},
		{"雨", Args{Style: Finals}, [][]string{{"v"}}},
		{"元", Args{Style: Tone2}, [][]string{{"yua2n"}}},
		{"元", Args{Style: Tone3}, [][]string{{"yuan2"}}},
		{"元", Args{Style: Finals}, [][]string{{"van"}}},
		// y, w 也不是拼音, yu的韵母是v, yi的韵母是i, wu的韵母是u
		{"呀", Args{Style: Initials}, [][]string{{""}}},
		{"呀", Args{Style: Tone2}, [][]string{{"ya"}}},
		{"呀", Args{Style: Tone3}, [][]string{{"ya"}}},
		{"呀", Args{Style: Finals}, [][]string{{"ia"}}},
		{"无", Args{Style: Initials}, [][]string{{""}}},
		{"无", Args{Style: Tone2}, [][]string{{"wu2"}}},
		{"无", Args{Style: Tone3}, [][]string{{"wu2"}}},
		{"无", Args{Style: Finals}, [][]string{{"u"}}},
		{"衣", Args{Style: Tone2}, [][]string{{"yi1"}}},
		{"衣", Args{Style: Tone3}, [][]string{{"yi1"}}},
		{"衣", Args{Style: Finals}, [][]string{{"i"}}},
		{"万", Args{Style: Tone2}, [][]string{{"wa4n"}}},
		{"万", Args{Style: Tone3}, [][]string{{"wan4"}}},
		{"万", Args{Style: Finals}, [][]string{{"uan"}}},
		// ju, qu, xu 的韵母应该是 v
		{"具", Args{Style: FinalsTone}, [][]string{{"ǜ"}}},
		{"具", Args{Style: FinalsTone2}, [][]string{{"v4"}}},
		{"具", Args{Style: FinalsTone3}, [][]string{{"v4"}}},
		{"具", Args{Style: Finals}, [][]string{{"v"}}},
		{"取", Args{Style: FinalsTone}, [][]string{{"ǚ"}}},
		{"取", Args{Style: FinalsTone2}, [][]string{{"v3"}}},
		{"取", Args{Style: FinalsTone3}, [][]string{{"v3"}}},
		{"取", Args{Style: Finals}, [][]string{{"v"}}},
		{"徐", Args{Style: FinalsTone}, [][]string{{"ǘ"}}},
		{"徐", Args{Style: FinalsTone2}, [][]string{{"v2"}}},
		{"徐", Args{Style: FinalsTone3}, [][]string{{"v2"}}},
		{"徐", Args{Style: Finals}, [][]string{{"v"}}},
		// # ń
		{"嗯", Args{Style: Normal}, [][]string{{"n"}}},
		{"嗯", Args{Style: Tone}, [][]string{{"ń"}}},
		{"嗯", Args{Style: Tone2}, [][]string{{"n2"}}},
		{"嗯", Args{Style: Tone3}, [][]string{{"n2"}}},
		{"嗯", Args{Style: Initials}, [][]string{{""}}},
		{"嗯", Args{Style: FirstLetter}, [][]string{{"n"}}},
		{"嗯", Args{Style: Finals}, [][]string{{"n"}}},
		{"嗯", Args{Style: FinalsTone}, [][]string{{"ń"}}},
		{"嗯", Args{Style: FinalsTone2}, [][]string{{"n2"}}},
		{"嗯", Args{Style: FinalsTone3}, [][]string{{"n2"}}},
		// # ḿ  \u1e3f  U+1E3F
		{"呣", Args{Style: Normal}, [][]string{{"m"}}},
		{"呣", Args{Style: Tone}, [][]string{{"ḿ"}}},
		{"呣", Args{Style: Tone2}, [][]string{{"m2"}}},
		{"呣", Args{Style: Tone3}, [][]string{{"m2"}}},
		{"呣", Args{Style: Initials}, [][]string{{""}}},
		{"呣", Args{Style: FirstLetter}, [][]string{{"m"}}},
		{"呣", Args{Style: Finals}, [][]string{{"m"}}},
		{"呣", Args{Style: FinalsTone}, [][]string{{"ḿ"}}},
		{"呣", Args{Style: FinalsTone2}, [][]string{{"m2"}}},
		{"呣", Args{Style: FinalsTone3}, [][]string{{"m2"}}},
		// 去除 0
		{"啊", Args{Style: Tone2}, [][]string{{"a"}}},
		{"啊", Args{Style: Tone3}, [][]string{{"a"}}},
		{"侵略", Args{Style: Tone2}, [][]string{{"qi1n"}, {"lve4"}}},
		{"侵略", Args{Style: FinalsTone2}, [][]string{{"i1n"}, {"ve4"}}},
		{"侵略", Args{Style: FinalsTone3}, [][]string{{"in1"}, {"ve4"}}},
	}
	testPinyinUpdate(t, testData, Pinyin)
}

func TestConvert(t *testing.T) {
	s := "中国人"
	v := Convert(s, nil)
	value := [][]string{{"zhong"}, {"guo"}, {"ren"}}
	if !reflect.DeepEqual(v, value) {
		t.Errorf("Expected %s, got %s", value, v)
	}

	a := NewArgs()
	v = Convert(s, &a)
	if !reflect.DeepEqual(v, value) {
		t.Errorf("Expected %s, got %s", value, v)
	}
}

func TestLazyConvert(t *testing.T) {
	s := "中国人"
	v := LazyConvert(s, nil)
	value := []string{"zhong", "guo", "ren"}
	if !reflect.DeepEqual(v, value) {
		t.Errorf("Expected %s, got %s", value, v)
	}

	a := NewArgs()
	v = LazyConvert(s, &a)
	if !reflect.DeepEqual(v, value) {
		t.Errorf("Expected %s, got %s", value, v)
	}
}

func TestPinyin_fallback_issue_35(t *testing.T) {
	a := NewArgs()
	a.Separator = ""
	a.Style = FirstLetter
	a.Fallback = func(r rune, a Args) []string {
		return []string{string(r)}
	}
	var s = "重。,a庆"
	v := Pinyin(s, a)
	expect := [][]string{{"z"}, {"。"}, {","}, {"a"}, {"q"}}
	if !reflect.DeepEqual(v, expect) {
		t.Errorf("Expected %s, got %s", expect, v)
	}
}
