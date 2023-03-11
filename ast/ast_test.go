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

var CommentTests = map[string]struct {
	code string
	want string
}{
	"Line": {
		code: "line // comment",
		want: "line",
	},
	"Block": {
		code: "/* comment */ block",
		want: "block",
	},
	"Multi Line Block": {
		code: "/*\n * multi\n* line\n*/ multi",
		want: "multi",
	},
}

type CommentMock struct {
	Value string `parser:"@Ident"`
}

func TestComment(t *testing.T) {
	parser := participle.MustBuild[CommentMock](
		participle.Lexer(polylang.Lexer),
	)

	for name, test := range CommentTests {
		t.Run(name, func(t *testing.T) {
			got, err := parser.ParseString("", test.code)
			if err != nil {
				t.Fatal("error: parsing polylang code: ", err)
			}

			if !reflect.DeepEqual(got.Value, test.want) {
				t.Fatal("error: value does not match")
			}
		})
	}
}

func BenchmarkComment(b *testing.B) {
	parser := participle.MustBuild[CommentMock](
		participle.Lexer(polylang.Lexer),
	)

	for name, test := range CommentTests {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				parser.ParseString("", test.code) //nolint:errcheck
			}
		})
	}
}
