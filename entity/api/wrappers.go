/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package api

// TransactionStatus is the blockchain status of a transaction.
type TransactionStatus struct {
    Code    uint32 `json:"code"`
    Message string `json:"message"`
}

// TransactionWrappers wraps a list of TransactionWrapper with the total transactions available in the blockchain.
// The API by default, will only returns 10 transactions.
type TransactionWrappers struct {
    Transactions []*TransactionWrapper `json:"transactions"`
    Total        int                   `json:"total"`
}

// TransactionWrapper wraps a Transaction with its blockchain status.
type TransactionWrapper struct {
    Transaction *Transaction       `json:"transaction"`
    Status      *TransactionStatus `json:"status"`
}
