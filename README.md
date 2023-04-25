# `durudex/go-polylang`

Implementation of the [Polylang](https://github.com/polybase/polylang) in Go programming language.

## Setup

To get the [`go-polylang`](https://github.com/durudex/go-polylang) module, you need to have or install [Go version >= 1.19](https://go.dev/dl/). To check your current version of Go, use the `go version` command.

**The command to get the module:**

```bash
go get github.com/durudex/go-polylang@latest
```

## Parser

To start using the Polylang parser, you need to follow the steps below.

1) Create a new parser instance.

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

    ...
}
```

2) Now you can parse the written code and get the AST.

### With file

```go
import (
    "os"

    ...
)

func main() {
    ...

    f, err := os.Open(" ... ")
    if err != nil { ... }
    defer f.Close()

    ast, err := parser.Parse("", f)
}
```

### With string

```go
import ( ... )

const code = " ... "

func main() {
    ...

    ast, err := parser.ParseString("", code)
}
```

## License

Copyright © 2022-2023 [Durudex](https://github.com/durudex). Released under the MIT license.
