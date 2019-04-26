/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package certify

// Certificate sets the default methods a real certificate must implement.
type Certificate interface {
    GetUuid() string
    GetCompanyChainID() string
    GetType() string
}
