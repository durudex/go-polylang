/*
 * Copyright © 2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package metadata

import "encoding/json"

type Directive struct {
	Name      string            `json:"name"`
	Arguments DirectiveArgument `json:"arguments"`
}

func (ca CollectionAttribute) Directive() (*Directive, bool, error) {
	if ca.Kind != "directive" {
		return nil, false, nil
	}

	var dr Directive
	if err := json.Unmarshal(ca.Value, &dr); err != nil {
		return nil, true, err
	}

	return &dr, true, nil
}

type DirectiveArgument AnyKind

func (da *DirectiveArgument) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, (*AnyKind)(da))
}

func (da DirectiveArgument) MarshalJSON() ([]byte, error) {
	return json.Marshal((AnyKind)(da))
}

type FieldReference struct {
	Path []string `json:"path"`
}

func (da DirectiveArgument) FieldReference() (*FieldReference, bool, error) {
	if da.Kind != "fieldreference" {
		return nil, false, nil
	}

	var fr FieldReference
	if err := json.Unmarshal(da.Value, &fr); err != nil {
		return nil, true, err
	}

	return &fr, true, nil
}
