# Changelog


## 0.6.0 (2016-05-14)

* **NEW** 命令行程序支持指定拼音风格:

  ```shell
  $ pinyin -s Normal 你好
  ni hao
  ```
* **Bugfixed** 解决韵母 i, u, ü 的问题：根据以下拼音方案，还原出正确的韵母
   [#8](https://github.com/mozillazg/go-pinyin/pull/8),  [python-pinyin#26](https://github.com/mozillazg/python-pinyin/pull/26)

    > i 行的韵母，前面没有声母的时候，写成：yi（衣），yɑ（呀），ye（耶），
    > yɑo（腰），you（忧），yɑn（烟），yin（因），yɑnɡ（央），yinɡ（英），yonɡ（雍）。
    >
    > u 行的韵母，前面没有声母的时候，写成wu（乌），wɑ（蛙），wo（窝），
    > wɑi（歪），wei（威），wɑn（弯），wen（温），wɑnɡ（汪），wenɡ（翁）。
    >
    > ü行的韵母跟声母j，q，x拼的时候，写成ju（居），qu（区），xu（虚），
    > ü上两点也省略；但是跟声母l，n拼的时候，仍然写成lü（吕），nü（女）。

    **注意** `y` 既不是声母也不是韵母。详见 [汉语拼音方案](http://www.edu.cn/20011114/3009777.shtml)

* **Bugfixed** 解决未正确处理鼻音 ḿ, ń, ň, ǹ 的问题：包含鼻音的拼音不应该有声母



## 0.5.0 (2016-03-12)

* **CHANGE** 改为使用来自 [pinyin-data](https://github.com/mozillazg/pinyin-data) 的拼音数据。
* **NEW** 命令行程序支持从标准输入读取数据（支持管道和重定向输入）:

  ```shell
  $ echo "你好" | pinyin
  nǐ hǎo
  $ pinyin < hello.txt
  nǐ hǎo
  ```


## 0.4.0 (2016-01-29)

* **NEW** `Args` 结构体新增 field: `Fallback func(r rune, a Args) []string`
  用于处理没有拼音的字符（默认忽略没有拼音的字符）:
  ```go
  a := pinyin.NewArgs()
  a.Fallback = func(r rune, a pinyin.Args) []string {
      return []string{string(r + 1)}
  }
  fmt.Println(pinyin.Pinyin("中国人abc", a))
  // Output: [[zhong] [guo] [ren] [b] [c] [d]]

  // or
  pinyin.Fallback = func(r rune, a pinyin.Args) []string {
      return []string{string(r)}
  }
  fmt.Println(pinyin.Pinyin("中国人abc", pinyin.NewArgs()))
  // Output: [[zhong] [guo] [ren] [a] [b] [c]]
  ```


## 0.3.0 (2015-12-29)

* fix "当字符串中有非中文的时候，会出现下标越界的情况"(影响 `pinyin.LazyPinyin` 和 `pinyin.Slug` ([#1](https://github.com/mozillazg/go-pinyin/issues/1)))
* 调整对非中文字符的处理：当遇到没有拼音的字符时，直接忽略
  ```go
  // before
  fmt.Println(pinyin.Pinyin("中国人abc", pinyin.NewArgs()))
  [[zhong] [guo] [ren] [] [] []]

  // after
  fmt.Println(pinyin.Pinyin("中国人abc", pinyin.NewArgs()))
  [[zhong] [guo] [ren]]
  ```


## 0.2.1 (2015-08-26)

* `yu`, `y`, `w` 不是声母


## 0.2.0 (2015-01-04)

* 新增 `func NewArgs() Args`
* 解决 `Args.Separator` 无法赋值为 `""` 的 BUG
* 规范命名:
    * `NORMAL` -> `Normal`
    * `TONE` -> `Tone`
    * `TONE2` -> `Tone2`
    * `INITIALS` -> `Initials`
    * `FIRST_LETTER` -> `FirstLetter`
    * `FINALS` -> `Finals`
    * `FINALS_TONE` -> `FinalsTone`
    * `FINALS_TONE2` -> `FinalsTone2`

## 0.1.1 (2014-12-07)
* 更新拼音库


## 0.1.0 (2014-11-23)
* Initial Release
