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

type DecoratorNameMock struct {
	Name ast.DecoratorName `parser:"@@"`
}

func TestDecoratorName(t *testing.T) {
	parser := participle.MustBuild[DecoratorNameMock](
		participle.Lexer(polylang.Lexer),
	)

	for name, want := range ast.StringToDecoratorName {
		t.Run(name, func(t *testing.T) {

			got, err := parser.ParseString("", name)
			if err != nil {
				t.Fatal("error: parsing decorator name: ", err)
			}

			if !reflect.DeepEqual(got.Name, want) {
				t.Fatal("error: decorator name does not match")
			}
		})
	}
}

func BenchmarkDecoratorName(b *testing.B) {
	parser := participle.MustBuild[DecoratorNameMock](
		participle.Lexer(polylang.Lexer),
	)

	for name := range ast.StringToDecoratorName {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parser.ParseString("", name) //nolint:errcheck
			}
		})
	}
}

var DecoratorTests = map[string]struct {
	code string
	want *ast.Decorator
}{
	"OK": {
		code: "@public",
		want: &ast.Decorator{Name: ast.Public},
	},
	"Argument": {
		code: "@call(owner)",
		want: &ast.Decorator{
			Name:      ast.Call,
			Arguments: []string{"owner"},
		},
	},
}

func TestDecorator(t *testing.T) {
	parser := participle.MustBuild[ast.Decorator](
		participle.Lexer(polylang.Lexer),
	)

	for name, test := range DecoratorTests {
		t.Run(name, func(t *testing.T) {
			got, err := parser.ParseString("", test.code)
			if err != nil {
				t.Fatal("error: parsing polylang code: ", err)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Fatal("error: decorator does not match")
			}
		})
	}
}

func BenchmarkDecorator(b *testing.B) {
	parser := participle.MustBuild[ast.Decorator](
		participle.Lexer(polylang.Lexer),
	)

	for name, test := range DecoratorTests {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parser.ParseString("", test.code) //nolint:errcheck
			}
		})
	}
}
