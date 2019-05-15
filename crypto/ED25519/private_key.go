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

// PrivateKey is an ED25519 private key wrapper (64 bytes).
type PrivateKey [ed25519.PrivateKeySize]byte

// PrivateKey constructor.
func NewPrivateKey(privateKeyBytes []byte) *PrivateKey {
    var privateKey PrivateKey
    copy(privateKey[:], privateKeyBytes[:])
    return &privateKey
}

// Sign accepts a message and returns its corresponding ED25519 signature.
func (pk *PrivateKey) Sign(message []byte) *Signature {
    var signature Signature
    copy(signature[:], ed25519.Sign(pk[:], message)[:])
    return &signature
}

// GetPublicKey returns the underlying ED25519 public key.
func (pk *PrivateKey) GetPublicKey() *PublicKey {
    var publicKeyBytes PublicKey
    copy(publicKeyBytes[:], pk[32:])
    return &publicKeyBytes
}

// Signature is an ED25519 signature wrapper (64 bytes).
type Signature [ed25519.SignatureSize]byte

// MarshalJSON returns the base64 value of an ED25519 signature.
func (s Signature) MarshalJSON() ([]byte, error) {
    return json.Marshal(s[:])
}

// UnmarshalJSON accepts a base64 value to load an ED25519 signature.
func (s *Signature) UnmarshalJSON(data []byte) error {
    var bytes []byte
    if err := json.Unmarshal(data, &bytes); err != nil {
        return err
    }
    copy(s[:], bytes)
    return nil
}
