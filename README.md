# koron-go/jaconv

[![PkgGoDev](https://pkg.go.dev/badge/github.com/koron-go/jaconv)](https://pkg.go.dev/github.com/koron-go/jaconv)
[![Actions/Go](https://github.com/koron-go/jaconv/workflows/Go/badge.svg)](https://github.com/koron-go/jaconv/actions?query=workflow%3AGo)
[![Go Report Card](https://goreportcard.com/badge/github.com/koron-go/jaconv)](https://goreportcard.com/report/github.com/koron-go/jaconv)

Package jaconv provides a set of converters related to Japanese.

Sub package | Description                                               | Example
------------|-----------------------------------------------------------|----------------------------
roma2hira   | Romaji to hiragana converter                              | `konnnichiha` → `こんにちは`
hira2kata   | Hiragana to katakana conveter                             | `こんにちは` → `コンニチハ`
wide2narrow | Full-width (zenkaku) to half-width (hankaku) converter    | `コンニチハ` → `ｺﾝﾆﾁﾊ`
narrow2wide | Half-width to (hankaku) to full-width (zenkaku) converter | `ｺﾝﾆﾁﾊ` → `コンニチハ`

## Install and Update

```console
$ go get github.com/koron-go/jaconv@latest
```

## Example codes

### Example: roma2hira

```go
import "github.com/koron-go/jaconv/roma2hira"

println(roma2hira.Convert("konnnichiha"))
// Output:
// こんにちは
```

### Example: hira2roma

```go
import "github.com/koron-go/jaconv/hira2roma"

println(hira2roma.Convert("こんにちは"))
// Output:
// コンニチハ
```

### Example: narrow2wide

```go
import "github.com/koron-go/jaconv/narrow2wide"

println(narrow2wide.Convert("コンニチハ"))
// Output:
// ｺﾝﾁﾆﾊ
```

### Example: wide2narrow

```go
import "github.com/koron-go/jaconv/wide2narrow"

println(wide2narrow.Convert("ｺﾝﾆﾁﾊ"))
// Output:
// コンニチハ
```
