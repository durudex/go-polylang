/*
 * Copyright © 2023 Durudex
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package metadata

import (
	"encoding/json"
	"os"
)

func Parse(data []byte) (*Root, error) {
	var root Root
	if err := json.Unmarshal(data, &root); err != nil {
		return nil, err
	}

	return &root, nil
}

func ParseFile(path string) (*Root, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var root Root
	if err := json.NewDecoder(f).Decode(&root); err != nil {
		return nil, err
	}

	return &root, nil
}
