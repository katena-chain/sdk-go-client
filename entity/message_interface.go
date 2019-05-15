/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package entity

// Message sets the default methods a real message must implement.
type Message interface {
    GetType() string
}
