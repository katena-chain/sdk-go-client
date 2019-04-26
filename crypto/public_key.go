/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package crypto

import (
    "encoding/json"

    "golang.org/x/crypto/ed25519"
)

// PublicKeyED25519 is an ED25519 public key wrapper (32 bytes).
type PublicKeyED25519 [ed25519.PublicKeySize]byte

// Verify indicates if a message and a signature match.
func (pk PublicKeyED25519) Verify(message []byte, signature []byte) bool {
    return ed25519.Verify(pk[:], message, signature)
}

// MarshalJSON returns the base64 value of an ED25519 public key.
func (pk PublicKeyED25519) MarshalJSON() ([]byte, error) {
    return json.Marshal(pk[:])
}

// UnmarshalJSON accepts a base64 value to load a ED25519 public key.
func (pk *PublicKeyED25519) UnmarshalJSON(data []byte) error {
    var bytes []byte
    if err := json.Unmarshal(data, &bytes); err != nil {
        return err
    }
    copy(pk[:], bytes)
    return nil
}
