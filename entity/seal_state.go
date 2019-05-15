/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package entity

import (
    "github.com/katena-chain/sdk-go-client/utils"
)

// SealState wraps a message and additional values in order to define the unique message state to be signed.
type SealState struct {
    Message   Message `json:"message"`
    ChainID   string  `json:"chain_id"`
    NonceTime *Time   `json:"nonce_time"`
}

// GetSignBytes returns the sorted and marshaled values of a seal state.
func (ss *SealState) GetSignBytes() ([]byte, error) {
    return utils.MarshalAndSortJSON(ss)
}
