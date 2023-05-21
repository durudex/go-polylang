/*
 * Copyright © 2022-2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package ast

type Collection struct {
	Decorators []*Decorator `parser:"( @@* )?"`
	Name       string       `parser:"'collection' @Ident"`
	Items      []*Item      `parser:"'{' @@* '}'"`
}

type Item struct {
	Decorators []*Decorator `parser:"( @@* )?"`
	Function   *Function    `parser:"@@"`
	Field      *Field       `parser:"| @@ ';'"`
	Index      *Index       `parser:"| @@ ';'"`
}

type Field struct {
	Name     string `parser:"@Ident"`
	Optional bool   `parser:"@'?'?"`
	Type     Type   `parser:"':' @@"`
}

type Index struct {
	Fields []*IndexField `parser:"'@' 'index' '(' ( @@ ( ',' @@ )* )? ')'"`
}

type IndexField struct {
	Name  string `parser:"( '[' )? ( @Ident )"`
	Order Order  `parser:"( ',' @@ ']' )?"`
}

type Function struct {
	Name       string       `parser:"( 'function' )? @Ident '('"`
	Parameters []*Field     `parser:"( @@ ( ',' @@ )* )? ')'"`
	ReturnType Type         `parser:"( ':' @@ )?"`
	Statements []*Statement `parser:"'{' ( @@* )? '}'"`
}
