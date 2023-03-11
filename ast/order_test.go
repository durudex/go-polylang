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

type OrderMock struct {
	Order ast.Order `parser:"@@"`
}

func TestOrder(t *testing.T) {
	parser := participle.MustBuild[OrderMock](
		participle.Lexer(polylang.Lexer),
	)

	for order, want := range ast.StringToOrder {
		t.Run(order, func(t *testing.T) {

			got, err := parser.ParseString("", order)
			if err != nil {
				t.Fatal("error: parsing order: ", err)
			}

			if !reflect.DeepEqual(got.Order, want) {
				t.Fatal("error: order does not match")
			}
		})
	}
}

func BenchmarkOrder(b *testing.B) {
	parser := participle.MustBuild[OrderMock](
		participle.Lexer(polylang.Lexer),
	)

	for order := range ast.StringToOrder {
		b.Run(order, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parser.ParseString("", order) //nolint:errcheck
			}
		})
	}
}
