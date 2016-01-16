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

func testPinyin(s string, d []testCase, f pinyinFunc) (t *testing.T) {
	for _, tc := range d {
		v := f(s, tc.args)
		if !reflect.DeepEqual(v, tc.result) {
			t.Errorf("Expected %s, got %s", tc.result, v)
		}
	}
	return t
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

	testPinyin(hans, testData, Pinyin)

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
	testPinyin(hans, testData, Pinyin)
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

// `yu`, `y`, `w` 不是声母
func TestNewInitials(t *testing.T) {
	hans := "鱼"
	testData := []testCase{
		testCase{
			Args{Style: Initials},
			[][]string{
				[]string{""},
			},
		},
		testCase{
			Args{Style: Finals},
			[][]string{
				[]string{"yu"},
			},
		},
	}
	testPinyin(hans, testData, Pinyin)

	hans = "五"
	testData = []testCase{
		testCase{
			Args{Style: Initials},
			[][]string{
				[]string{""},
			},
		},
		testCase{
			Args{Style: Finals},
			[][]string{
				[]string{"wu"},
			},
		},
	}
	testPinyin(hans, testData, Pinyin)
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
	testPinyin(hans, testData, Pinyin)
}
