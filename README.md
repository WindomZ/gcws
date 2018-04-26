# gcws

[![Build Status](https://travis-ci.org/WindomZ/gcws.svg?branch=master)](https://travis-ci.org/WindomZ/gcws)
[![Coverage Status](https://coveralls.io/repos/github/WindomZ/gcws/badge.svg?branch=master)](https://coveralls.io/github/WindomZ/gcws?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/WindomZ/gcws)](https://goreportcard.com/report/github.com/WindomZ/gcws)
[![GoDoc](https://godoc.org/github.com/WindomZ/gcws?status.svg)](https://godoc.org/github.com/WindomZ/gcws)

> gcws是golang版本的CWS(Chinese Word Segmentation) - 一个开源中文分词集成适配管理器

[English](README_en.md)

## 安装
```bash
go get github.com/WindomZ/gcws/...
```

## 支持
- [x] [sego](https://github.com/WindomZ/gcws/tree/master/sego) - Go中文分词，用双数组trie（Double-Array Trie）实现[[GitHub]](https://github.com/huichen/sego)
- [x] [jieba](https://github.com/WindomZ/gcws/tree/master/jieba) - "结巴"中文分词的Golang版本[[GitHub]](https://github.com/yanyiwu/gojieba)
- [x] [cwsharp](https://github.com/WindomZ/gcws/tree/master/cwsharp) - Golang中文分词库，支持多种分词模式，支持自定义字典和扩展[[GitHub]](https://github.com/zhengchun/cwsharp-go)
- [x] [segment](https://github.com/WindomZ/gcws/tree/master/segment) - golang 版中文分词包, inspired from 盘古分词[[GitHub]](https://github.com/WindomZ/gosegment)
- [x] [gse](https://github.com/WindomZ/gcws/tree/master/gse) - Go 语言高效分词, 支持英文、中文、日文等[[GitHub]](https://github.com/go-ego/gse)

## 用法
导入
```
import (
    "github.com/WindomZ/gcws"
)
```

初始化(以`jieba`为例)
```
import (
    _ "github.com/WindomZ/gcws/jieba"
)
...
cws, err := gcws.NewCWS("jieba")
```

简单易用
```
cws.Tokenize("喜欢就坚持，爱就别放弃") // 返回[]string{...}
```

## 模式
- ModeDefault - 默认分词模式
- ModeSearch - 搜索分词模式，`sego`, `jieba`, `segment`, `gse`支持
- ModeFast - 快速分词模式，`cwsharp`支持
- ModeEnglish - 英文分词模式，`sego`, `jieba`支持
