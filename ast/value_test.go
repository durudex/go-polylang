/*
 * Copyright © 2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package ast_test

import (
	"reflect"
	"testing"

	"github.com/durudex/go-polylang"
	"github.com/durudex/go-polylang/ast"

	"github.com/alecthomas/participle/v2"
)

var ValueTests = map[string]struct {
	code string
	want *ast.Value
}{
	"Number": {
		code: "10",
		want: &ast.Value{
			Number: func(v int) *int { return &v }(10),
		},
	},
	"Single Quoted String": {
		code: "'Durudex'",
		want: &ast.Value{
			String: func(v string) *string { return &v }("'Durudex'"),
		},
	},
	"Double Quoted String": {
		code: "\"Durudex\"",
		want: &ast.Value{
			String: func(v string) *string { return &v }("\"Durudex\""),
		},
	},
	"True": {
		code: "true",
		want: &ast.Value{Boolean: true},
	},
	"False": {
		code: "false",
		want: &ast.Value{Boolean: false},
	},
	"Ident": {
		code: "Durudex",
		want: &ast.Value{
			Ident: func(v string) *string { return &v }("Durudex"),
		},
	},
	"Sub Expression": {
		code: "(this.id == id)",
		want: &ast.Value{
			Sub: &ast.Expression{
				Left: &ast.Value{
					Ident: func(v string) *string { return &v }("this.id"),
				},
				Operator: ast.Equal,
				Right: &ast.Value{
					Ident: func(v string) *string { return &v }("id"),
				},
			},
		},
	},
}

func TestValue(t *testing.T) {
	parser := participle.MustBuild[ast.Value](
		participle.Lexer(polylang.Lexer),
	)

	for name, test := range ValueTests {
		t.Run(name, func(t *testing.T) {
			got, err := parser.ParseString("", test.code)
			if err != nil {
				t.Fatal("error: parsing polylang code: ", err)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Fatal("error: value does not match")
			}
		})
	}
}

func BenchmarkValue(b *testing.B) {
	parser := participle.MustBuild[ast.Value](
		participle.Lexer(polylang.Lexer),
	)

	for name, test := range ValueTests {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parser.ParseString("", test.code) //nolint:errcheck
			}
		})
	}
}
