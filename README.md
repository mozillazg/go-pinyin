go-pinyin
=========

[![Build Status](https://github.com/mozillazg/go-pinyin/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/mozillazg/go-pinyin/actions/workflows/ci.yml)
[![Coverage Status](https://coveralls.io/repos/mozillazg/go-pinyin/badge.svg?branch=master)](https://coveralls.io/r/mozillazg/go-pinyin?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/mozillazg/go-pinyin)](https://goreportcard.com/report/github.com/mozillazg/go-pinyin)
[![GoDoc](https://godoc.org/github.com/mozillazg/go-pinyin?status.svg)](https://godoc.org/github.com/mozillazg/go-pinyin)

汉语拼音转换工具 Go 版。


Installation
------------

```
go get github.com/mozillazg/go-pinyin
```

install CLI tool:

```
# go version>=1.17
go install github.com/mozillazg/go-pinyin/cli/pinyin@latest

# go version<1.17
go get -u github.com/mozillazg/go-pinyin/cli/pinyin

$ pinyin 中国人
zhōng guó rén
```


Documentation
--------------

API documentation can be found here:
https://godoc.org/github.com/mozillazg/go-pinyin


Usage
------

### 姓名拼音

`Name(s string, a Args) [][]string` 为姓在前、名在后的单个中文姓名选择姓氏读音：

```go
a := pinyin.NewArgs()
pinyin.Name("单明", a)   // [[shan] [ming]]
pinyin.Name("万俟明", a) // [[mo] [qi] [ming]]
a.Style = pinyin.Tone
pinyin.Name("单明", a)   // [[shàn] [míng]]
```

内置 29 个单姓、11 个复姓规则（简繁体写法合并计数）：

| 类型 | 姓氏 |
| --- | --- |
| 单姓 | 单（單）、解、仇、曾、区（區）、查、朴、乐（樂）、缪（繆）、种、任、华（華）、纪（紀）、应（應）、都、盛、曲、燕、蔚、盖（蓋）、繁、员（員）、褚、卜、宿、阚（闞）、乜、过（過）、折 |
| 复姓 | 尉迟（尉遲）、万俟、单于（單于）、长孙（長孫）、皇甫、澹台（澹臺）、鲜于（鮮于）、乐正（樂正）、拓跋、令狐、子车（子車） |

复姓优先于单姓。
宁（寧）、贲（賁）不添加姓氏选音规则，继续沿用通用字典；这不代表字典的首个读音适合具体本人。
折、子车（子車）分别采用默认读音 `shé`、`zǐ jū`，不代表覆盖所有家族异读。

匹配的姓氏只返回规则指定的读音；名字部分及未匹配的姓名沿用 `Pinyin`，
支持全部 `Style`、`Heteronym` 和 `Fallback` 配置。`Heteronym = true` 时，
名字仍返回字典中的候选读音。此函数不推断本人名字的实际读音，也不覆盖全部姓氏、
地域或家族读音（例如乐姓采用 `yuè`）。输入不会自动去除空白，前缀匹配无法区分
复姓与恰好同字的单姓加名字。普通文本请继续使用 `Pinyin`。

### 通用拼音

```go
package main

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
)

func main() {
	hans := "中国人"

	// 默认
	a := pinyin.NewArgs()
	fmt.Println(pinyin.Pinyin(hans, a))
	// [[zhong] [guo] [ren]]

	// 包含声调
	a.Style = pinyin.Tone
	fmt.Println(pinyin.Pinyin(hans, a))
	// [[zhōng] [guó] [rén]]

	// 声调用数字表示
	a.Style = pinyin.Tone2
	fmt.Println(pinyin.Pinyin(hans, a))
	// [[zho1ng] [guo2] [re2n]]

	// 开启多音字模式
	a = pinyin.NewArgs()
	a.Heteronym = true
	fmt.Println(pinyin.Pinyin(hans, a))
	// [[zhong] [guo] [ren]]
	a.Style = pinyin.Tone2
	fmt.Println(pinyin.Pinyin(hans, a))
	// [[zho1ng zho4ng] [guo2] [re2n]]

	fmt.Println(pinyin.LazyPinyin(hans, pinyin.NewArgs()))
	// [zhong guo ren]

	fmt.Println(pinyin.Convert(hans, nil))
	// [[zhong] [guo] [ren]]

	fmt.Println(pinyin.LazyConvert(hans, nil))
	// [zhong guo ren]
}
```

注意：

* 默认情况下会忽略没有拼音的字符（可以通过自定义 `Fallback` 参数的值来自定义如何处理没有拼音的字符，
  详见 [示例](https://godoc.org/github.com/mozillazg/go-pinyin#example-Pinyin--FallbackCustom1)）。
* 根据 [《汉语拼音方案》](http://www.moe.gov.cn/s78/A19/yxs_left/moe_810/s230/195802/t19580201_186000.html) y，w，ü (yu) 都不是声母，
  以及不是所有拼音都有声母，如果这不是你预期的话，你可能需要的是首字母风格 `FirstLetter`
  （ [详细信息](https://github.com/mozillazg/python-pinyin#%E4%B8%BA%E4%BB%80%E4%B9%88%E6%B2%A1%E6%9C%89-y-w-yu-%E5%87%A0%E4%B8%AA%E5%A3%B0%E6%AF%8D) ）。


Related Projects
-----------------

* [hotoo/pinyin](https://github.com/hotoo/pinyin): 汉语拼音转换工具 Node.js/JavaScript 版。
* [mozillazg/python-pinyin](https://github.com/mozillazg/python-pinyin): 汉语拼音转换工具 Python 版。
* [mozillazg/rust-pinyin](https://github.com/mozillazg/rust-pinyin): 汉语拼音转换工具 Rust 版。


pinyin data
-----------------

* 使用 [pinyin-data](https://github.com/mozillazg/pinyin-data) 的拼音数据


License
---------

Under the MIT License.
