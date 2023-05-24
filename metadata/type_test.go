/*
 * Copyright © 2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package metadata_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/durudex/go-polylang/metadata"
)

func TestType_Primitive(t *testing.T) {
	raw := []byte("{\"kind\":\"primitive\",\"value\":\"string\"}")
	want := &metadata.Primitive{Value: metadata.PrimitiveTypeString}
	tp := metadata.Type{}

	if err := tp.UnmarshalJSON(raw); err != nil {
		t.Fatal("error: unmarshal json: ", err)
	}

	got, status, err := tp.Primitive()
	if err != nil {
		t.Fatal("error: unmarshal primitive: ", err)
	} else if !status {
		t.Fatal("error: type kind is not primitive")
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatal("error: primitive does not match")
	}
}

func TestType_Array(t *testing.T) {
	raw := []byte("{\"kind\":\"array\",\"value\":" +
		"{\"kind\":\"primitive\",\"value\":\"string\"}}",
	)
	want := &metadata.Array{Value: metadata.Type{
		Kind: "primitive",
		Value: json.RawMessage(
			"{\"kind\":\"primitive\",\"value\":\"string\"}",
		),
	}}
	tp := metadata.Type{}

	if err := tp.UnmarshalJSON(raw); err != nil {
		t.Fatal("error: unmarshal json: ", err)
	}

	got, status, err := tp.Array()
	if err != nil {
		t.Fatal("error: unmarshal array: ", err)
	} else if !status {
		t.Fatal("error: type kind is not array")
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatal("error: array does not match")
	}
}
func TestType_Map(t *testing.T) {
	raw := []byte("{\"kind\":\"map\",\"key\":" +
		"{\"kind\":\"primitive\",\"value\":\"string\"}," +
		"\"value\":{\"kind\":\"primitive\",\"value\":\"string\"}}",
	)
	want := &metadata.Map{
		Key: metadata.Type{
			Kind: "primitive",
			Value: json.RawMessage(
				"{\"kind\":\"primitive\",\"value\":\"string\"}",
			),
		},
		Value: metadata.Type{
			Kind: "primitive",
			Value: json.RawMessage(
				"{\"kind\":\"primitive\",\"value\":\"string\"}",
			),
		},
	}
	tp := metadata.Type{}

	if err := tp.UnmarshalJSON(raw); err != nil {
		t.Fatal("error: unmarshal json: ", err)
	}

	got, status, err := tp.Map()
	if err != nil {
		t.Fatal("error: unmarshal map: ", err)
	} else if !status {
		t.Fatal("error: type kind is not map")
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatal("error: map does not match")
	}
}

func TestType_Object(t *testing.T) {
	raw := []byte("{\"kind\":\"object\",\"fields\":[" +
		"{\"name\":\"test\",\"type\":{\"kind\":\"primitive\"," +
		"\"value\":\"string\"},\"required\":true}]}",
	)
	want := &metadata.Object{Fields: []metadata.ObjectField{
		{
			Name: "test",
			Type: metadata.Type{
				Kind: "primitive",
				Value: json.RawMessage(
					"{\"kind\":\"primitive\",\"value\":\"string\"}",
				),
			},
			Required: true,
		},
	}}
	tp := metadata.Type{}

	if err := tp.UnmarshalJSON(raw); err != nil {
		t.Fatal("error: unmarshal json: ", err)
	}

	got, status, err := tp.Object()
	if err != nil {
		t.Fatal("error: unmarshal object: ", err)
	} else if !status {
		t.Fatal("error: type kind is not object")
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatal("error: object does not match")
	}
}

func TestRecord_MarshalJSON(t *testing.T) {
	raw := []byte("{\"kind\":\"example\"}")
	want := []byte("{}")

	var rec metadata.Record
	if err := json.Unmarshal(raw, &rec); err != nil {
		t.Fatal("error: unmarshal json: ", err)
	}

	got, err := json.Marshal(rec)
	if err != nil {
		t.Fatal("error: marshal json: ", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatal("error: record does not match")
	}
}

func TestType_Record(t *testing.T) {
	raw := []byte("{\"kind\":\"record\",\"value\":\"example\"}")
	want := &metadata.Record{}
	tp := metadata.Type{}

	if err := tp.UnmarshalJSON(raw); err != nil {
		t.Fatal("error: unmarshal json: ", err)
	}

	got, status, err := tp.Record()
	if err != nil {
		t.Fatal("error: unmarshal record: ", err)
	} else if !status {
		t.Fatal("error: type kind is not record")
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatal("error: record does not match")
	}
}

func TestType_ForeignRecord(t *testing.T) {
	raw := []byte("{\"kind\":\"foreignrecord\"," +
		"\"collection\":\"example\"}")
	want := &metadata.ForeignRecord{Collection: "example"}
	tp := metadata.Type{}

	if err := tp.UnmarshalJSON(raw); err != nil {
		t.Fatal("error: unmarshal json: ", err)
	}

	got, status, err := tp.ForeignRecord()
	if err != nil {
		t.Fatal("error: unmarshal foreign record: ", err)
	} else if !status {
		t.Fatal("error: type kind is not foreign record")
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatal("error: foreign record does not match")
	}
}

func TestType_PublicKey(t *testing.T) {
	raw := []byte("{\"kind\":\"publickey\"}")
	want := &metadata.PublicKey{}
	tp := metadata.Type{}

	if err := tp.UnmarshalJSON(raw); err != nil {
		t.Fatal("error: unmarshal json: ", err)
	}

	got, status, err := tp.PublicKey()
	if err != nil {
		t.Fatal("error: unmarshal public key: ", err)
	} else if !status {
		t.Fatal("error: type kind is not public key")
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatal("error: public key does not match")
	}
}
