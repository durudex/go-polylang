/*
 * Copyright © 2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package metadata

import "encoding/json"

type kind struct {
	Kind string `json:"kind"`
}

type AnyKind struct {
	Kind  string
	Value json.RawMessage
}

func (ak *AnyKind) UnmarshalJSON(data []byte) error {
	var k kind
	if err := json.Unmarshal(data, &k); err != nil {
		return err
	}

	ak.Kind = k.Kind
	ak.Value = data

	return nil
}

func (ak AnyKind) MarshalJSON() ([]byte, error) {
	return ak.Value, nil
}

type Root []Node

type Node AnyKind

func (n *Node) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, (*AnyKind)(n))
}

func (n Node) MarshalJSON() ([]byte, error) {
	return json.Marshal((AnyKind)(n))
}

func (n Node) Collection() (*Collection, bool, error) {
	if n.Kind != "collection" {
		return nil, false, nil
	}

	var coll Collection
	if err := json.Unmarshal(n.Value, &coll); err != nil {
		return nil, false, err
	}

	return &coll, true, nil
}
