/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package entity

import (
    "encoding/json"
    "time"
)

// Time is a time.Time wrapper.
type Time time.Time

// MarshalJSON converts a Time into a time.Time and marshals its UTC value.
func (t Time) MarshalJSON() ([]byte, error) {
    return json.Marshal((*time.Time)(&t).Round(0).UTC())
}
