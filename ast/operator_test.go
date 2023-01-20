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

func Test_Operator(t *testing.T) {
	type Mock struct {
		Operator ast.Operator `parser:"@@"`
	}

	parser := participle.MustBuild[Mock](
		participle.Lexer(polylang.Lexer),
	)

	for i, want := range ast.StringToOperator {
		t.Run(i, func(t *testing.T) {

			got, err := parser.ParseString("", i)
			if err != nil {
				t.Fatal("error: parsing operator: ", err)
			}

			if !reflect.DeepEqual(got.Operator, want) {
				t.Fatal("error: operator does not match")
			}
		})
	}
}
