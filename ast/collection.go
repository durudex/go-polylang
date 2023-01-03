/*
 * Copyright © 2022-2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package ast

import "github.com/alecthomas/participle/v2/lexer"

type Collection struct {
	Pos lexer.Position

	Name  string  `parser:"'collection' @Ident"`
	Items []*Item `parser:"'{' @@* '}'"`
}

type Item struct {
	Pos lexer.Position

	Field    *Field    `parser:"@@ ';'"`
	Function *Function `parser:"| @@"`
	Index    *Index    `parser:"| @@ ';'"`
}

type Field struct {
	Pos lexer.Position

	Name     string `parser:"@Ident"`
	Optional bool   `parser:"@'?'?"`
	Type     string `parser:"':' @Ident"`
}

type Index struct {
	Pos lexer.Position

	Unique bool          `parser:"'@' ( @'unique' | 'index' )"`
	Fields []*IndexField `parser:"'(' ( @@ ( ',' @@ )* )? ')'"`
}

type IndexField struct {
	Pos lexer.Position

	Name  string `parser:"( '[' )? ( @Ident )"`
	Order Order  `parser:"( ',' @@ ']' )?"`
}

type Function struct {
	Pos lexer.Position

	Name       string       `parser:"'function' @Ident '('"`
	Parameters []*Field     `parser:"( @@ ( ',' @@ )* )? ')'"`
	ReturnType string       `parser:"( ':' @Ident )?"`
	Statements []*Statement `parser:"'{' ( @@* )? '}'"`
}
