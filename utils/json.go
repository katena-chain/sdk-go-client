/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package utils

import (
    "encoding/json"
)

// JSONWrapper wraps an interface with its corresponding custom type.
type JSONWrapper struct {
    Value json.RawMessage `json:"value"`
    Type  string          `json:"type"`
}

// MarshalAndSortJSON sorts alphabetically the json representation of an interface and returns its marshaled value.
func MarshalAndSortJSON(jsonValue interface{}) ([]byte, error) {
    jsonBytes, err := json.Marshal(jsonValue)
    if err != nil {
        return nil, err
    }
    // The json package sorts by structure's fields order.
    // Marshal and Unmarshal back to an interface do the trick.
    var sortedJsonValue interface{}
    err = json.Unmarshal(jsonBytes, &sortedJsonValue)
    if err != nil {
        return nil, err
    }
    return json.Marshal(sortedJsonValue)
}
