# gcws

[![Build Status](https://travis-ci.org/WindomZ/gcws.svg?branch=master)](https://travis-ci.org/WindomZ/gcws)
[![Coverage Status](https://coveralls.io/repos/github/WindomZ/gcws/badge.svg?branch=master)](https://coveralls.io/github/WindomZ/gcws?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/WindomZ/gcws)](https://goreportcard.com/report/github.com/WindomZ/gcws)
[![GoDoc](https://godoc.org/github.com/WindomZ/gcws?status.svg)](https://godoc.org/github.com/WindomZ/gcws)

> gcws is CWS(Chinese Word Segmentation) for golang - some cws adapters manager.

The repo is inspired by `database/sql`.

[中文说明](README.md)

## Install
```bash
go get github.com/WindomZ/gcws/...
```

## Supported
- [x] [sego](https://github.com/huichen/sego)
- [x] [jieba](https://github.com/yanyiwu/gojieba)
- [x] [cwsharp](https://github.com/zhengchun/cwsharp-go)

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
