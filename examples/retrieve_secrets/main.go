package main

import (
    "encoding/base64"
    "fmt"

    "github.com/katena-chain/sdk-go-client/client"
    "github.com/katena-chain/sdk-go-client/utils"
)

func main() {

    // Common Katena network informations
    apiUrl := "https://api.test.katena.transchain.io/api/v1"

    // Your Katena network informations
    companyChainID := "abcdef"

    // Create a Katena API helper
    transactor := client.NewTransactor(apiUrl, "", nil, companyChainID)

    // Your decryption private key
    recipientPrivateKeyX25519Base64 := "/HYK9/xU3SSKNtylLEQs/MrjujgrxYkWuDFQ4A2QayQ="
    recipientPrivateKey, err := utils.CreatePrivateKeyX25519FromBase64(recipientPrivateKeyX25519Base64)
    if err != nil {
        panic(err)
    }

    // Certificate uuid you want to retrieve secrets
    certificateUuid := "2075c941-6876-405b-87d5-13791c0dc53a"

    // Retrieve version 1 of secrets from Katena blockchain
    secretV1Wrappers, err := transactor.RetrieveSecretsV1(certificateUuid)
    if err != nil {
        panic(err)
    }

    for _, secretV1Wrapper := range secretV1Wrappers.Secrets {
        fmt.Println("Transaction status")
        fmt.Println(fmt.Sprintf("  Code    : %d", secretV1Wrapper.Status.Code))
        fmt.Println(fmt.Sprintf("  Message : %s", secretV1Wrapper.Status.Message))

        fmt.Println("SecretV1")
        fmt.Println(fmt.Sprintf("  Certificate uuid  : %s", secretV1Wrapper.Secret.CertificateUuid))
        fmt.Println(fmt.Sprintf("  Company chain id  : %s", secretV1Wrapper.Secret.CompanyChainID))
        fmt.Println(fmt.Sprintf("  Lock encryptor    : %s", base64.StdEncoding.EncodeToString(secretV1Wrapper.Secret.Lock.Encryptor[:])))
        fmt.Println(fmt.Sprintf("  Lock nonce        : %s", base64.StdEncoding.EncodeToString(secretV1Wrapper.Secret.Lock.Nonce[:])))
        fmt.Println(fmt.Sprintf("  Lock content      : %s", base64.StdEncoding.EncodeToString(secretV1Wrapper.Secret.Lock.Content)))

        // Try to decrypt the content
        decryptedContent, ok := recipientPrivateKey.Open(
            secretV1Wrapper.Secret.Lock.Content,
            secretV1Wrapper.Secret.Lock.Encryptor,
            secretV1Wrapper.Secret.Lock.Nonce,
        )

        if !ok {
            decryptedContent = []byte("Unable to decrypt")
        }
        fmt.Println(fmt.Sprintf("  Decrypted content : %s", decryptedContent))
        fmt.Println()
    }

}
