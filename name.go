package pinyin

// surnamePinyin 只保存需要指定读音的姓氏，使用带声调拼音复用现有风格转换。
// 覆盖范围见 README 的姓名拼音说明。
// 此表仅供包内读取，不修改通用的 PinyinDict。
var surnamePinyin = map[string][]string{
	"单":  {"shàn"},
	"單":  {"shàn"},
	"解":  {"xiè"},
	"仇":  {"qiú"},
	"曾":  {"zēng"},
	"区":  {"ōu"},
	"區":  {"ōu"},
	"查":  {"zhā"},
	"朴":  {"piáo"},
	"乐":  {"yuè"},
	"樂":  {"yuè"},
	"缪":  {"miào"},
	"繆":  {"miào"},
	"种":  {"chóng"},
	"任":  {"rén"},
	"华":  {"huà"},
	"華":  {"huà"},
	"纪":  {"jǐ"},
	"紀":  {"jǐ"},
	"应":  {"yīng"},
	"應":  {"yīng"},
	"都":  {"dū"},
	"盛":  {"shèng"},
	"曲":  {"qū"},
	"燕":  {"yān"},
	"蔚":  {"yù"},
	"盖":  {"gě"},
	"蓋":  {"gě"},
	"繁":  {"pó"},
	"员":  {"yùn"},
	"員":  {"yùn"},
	"褚":  {"chǔ"},
	"卜":  {"bǔ"},
	"宿":  {"sù"},
	"阚":  {"kàn"},
	"闞":  {"kàn"},
	"乜":  {"niè"},
	"过":  {"guō"},
	"過":  {"guō"},
	"折":  {"shé"},
	"尉迟": {"yù", "chí"},
	"尉遲": {"yù", "chí"},
	"万俟": {"mò", "qí"},
	"单于": {"chán", "yú"},
	"單于": {"chán", "yú"},
	"长孙": {"zhǎng", "sūn"},
	"長孫": {"zhǎng", "sūn"},
	"皇甫": {"huáng", "fǔ"},
	"澹台": {"tán", "tái"},
	"澹臺": {"tán", "tái"},
	"鲜于": {"xiān", "yú"},
	"鮮于": {"xiān", "yú"},
	"乐正": {"yuè", "zhèng"},
	"樂正": {"yuè", "zhèng"},
	"拓跋": {"tuò", "bá"},
	"令狐": {"líng", "hú"},
	"子车": {"zǐ", "jū"},
	"子車": {"zǐ", "jū"},
}

// Name 将姓在前、名在后的单个中文姓名转换为拼音。
// 优先匹配内置双字复姓，再匹配单姓；匹配的姓氏只返回指定读音。
// 名字和未匹配的姓名沿用 Pinyin，包括 Style、Heteronym 和 Fallback 行为。
// 不去除空白，不推断名字中的多音字，也不保证覆盖所有姓氏或家族读音。
func Name(s string, a Args) [][]string {
	runes := []rune(s)
	for _, size := range []int{2, 1} {
		if len(runes) < size {
			continue
		}
		if surname, ok := surnamePinyin[string(runes[:size])]; ok {
			pys := make([][]string, 0, len(runes))
			for _, py := range surname {
				pys = append(pys, applyStyle([]string{py}, a))
			}
			return append(pys, Pinyin(string(runes[size:]), a)...)
		}
	}
	return Pinyin(s, a)
}
