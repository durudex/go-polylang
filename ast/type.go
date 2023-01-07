/*
 * Copyright © 2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package ast

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

type BasicType int

const (
	String BasicType = iota + 1
	Number
	Boolean
)

var (
	typeToString = map[BasicType]string{String: "string", Number: "number", Boolean: "boolean"}
	stringToType = map[string]BasicType{"string": String, "number": Number, "boolean": Boolean}
)

type Type struct {
	Basic  BasicType `parser:"@@"`
	Array  bool      `parser:"@( '[' ']' )?"`
	Map    *Map      `parser:"| @@"`
	Object []*Field  `parser:"| '{' ( ( @@ ';' )* )? '}'"`
}

type Map struct {
	Key   BasicType `parser:"'map' '<' @@ ','"`
	Value Type      `parser:"@@ '>'"`
}

func (t BasicType) String() string { return typeToString[t] }

func (t *BasicType) Parse(lex *lexer.PeekingLexer) error {
	token := lex.Peek()

	v, ok := stringToType[token.Value]
	if !ok {
		return participle.NextMatch
	}
	lex.Next()

	*t = v

	return nil
}
