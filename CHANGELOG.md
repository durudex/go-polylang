# Change Log

## [v0.0.3] - 2023-05-31

### Added

- Added [`metadata`](https://pkg.go.dev/github.com/durudex/go-polylang/metadata) package.
- Added [`parser`](https://pkg.go.dev/github.com/durudex/go-polylang/parser) package.

### Fixed

- Fixed a lexer multi-line block comment bug.
- Fixed a decorator parsing bug when selecting an item.

## [v0.0.2] - 2023-01-25

### Added

- Added AST [Decorator](https://pkg.go.dev/github.com/durudex/go-polylang/ast#Decorator) and [Value](https://pkg.go.dev/github.com/durudex/go-polylang/ast#Value).
- Added AST [Bytes](https://pkg.go.dev/github.com/durudex/go-polylang/ast#Bytes) and [PublicKey](https://pkg.go.dev/github.com/durudex/go-polylang/ast#PublicKey).

### Fixed

- Fixed AST Operator empty value.
- Fixed comment lexer rule.

## [v0.0.1] - 2023-01-10

### Added

- Added basic [AST](https://pkg.go.dev/github.com/durudex/go-polylang/ast).
- Added basic [lexer rules](https://pkg.go.dev/github.com/durudex/go-polylang#Lexer).
