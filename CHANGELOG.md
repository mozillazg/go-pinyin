# Changelog


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
