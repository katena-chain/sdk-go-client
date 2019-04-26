/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package api

import (
    entityApi "github.com/katena-chain/sdk-go-client/entity/api"
    "github.com/katena-chain/sdk-go-client/utils"
    utilsApi "github.com/katena-chain/sdk-go-client/utils/api"
)

const certificateRoute = "certificates/certify"

// Handler provides helper methods to send and retrieve transactions without directly interacting with the HTTP Client.
type Handler struct {
    apiClient *utilsApi.Client
}

// Handler constructor.
func NewHandler(apiUrl string, apiUrlSuffix string) (*Handler, error) {
    fullApiUrl, err := utils.GetUri(apiUrl, []string{apiUrlSuffix}, nil)
    if err != nil {
        return nil, err
    }
    return &Handler{
        apiClient: utilsApi.NewClient(fullApiUrl.String()),
    }, nil
}

// SendCertificate accepts a transaction and sends it to the appropriate API route.
func (h *Handler) SendCertificate(transaction *entityApi.Transaction) (*utilsApi.Response, error) {
    return h.apiClient.Post(certificateRoute, nil, transaction)
}
