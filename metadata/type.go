/*
 * Copyright © 2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package metadata

import "encoding/json"

type Type AnyKind

func (t *Type) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, (*AnyKind)(t))
}

func (t Type) MarshalJSON() ([]byte, error) {
	return json.Marshal((AnyKind)(t))
}

type PrimitiveType string

const (
	PrimitiveTypeString  PrimitiveType = "string"
	PrimitiveTypeNumber  PrimitiveType = "number"
	PrimitiveTypeBoolean PrimitiveType = "boolean"
	PrimitiveTypeBytes   PrimitiveType = "bytes"
)

type Primitive struct {
	Value PrimitiveType `json:"value"`
}

func (t Type) Primitive() (*Primitive, bool, error) {
	if t.Kind != "primitive" {
		return nil, false, nil
	}

	var prm Primitive
	if err := json.Unmarshal(t.Value, &prm); err != nil {
		return nil, true, err
	}

	return &prm, true, nil
}

type Array struct {
	Value Type `json:"value"`
}

func (t Type) Array() (*Array, bool, error) {
	if t.Kind != "array" {
		return nil, false, nil
	}

	var array Array
	if err := json.Unmarshal(t.Value, &array); err != nil {
		return nil, true, err
	}

	return &array, true, nil
}

type Map struct {
	Key   Type `json:"key"`
	Value Type `json:"value"`
}

func (t Type) Map() (*Map, bool, error) {
	if t.Kind != "map" {
		return nil, false, nil
	}

	var mp Map
	if err := json.Unmarshal(t.Value, &mp); err != nil {
		return nil, true, err
	}

	return &mp, true, nil
}

type Object struct {
	Fields []ObjectField `json:"fields"`
}

type ObjectField struct {
	Name     string `json:"name"`
	Type     Type   `json:"type"`
	Required bool   `json:"required"`
}

func (t Type) Object() (*Object, bool, error) {
	if t.Kind != "object" {
		return nil, false, nil
	}

	var obj Object
	if err := json.Unmarshal(t.Value, &obj); err != nil {
		return nil, true, err
	}

	return &obj, true, nil
}

type Record struct{}

func (r Record) MarshalJSON() ([]byte, error) {
	return []byte("{}"), nil
}

func (t Type) Record() (*Record, bool, error) {
	if t.Kind != "record" {
		return nil, false, nil
	}

	var rec Record
	if err := json.Unmarshal(t.Value, &rec); err != nil {
		return nil, true, err
	}

	return &rec, true, nil
}

type ForeignRecord struct {
	Collection string `json:"collection"`
}

func (t Type) ForeignRecord() (*ForeignRecord, bool, error) {
	if t.Kind != "foreignrecord" {
		return nil, false, nil
	}

	var rec ForeignRecord
	if err := json.Unmarshal(t.Value, &rec); err != nil {
		return nil, true, err
	}

	return &rec, true, nil
}

type PublicKey struct{}

func (t Type) PublicKey() (*PublicKey, bool, error) {
	if t.Kind != "publickey" {
		return nil, false, nil
	}

	var pub PublicKey
	if err := json.Unmarshal(t.Value, &pub); err != nil {
		return nil, true, err
	}

	return &pub, true, nil
}
