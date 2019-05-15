/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package certify

import (
    "encoding/json"

    "github.com/katena-chain/sdk-go-client/crypto/X25519"
    "github.com/katena-chain/sdk-go-client/utils"
)

// Secret sets the default methods a real secret must implement.
type Secret interface {
    GetCertificateUuid() string
    GetCompanyChainID() string
    GetType() string
}

// lock is a wrapper to an X25519 encryptor (32 bytes), its corresponding nonce (24 bytes) and the raw encrypted content
// (16 < x < 128 bytes) to perform an ECDH shared key agreement.
type lock struct {
    Encryptor *X25519.PublicKey    `json:"encryptor"`
    Nonce     *X25519.NaclBoxNonce `json:"nonce"`
    Content   []byte               `json:"content"`
}

// SecretV1 is the first version of a secret to send in a transaction's message.
// It should implement the Secret interface.
type SecretV1 struct {
    CertificateUuid string `json:"certificate_uuid"`
    CompanyChainID  string `json:"company_chain_id"`
    Lock            *lock  `json:"lock"`
}

// SecretV1 constructor.
func NewSecretV1(
    certificateUuid string,
    companyChainID string,
    lockEncryptor *X25519.PublicKey,
    lockNonce *X25519.NaclBoxNonce,
    lockContent []byte,
) *SecretV1 {

    return &SecretV1{
        CertificateUuid: certificateUuid,
        CompanyChainID:  companyChainID,
        Lock: &lock{
            Encryptor: lockEncryptor,
            Nonce:     lockNonce,
            Content:   lockContent,
        },
    }
}

// GetCertificateUuid returns the certificate uuid value (Secret interface requirement).
func (c SecretV1) GetCertificateUuid() string {
    return c.CertificateUuid
}

// GetCompanyChainID returns the company chain id value (Secret interface requirement).
func (c SecretV1) GetCompanyChainID() string {
    return c.CompanyChainID
}

// GetType returns the type value (Secret interface requirement).
func (c SecretV1) GetType() string {
    return SecretV1Type
}

// MarshalJSON converts a SecretV1 into a jsonSecretV1 and wraps
// it in a JSONWrapper to indicate which version of a secret should be
// unmarshaled back.
func (c SecretV1) MarshalJSON() ([]byte, error) {
    type jsonAlias SecretV1
    value, _ := json.Marshal((*jsonAlias)(&c))
    return json.Marshal(&utils.JSONWrapper{
        Type:  c.GetType(),
        Value: value,
    })
}
