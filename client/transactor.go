/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package client

import (
    "errors"
    "fmt"
    "time"

    "github.com/katena-chain/sdk-go-client/api"
    "github.com/katena-chain/sdk-go-client/crypto/ED25519"
    "github.com/katena-chain/sdk-go-client/crypto/X25519"
    "github.com/katena-chain/sdk-go-client/entity"
    entityApi "github.com/katena-chain/sdk-go-client/entity/api"
    "github.com/katena-chain/sdk-go-client/entity/certify"
    "github.com/katena-chain/sdk-go-client/entity/client"
)

// Transactor provides helper function to hide the complexity of Transaction creation, signature and API dialog.
type Transactor struct {
    apiHandler     *api.Handler
    chainID        string
    msgSigner      *ED25519.PrivateKey
    companyChainID string
}

// Transactor constructor.
func NewTransactor(
    apiUrl string,
    chainID string,
    msgSigner *ED25519.PrivateKey,
    companyChainID string,
) *Transactor {

    return &Transactor{
        apiHandler:     api.NewHandler(apiUrl),
        chainID:        chainID,
        msgSigner:      msgSigner,
        companyChainID: companyChainID,
    }
}

// SendCertificateV1 wraps a CertificateV1 in a MsgCreateCertificate, creates a transaction and sends it to the API.
func (t Transactor) SendCertificateV1(
    uuid string,
    dataSignature []byte,
    dataSigner []byte,
) (*entityApi.TransactionStatus, error) {

    certificate := certify.NewCertificateV1(uuid, t.companyChainID, dataSignature, dataSigner)
    message := &certify.MsgCreateCertificate{
        Certificate: certificate,
    }

    transaction, err := t.getTransaction(message)
    if err != nil {
        return nil, err
    }

    return t.apiHandler.SendCertificate(transaction)
}

// RetrieveCertificateV1 fetches the API to find the corresponding transaction and converts its content to a
// CertificateV1 with its blockchain status.
func (t Transactor) RetrieveCertificateV1(uuid string) (*client.CertificateV1Wrapper, error) {
    transactionWrapper, err := t.apiHandler.RetrieveCertificate(t.companyChainID, uuid)
    if err != nil {
        return nil, err
    }
    if transactionWrapper.Transaction.Message.GetType() == certify.MsgCreateCertificateType {
        message := transactionWrapper.Transaction.Message.(*certify.MsgCreateCertificate)
        if message.Certificate.GetType() == certify.CertificateV1Type {
            return &client.CertificateV1Wrapper{
                Certificate: message.Certificate.(*certify.CertificateV1),
                Status:      transactionWrapper.Status,
            }, nil
        } else {
            return nil, errors.New(fmt.Sprintf("bad certificate type: %s", message.Certificate.GetType()))
        }
    } else {
        return nil, errors.New(
            fmt.Sprintf("bad message type: %s", transactionWrapper.Transaction.Message.GetType()),
        )
    }

}

// SendSecretV1 wraps a SecretV1 in a MsgCreateSecret, creates a transaction and sends it to the API.
func (t Transactor) SendSecretV1(
    certificateUuid string,
    lockEncryptor *X25519.PublicKey,
    lockNonce *X25519.NaclBoxNonce,
    lockContent []byte,
) (*entityApi.TransactionStatus, error) {

    secret := certify.NewSecretV1(
        certificateUuid,
        t.companyChainID,
        lockEncryptor,
        lockNonce,
        lockContent,
    )
    message := &certify.MsgCreateSecret{
        Secret: secret,
    }

    transaction, err := t.getTransaction(message)
    if err != nil {
        return nil, err
    }

    return t.apiHandler.SendSecret(transaction, t.companyChainID, certificateUuid)

}

// RetrieveSecretsV1 fetches the API to find the corresponding transactions and converts their content to a SecretV1
// with its blockchain status.
func (t Transactor) RetrieveSecretsV1(uuid string) (*client.SecretV1Wrappers, error) {
    transactionWrappers, err := t.apiHandler.RetrieveSecrets(t.companyChainID, uuid)
    if err != nil {
        return nil, err
    }
    var secretV1Wrappers client.SecretV1Wrappers
    for _, transactionWrapper := range transactionWrappers.Transactions {
        if transactionWrapper.Transaction.Message.GetType() == certify.MsgCreateSecretType {
            message := transactionWrapper.Transaction.Message.(*certify.MsgCreateSecret)
            if message.Secret.GetType() == certify.SecretV1Type {
                secretV1Wrappers.Secrets = append(secretV1Wrappers.Secrets, &client.SecretV1Wrapper{
                    Secret: message.Secret.(*certify.SecretV1),
                    Status: transactionWrapper.Status,
                })
            } else {
                return nil, errors.New(fmt.Sprintf("bad secret type: %s", message.Secret.GetType()))
            }
        } else {
            return nil, errors.New(
                fmt.Sprintf("bad message type: %s", transactionWrapper.Transaction.Message.GetType()),
            )
        }
    }
    return &secretV1Wrappers, nil
}

// getTransaction signs a message and returns a new transaction ready to be sent.
func (t Transactor) getTransaction(message entity.Message) (*entityApi.Transaction, error) {
    nonceTime := entity.Time{
        Time: time.Now(),
    }
    sealState := &entity.SealState{
        Message:   message,
        ChainID:   t.chainID,
        NonceTime: &nonceTime,
    }
    sealStateBytes, err := sealState.GetSignBytes()
    if err != nil {
        return nil, err
    }
    if t.msgSigner == nil {
        return nil, errors.New("impossible to create transactions without a private key")
    }
    msgSignature := t.msgSigner.Sign(sealStateBytes)

    return entityApi.NewTransaction(message, msgSignature, t.msgSigner.GetPublicKey(), &nonceTime), nil
}
