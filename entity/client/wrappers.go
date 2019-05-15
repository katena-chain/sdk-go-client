/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package client

import (
    "github.com/katena-chain/sdk-go-client/entity/api"
    "github.com/katena-chain/sdk-go-client/entity/certify"
)

// CertificateV1Wrapper wraps a CertificateV1 with its blockchain status.
type CertificateV1Wrapper struct {
    Certificate *certify.CertificateV1
    Status      *api.TransactionStatus
}

// SecretV1Wrappers wraps a list of SecretV1Wrapper with the total transactions available in the blockchain.
// The API by default, will only returns 10 transactions.
type SecretV1Wrappers struct {
    Secrets []*SecretV1Wrapper
    Total   int
}

// CertificateV1Wrapper wraps a CertificateV1 with its blockchain status.
type SecretV1Wrapper struct {
    Secret *certify.SecretV1
    Status *api.TransactionStatus
}
