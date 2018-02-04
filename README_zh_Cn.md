# gcws

[![Build Status](https://travis-ci.org/WindomZ/gcws.svg?branch=master)](https://travis-ci.org/WindomZ/gcws)
[![Coverage Status](https://coveralls.io/repos/github/WindomZ/gcws/badge.svg?branch=master)](https://coveralls.io/github/WindomZ/gcws?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/WindomZ/gcws)](https://goreportcard.com/report/github.com/WindomZ/gcws)
[![GoDoc](https://godoc.org/github.com/WindomZ/gcws?status.svg)](https://godoc.org/github.com/WindomZ/gcws)

> gcws是golang版本的CWS(Chinese Word Segmentation) - 一个开源中文分词适配管理器

[English](./README.md)

## 安装
```bash
go get github.com/WindomZ/gcws/...
```

## 支持
- [x] jieba
- [x] sego

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

简单使用
```
cws.Tokenize("喜欢就坚持，爱就别放弃")
```
