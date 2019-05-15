/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package errors

import (
    "fmt"
)

// ApiError allows to wrap API errors.
type ApiError struct {
    Code    uint32 `json:"code"`
    Message string `json:"message"`
}

// Error the ApiError formatted as a string (error interface requirement).
func (e ApiError) Error() string {
    return fmt.Sprintf(`api error:
  Code    : %d
  Message : %s`, e.Code, e.Message)
}
