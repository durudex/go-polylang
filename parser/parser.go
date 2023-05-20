/*
 * Copyright © 2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package parser

import (
	"os"
	"path/filepath"

	"github.com/durudex/go-polylang"
	"github.com/durudex/go-polylang/ast"

	"github.com/alecthomas/participle/v2"
)

var Must = participle.MustBuild[ast.Program](
	participle.Lexer(polylang.Lexer),
)

func Parse(path string) (*ast.Program, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if info.IsDir() {
		return ParseDir(path)
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return Must.Parse(info.Name(), f)
}

func ParseDir(path string) (*ast.Program, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	entries, err := f.ReadDir(0)
	if err != nil {
		return nil, err
	}

	var ast ast.Program

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		fp := filepath.Join(path, entry.Name())

		ext := filepath.Ext(fp)
		if ext != ".polylang" {
			continue
		}

		f, err := os.Open(fp)
		if err != nil {
			return &ast, err
		}
		defer f.Close()

		fileAst, err := Must.Parse(f.Name(), f)
		if err != nil {
			return &ast, err
		}

		ast.Nodes = append(ast.Nodes, fileAst.Nodes...)
	}

	return &ast, nil
}
