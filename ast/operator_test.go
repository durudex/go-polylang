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

type OperatorMock struct {
	Operator ast.Operator `parser:"@@"`
}

func TestOperator(t *testing.T) {
	parser := participle.MustBuild[OperatorMock](
		participle.Lexer(polylang.Lexer),
	)

	for operator, want := range ast.StringToOperator {
		t.Run(operator, func(t *testing.T) {

			got, err := parser.ParseString("", operator)
			if err != nil {
				t.Fatal("error: parsing operator: ", err)
			}

			if !reflect.DeepEqual(got.Operator, want) {
				t.Fatal("error: operator does not match")
			}
		})
	}
}

func BenchmarkOperator(b *testing.B) {
	parser := participle.MustBuild[OperatorMock](
		participle.Lexer(polylang.Lexer),
	)

	for operator := range ast.StringToOperator {
		b.Run(operator, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parser.ParseString("", operator) //nolint:errcheck
			}
		})
	}
}
