# gcws

[![Build Status](https://travis-ci.org/WindomZ/gcws.svg?branch=master)](https://travis-ci.org/WindomZ/gcws)
[![Coverage Status](https://coveralls.io/repos/github/WindomZ/gcws/badge.svg?branch=master)](https://coveralls.io/github/WindomZ/gcws?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/WindomZ/gcws)](https://goreportcard.com/report/github.com/WindomZ/gcws)
[![GoDoc](https://godoc.org/github.com/WindomZ/gcws?status.svg)](https://godoc.org/github.com/WindomZ/gcws)

> gcws is CWS(Chinese Word Segmentation) for golang - many cws adapters manager.

The repo is inspired by `database/sql`.

[中文说明](README.md)

## Install
```bash
go get github.com/WindomZ/gcws/...
```

## Supported
- [x] [sego](https://github.com/WindomZ/gcws/tree/master/sego) - Go中文分词，用双数组trie（Double-Array Trie）实现[[GitHub]](https://github.com/huichen/sego)
- [x] [jieba](https://github.com/WindomZ/gcws/tree/master/jieba) - "结巴"中文分词的Golang版本[[GitHub]](https://github.com/yanyiwu/gojieba)
- [x] [cwsharp](https://github.com/WindomZ/gcws/tree/master/cwsharp) - Golang中文分词库，支持多种分词模式，支持自定义字典和扩展[[GitHub]](https://github.com/zhengchun/cwsharp-go)
- [x] [segment](https://github.com/WindomZ/gcws/tree/master/segment) - golang 版中文分词包, inspired from 盘古分词[[GitHub]](https://github.com/WindomZ/gosegment)
- [x] [gse](https://github.com/WindomZ/gcws/tree/master/gse) - Go efficient text segmentation; support english, chinese, japanese and other.[[GitHub]](https://github.com/go-ego/gse)

## Usage
Import it
```
import (
    "github.com/WindomZ/gcws"
)
```

Init it (example with `jieba`)
```
import (
    _ "github.com/WindomZ/gcws/jieba"
)
...
cws, err := gcws.NewCWS("jieba")
```

Use it
```
cws.Tokenize("For man is man and master of his fate.") // return []string{...}
```

## Mode
- ModeDefault - default mode
- ModeSearch - search optimization, support `sego`, `jieba`, `segment`, `gse`
- ModeFast - run fast, support `cwsharp`
- ModeEnglish - optimization for English, support`sego`, `jieba`
