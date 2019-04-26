/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package api

import (
    "github.com/katena-chain/sdk-go-client/crypto"
    "github.com/katena-chain/sdk-go-client/entity"
)

// Transaction wraps a message, its signature infos and the nonce time used to sign the message.
type Transaction struct {
    Message   entity.Message `json:"message"`
    Seal      *entity.Seal   `json:"seal"`
    NonceTime *entity.Time   `json:"nonce_time"`
}

// NewTransaction constructor.
func NewTransaction(
    message entity.Message,
    msgSignature *crypto.SignatureED25519,
    msgSigner *crypto.PublicKeyED25519,
    nonceTime *entity.Time,
) *Transaction {
    return &Transaction{
        Message:   message,
        Seal:      entity.NewSeal(msgSignature, msgSigner),
        NonceTime: nonceTime,
    }
}
