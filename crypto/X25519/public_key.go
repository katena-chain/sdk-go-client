/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package X25519

import (
    "encoding/json"
)

const PublicKeySize = 32

// PublicKey is an X25519 public key wrapper (32 bytes).
type PublicKey [PublicKeySize]byte

// PublicKey constructor.
func NewPublicKey(publicKeyBytes []byte) *PublicKey {
    var publicKey PublicKey
    copy(publicKey[:], publicKeyBytes[:])
    return &publicKey
}

// MarshalJSON returns the base64 value of an X25519 public key.
func (pk PublicKey) MarshalJSON() ([]byte, error) {
    return json.Marshal(pk[:])
}

// UnmarshalJSON accepts a base64 value to load an X25519 public key.
func (pk *PublicKey) UnmarshalJSON(data []byte) error {
    var bytes []byte
    if err := json.Unmarshal(data, &bytes); err != nil {
        return err
    }
    copy(pk[:], bytes)
    return nil
}
