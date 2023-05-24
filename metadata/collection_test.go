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

var NamespaceTests = map[string]struct {
	raw  []byte
	data *metadata.Namespace
}{
	"OK": {
		raw:  []byte("{\"kind\":\"namespace\",\"value\":\"hello\"}"),
		data: &metadata.Namespace{Value: "hello"},
	},
}

func TestNamespace_UnmarshalJSON(t *testing.T) {
	for name, test := range NamespaceTests {
		t.Run(name, func(t *testing.T) {
			ns := &metadata.Namespace{}
			if err := ns.UnmarshalJSON(test.raw); err != nil {
				t.Fatal("error: unmarshal json: ", err)
			}

			if !reflect.DeepEqual(ns, test.data) {
				t.Fatal("error: type does not match")
			}
		})
	}
}

func TestNamespace_MarshalJSON(t *testing.T) {
	for name, test := range NamespaceTests {
		t.Run(name, func(t *testing.T) {
			got, err := test.data.MarshalJSON()
			if err != nil {
				t.Fatal("error: marshal json: ", err)
			}

			if !reflect.DeepEqual(got, test.raw) {
				t.Fatal("error: type does not match")
			}
		})
	}
}

func TestCollectionAttribute_Property(t *testing.T) {
	raw := []byte("{\"kind\":\"property\",\"name\":\"test\"," +
		"\"type\":{\"kind\":\"primitive\",\"value\":\"string\"}," +
		"\"directives\":[],\"required\":true}",
	)
	want := &metadata.Property{
		Name: "test",
		Type: metadata.Type{
			Kind: "primitive",
			Value: json.RawMessage(
				"{\"kind\":\"primitive\",\"value\":\"string\"}",
			),
		},
		Directives: []metadata.Directive{},
		Required:   true,
	}
	ca := metadata.CollectionAttribute{}

	if err := ca.UnmarshalJSON(raw); err != nil {
		t.Fatal("error: unmarshal json: ", err)
	}

	got, status, err := ca.Property()
	if err != nil {
		t.Fatal("error: unmarshal property: ", err)
	} else if !status {
		t.Fatal("error: type kind is not property")
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatal("error: property does not match")
	}
}
