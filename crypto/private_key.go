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

// PrivateKeyED25519 is an ED25519 private key wrapper (64 bytes).
type PrivateKeyED25519 [ed25519.PrivateKeySize]byte

// PrivateKeyED25519 constructor.
func NewPrivateKeyED25519(privateKeyBytes []byte) *PrivateKeyED25519 {
    var privateKey PrivateKeyED25519
    copy(privateKey[:], privateKeyBytes[:])
    return &privateKey
}

// Sign accepts a message and returns its corresponding signature.
func (pk *PrivateKeyED25519) Sign(message []byte) *SignatureED25519 {
    var signature SignatureED25519
    copy(signature[:], ed25519.Sign(pk[:], message)[:])
    return &signature
}

// GetPublicKey returns the underlying ED25519 public key.
func (pk *PrivateKeyED25519) GetPublicKey() *PublicKeyED25519 {
    var publicKeyBytes PublicKeyED25519
    copy(publicKeyBytes[:], pk[32:])
    return &publicKeyBytes
}

// SignatureED25519 is an ED25519 private key wrapper (64 bytes).
type SignatureED25519 [ed25519.SignatureSize]byte

// MarshalJSON returns the base64 value of an ED25519 signature.
func (s SignatureED25519) MarshalJSON() ([]byte, error) {
    return json.Marshal(s[:])
}

// UnmarshalJSON accepts a base64 value to load an ED25519 signature.
func (s *SignatureED25519) UnmarshalJSON(data []byte) error {
    var bytes []byte
    if err := json.Unmarshal(data, &bytes); err != nil {
        return err
    }
    copy(s[:], bytes)
    return nil
}
