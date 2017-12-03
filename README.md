# base256
[![Travis](https://img.shields.io/travis/Equim-chan/base256.svg)](https://travis-ci.org/Equim-chan/base256)
[![Coverage Status](https://img.shields.io/codecov/c/gh/Equim-chan/base256.svg)](https://codecov.io/gh/Equim-chan/base256)
[![Release](https://img.shields.io/github/release/Equim-chan/base256.svg)](https://github.com/Equim-chan/base256/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/Equim-chan/base256)](https://goreportcard.com/report/github.com/Equim-chan/base256)
[![License](https://img.shields.io/badge/BSD-3-blue.svg)](https://github.com/Equim-chan/base256/blob/master/LICENSE)
[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg)](https://godoc.org/ekyu.moe/base256)

:cyclone: Binary to emoji encoding.

## Install
```bash
$ go get -u ekyu.moe/base256
# or better
$ dep ensure -add ekyu.moe/base256
```

## Example
```go
package main

import (
    "fmt"

    "ekyu.moe/base256"
)

func main() {
    fmt.Println(base256.EncodeToString([]byte("Hello, 世界")))   //=> 👾🍧🙆🍬🙇🌱😌🚟💦🏥🐴🏤👈
    fmt.Println(string(base256.DecodeString("👾🍧🙆🍬🙇🌱😌🚟💦🏥🐴🏤👈"))) //=> Hello, 世界
}
```

Check [godoc](https://godoc.org/ekyu.moe/base256) for further documents.

A CLI utility is also available with `go get ekyu.moe/base256/cmd/base256`.

## Specification
base256 encoding uses two tables to transform a single byte into a single emoji. The tables can be found in [table.go](https://github.com/Equim-chan/base256/blob/master/table.go). In this implementation, these tables are utilized circularly when encoding, however such rule does not apply to decoding.

To achieve the best compatibility, all the emojis are picked from the [classic version](http://classic.getemoji.com/) and each of them is 4 bytes long in UTF-8.

## License
[BSD-3-clause](https://github.com/Equim-chan/base256/blob/master/LICENSE)
