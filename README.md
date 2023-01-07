<h1 align="center">Go Polylang</h1>

<p align="center">
    Implementation of the <a href="https://github.com/polybase/polylang">Polylang</a> language on Go.
</p>

## Setup

```bash
go get github.com/durudex/go-polylang@latest
```

## Usage

1) To get started, you need to create a parser.

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

**Parse file:**

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

**Parse string:**

```go
import (
    "strings"

    ...
)

const code = " ... "

func main() {
    ...

    ast, err := parser.ParseString("", code)
}
```

## ⚠️ License

Copyright © 2022-2023 [Durudex](https://github.com/durudex). Released under the MIT license.
