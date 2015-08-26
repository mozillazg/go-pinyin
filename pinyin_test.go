package pinyin

import (
	"reflect"
	"testing"
)

type pinyinFunc func(string, Args) [][]string

func testPinyin(s string, d map[Args][][]string, f pinyinFunc) (t *testing.T) {
	for a := range d {
		value, _ := d[a]
		v := f(s, a)
		if !reflect.DeepEqual(v, value) {
			t.Errorf("Expected %s, got %s", value, v)
		}
	}
	return t
}

func TestPinyin(t *testing.T) {
	hans := "中国人"
	testData := map[Args][][]string{
		Args{}: [][]string{
			[]string{"zhong"},
			[]string{"guo"},
			[]string{"ren"},
		},
		Args{Style: Normal}: [][]string{
			[]string{"zhong"},
			[]string{"guo"},
			[]string{"ren"},
		},
		Args{Style: Tone}: [][]string{
			[]string{"zhōng"},
			[]string{"guó"},
			[]string{"rén"},
		},
		Args{Style: Tone2}: [][]string{
			[]string{"zho1ng"},
			[]string{"guo2"},
			[]string{"re2n"},
		},
		Args{Style: Initials}: [][]string{
			[]string{"zh"},
			[]string{"g"},
			[]string{"r"},
		},
		Args{Style: FirstLetter}: [][]string{
			[]string{"z"},
			[]string{"g"},
			[]string{"r"},
		},
		Args{Style: Finals}: [][]string{
			[]string{"ong"},
			[]string{"uo"},
			[]string{"en"},
		},
		Args{Style: FinalsTone}: [][]string{
			[]string{"ōng"},
			[]string{"uó"},
			[]string{"én"},
		},
		Args{Style: FinalsTone2}: [][]string{
			[]string{"o1ng"},
			[]string{"uo2"},
			[]string{"e2n"},
		},
		Args{Heteronym: true}: [][]string{
			[]string{"zhong", "zhong"},
			[]string{"guo"},
			[]string{"ren"},
		},
	}

	testPinyin(hans, testData, Pinyin)

	// 测试 Heteronym
	hans = "你"
	testData = map[Args][][]string{
		Args{}: [][]string{
			[]string{"ni"},
		},
		Args{Heteronym: true}: [][]string{
			[]string{"ni"},
		},
	}
	testPinyin(hans, testData, Pinyin)
}

func TestNoneHans(t *testing.T) {
	s := "abc"
	v := Pinyin(s, NewArgs())
	value := [][]string{[]string{}, []string{}, []string{}}
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
}

func TestFinal(t *testing.T) {
	value := "an"
	v := final("an")
	if v != value {
		t.Errorf("Expected %s, got %s", value, v)
	}
}

func TestNewInitials(t *testing.T) {
	hans := "鱼"
	testData := map[Args][][]string{
		Args{Style: Initials}: [][]string{
			[]string{""},
		},
		Args{Style: Finals}: [][]string{
			[]string{"yu"},
		},
	}
	testPinyin(hans, testData, Pinyin)

	hans = "五"
	testData = map[Args][][]string{
		Args{Style: Initials}: [][]string{
			[]string{""},
		},
		Args{Style: Finals}: [][]string{
			[]string{"wu"},
		},
	}
	testPinyin(hans, testData, Pinyin)
}
