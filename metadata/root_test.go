/*
 * Copyright © 2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package metadata_test

import (
	"reflect"
	"testing"

	"github.com/durudex/go-polylang/metadata"
)

var AnyKindTests = map[string]struct {
	raw  []byte
	data *metadata.AnyKind
}{
	"OK": {
		raw: []byte("{\"kind\":\"string\",\"value\":\"hello\"}"),
		data: &metadata.AnyKind{
			Kind:  "string",
			Value: []byte("{\"kind\":\"string\",\"value\":\"hello\"}"),
		},
	},
}

func TestAnyKind_UnmarshalJSON(t *testing.T) {
	for name, test := range AnyKindTests {
		t.Run(name, func(t *testing.T) {
			ak := &metadata.AnyKind{}
			if err := ak.UnmarshalJSON(test.raw); err != nil {
				t.Fatal("error: unmarshal json: ", err)
			}

			if !reflect.DeepEqual(ak, test.data) {
				t.Fatal("error: type does not match")
			}
		})
	}
}

func TestAnyKind_MarshalJSON(t *testing.T) {
	for name, test := range AnyKindTests {
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

func TestNode_Collection(t *testing.T) {
	raw := []byte("{\"kind\":\"collection\",\"value\":{}}")
	want := &metadata.Collection{}
	node := metadata.Node{}

	if err := node.UnmarshalJSON(raw); err != nil {
		t.Fatal("error: unmarshal json: ", err)
	}

	got, status, err := node.Collection()
	if err != nil {
		t.Fatal("error: unmarshal collection: ", err)
	} else if !status {
		t.Fatal("error: node is not collection:")
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatal("error: collection does not match")
	}
}
