package pinyin

import (
	"reflect"
	"strings"
	"testing"
)

// TestNameExpandedSurnames 验证扩展姓氏及其繁体写法的指定读音，
// 并确认多音字配置及名字中的同字仍沿用通用转换行为。
func TestNameExpandedSurnames(t *testing.T) {
	tests := []struct {
		surnames string
		reading  string
	}{
		{"任", "rén"},
		{"华/華", "huà"},
		{"纪/紀", "jǐ"},
		{"应/應", "yīng"},
		{"都", "dū"},
		{"盛", "shèng"},
		{"曲", "qū"},
		{"燕", "yān"},
		{"蔚", "yù"},
		{"盖/蓋", "gě"},
		{"繁", "pó"},
		{"员/員", "yùn"},
		{"褚", "chǔ"},
		{"卜", "bǔ"},
		{"宿", "sù"},
		{"阚/闞", "kàn"},
		{"乜", "niè"},
		{"过/過", "guō"},
		{"皇甫", "huáng fǔ"},
		{"澹台/澹臺", "tán tái"},
		{"鲜于/鮮于", "xiān yú"},
		{"乐正/樂正", "yuè zhèng"},
		{"拓跋", "tuò bá"},
		{"令狐", "líng hú"},
		{"折", "shé"},
		{"子车/子車", "zǐ jū"},
	}
	for _, tt := range tests {
		for _, surname := range strings.Split(tt.surnames, "/") {
			t.Run(surname, func(t *testing.T) {
				for _, heteronym := range []bool{false, true} {
					a := Args{Style: Tone, Heteronym: heteronym}
					want := [][]string{}
					for _, syllable := range strings.Fields(tt.reading) {
						want = append(want, []string{syllable})
					}
					if got := Name(surname, a); !reflect.DeepEqual(got, want) {
						t.Errorf("surname, heteronym %v: got %v, want %v", heteronym, got, want)
					}
					want = append(want, Pinyin("乐", a)...)
					if got := Name(surname+"乐", a); !reflect.DeepEqual(got, want) {
						t.Errorf("full name, heteronym %v: got %v, want %v", heteronym, got, want)
					}
					if got, want := Name("李"+surname, a), Pinyin("李"+surname, a); !reflect.DeepEqual(got, want) {
						t.Errorf("given name changed: got %v, want %v", got, want)
					}
				}
			})
		}
	}
}

// TestNameDeferredSurnames 验证宁、贲及其繁体写法沿用通用字典读音。
func TestNameDeferredSurnames(t *testing.T) {
	for _, heteronym := range []bool{false, true} {
		a := Args{Style: Tone, Heteronym: heteronym}
		for _, surname := range []string{"宁", "寧", "贲", "賁"} {
			name := surname + "明"
			if got, want := Name(name, a), Pinyin(name, a); !reflect.DeepEqual(got, want) {
				t.Errorf("deferred %q, heteronym %v: got %v, want %v", surname, heteronym, got, want)
			}
		}
	}
}

// TestNameSurnames 验证单姓、复姓及仅含姓氏的输入使用指定读音，复姓优先匹配。
func TestNameSurnames(t *testing.T) {
	tests := []struct {
		name string
		want [][]string
	}{
		{"单明", [][]string{{"shàn"}, {"míng"}}},
		{"單明", [][]string{{"shàn"}, {"míng"}}},
		{"解明", [][]string{{"xiè"}, {"míng"}}},
		{"仇明", [][]string{{"qiú"}, {"míng"}}},
		{"曾明", [][]string{{"zēng"}, {"míng"}}},
		{"区明", [][]string{{"ōu"}, {"míng"}}},
		{"區明", [][]string{{"ōu"}, {"míng"}}},
		{"查明", [][]string{{"zhā"}, {"míng"}}},
		{"朴明", [][]string{{"piáo"}, {"míng"}}},
		{"乐明", [][]string{{"yuè"}, {"míng"}}},
		{"樂明", [][]string{{"yuè"}, {"míng"}}},
		{"缪明", [][]string{{"miào"}, {"míng"}}},
		{"繆明", [][]string{{"miào"}, {"míng"}}},
		{"种明", [][]string{{"chóng"}, {"míng"}}},
		{"尉迟明", [][]string{{"yù"}, {"chí"}, {"míng"}}},
		{"尉遲明", [][]string{{"yù"}, {"chí"}, {"míng"}}},
		{"万俟明", [][]string{{"mò"}, {"qí"}, {"míng"}}},
		{"单于明", [][]string{{"chán"}, {"yú"}, {"míng"}}},
		{"單于明", [][]string{{"chán"}, {"yú"}, {"míng"}}},
		{"长孙明", [][]string{{"zhǎng"}, {"sūn"}, {"míng"}}},
		{"長孫明", [][]string{{"zhǎng"}, {"sūn"}, {"míng"}}},
		{"单", [][]string{{"shàn"}}},
		{"单于", [][]string{{"chán"}, {"yú"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewArgs()
			a.Style = Tone
			if got := Name(tt.name, a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Name(%q) = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

// TestNameStyles 验证指定的姓氏读音支持全部拼音输出风格。
func TestNameStyles(t *testing.T) {
	tests := []struct {
		style int
		want  string
	}{
		{Normal, "shan"}, {Tone, "shàn"}, {Tone2, "sha4n"},
		{Tone3, "shan4"}, {Initials, "sh"}, {FirstLetter, "s"},
		{Finals, "an"}, {FinalsTone, "àn"}, {FinalsTone2, "a4n"},
		{FinalsTone3, "an4"},
	}
	for _, tt := range tests {
		got := Name("单", Args{Style: tt.style})
		want := [][]string{{tt.want}}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("style %d: got %v, want %v", tt.style, got, want)
		}
	}
}

// TestNameGivenNameAndFallback 验证名字和未匹配的输入沿用 Pinyin，
// 并确认自定义 Fallback 接收原始配置并保留非汉字字符。
func TestNameGivenNameAndFallback(t *testing.T) {
	for _, heteronym := range []bool{false, true} {
		a := NewArgs()
		a.Style = Tone
		a.Heteronym = heteronym
		for _, given := range []string{"乐", "单", "曾", "万俟", "", "🙂A"} {
			want := append([][]string{{"shàn"}}, Pinyin(given, a)...)
			if got := Name("单"+given, a); !reflect.DeepEqual(got, want) {
				t.Errorf("given %q, heteronym %v: got %v, want %v", given, heteronym, got, want)
			}
		}
		for _, name := range []string{"", "🙂", "王单", "李乐", "欧阳明", "A单", " 单"} {
			if got, want := Name(name, a), Pinyin(name, a); !reflect.DeepEqual(got, want) {
				t.Errorf("unmatched %q: got %v, want %v", name, got, want)
			}
		}
		// Fallback 校验传入的拼音配置，并将无法转换的字符原样返回。
		a.Fallback = func(r rune, args Args) []string {
			if args.Style != a.Style || args.Heteronym != a.Heteronym {
				t.Error("fallback received changed configuration")
			}
			return []string{string(r)}
		}
		want := [][]string{{"shàn"}, {"🙂"}, {"A"}}
		if got := Name("单🙂A", a); !reflect.DeepEqual(got, want) {
			t.Errorf("custom fallback: got %v, want %v", got, want)
		}
	}
}

// TestNameDoesNotChangePinyin 验证修改姓名转换结果不会影响后续调用或通用拼音转换。
func TestNameDoesNotChangePinyin(t *testing.T) {
	a := NewArgs()
	before := Pinyin("单于乐", a)
	got := Name("单明", a)
	got[0][0] = "changed"
	if want := [][]string{{"shan"}, {"ming"}}; !reflect.DeepEqual(Name("单明", a), want) {
		t.Error("mutating a result affected subsequent calls")
	}
	if after := Pinyin("单于乐", a); !reflect.DeepEqual(before, after) {
		t.Errorf("Pinyin changed from %v to %v", before, after)
	}
}
