/*
 * Copyright © 2022-2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package ast_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/durudex/go-polylang"
	"github.com/durudex/go-polylang/ast"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

func Test_Field(t *testing.T) {
	parser := participle.MustBuild[ast.Field](
		participle.Lexer(polylang.Lexer),
	)

	tests := []struct {
		name string
		code string
		want *ast.Field
	}{
		{
			name: "OK",
			code: "id: string",
			want: &ast.Field{
				Pos:  lexer.Position{Line: 1, Column: 1},
				Name: "id",
				Type: "string",
			},
		},
		{
			name: "Optional Field",
			code: "name?: string",
			want: &ast.Field{
				Pos:      lexer.Position{Line: 1, Column: 1},
				Name:     "name",
				Optional: true,
				Type:     "string",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parser.Parse("", strings.NewReader(tt.code))
			if err != nil {
				t.Fatal("error: parsing polylang code: ", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatal("error: field does not match")
			}
		})
	}
}

func Test_Index(t *testing.T) {
	parser := participle.MustBuild[ast.Index](
		participle.Lexer(polylang.Lexer),
	)

	tests := []struct {
		name string
		code string
		want *ast.Index
	}{
		{
			name: "OK",
			code: "@index()",
			want: &ast.Index{
				Pos: lexer.Position{Line: 1, Column: 1},
			},
		},
		{
			name: "Unique",
			code: "@unique()",
			want: &ast.Index{
				Pos:    lexer.Position{Line: 1, Column: 1},
				Unique: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parser.Parse("", strings.NewReader(tt.code))
			if err != nil {
				t.Fatal("error: parsing polylang code: ", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatal("error: index does not match")
			}
		})
	}
}

func Test_IndexField(t *testing.T) {
	parser := participle.MustBuild[ast.IndexField](
		participle.Lexer(polylang.Lexer),
	)

	tests := []struct {
		name string
		code string
		want *ast.IndexField
	}{
		{
			name: "OK",
			code: "[id, desc]",
			want: &ast.IndexField{
				Pos:   lexer.Position{Line: 1, Column: 1},
				Name:  "id",
				Order: ast.Desc,
			},
		},
		{
			name: "Simple Field",
			code: "id",
			want: &ast.IndexField{
				Pos:   lexer.Position{Line: 1, Column: 1},
				Name:  "id",
				Order: ast.Asc,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parser.Parse("", strings.NewReader(tt.code))
			if err != nil {
				t.Fatal("error: parsing polylang code: ", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatal("error: index field does not match")
			}
		})
	}
}

func Test_Function(t *testing.T) {
	parser := participle.MustBuild[ast.Function](
		participle.Lexer(polylang.Lexer),
	)

	tests := []struct {
		name string
		code string
		want *ast.Function
	}{
		{
			name: "OK",
			code: "function Test() {}",
			want: &ast.Function{
				Pos:  lexer.Position{Line: 1, Column: 1},
				Name: "Test",
			},
		},
		{
			name: "Return Type",
			code: "function Test(): string {}",
			want: &ast.Function{
				Pos:        lexer.Position{Line: 1, Column: 1},
				Name:       "Test",
				ReturnType: "string",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parser.Parse("", strings.NewReader(tt.code))
			if err != nil {
				t.Fatal("error: parsing polylang code: ", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatal("error: function does not match")
			}
		})
	}
}
