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

func Test_Collection(t *testing.T) {
	parser := participle.MustBuild[ast.Collection](
		participle.Lexer(polylang.Lexer),
	)

	tests := []struct {
		name string
		file string
		want *ast.Collection
	}{
		{
			name: "OK",
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := os.Open(tt.file)
			if err != nil {
				t.Fatal("error: opening fixtures file: ", err)
			}
			defer f.Close()

			got, err := parser.Parse("", f)
			if err != nil {
				t.Fatal("error: parsing polylang code: ", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatal("error: collection does not match")
			}
		})
	}
}

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
				Name: "id",
				Type: ast.Type{Basic: ast.String},
			},
		},
		{
			name: "Optional",
			code: "name?: string",
			want: &ast.Field{
				Name:     "name",
				Optional: true,
				Type:     ast.Type{Basic: ast.String},
			},
		},
		{
			name: "Decorator",
			code: "@read name: string",
			want: &ast.Field{
				Decorators: []*ast.Decorator{{Name: ast.Read}},
				Name:       "name",
				Type:       ast.Type{Basic: ast.String},
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
			want: &ast.Index{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parser.ParseString("", tt.code)
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
				Name:  "id",
				Order: ast.Desc,
			},
		},
		{
			name: "Simple Field",
			code: "id",
			want: &ast.IndexField{
				Name:  "id",
				Order: ast.Asc,
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
			code: "function test() {}",
			want: &ast.Function{Name: "test"},
		},
		{
			name: "Return Type",
			code: "function test(): string {}",
			want: &ast.Function{
				Name:       "test",
				ReturnType: ast.Type{Basic: ast.String},
			},
		},
		{
			name: "Short",
			code: "test() {}",
			want: &ast.Function{Name: "test"},
		},
		{
			name: "Decorator",
			code: "@call(owner) function test() {}",
			want: &ast.Function{
				Decorators: []*ast.Decorator{
					{Name: ast.Call, Arguments: []string{"owner"}},
				},
				Name: "test",
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
				t.Fatal("error: function does not match")
			}
		})
	}
}
