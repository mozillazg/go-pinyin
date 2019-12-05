package pinyin

import (
	"testing"
)

func BenchmarkParagraph(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// about 0.06ms/op
		Paragraph("这条恶狗真可恶 满身臭味 让人闻了就恶心 让人厌恶 像恶魔让人做恶梦")
	}
}

func TestParagraph(t *testing.T) {
	expects := map[string]string{
		"天府大道北段18号高新国际广场A-3号!":      "tian fu da dao bei duan 18 hao gao xin guo ji guang chang A-3 hao!",
		"人民银行旁边一行人abc字母路牌，平行宇宙发行股票": "ren min yin xing pang bian yi xing ren abc zi mu lu pai, ping xing yu zhou fa xing gu piao",
		"我的大王！": "wo de dai wang!",
		"A abc&1234个人. 银行旁边,一行人?":           "A abc&1234 ge ren. yin xing pang bian, yi xing ren?",
		"A abc&1234个人，中国（银行）旁边；一“行”人？":      "A abc&1234 ge ren, zhong guo (yin xing) pang bian; yi xing ren?",
		"字符串长桥: aBc多音字&123":                 "zi fu chuan chang qiao aBc duo yin zi &123",
		"【腾讯(00700)拟2.305亿元出售应收账款 赚1720万元】": "[teng xun (00700) ni 2.305 yi yuan chu shou ying shou zhang kuan zhuan 1720 wan yuan]",
		"地址：重庆市江北区重工业？":                     "di zhi: chong qing shi jiang bei qu zhong gong ye?",
		"交给团长，告诉他我们给予期望。前线的供给一定要能自给自足！":     "jiao gei tuan zhang, gao su ta wo men ji yu qi wang. qian xian de gong ji yi ding yao neng zi ji zi zu!",
		"abc123":  "abc123",
		"義灬骉驫芔淼㴇": "yi biao biao biao hui miao she",
		"":        "",
	}

	for source, expect := range expects {
		actual := Paragraph(source)
		if expect != actual {
			t.Errorf("\nexpect: %s\nactual: %s\n", expect, actual)
			break
		}
	}

}
