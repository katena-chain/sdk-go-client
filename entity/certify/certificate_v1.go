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

// dataSeal is a wrapper to a raw signature (16 < x < 128 bytes) and its corresponding raw signer (16 < x < 128 bytes).
type dataSeal struct {
    Signature []byte `json:"signature"`
    Signer    []byte `json:"signer"`
}

// CertificateV1 is the first version of a data certificate to send in a transaction's message.
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
    return certificateV1Type
}

// jsonCertificateV1 wraps a CertificateV1 for its json marshaled value.
type jsonCertificateV1 CertificateV1

// MarshalJSON converts a CertificateV1 into a jsonCertificateV1 and wraps
// it in a JSONWrapper to indicate which version of a certificate should be
// unmarshaled back.
func (c CertificateV1) MarshalJSON() ([]byte, error) {
    return json.Marshal(&utils.JSONWrapper{
        Type:  c.GetType(),
        Value: (*jsonCertificateV1)(&c),
    })
}

// UnmarshalJSON converts a jsonCertificateV1 wrapped in a JSONWrapper into a CertificateV1.
func (c *CertificateV1) UnmarshalJSON(data []byte) error {
    jsonValue := &utils.JSONWrapper{
        Value: (*jsonCertificateV1)(c),
    }
    if err := json.Unmarshal(data, jsonValue); err != nil {
        return err
    }
    value := (jsonValue.Value).(*jsonCertificateV1)
    c.CompanyChainID = value.CompanyChainID
    c.Uuid = value.Uuid
    c.Seal = value.Seal
    return nil
}
