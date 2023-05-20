/*
 * Copyright © 2022-2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package ast

type Program struct {
	Nodes []*Node `parser:"@@*" json:"nodes,omitempty"`
}

type Node struct {
	Collection *Collection `parser:"@@"   json:"collection,omitempty"`
	Function   *Function   `parser:"| @@" json:"function,omitempty"`
}
