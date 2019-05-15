package main

import (
    "fmt"

    "github.com/katena-chain/sdk-go-client/client"
    "github.com/katena-chain/sdk-go-client/utils"
)

func main() {

    // Common Katena network information
    apiUrl := "https://api.test.katena.transchain.io/api/v1"
    chainID := "katena-chain-test"

    // Your Katena network information
    privateKeyED25519Base64 := "7C67DeoLnhI6jvsp3eMksU2Z6uzj8sqZbpgwZqfIyuCZbfoPcitCiCsSp2EzCfkY52Mx58xDOyQLb1OhC7cL5A=="
    companyChainID := "abcdef"

    // Convert your private key
    privateKey, err := utils.CreatePrivateKeyED25519FromBase64(privateKeyED25519Base64)
    if err != nil {
        panic(err)
    }

    // Create a Katena API helper
    transactor := client.NewTransactor(apiUrl, chainID, privateKey, companyChainID)

    // Off chain information you want to send
    certificateUuid := "2075c941-6876-405b-87d5-13791c0dc53a"
    dataSignature := []byte("off_chain_data_signature_from_go")
    dataSigner := []byte("off_chain_data_signer_from_go")

    // Send a version 1 of a certificate on Katena blockchain
    transactionStatus, err := transactor.SendCertificateV1(certificateUuid, dataSignature, dataSigner)
    if err != nil {
        panic(err)
    }

    fmt.Println("Transaction status")
    fmt.Println(fmt.Sprintf("  Code    : %d", transactionStatus.Code))
    fmt.Println(fmt.Sprintf("  Message : %s", transactionStatus.Message))

}
