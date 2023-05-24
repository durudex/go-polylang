/*
 * Copyright © 2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package metadata

import "encoding/json"

type Order string

type IndexField struct {
	Direction Order    `json:"direction"`
	FieldPath []string `json:"fieldPath"`
}

type Index struct {
	Fields []IndexField `json:"fields"`
}

func (ca CollectionAttribute) Index() (*Index, bool, error) {
	if ca.Kind != "index" {
		return nil, false, nil
	}

	var idx Index
	if err := json.Unmarshal(ca.Value, &idx); err != nil {
		return nil, true, err
	}

	return &idx, true, nil
}
