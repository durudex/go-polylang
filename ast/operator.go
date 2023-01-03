/*
 * Copyright © 2022-2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package ast

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

type Operator int

const (
	Not Operator = iota
	BitNot
	Exponent
	Multiply
	Divide
	Modulo
	Add
	Subtract
	ShiftLeft
	ShiftRight
	BitAnd
	BitXor
	BitOr
	LessThan
	GreaterThan
	LessThanOrEqual
	GreaterThanOrEqual
	Equal
	NotEqual
	And
	Or
	AssignSub
	AssignAdd
	Assign
)

var (
	operatorToString = map[Operator]string{
		Not: "!", BitNot: "~", Exponent: "**", Multiply: "*", Divide: "/", Modulo: "%",
		Add: "+", Subtract: "-", ShiftLeft: "<<", ShiftRight: ">>", BitAnd: "&", BitXor: "^",
		BitOr: "|", LessThan: "<", GreaterThan: ">", LessThanOrEqual: "<=", GreaterThanOrEqual: ">=",
		Equal: "==", NotEqual: "!=", And: "&&", Or: "||", AssignSub: "-=", AssignAdd: "+=", Assign: "=",
	}
	stringToOperator = map[string]Operator{
		"!": Not, "~": BitNot, "**": Exponent, "*": Multiply, "/": Divide, "%": Modulo,
		"+": Add, "-": Subtract, "<<": ShiftLeft, ">>": ShiftRight, "&": BitAnd, "^": BitXor,
		"|": BitOr, "<": LessThan, ">": GreaterThan, "<=": LessThanOrEqual, ">=": GreaterThanOrEqual,
		"==": Equal, "!=": NotEqual, "&&": And, "||": Or, "-=": AssignSub, "+=": AssignAdd, "=": Assign,
	}
)

func (o Operator) GoString() string { return operatorToString[o] }

func (o *Operator) Parse(lex *lexer.PeekingLexer) error {
	token := lex.Peek()

	one, ok := stringToOperator[token.Value]
	if !ok {
		return participle.NextMatch
	} else {
		lex.Next()

		next := lex.Peek()

		two, ok := stringToOperator[token.Value+next.Value]
		if !ok {
			*o = one
		} else {
			lex.Next()

			*o = two
		}

		return nil
	}
}
