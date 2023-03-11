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

var SmallStatementTests = map[string]struct {
	code string
	want *ast.SmallStatement
}{
	"Break": {
		code: "break",
		want: &ast.SmallStatement{Break: true},
	},
	"Return": {
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
	"Throw": {
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

func TestSmallStatement(t *testing.T) {
	parser := participle.MustBuild[ast.SmallStatement](
		participle.Lexer(polylang.Lexer),
	)

	for name, test := range SmallStatementTests {
		t.Run(name, func(t *testing.T) {
			got, err := parser.ParseString("", test.code)
			if err != nil {
				t.Fatal("error: parsing polylang code: ", err)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Fatal("error: small statement does not match")
			}
		})
	}
}

func BenchmarkSmallStatement(b *testing.B) {
	parser := participle.MustBuild[ast.SmallStatement](
		participle.Lexer(polylang.Lexer),
	)

	for name, test := range SmallStatementTests {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parser.ParseString("", test.code) //nolint:errcheck
			}
		})
	}
}

var ExpressionTests = map[string]struct {
	code string
	want *ast.Expression
}{
	"OK": {
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

func TestExpression(t *testing.T) {
	parser := participle.MustBuild[ast.Expression](
		participle.Lexer(polylang.Lexer),
	)

	for name, test := range ExpressionTests {
		t.Run(name, func(t *testing.T) {
			got, err := parser.ParseString("", test.code)
			if err != nil {
				t.Fatal("error: parsing polylang code: ", err)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Fatal("error: expression does not match")
			}
		})
	}
}

func BenchmarkExpression(b *testing.B) {
	parser := participle.MustBuild[ast.Expression](
		participle.Lexer(polylang.Lexer),
	)

	for name, test := range ExpressionTests {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parser.ParseString("", test.code) //nolint:errcheck
			}
		})
	}
}

var IfTests = map[string]struct {
	code string
	want *ast.If
}{
	"OK": {
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
			Statement: &ast.StatementsOrSimple{
				Statements: []*ast.Statement{
					{
						Simple: &ast.SimpleStatement{
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
	},
	"Else": {
		code: "if (this.name) {} else { this.age = age; }",
		want: &ast.If{
			Condition: &ast.Expression{
				Left: &ast.Value{
					Ident: func(v string) *string {
						return &v
					}("this.name"),
				},
			},
			Statement: &ast.StatementsOrSimple{},
			Else: &ast.StatementsOrSimple{
				Statements: []*ast.Statement{
					{
						Simple: &ast.SimpleStatement{
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
	},
	"Simple": {
		code: "if (name) return 123;",
		want: &ast.If{
			Condition: &ast.Expression{
				Left: &ast.Value{
					Ident: func(v string) *string { return &v }("name"),
				},
			},
			Statement: &ast.StatementsOrSimple{
				Simple: &ast.SimpleStatement{
					Small: &ast.SmallStatement{
						Return: &ast.Expression{
							Left: &ast.Value{
								Number: func(v int) *int {
									return &v
								}(123),
							},
						},
					},
				},
			},
		},
	},
}

func TestIf(t *testing.T) {
	parser := participle.MustBuild[ast.If](
		participle.Lexer(polylang.Lexer),
	)

	for name, test := range IfTests {
		t.Run(name, func(t *testing.T) {
			got, err := parser.ParseString("", test.code)
			if err != nil {
				t.Fatal("error: parsing polylang code: ", err)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Fatal("error: if statement does not match")
			}
		})
	}
}

func BenchmarkIf(b *testing.B) {
	parser := participle.MustBuild[ast.If](
		participle.Lexer(polylang.Lexer),
	)

	for name, test := range IfTests {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parser.ParseString("", test.code) //nolint:errcheck
			}
		})
	}
}

var WhileTests = map[string]struct {
	code string
	want *ast.While
}{
	"OK": {
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
					Simple: &ast.SimpleStatement{
						Small: &ast.SmallStatement{
							Break: true,
						},
					},
				},
			},
		},
	},
}

func TestWhile(t *testing.T) {
	parser := participle.MustBuild[ast.While](
		participle.Lexer(polylang.Lexer),
	)

	for name, test := range WhileTests {
		t.Run(name, func(t *testing.T) {
			got, err := parser.ParseString("", test.code)
			if err != nil {
				t.Fatal("error: parsing polylang code: ", err)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Fatal("error: while statement does not match")
			}
		})
	}
}

func BenchmarkWhile(b *testing.B) {
	parser := participle.MustBuild[ast.While](
		participle.Lexer(polylang.Lexer),
	)

	for name, test := range WhileTests {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parser.ParseString("", test.code) //nolint:errcheck
			}
		})
	}
}

var LetTests = map[string]struct {
	code string
	want *ast.Let
}{
	"OK": {
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

func TestLet(t *testing.T) {
	parser := participle.MustBuild[ast.Let](
		participle.Lexer(polylang.Lexer),
	)

	for name, test := range LetTests {
		t.Run(name, func(t *testing.T) {
			got, err := parser.ParseString("", test.code)
			if err != nil {
				t.Fatal("error: parsing polylang code: ", err)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Fatal("error: let statement does not match")
			}
		})
	}
}

func BenchmarkLet(b *testing.B) {
	parser := participle.MustBuild[ast.Let](
		participle.Lexer(polylang.Lexer),
	)

	for name, test := range LetTests {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parser.ParseString("", test.code) //nolint:errcheck
			}
		})
	}
}

var ForTests = map[string]struct {
	code string
	want *ast.For
}{
	"OK": {
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
					Simple: &ast.SimpleStatement{
						Small: &ast.SmallStatement{
							Break: true,
						},
					},
				},
			},
		},
	},
}

func TestFor(t *testing.T) {
	parser := participle.MustBuild[ast.For](
		participle.Lexer(polylang.Lexer),
	)

	for name, test := range ForTests {
		t.Run(name, func(t *testing.T) {
			got, err := parser.ParseString("", test.code)
			if err != nil {
				t.Fatal("error: parsing polylang code: ", err)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Fatal("error: for statement does not match")
			}
		})
	}
}

func BenchmarkFor(b *testing.B) {
	parser := participle.MustBuild[ast.For](
		participle.Lexer(polylang.Lexer),
	)

	for name, test := range ForTests {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parser.ParseString("", test.code) //nolint:errcheck
			}
		})
	}
}
