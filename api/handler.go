/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package api

import (
    "encoding/json"
    "fmt"

    "github.com/valyala/fasthttp"

    entityApi "github.com/katena-chain/sdk-go-client/entity/api"
    "github.com/katena-chain/sdk-go-client/errors"
)

const certificatesRoute = "certificates"
const certificateRoute = certificatesRoute + "/%s-%s"
const certificateCertifyRoute = certificatesRoute + "/certify"
const secretsRoute = certificateRoute + "/secrets"
const secretCertifyRoute = secretsRoute + "/certify"

// Handler provides helper methods to send and retrieve transactions without directly interacting with the HTTP Client.
type Handler struct {
    apiClient *Client
}

// Handler constructor.
func NewHandler(apiUrl string) *Handler {
    return &Handler{
        apiClient: NewClient(apiUrl),
    }
}

// SendCertificate accepts a transaction and sends it to the appropriate certificate API route.
func (h *Handler) SendCertificate(transaction *entityApi.Transaction) (*entityApi.TransactionStatus, error) {
    return h.SendTransaction(certificateCertifyRoute, transaction)
}

// SendSecret accepts a transaction and sends it to the appropriate secret API route.
func (h *Handler) SendSecret(
    transaction *entityApi.Transaction,
    companyChainID string,
    certificateUuid string,
) (*entityApi.TransactionStatus, error) {

    return h.SendTransaction(fmt.Sprintf(secretCertifyRoute, companyChainID, certificateUuid), transaction)
}

// RetrieveCertificate fetches the API and returns a transaction wrapper or an error.
func (h *Handler) RetrieveCertificate(companyChainID string, uuid string) (*entityApi.TransactionWrapper, error) {
    apiResponse, err := h.apiClient.Get(fmt.Sprintf(certificateRoute, companyChainID, uuid), nil)
    if err != nil {
        return nil, err
    }
    var transactionWrapper entityApi.TransactionWrapper
    if err := unmarshalApiResponse(apiResponse, &transactionWrapper); err != nil {
        return nil, err
    }
    return &transactionWrapper, nil
}

// RetrieveSecrets fetches the API and returns a transaction wrapper list or an error.
func (h *Handler) RetrieveSecrets(
    companyChainID string,
    certificateUuid string,
) (*entityApi.TransactionWrappers, error) {

    apiResponse, err := h.apiClient.Get(fmt.Sprintf(secretsRoute, companyChainID, certificateUuid), nil)
    if err != nil {
        return nil, err
    }
    var transactionWrappers entityApi.TransactionWrappers
    if err := unmarshalApiResponse(apiResponse, &transactionWrappers); err != nil {
        return nil, err
    }
    return &transactionWrappers, nil
}

// SendTransaction tries to send a transaction to the API and returns a transaction status or an error.
func (h *Handler) SendTransaction(
    route string,
    transaction *entityApi.Transaction,
) (*entityApi.TransactionStatus, error) {

    apiResponse, err := h.apiClient.Post(route, nil, transaction)
    if err != nil {
        return nil, err
    }
    var transactionStatus entityApi.TransactionStatus
    if err := unmarshalApiResponse(apiResponse, &transactionStatus); err != nil {
        return nil, err
    }
    return &transactionStatus, nil
}

// unmarshalApiResponse tries to parse the api response body if the API returns a 2xx HTTP code.
func unmarshalApiResponse(apiResponse *RawResponse, dest interface{}) error {
    if apiResponse.StatusCode == fasthttp.StatusOK || apiResponse.StatusCode == fasthttp.StatusAccepted {
        if err := json.Unmarshal(apiResponse.Body, dest); err != nil {
            return err
        }
        return nil
    } else {
        var apiError errors.ApiError
        if err := json.Unmarshal(apiResponse.Body, &apiError); err != nil {
            return err
        }
        return apiError
    }
}
