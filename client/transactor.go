/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package client

import (
    "time"

    "github.com/katena-chain/sdk-go-client/api"
    "github.com/katena-chain/sdk-go-client/crypto"
    "github.com/katena-chain/sdk-go-client/entity"
    entityApi "github.com/katena-chain/sdk-go-client/entity/api"
    "github.com/katena-chain/sdk-go-client/entity/certify"
    "github.com/katena-chain/sdk-go-client/utils"
    utilsApi "github.com/katena-chain/sdk-go-client/utils/api"
)

// Transactor provides helper function to hide the complexity of Transaction creation, signature and API dialog.
type Transactor struct {
    apiHandler     *api.Handler
    chainID        string
    privateKey     *crypto.PrivateKeyED25519
    companyChainID string
}

// Transactor constructor.
func NewTransactor(
    apiUrl string,
    apiUrlSuffix string,
    chainID string,
    privateKeyBase64 string,
    companyChainID string,
) (*Transactor, error) {
    privateKey, err := utils.CreatePrivateKeyED25519FromBase64(privateKeyBase64)
    if err != nil {
        return nil, err
    }
    apiHandler, err := api.NewHandler(apiUrl, apiUrlSuffix)
    if err != nil {
        return nil, err
    }
    return &Transactor{
        apiHandler:     apiHandler,
        chainID:        chainID,
        privateKey:     privateKey,
        companyChainID: companyChainID,
    }, nil
}

// SendCertificate creates a CertificateV1 wrapped in a MsgCreateCertificate, signs it and sends it to the API.
func (t Transactor) SendCertificate(uuid string, dataSignature string, dataSigner string) (*utilsApi.Response, error) {
    certificate := certify.NewCertificateV1(uuid, t.companyChainID, []byte(dataSignature), []byte(dataSigner))
    nonceTime := entity.Time(time.Now())
    message := certify.NewMsgCreateCertificate(certificate)

    sealState := entity.NewSealState(message, t.chainID, &nonceTime)
    sealStateBytes, err := sealState.GetSignBytes()
    if err != nil {
        return nil, err
    }
    messageSignature := t.privateKey.Sign(sealStateBytes)

    transaction := entityApi.NewTransaction(message, messageSignature, t.privateKey.GetPublicKey(), &nonceTime)
    apiResponse, err := t.apiHandler.SendCertificate(transaction)
    if err != nil {
        return nil, err
    }
    return apiResponse, nil
}
