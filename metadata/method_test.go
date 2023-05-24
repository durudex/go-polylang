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

func TestCollectionAttribute_Method(t *testing.T) {
	raw := []byte("{\"kind\":\"method\",\"name\":\"hello\"," +
		"\"attributes\":[],\"code\":\"if(this.id!=1970)\"}",
	)
	want := &metadata.Method{
		Name:       "hello",
		Attributes: []metadata.MethodAttribute{},
		Code:       "if(this.id!=1970)",
	}
	ma := metadata.CollectionAttribute{}

	if err := ma.UnmarshalJSON(raw); err != nil {
		t.Fatal("error: unmarshal json: ", err)
	}

	got, status, err := ma.Method()
	if err != nil {
		t.Fatal("error: unmarshal method: ", err)
	} else if !status {
		t.Fatal("error: type kind is not method")
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatal("error: method does not match")
	}
}

func TestMethodAttribute_Directive(t *testing.T) {
	raw := []byte("{\"kind\":\"directive\",\"name\":\"call\"," +
		"\"arguments\":{\"kind\":\"fieldreference\",\"path\":" +
		"[\"owner\"]}}",
	)
	want := &metadata.Directive{
		Name: "call",
		Arguments: metadata.DirectiveArgument{
			Kind: "fieldreference",
			Value: json.RawMessage(
				"{\"kind\":\"fieldreference\",\"path\":[\"owner\"]}",
			),
		},
	}
	ma := metadata.MethodAttribute{}

	if err := ma.UnmarshalJSON(raw); err != nil {
		t.Fatal("error: unmarshal json: ", err)
	}

	got, status, err := ma.Directive()
	if err != nil {
		t.Fatal("error: unmarshal directive: ", err)
	} else if !status {
		t.Fatal("error: type kind is not directive")
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatal("error: directive does not match")
	}
}

func TestMethodAttribute_Parameter(t *testing.T) {
	raw := []byte("{\"kind\":\"parameter\",\"name\":\"id\"," +
		"\"type\":{\"kind\":\"primitive\",\"value\":\"string\"}," +
		"\"required\":true}",
	)
	want := &metadata.Parameter{
		Name: "id",
		Type: metadata.Type{
			Kind: "primitive",
			Value: json.RawMessage(
				"{\"kind\":\"primitive\",\"value\":\"string\"}",
			),
		},
		Required: true,
	}
	ma := metadata.MethodAttribute{}

	if err := ma.UnmarshalJSON(raw); err != nil {
		t.Fatal("error: unmarshal json: ", err)
	}

	got, status, err := ma.Parameter()
	if err != nil {
		t.Fatal("error: unmarshal parameter: ", err)
	} else if !status {
		t.Fatal("error: type kind is not parameter")
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatal("error: parameter does not match")
	}
}

func TestMethodAttribute_ReturnValue(t *testing.T) {
	raw := []byte("{\"kind\":\"returnvalue\",\"name\":\"id\"," +
		"\"type\":{\"kind\":\"primitive\",\"value\":\"string\"}}",
	)
	want := &metadata.ReturnValue{
		Name: "id",
		Type: metadata.Type{
			Kind: "primitive",
			Value: json.RawMessage(
				"{\"kind\":\"primitive\",\"value\":\"string\"}",
			),
		},
	}
	ma := metadata.MethodAttribute{}

	if err := ma.UnmarshalJSON(raw); err != nil {
		t.Fatal("error: unmarshal json: ", err)
	}

	got, status, err := ma.ReturnValue()
	if err != nil {
		t.Fatal("error: unmarshal return value: ", err)
	} else if !status {
		t.Fatal("error: type kind is not return value")
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatal("error: return value does not match")
	}
}
