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

func Test_DecoratorName(t *testing.T) {
	type Mock struct {
		Name ast.DecoratorName `parser:"@@"`
	}

	parser := participle.MustBuild[Mock](
		participle.Lexer(polylang.Lexer),
	)

	for i, want := range ast.StringToDecoratorName {
		t.Run(i, func(t *testing.T) {

			got, err := parser.ParseString("", i)
			if err != nil {
				t.Fatal("error: parsing decorator name: ", err)
			}

			if !reflect.DeepEqual(got.Name, want) {
				t.Fatal("error: decorator name does not match")
			}
		})
	}
}
