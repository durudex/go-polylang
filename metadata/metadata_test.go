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

var ParseTests = map[string]struct {
	data []byte
	want *metadata.Root
}{
	"OK": {
		data: []byte("[{\"kind\":\"collection\",\"value\":{}}]"),
		want: &metadata.Root{
			{
				Kind:  "collection",
				Value: json.RawMessage("{\"kind\":\"collection\",\"value\":{}}"),
			},
		},
	},
}

func TestParse(t *testing.T) {
	for name, test := range ParseTests {
		t.Run(name, func(t *testing.T) {
			got, err := metadata.Parse(test.data)
			if err != nil {
				t.Fatal("error: parsing metadata: ", err)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Fatal("error: root does not match")
			}
		})
	}
}
