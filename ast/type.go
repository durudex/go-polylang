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
	Record
)

var (
	TypeToString = map[BasicType]string{
		String: "string", Number: "number", Boolean: "boolean", Record: "record",
	}
	StringToType = map[string]BasicType{
		"string": String, "number": Number, "boolean": Boolean, "record": Record,
	}
)

type Type struct {
	Basic   BasicType `parser:"@@"`
	Array   bool      `parser:"@( '[' ']' )?"`
	Map     *Map      `parser:"| @@"`
	Object  []*Field  `parser:"| '{' ( ( @@ ';' )* )? '}'"`
	Foreign string    `parser:"| @Ident"`
}

type Map struct {
	Key   BasicType `parser:"'map' '<' @@ ','"`
	Value Type      `parser:"@@ '>'"`
}

func (t BasicType) String() string { return TypeToString[t] }

func (t *BasicType) Parse(lex *lexer.PeekingLexer) error {
	token := lex.Peek()

	v, ok := StringToType[token.Value]
	if !ok {
		return participle.NextMatch
	}
	lex.Next()

	*t = v

	return nil
}
