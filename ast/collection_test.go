/*
 * Copyright © 2022-2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package ast_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/durudex/go-polylang"
	"github.com/durudex/go-polylang/ast"

	"github.com/alecthomas/participle/v2"
)

var CollectionTests = map[string]struct {
	file string
	want *ast.Collection
}{
	"OK": {
		file: "fixtures/article.polylang",
		want: &ast.Collection{
			Decorators: []*ast.Decorator{{Name: ast.Public}},
			Name:       "Article",
			Items: []*ast.Item{
				{
					Field: &ast.Field{
						Name: "id",
						Type: ast.Type{Basic: ast.String},
					},
				},
				{
					Field: &ast.Field{
						Name: "title",
						Type: ast.Type{Basic: ast.String},
					},
				},
				{
					Field: &ast.Field{
						Name: "info",
						Type: ast.Type{
							Object: []*ast.Field{
								{
									Name: "author",
									Type: ast.Type{Basic: ast.String},
								},
								{
									Name:     "sponsor",
									Optional: true,
									Type:     ast.Type{Basic: ast.String},
								},
							},
						},
					},
				},
				{
					Function: &ast.Function{
						Name: "constructor",
						Parameters: []*ast.Field{
							{
								Name: "id",
								Type: ast.Type{Basic: ast.String},
							},
							{
								Name: "title",
								Type: ast.Type{Basic: ast.String},
							},
						},
						Statements: []*ast.Statement{
							{
								Simple: &ast.SimpleStatement{
									Small: &ast.SmallStatement{
										Expression: &ast.Expression{
											Left: &ast.Value{
												Ident: func(v string) *string {
													return &v
												}("this.id"),
											},
											Operator: ast.Assign,
											Right: &ast.Value{
												Ident: func(v string) *string {
													return &v
												}("id"),
											},
										},
									},
								},
							},
							{
								Simple: &ast.SimpleStatement{
									Small: &ast.SmallStatement{
										Expression: &ast.Expression{
											Left: &ast.Value{
												Ident: func(v string) *string {
													return &v
												}("this.title"),
											},
											Operator: ast.Assign,
											Right: &ast.Value{
												Ident: func(v string) *string {
													return &v
												}("title"),
											},
										},
									},
								},
							},
						},
					},
				},
				{
					Function: &ast.Function{
						Name: "del",
						Statements: []*ast.Statement{
							{
								Simple: &ast.SimpleStatement{
									Small: &ast.SmallStatement{
										Expression: &ast.Expression{
											Left: &ast.Value{
												Ident: func(v string) *string {
													return &v
												}("selfdestruct"),
											},
											Right: &ast.Value{},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	},
}

func TestCollection(t *testing.T) {
	parser := participle.MustBuild[ast.Collection](
		participle.Lexer(polylang.Lexer),
	)

	for name, test := range CollectionTests {
		t.Run(name, func(t *testing.T) {
			f, err := os.Open(test.file)
			if err != nil {
				t.Fatal("error: opening fixtures file: ", err)
			}
			defer f.Close()

			got, err := parser.Parse("", f)
			if err != nil {
				t.Fatal("error: parsing polylang code: ", err)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Fatal("error: collection does not match")
			}
		})
	}
}

func BenchmarkCollection(b *testing.B) {
	parser := participle.MustBuild[ast.Collection](
		participle.Lexer(polylang.Lexer),
	)

	for name, test := range CollectionTests {
		b.Run(name, func(b *testing.B) {
			f, err := os.Open(test.file)
			if err != nil {
				b.Fatal("error: opening fixtures file: ", err)
			}
			defer f.Close()

			for i := 0; i < b.N; i++ {
				parser.Parse("", f) //nolint:errcheck
			}
		})
	}
}

var FieldTests = map[string]struct {
	code string
	want *ast.Field
}{
	"OK": {
		code: "id: string",
		want: &ast.Field{
			Name: "id",
			Type: ast.Type{Basic: ast.String},
		},
	},
	"Optional": {
		code: "name?: string",
		want: &ast.Field{
			Name:     "name",
			Optional: true,
			Type:     ast.Type{Basic: ast.String},
		},
	},
}

func TestField(t *testing.T) {
	parser := participle.MustBuild[ast.Field](
		participle.Lexer(polylang.Lexer),
	)

	for name, test := range FieldTests {
		t.Run(name, func(t *testing.T) {
			got, err := parser.ParseString("", test.code)
			if err != nil {
				t.Fatal("error: parsing polylang code: ", err)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Fatal("error: field does not match")
			}
		})
	}
}

func BenchmarkField(b *testing.B) {
	parser := participle.MustBuild[ast.Field](
		participle.Lexer(polylang.Lexer),
	)

	for name, test := range FieldTests {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parser.ParseString("", test.code) //nolint:errcheck
			}
		})
	}
}

var IndexTests = map[string]struct {
	code string
	want *ast.Index
}{
	"OK": {
		code: "@index()",
		want: &ast.Index{},
	},
}

func TestIndex(t *testing.T) {
	parser := participle.MustBuild[ast.Index](
		participle.Lexer(polylang.Lexer),
	)

	for name, test := range IndexTests {
		t.Run(name, func(t *testing.T) {
			got, err := parser.ParseString("", test.code)
			if err != nil {
				t.Fatal("error: parsing polylang code: ", err)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Fatal("error: index does not match")
			}
		})
	}
}

func BenchmarkIndex(b *testing.B) {
	parser := participle.MustBuild[ast.Index](
		participle.Lexer(polylang.Lexer),
	)

	for name, test := range IndexTests {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parser.ParseString("", test.code) //nolint:errcheck
			}
		})
	}
}

var IndexFieldTests = map[string]struct {
	code string
	want *ast.IndexField
}{
	"OK": {
		code: "[id, desc]",
		want: &ast.IndexField{
			Name:  "id",
			Order: ast.Desc,
		},
	},
	"Simple Field": {
		code: "id",
		want: &ast.IndexField{
			Name:  "id",
			Order: ast.Asc,
		},
	},
}

func TestIndexField(t *testing.T) {
	parser := participle.MustBuild[ast.IndexField](
		participle.Lexer(polylang.Lexer),
	)

	for name, test := range IndexFieldTests {
		t.Run(name, func(t *testing.T) {
			got, err := parser.ParseString("", test.code)
			if err != nil {
				t.Fatal("error: parsing polylang code: ", err)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Fatal("error: index field does not match")
			}
		})
	}
}

func BenchmarkIndexField(b *testing.B) {
	parser := participle.MustBuild[ast.IndexField](
		participle.Lexer(polylang.Lexer),
	)

	for name, test := range IndexFieldTests {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parser.ParseString("", test.code) //nolint:errcheck
			}
		})
	}
}

var FunctionTests = map[string]struct {
	code string
	want *ast.Function
}{
	"OK": {
		code: "function test() {}",
		want: &ast.Function{Name: "test"},
	},
	"Return Type": {
		code: "function test(): string {}",
		want: &ast.Function{
			Name:       "test",
			ReturnType: ast.Type{Basic: ast.String},
		},
	},
	"Short": {
		code: "test() {}",
		want: &ast.Function{Name: "test"},
	},
}

func TestFunction(t *testing.T) {
	parser := participle.MustBuild[ast.Function](
		participle.Lexer(polylang.Lexer),
	)

	for name, test := range FunctionTests {
		t.Run(name, func(t *testing.T) {
			got, err := parser.ParseString("", test.code)
			if err != nil {
				t.Fatal("error: parsing polylang code: ", err)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Fatal("error: function does not match")
			}
		})
	}
}

func BenchmarkFunction(b *testing.B) {
	parser := participle.MustBuild[ast.Function](
		participle.Lexer(polylang.Lexer),
	)

	for name, test := range FunctionTests {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parser.ParseString("", test.code) //nolint:errcheck
			}
		})
	}
}
