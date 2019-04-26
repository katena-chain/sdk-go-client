/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package entity

import (
    "github.com/katena-chain/sdk-go-client/crypto"
)

// Seal is a wrapper to an ED25519 signature and its corresponding ED25519 public key.
type Seal struct {
    Signature *crypto.SignatureED25519 `json:"signature"`
    Signer    *crypto.PublicKeyED25519 `json:"signer"`
}

// Seal constructor.
func NewSeal(signature *crypto.SignatureED25519, signer *crypto.PublicKeyED25519) *Seal {
    return &Seal{
        Signature: signature,
        Signer:    signer,
    }
}
