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

func TestCollectionAttribute_Directive(t *testing.T) {
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
	ca := metadata.CollectionAttribute{}

	if err := ca.UnmarshalJSON(raw); err != nil {
		t.Fatal("error: unmarshal json: ", err)
	}

	got, status, err := ca.Directive()
	if err != nil {
		t.Fatal("error: unmarshal directive: ", err)
	} else if !status {
		t.Fatal("error: type kind is not directive")
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatal("error: directive does not match")
	}
}

func TestDirectiveArgument_FieldReference(t *testing.T) {
	raw := []byte("{\"kind\":\"fieldreference\",\"path\":[\"owner\"]}")
	want := &metadata.FieldReference{
		Path: []string{"owner"},
	}
	da := metadata.DirectiveArgument{}

	if err := da.UnmarshalJSON(raw); err != nil {
		t.Fatal("error: unmarshal json: ", err)
	}

	got, status, err := da.FieldReference()
	if err != nil {
		t.Fatal("error: unmarshal field reference: ", err)
	} else if !status {
		t.Fatal("error: type kind is not field reference")
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatal("error: field reference does not match")
	}
}
