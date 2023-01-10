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

func Test_Type(t *testing.T) {
	parser := participle.MustBuild[ast.Type](
		participle.Lexer(polylang.Lexer),
	)

	tests := []struct {
		name string
		code string
		want *ast.Type
	}{
		{
			name: "OK",
			code: "string",
			want: &ast.Type{Basic: ast.String},
		},
		{
			name: "Array",
			code: "string[]",
			want: &ast.Type{Basic: ast.String, Array: true},
		},
		{
			name: "Map",
			code: "map<string, number>",
			want: &ast.Type{
				Map: &ast.Map{
					Key:   ast.String,
					Value: ast.Type{Basic: ast.Number},
				},
			},
		},
		{
			name: "Object",
			code: "{name: string; website?: string;}",
			want: &ast.Type{
				Object: []*ast.Field{
					{
						Name: "name",
						Type: ast.Type{Basic: ast.String},
					},
					{
						Name:     "website",
						Optional: true,
						Type:     ast.Type{Basic: ast.String},
					},
				},
			},
		},
		{
			name: "Foreign",
			code: "ForeignCollection",
			want: &ast.Type{Foreign: "ForeignCollection"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parser.ParseString("", tt.code)
			if err != nil {
				t.Fatal("error: parsing polylang code: ", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatal("error: type does not match")
			}
		})
	}
}
