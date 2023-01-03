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
)

func Test_SmallStatement(t *testing.T) {
	parser := participle.MustBuild[ast.SmallStatement](
		participle.Lexer(polylang.Lexer),
	)

	tests := []struct {
		name string
		code string
		want *ast.SmallStatement
	}{
		{
			name: "Break",
			code: "break",
			want: &ast.SmallStatement{Break: true},
		},
		{
			name: "Throw",
			code: "throw error('error message')",
			want: &ast.SmallStatement{
				Throw: &ast.Expression{
					Left: "error",
					Expression: &ast.Expression{
						Left: "'error message'",
					},
				},
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
				t.Fatal("error: small statement does not match")
			}
		})
	}
}

func Test_Expression(t *testing.T) {
	parser := participle.MustBuild[ast.Expression](
		participle.Lexer(polylang.Lexer),
	)

	tests := []struct {
		name string
		code string
		want *ast.Expression
	}{
		{
			name: "OK",
			code: "this.id == id",
			want: &ast.Expression{
				Left:     "this.id",
				Operator: ast.Equal,
				Right:    "id",
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
				t.Fatal("error: expression does not match")
			}
		})
	}
}

func Test_If(t *testing.T) {
	parser := participle.MustBuild[ast.If](
		participle.Lexer(polylang.Lexer),
	)

	tests := []struct {
		name string
		code string
		want *ast.If
	}{
		{
			name: "OK",
			code: "if (this.id != id) { this.name = name; }",
			want: &ast.If{
				Condition: &ast.Expression{
					Left:     "this.id",
					Operator: ast.NotEqual,
					Right:    "id",
				},
				Statements: []*ast.Statement{
					{
						SimpleStatement: ast.SimpleStatement{
							Small: &ast.SmallStatement{
								Expression: &ast.Expression{
									Left:     "this.name",
									Operator: ast.Assign,
									Right:    "name",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Else",
			code: "if (this.name) {} else { this.age = age; }",
			want: &ast.If{
				Condition: &ast.Expression{Left: "this.name"},
				Else: []*ast.Statement{
					{
						SimpleStatement: ast.SimpleStatement{
							Small: &ast.SmallStatement{
								Expression: &ast.Expression{
									Left:     "this.age",
									Operator: ast.Assign,
									Right:    "age",
								},
							},
						},
					},
				},
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
				t.Fatal("error: if statement does not match")
			}
		})
	}
}

func Test_While(t *testing.T) {
	parser := participle.MustBuild[ast.While](
		participle.Lexer(polylang.Lexer),
	)

	tests := []struct {
		name string
		code string
		want *ast.While
	}{
		{
			name: "OK",
			code: "while (this.balance < balance) { break; }",
			want: &ast.While{
				Condition: &ast.Expression{
					Left:     "this.balance",
					Operator: ast.LessThan,
					Right:    "balance",
				},
				Statements: []*ast.Statement{
					{
						SimpleStatement: ast.SimpleStatement{
							Small: &ast.SmallStatement{
								Break: true,
							},
						},
					},
				},
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
				t.Fatal("error: while statement does not match")
			}
		})
	}
}
