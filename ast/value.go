/*
 * Copyright © 2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package ast

type Value struct {
	Number  *int        `parser:"@Number"`
	String  *string     `parser:"| @String"`
	Boolean bool        `parser:"| ( @'true' | 'false' )"`
	Ident   *string     `parser:"| @Ident"`
	Sub     *Expression `parser:"| '(' ( @@ )? ')'"`
}
