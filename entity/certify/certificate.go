/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package certify

import (
    "encoding/json"

    "github.com/katena-chain/sdk-go-client/utils"
)

// Certificate sets the default methods a real certificate must implement.
type Certificate interface {
    GetUuid() string
    GetCompanyChainID() string
    GetType() string
}

// dataSeal is a wrapper to a raw signature (16 < x < 128 bytes) and its corresponding raw signer (16 < x < 128 bytes).
type dataSeal struct {
    Signature []byte `json:"signature"`
    Signer    []byte `json:"signer"`
}

// CertificateV1 is the first version of a certificate to send in a transaction's message.
// It should implement the Certificate interface.
type CertificateV1 struct {
    Uuid           string    `json:"uuid"`
    CompanyChainID string    `json:"company_chain_id"`
    Seal           *dataSeal `json:"seal"`
}

// CertificateV1 constructor.
func NewCertificateV1(uuid string, companyChainID string, dataSignature []byte, dataSigner []byte) *CertificateV1 {
    return &CertificateV1{
        Uuid:           uuid,
        CompanyChainID: companyChainID,
        Seal: &dataSeal{
            Signature: dataSignature,
            Signer:    dataSigner,
        },
    }
}

// GetUuid returns the uuid value (Certificate interface requirement).
func (c CertificateV1) GetUuid() string {
    return c.Uuid
}

// GetCompanyChainID returns the company chain id value (Certificate interface requirement).
func (c CertificateV1) GetCompanyChainID() string {
    return c.CompanyChainID
}

// GetType returns the type value (Certificate interface requirement).
func (c CertificateV1) GetType() string {
    return CertificateV1Type
}

// MarshalJSON converts a CertificateV1 into a jsonCertificateV1 and wraps
// it in a JSONWrapper to indicate which version of a certificate should be
// unmarshaled back.
func (c CertificateV1) MarshalJSON() ([]byte, error) {
    type jsonAlias CertificateV1
    value, _ := json.Marshal((*jsonAlias)(&c))
    return json.Marshal(&utils.JSONWrapper{
        Type:  c.GetType(),
        Value: value,
    })
}
