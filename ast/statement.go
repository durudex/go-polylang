/*
 * Copyright © 2022-2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package ast

type Statement struct {
	Compound *CompoundStatement `parser:"@@"`
	Simple   *SimpleStatement   `parser:"| @@"`
}

type CompoundStatement struct {
	If    *If    `parser:"@@"`
	While *While `parser:"| @@"`
	For   *For   `parser:"| @@"`
}

type SimpleStatement struct {
	Small *SmallStatement `parser:"@@ ';'"`
}

type SmallStatement struct {
	Break      bool        `parser:"( @'break'? )!"`
	Return     *Expression `parser:"| 'return' @@"`
	Throw      *Expression `parser:"| 'throw' @@"`
	Let        *Let        `parser:"| @@"`
	Expression *Expression `parser:"| @@"`
}

type StatementsOrSimple struct {
	Statements []*Statement     `parser:"'{' @@* '}'"`
	Simple     *SimpleStatement `parser:"| @@"`
}

type Expression struct {
	Left     *Value   `parser:"@@"`
	Operator Operator `parser:"( @@ )?"`
	Right    *Value   `parser:"( @@ )?"`
}

type If struct {
	Condition *Expression         `parser:"'if' '(' @@ ')'"`
	Statement *StatementsOrSimple `parser:"( @@ )?"`
	Else      *StatementsOrSimple `parser:"( 'else' @@ )?"`
}

type While struct {
	Condition  *Expression  `parser:"'while' '(' @@ ')'"`
	Statements []*Statement `parser:"'{' @@* '}'"`
}

type Let struct {
	Ident      string      `parser:"'let' @Ident '='"`
	Expression *Expression `parser:"@@"`
}

type For struct {
	Initial    *ForInitial  `parser:"'for' '(' @@ ';'"`
	Condition  *Expression  `parser:"@@ ';'"`
	Post       *Expression  `parser:"@@ ')'"`
	Statements []*Statement `parser:"'{' @@* '}'"`
}

type ForInitial struct {
	Let        *Let        `parser:"@@"`
	Expression *Expression `parser:"| @@"`
}
