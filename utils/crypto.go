/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package utils

import (
    "crypto/rand"
    "encoding/base64"

    "golang.org/x/crypto/nacl/box"

    "github.com/katena-chain/sdk-go-client/crypto/ED25519"
    "github.com/katena-chain/sdk-go-client/crypto/X25519"
)

// CreatePrivateKeyED25519FromBase64 accepts a base64 encoded ED25519 private key (88 chars) and returns an ED25519
// private key.
func CreatePrivateKeyED25519FromBase64(privateKeyBase64 string) (*ED25519.PrivateKey, error) {
    privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKeyBase64)
    if err != nil {
        return nil, err
    }
    return ED25519.NewPrivateKey(privateKeyBytes), nil
}

// CreatePublicKeyX25519FromBase64 accepts a base64 encoded X25519 public key (44 chars) and returns an X25519 public
// key.
func CreatePublicKeyX25519FromBase64(publicKeyBase64 string) (*X25519.PublicKey, error) {
    publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKeyBase64)
    if err != nil {
        return nil, err
    }
    return X25519.NewPublicKey(publicKeyBytes), nil
}

// CreatePrivateKeyX25519FromBase64 accepts a base64 encoded X25519 private key (44 chars) and returns an X25519
// private key.
func CreatePrivateKeyX25519FromBase64(privateKeyBase64 string) (*X25519.PrivateKey, error) {
    privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKeyBase64)
    if err != nil {
        return nil, err
    }
    return X25519.NewPrivateKey(privateKeyBytes), nil
}

// CreateNewKeysX25519 returns a new X25519 key pair.
func CreateNewKeysX25519() (*X25519.PublicKey, *X25519.PrivateKey, error) {
    publicKey, privateKey, err := box.GenerateKey(rand.Reader)
    return (*X25519.PublicKey)(publicKey), (*X25519.PrivateKey)(privateKey), err
}
