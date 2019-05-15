/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package X25519

import (
    "crypto/rand"
    "encoding/json"
    "io"

    "golang.org/x/crypto/nacl/box"
)

const PrivateKeySize = 32
const NaclBoxNonceSize = 24

// PrivateKey is an X25519 private key wrapper (32 bytes).
type PrivateKey [PrivateKeySize]byte

// PrivateKey constructor.
func NewPrivateKey(privateKeyBytes []byte) *PrivateKey {
    var privateKey PrivateKey
    copy(privateKey[:], privateKeyBytes[:])
    return &privateKey
}

// Seal encrypts a plain text message decipherable afterwards by the recipient public key.
func (pk *PrivateKey) Seal(message []byte, recipientPublicKey *PublicKey) (*NaclBoxNonce, []byte, error) {
    var nonce [24]byte
    if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
        return nil, nil, err
    }
    encryptedMessage := box.Seal(nil, message, &nonce, (*[32]byte)(recipientPublicKey), (*[32]byte)(pk))
    return (*NaclBoxNonce)(&nonce), encryptedMessage, nil
}

// Open decrypts an encrypted message with the appropriate sender information.
func (pk *PrivateKey) Open(encryptedMessage []byte, senderPublicKey *PublicKey, nonce *NaclBoxNonce) ([]byte, bool) {
    return box.Open(nil, encryptedMessage, (*[24]byte)(nonce), (*[32]byte)(senderPublicKey), (*[32]byte)(pk))
}

// NaclBoxNonce is a Nacl box nonce wrapper (24 bytes)
type NaclBoxNonce [NaclBoxNonceSize]byte

// MarshalJSON returns the base64 value of a nacl box nonce.
func (nbn NaclBoxNonce) MarshalJSON() ([]byte, error) {
    return json.Marshal(nbn[:])
}

// UnmarshalJSON accepts a base64 value to load a nacl box nonce.
func (nbn *NaclBoxNonce) UnmarshalJSON(data []byte) error {
    var bytes []byte
    if err := json.Unmarshal(data, &bytes); err != nil {
        return err
    }
    copy(nbn[:], bytes)
    return nil
}
