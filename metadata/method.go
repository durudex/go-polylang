/*
 * Copyright © 2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package metadata

import "encoding/json"

type Method struct {
	Name       string            `json:"name"`
	Attributes []MethodAttribute `json:"attributes"`
	Code       string            `json:"code"`
}

func (ca CollectionAttribute) Method() (*Method, bool, error) {
	if ca.Kind != "method" {
		return nil, false, nil
	}

	var mt Method
	if err := json.Unmarshal(ca.Value, &mt); err != nil {
		return nil, true, err
	}

	return &mt, true, nil
}

type MethodAttribute AnyKind

func (ma *MethodAttribute) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, (*AnyKind)(ma))
}

func (ma MethodAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal((AnyKind)(ma))
}

func (ma MethodAttribute) Directive() (*Directive, bool, error) {
	if ma.Kind != "directive" {
		return nil, false, nil
	}

	var dr Directive
	if err := json.Unmarshal(ma.Value, &dr); err != nil {
		return nil, true, err
	}

	return &dr, true, nil
}

type Parameter struct {
	Name     string `json:"name"`
	Type     Type   `json:"type"`
	Required bool   `json:"required"`
}

func (ma MethodAttribute) Parameter() (*Parameter, bool, error) {
	if ma.Kind != "parameter" {
		return nil, false, nil
	}

	var pr Parameter
	if err := json.Unmarshal(ma.Value, &pr); err != nil {
		return nil, false, err
	}

	return &pr, true, nil
}

type ReturnValue struct {
	Name string `json:"name"`
	Type Type   `json:"type"`
}

func (ma MethodAttribute) ReturnValue() (*ReturnValue, bool, error) {
	if ma.Kind != "returnvalue" {
		return nil, false, nil
	}

	var rv ReturnValue
	if err := json.Unmarshal(ma.Value, &rv); err != nil {
		return nil, false, err
	}

	return &rv, true, nil
}
