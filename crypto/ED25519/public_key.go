/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package ED25519

import (
    "encoding/json"

    "golang.org/x/crypto/ed25519"
)

// PublicKey is an ED25519 public key wrapper (32 bytes).
type PublicKey [ed25519.PublicKeySize]byte

// Verify indicates if a message and a signature match.
func (pk PublicKey) Verify(message []byte, signature *Signature) bool {
    return ed25519.Verify(pk[:], message, signature[:])
}

// MarshalJSON returns the base64 value of an ED25519 public key.
func (pk PublicKey) MarshalJSON() ([]byte, error) {
    return json.Marshal(pk[:])
}

// UnmarshalJSON accepts a base64 value to load a ED25519 public key.
func (pk *PublicKey) UnmarshalJSON(data []byte) error {
    var bytes []byte
    if err := json.Unmarshal(data, &bytes); err != nil {
        return err
    }
    copy(pk[:], bytes)
    return nil
}
