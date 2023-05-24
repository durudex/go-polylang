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

func TestCollectionAttribute_Index(t *testing.T) {
	raw := []byte("{\"kind\":\"index\",\"fields\":[{" +
		"\"direction\":\"asc\",\"fieldPath\":[" +
		"\"firstName\",\"lastName\"]}]}",
	)
	want := &metadata.Index{Fields: []metadata.IndexField{
		{
			Direction: "asc",
			FieldPath: []string{"firstName", "lastName"},
		},
	}}
	ca := metadata.CollectionAttribute{}

	if err := ca.UnmarshalJSON(raw); err != nil {
		t.Fatal("error: unmarshal json: ", err)
	}

	got, status, err := ca.Index()
	if err != nil {
		t.Fatal("error: unmarshal index: ", err)
	} else if !status {
		t.Fatal("error: type kind is not index")
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatal("error: index does not match")
	}
}
