package main

import (
    "fmt"

    "github.com/katena-chain/sdk-go-client/client"
)

func main() {

    // Common Katena network informations
    apiUrl := "https://api.test.katena.transchain.io/api/v1"

    // Your Katena network informations
    companyChainID := "abcdef"

    // Create a Katena API helper
    transactor := client.NewTransactor(apiUrl, "", nil, companyChainID)

    // Certificate uuid you want to retrieve
    certificateUuid := "2075c941-6876-405b-87d5-13791c0dc53a"

    // Retrieve a version 1 of a certificate from Katena blockchain
    certificateV1Wrapper, err := transactor.RetrieveCertificateV1(certificateUuid)
    if err != nil {
        panic(err)
    }

    fmt.Println("Transaction status")
    fmt.Println(fmt.Sprintf("  Code    : %d", certificateV1Wrapper.Status.Code))
    fmt.Println(fmt.Sprintf("  Message : %s", certificateV1Wrapper.Status.Message))

    fmt.Println("CertificateV1")
    fmt.Println(fmt.Sprintf("  Uuid             : %s", certificateV1Wrapper.Certificate.Uuid))
    fmt.Println(fmt.Sprintf("  Company chain id : %s", certificateV1Wrapper.Certificate.CompanyChainID))
    fmt.Println(fmt.Sprintf("  Data signer      : %s", certificateV1Wrapper.Certificate.Seal.Signer))
    fmt.Println(fmt.Sprintf("  Data signature   : %s", certificateV1Wrapper.Certificate.Seal.Signature))

}
