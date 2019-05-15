/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package entity

import (
    "github.com/katena-chain/sdk-go-client/crypto/ED25519"
)

// Seal is a wrapper to an ED25519 signature and its corresponding ED25519 public key.
type Seal struct {
    Signature *ED25519.Signature `json:"signature"`
    Signer    *ED25519.PublicKey `json:"signer"`
}
