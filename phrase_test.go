package pinyin

import (
	"strings"
	"testing"
)

func Benchmark_pinyinPhrase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// about 0.015ms/op
		pinyinPhrase("人民银行旁边一行人abc字母路牌的平行宇宙")
	}
}

func Test_pinyinPhrase(t *testing.T) {
	expects := map[string]string{
		"人民银行旁边一行人abc字母路牌的平行宇宙发行股票":           "人民银行 pang bian yi xing 人abc字母路牌的 ping xing 宇宙 fa xing gu piao",
		"重庆市江北区重工业":                           "chong qing 市 jiang bei 区 zhong gong 业",
		"矿下的巷道与北京四合院的小巷有点相似":                  "矿下的 hang dao 与 bei jing si he yuan 的小巷 you dian xiang si",
		"小明在宿舍说了一宿有关星宿的常识":                    "小明在 su she 说了一宿 you guan xing xiu 的常识",
		"西藏的布达拉宫是收藏大藏经的宝藏":                    "xi zang 的 bu da la gong 是 shou cang da zang 经的 bao zang",
		"出现矛盾要先调查然后调解":                        "出现矛盾要先 diao cha 然后 tiao jie",
		"这两件瓷器模样很相似 像是由一个模型做出来的":              "这两件瓷器 mu yang 很 xiang si 像是由 yi ge mo xing 做出来的",
		"这条恶狗真可恶 满身臭味 让人闻了就恶心 让人厌恶 像恶魔让人做恶梦":  "这条 e gou 真 ke wu 满身 chou wei 让人闻了就 e xin 让人 yan wu 像 e mo 让人做 e meng",
		"薄荷油味不薄 薄薄的一片 很受欢迎 但要薄利多销":            "bo he 油味 bu bao bao bao 的 yi pian 很受欢迎 但要 bo li 多销",
		"单老师说 单于只会骑马 不会骑单车":                   "单老师说 chan yu 只会骑马 bu hui 骑 dan che",
		"这两批货物都打折出售 严重折本 他再也经不起这样折腾了 不如折断":    "这两批货物都 da zhe 出售 yan zhong she ben 他再也经不起 zhe yang zhe teng 了 bu ru zhe duan",
		"武松大喝一声：快拿酒来！我要喝十二碗。博得众食客一阵喝彩":        "武松大喝一声：快拿酒来！我要喝十二碗。 bo de 众 shi ke 一阵 he cai",
		"你这着真绝 让他干着急 又无法着手应付 心里老是悬着 着眼未来":     "你这着真绝 让他 gan zhao ji 又无法 zhuo shou ying fu 心里老是悬着 zhuo yan 未来",
		"一曝十寒的事情在校会上被曝光":                      "yi pu shi han 的事情在校 hui shang 被 bao guang",
		"省长李某如能早些省悟 自我反省、自省；就不致于丢官弃职 气得不省人事了": "省长李某如能早些 xing wu 自我 fan xing 、 zi xing ；就 bu zhi 于丢官弃职 气得 bu xing ren shi 了",
		"宽宏大度 一向度德量力 从不以己度人":                  "kuan hong da du 一向 duo de liang li cong bu yi ji duo ren",
		"他每次出差差不多都要出点差错":                      "他每次 chu chai cha bu duo 都要出点 cha cuo",
		"鱼不停的挣扎 鱼刺扎破了手 他随意包扎一下":               "鱼不停的 zheng zha yu ci 扎破了手 他随意 bao za 一下",
		"他自己懒散，却总是埋怨别人埋头工作":                   "他自己 lan san ，却总是 man yuan bie ren mai tou 工作",
		"是汉代有各种传说，是传记而不是唐代传奇 远距离传输":           "是汉代有 ge zhong chuan shuo ，是 zhuan ji 而 bu shi 唐代 chuan qi yuan ju li chuan shu",
		"奇怪的数学中奇数是最奇妙的":                       "qi guai 的 shu xue 中 ji shu 是最 qi miao 的",
		"交给团长 告诉他我们给予期望 前线的供给一定要能自给自足":        "jiao gei tuan zhang 告诉他我们 ji yu qi wang 前线的 gong ji 一定要能 zi ji 自足",
		"长安街边有条长桥，长辈和长头发的长老们":                 "chang an 街边有条 chang qiao ， zhang bei 和长 tou fa 的 zhang lao 们",
	}

	for source, expect := range expects {
		actual := splacesRegexp.ReplaceAllString(strings.TrimSpace(pinyinPhrase(source)), " ")
		if expect != actual {
			t.Errorf("\nexpect: %s\nactual: %s\n", expect, actual)
			break
		}
	}
}
