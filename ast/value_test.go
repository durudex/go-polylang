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

func Test_Value(t *testing.T) {
	parser := participle.MustBuild[ast.Value](
		participle.Lexer(polylang.Lexer),
	)

	tests := []struct {
		name string
		code string
		want *ast.Value
	}{
		{
			name: "Number",
			code: "10",
			want: &ast.Value{
				Number: func(v int) *int { return &v }(10),
			},
		},
		{
			name: "Single Quoted String",
			code: "'Durudex'",
			want: &ast.Value{
				String: func(v string) *string { return &v }("'Durudex'"),
			},
		},
		{
			name: "Double Quoted String",
			code: "\"Durudex\"",
			want: &ast.Value{
				String: func(v string) *string { return &v }("\"Durudex\""),
			},
		},
		{
			name: "Ident",
			code: "Durudex",
			want: &ast.Value{
				Ident: func(v string) *string { return &v }("Durudex"),
			},
		},
		{
			name: "Sub Expression",
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parser.ParseString("", tt.code)
			if err != nil {
				t.Fatal("error: parsing polylang code: ", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatal("error: value does not match")
			}
		})
	}
}
