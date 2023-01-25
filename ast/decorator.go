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

type DecoratorName int

const (
	Public DecoratorName = iota + 1
	Read
	Call
	Delegate
)

var (
	DecoratorNameToString = map[DecoratorName]string{
		Public: "public", Read: "read", Call: "call", Delegate: "delegate",
	}
	StringToDecoratorName = map[string]DecoratorName{
		"public": Public, "read": Read, "call": Call, "delegate": Delegate,
	}
)

func (d DecoratorName) String() string { return DecoratorNameToString[d] }

func (d *DecoratorName) Parse(lex *lexer.PeekingLexer) error {
	token := lex.Peek()

	v, ok := StringToDecoratorName[token.Value]
	if !ok {
		return participle.NextMatch
	}
	lex.Next()

	*d = v

	return nil
}

type Decorator struct {
	Name      DecoratorName `parser:"'@' @@"`
	Arguments []string      `parser:"( '(' @Ident ')' )?"`
}
