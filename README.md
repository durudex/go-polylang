# `go-polylang`

Implementation of the [Polylang](https://github.com/polybase/polylang) in Go programming language.

## Setup

To get the [`go-polylang`](https://github.com/durudex/go-polylang) module, you need to have or install [Go version >= 1.19](https://go.dev/dl/). To check your current version of Go, use the `go version` command.

**The command to get the module:**

```bash
go get github.com/durudex/go-polylang@latest
```

## Parser

You can use the [`parser`](https://pkg.go.dev/github.com/durudex/go-polylang/parser) package to parse your code. It can be used follow:

### Parsing files

To parse a file or directory of files, you can use the [`parser.Parse()`](https://pkg.go.dev/github.com/durudex/go-polylang/parser#Parse) function by passing the required file or directory path to its arguments.

```go
import "github.com/durudex/go-polylang/parser"

func main() {
    ast, err := parser.Parse("filename.polylang")
    if err != nil { /* ... */ }
}
```

### Custom parser

Currently, we are using the [`participle`](github.com/alecthomas/participle) library for code parsing. However, you can create your own parser by configuring it to meet your specific needs.

```go
import (
    "github.com/durudex/go-polylang"
    "github.com/durudex/go-polylang/ast"

    "github.com/alecthomas/participle/v2"
)

func main() {
    parser := participle.MustBuild[ast.Program](
        participle.Lexer(polylang.Lexer),
    )

    // ...
}
```

> Note:
> If you want to use all the features of the library, you can use our ready-made variable [`Must`](https://pkg.go.dev/github.com/durudex/go-polylang/parser#Must), which contains all of the necessary settings for using the library.

## License

Copyright © 2022-2023 [Durudex](https://github.com/durudex). Released under the MIT license.
