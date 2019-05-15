/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package api

import (
    "encoding/json"
    "errors"
    "fmt"

    "github.com/katena-chain/sdk-go-client/crypto/ED25519"
    "github.com/katena-chain/sdk-go-client/entity"
    "github.com/katena-chain/sdk-go-client/entity/certify"
    "github.com/katena-chain/sdk-go-client/utils"
)

var messagesFactory = map[string]func() entity.Message{
    certify.MsgCreateCertificateType: func() entity.Message { return &certify.MsgCreateCertificate{} },
    certify.MsgCreateSecretType:      func() entity.Message { return &certify.MsgCreateSecret{} },
}

// Transaction wraps a message, its signature infos and the nonce time used to sign the message.
type Transaction struct {
    Message   entity.Message `json:"message"`
    Seal      *entity.Seal   `json:"seal"`
    NonceTime *entity.Time   `json:"nonce_time"`
}

// NewTransaction constructor.
func NewTransaction(
    message entity.Message,
    msgSignature *ED25519.Signature,
    msgSigner *ED25519.PublicKey,
    nonceTime *entity.Time,
) *Transaction {
    return &Transaction{
        Message: message,
        Seal: &entity.Seal{
            Signature: msgSignature,
            Signer:    msgSigner,
        },
        NonceTime: nonceTime,
    }
}

// UnmarshalJSON converts a byte array a Transaction.
// It handles the Message interface conversion as well.
func (t *Transaction) UnmarshalJSON(data []byte) error {
    type JsonAlias Transaction
    var jsonTransaction struct {
        *JsonAlias
        Message *utils.JSONWrapper `json:"message"`
    }

    if err := json.Unmarshal(data, &jsonTransaction); err != nil {
        return err
    }
    if typeFactory, ok := messagesFactory[jsonTransaction.Message.Type]; ok {
        concreteValue := typeFactory()
        if err := json.Unmarshal(jsonTransaction.Message.Value, concreteValue); err != nil {
            return err
        }
        t.Message = concreteValue
    } else {
        return errors.New(fmt.Sprintf("unknown message type: %s", jsonTransaction.Message.Type))
    }
    t.Seal = jsonTransaction.Seal
    t.NonceTime = jsonTransaction.NonceTime
    return nil
}
