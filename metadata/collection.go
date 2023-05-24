/*
 * Copyright © 2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package metadata

import (
	"encoding/json"
	"fmt"
)

type Collection struct {
	Namespace  Namespace             `json:"namespace"`
	Name       string                `json:"name"`
	Attributes []CollectionAttribute `json:"attributes"`
}

type Namespace struct {
	Value string `json:"value"`
}

type namespace struct {
	Kind  string `json:"kind"`
	Value string `json:"value"`
}

func (n *Namespace) UnmarshalJSON(data []byte) error {
	var ns namespace
	if err := json.Unmarshal(data, &ns); err != nil {
		return err
	}

	if ns.Kind != "namespace" {
		return fmt.Errorf("invalid '%s' kind type", ns.Kind)
	}

	n.Value = ns.Value

	return nil
}

func (n Namespace) MarshalJSON() ([]byte, error) {
	return json.Marshal(&namespace{
		Kind:  "namespace",
		Value: n.Value,
	})
}

type CollectionAttribute AnyKind

func (ca *CollectionAttribute) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, (*AnyKind)(ca))
}

func (ca CollectionAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal((AnyKind)(ca))
}

type Property struct {
	Name       string      `json:"name"`
	Type       Type        `json:"type"`
	Directives []Directive `json:"directives"`
	Required   bool        `json:"required"`
}

func (ca CollectionAttribute) Property() (*Property, bool, error) {
	if ca.Kind != "property" {
		return nil, false, nil
	}

	var prop Property
	if err := json.Unmarshal(ca.Value, &prop); err != nil {
		return nil, true, err
	}

	return &prop, true, nil
}
