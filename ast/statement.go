/*
 * Copyright © 2022-2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package ast

type Statement struct {
	CompoundStatement
	SimpleStatement
}

type CompoundStatement struct {
	If    *If    `parser:"@@"`
	While *While `parser:"| @@"`
}

type SimpleStatement struct {
	Small *SmallStatement `parser:"| @@ ';'"`
}

type SmallStatement struct {
	Break      bool        `parser:"( @'break'? )!"`
	Throw      *Expression `parser:"| 'throw' @@"`
	Expression *Expression `parser:"| @@"`
}

type Expression struct {
	Left       string      `parser:"@( Ident | String )"`
	Operator   Operator    `parser:"( @@ )?"`
	Expression *Expression `parser:"( '(' @@ ')' )?"`
	Right      string      `parser:"( @( Ident | String ) )?"`
}

type If struct {
	Condition  *Expression  `parser:"'if' '(' @@ ')'"`
	Statements []*Statement `parser:"'{' @@* '}'"`
	Else       []*Statement `parser:"( 'else' '{' @@* '}' )?"`
}

type While struct {
	Condition  *Expression  `parser:"'while' '(' @@ ')'"`
	Statements []*Statement `parser:"'{' @@* '}'"`
}
