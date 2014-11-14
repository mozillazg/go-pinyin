go-pinyin
=========

[![Build Status](https://travis-ci.org/mozillazg/go-pinyin.svg)](https://travis-ci.org/mozillazg/go-pinyin)
[![GoDoc](https://godoc.org/github.com/mozillazg/go-pinyin?status.svg)](https://godoc.org/github.com/mozillazg/go-pinyin)

汉语拼音转换工具 Go 版。


Installation
------------

```
go get -u github.com/mozillazg/go-pinyin
```

install CLI tool:

```
go get -u github.com/mozillazg/go-pinyin/pinyin
pinyin --help
```

Usage
------

```go
package main

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
)

func main() {
	hans := "中国人"
    // 默认输出 [[zhōng] [guó] [rén]]
	fmt.Println(pinyin.Pinyin(hans, pinyin.Args{}))
    // 开启多音字模式 [[zhōng zhòng] [guó] [rén]]
	fmt.Println(pinyin.Pinyin(hans, pinyin.Args{Heteronym: true}))
}
```
