/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package utils

import (
    "encoding/base64"

    "github.com/katena-chain/sdk-go-client/crypto"
)

// CreatePrivateKeyED25519FromBase64 accepts a base64 encoded ed25519 private key (88 chars) and returns an ED25519
// private key.
func CreatePrivateKeyED25519FromBase64(privateKeyBase64 string) (*crypto.PrivateKeyED25519, error) {
    privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKeyBase64)
    if err != nil {
        return nil, err
    }
    return crypto.NewPrivateKeyED25519(privateKeyBytes), nil
}
