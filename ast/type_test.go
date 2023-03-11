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

type BasicTypeMock struct {
	Type ast.BasicType `parser:"@@"`
}

func TestBasicType(t *testing.T) {
	parser := participle.MustBuild[BasicTypeMock](
		participle.Lexer(polylang.Lexer),
	)

	for basic, want := range ast.StringToType {
		t.Run(basic, func(t *testing.T) {

			got, err := parser.ParseString("", basic)
			if err != nil {
				t.Fatal("error: parsing basic type: ", err)
			}

			if !reflect.DeepEqual(got.Type, want) {
				t.Fatal("error: basic type does not match")
			}
		})
	}
}

func BenchmarkBasicType(b *testing.B) {
	parser := participle.MustBuild[BasicTypeMock](
		participle.Lexer(polylang.Lexer),
	)

	for basic := range ast.StringToType {
		b.Run(basic, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parser.ParseString("", basic) //nolint:errcheck
			}
		})
	}
}

var TypeTests = map[string]struct {
	code string
	want *ast.Type
}{
	"Array": {
		code: "string[]",
		want: &ast.Type{Basic: ast.String, Array: true},
	},
	"Map": {
		code: "map<string, number>",
		want: &ast.Type{
			Map: &ast.Map{
				Key:   ast.String,
				Value: ast.Type{Basic: ast.Number},
			},
		},
	},
	"Object": {
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
	"Foreign": {
		code: "ForeignCollection",
		want: &ast.Type{Foreign: "ForeignCollection"},
	},
}

func TestType(t *testing.T) {
	parser := participle.MustBuild[ast.Type](
		participle.Lexer(polylang.Lexer),
	)

	for name, test := range TypeTests {
		t.Run(name, func(t *testing.T) {
			got, err := parser.ParseString("", test.code)
			if err != nil {
				t.Fatal("error: parsing polylang code: ", err)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Fatal("error: type does not match")
			}
		})
	}
}

func BenchmarkType(b *testing.B) {
	parser := participle.MustBuild[ast.Type](
		participle.Lexer(polylang.Lexer),
	)

	for name, test := range TypeTests {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parser.ParseString("", test.code) //nolint:errcheck
			}
		})
	}
}
