/*
 * Copyright © 2022-2023 Durudex
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
			name: "Return",
			code: "return this.name == name",
			want: &ast.SmallStatement{
				Return: &ast.Expression{
					Left: &ast.Value{
						Ident: func(v string) *string {
							return &v
						}("this.name"),
					},
					Operator: ast.Equal,
					Right: &ast.Value{
						Ident: func(v string) *string {
							return &v
						}("name"),
					},
				},
			},
		},
		{
			name: "Throw",
			code: "throw error('error message')",
			want: &ast.SmallStatement{
				Throw: &ast.Expression{
					Left: &ast.Value{
						Ident: func(v string) *string {
							return &v
						}("error"),
					},
					Right: &ast.Value{
						Sub: &ast.Expression{
							Left: &ast.Value{
								String: func(v string) *string {
									return &v
								}("'error message'"),
							},
						},
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
				Left: &ast.Value{
					Ident: func(v string) *string {
						return &v
					}("this.id"),
				},
				Operator: ast.Equal,
				Right: &ast.Value{
					Ident: func(v string) *string {
						return &v
					}("id"),
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
					Left: &ast.Value{
						Ident: func(v string) *string {
							return &v
						}("this.id"),
					},
					Operator: ast.NotEqual,
					Right: &ast.Value{
						Ident: func(v string) *string {
							return &v
						}("id"),
					},
				},
				Statements: []*ast.Statement{
					{
						SimpleStatement: ast.SimpleStatement{
							Small: &ast.SmallStatement{
								Expression: &ast.Expression{
									Left: &ast.Value{
										Ident: func(v string) *string {
											return &v
										}("this.name"),
									},
									Operator: ast.Assign,
									Right: &ast.Value{
										Ident: func(v string) *string {
											return &v
										}("name"),
									},
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
				Condition: &ast.Expression{
					Left: &ast.Value{
						Ident: func(v string) *string {
							return &v
						}("this.name"),
					}},
				Else: []*ast.Statement{
					{
						SimpleStatement: ast.SimpleStatement{
							Small: &ast.SmallStatement{
								Expression: &ast.Expression{
									Left: &ast.Value{
										Ident: func(v string) *string {
											return &v
										}("this.age"),
									},
									Operator: ast.Assign,
									Right: &ast.Value{
										Ident: func(v string) *string {
											return &v
										}("age"),
									},
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
			got, err := parser.ParseString("", tt.code)
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
					Left: &ast.Value{
						Ident: func(v string) *string {
							return &v
						}("this.balance"),
					},
					Operator: ast.LessThan,
					Right: &ast.Value{
						Ident: func(v string) *string {
							return &v
						}("balance"),
					},
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
			got, err := parser.ParseString("", tt.code)
			if err != nil {
				t.Fatal("error: parsing polylang code: ", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatal("error: while statement does not match")
			}
		})
	}
}

func Test_Let(t *testing.T) {
	parser := participle.MustBuild[ast.Let](
		participle.Lexer(polylang.Lexer),
	)

	tests := []struct {
		name string
		code string
		want *ast.Let
	}{
		{
			name: "OK",
			code: "let i = 10",
			want: &ast.Let{
				Ident: "i",
				Expression: &ast.Expression{
					Left: &ast.Value{
						Number: func(v int) *int { return &v }(10),
					}},
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
				t.Fatal("error: let statement does not match")
			}
		})
	}
}

func Test_For(t *testing.T) {
	parser := participle.MustBuild[ast.For](
		participle.Lexer(polylang.Lexer),
	)

	tests := []struct {
		name string
		code string
		want *ast.For
	}{
		{
			name: "OK",
			code: "for (let i = 0; i < 100; i + 1) { break; }",
			want: &ast.For{
				Initial: &ast.ForInitial{
					Let: &ast.Let{
						Ident: "i",
						Expression: &ast.Expression{
							Left: &ast.Value{
								Number: func(v int) *int {
									return &v
								}(0),
							},
						},
					},
				},
				Condition: &ast.Expression{
					Left: &ast.Value{
						Ident: func(v string) *string {
							return &v
						}("i"),
					},
					Operator: ast.LessThan,
					Right: &ast.Value{
						Number: func(v int) *int {
							return &v
						}(100),
					},
				},
				Post: &ast.Expression{
					Left: &ast.Value{
						Ident: func(v string) *string {
							return &v
						}("i"),
					},
					Operator: ast.Add,
					Right: &ast.Value{
						Number: func(v int) *int {
							return &v
						}(1),
					},
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
			got, err := parser.ParseString("", tt.code)
			if err != nil {
				t.Fatal("error: parsing polylang code: ", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatal("error: for statement does not match")
			}
		})
	}
}
