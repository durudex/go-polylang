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

type Order int

const (
	Asc Order = iota
	Desc
)

var (
	OrderToString = map[Order]string{Asc: "asc", Desc: "desc"}
	StringToOrder = map[string]Order{"asc": Asc, "desc": Desc}
)

func (o Order) String() string { return OrderToString[o] }

func (o *Order) Parse(lex *lexer.PeekingLexer) error {
	token := lex.Peek()

	v, ok := StringToOrder[token.Value]
	if !ok {
		return participle.NextMatch
	}
	lex.Next()

	*o = v

	return nil
}
