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

	"github.com/alecthomas/participle/v2"
)

func Test_Comment(t *testing.T) {
	type Mock struct {
		Value string `parser:"@Ident"`
	}

	parser := participle.MustBuild[Mock](
		participle.Lexer(polylang.Lexer),
	)

	tests := []struct {
		name string
		code string
		want string
	}{
		{
			name: "Line",
			code: "line // comment",
			want: "line",
		},
		{
			name: "Block",
			code: "/* comment */ block",
			want: "block",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parser.ParseString("", tt.code)
			if err != nil {
				t.Fatal("error: parsing polylang code: ", err)
			}

			if !reflect.DeepEqual(got.Value, tt.want) {
				t.Fatal("error: value does not match")
			}
		})
	}
}
