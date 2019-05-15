package main

import (
    "fmt"

    "github.com/katena-chain/sdk-go-client/client"
    "github.com/katena-chain/sdk-go-client/utils"
)

func main() {

    // Common Katena network informations
    apiUrl := "https://api.test.katena.transchain.io/api/v1"
    chainID := "katena-chain-test"

    // Your Katena network informations
    privateKeyED25519Base64 := "7C67DeoLnhI6jvsp3eMksU2Z6uzj8sqZbpgwZqfIyuCZbfoPcitCiCsSp2EzCfkY52Mx58xDOyQLb1OhC7cL5A=="
    companyChainID := "abcdef"

    // Convert your private key
    privateKey, err := utils.CreatePrivateKeyED25519FromBase64(privateKeyED25519Base64)
    if err != nil {
        panic(err)
    }

    // Create a Katena API helper
    transactor := client.NewTransactor(apiUrl, chainID, privateKey, companyChainID)

    // Secret information about a certificate you want to send
    certificateUuid := "2075c941-6876-405b-87d5-13791c0dc53a"
    content := []byte("off_chain_data_aes_encryption_key_from_go")

    // The recipient public key able to decrypt the secret later
    recipientPublicKeyX25519Base64 := "CgguJuEb+/cSHD4Jo8JcVRpwDlt834pFijvd2AdWIgE="
    recipientPublicKey, err := utils.CreatePublicKeyX25519FromBase64(recipientPublicKeyX25519Base64)
    if err != nil {
        panic(err)
    }

    // Ephemeral key pair (recommended) to encrypt the secret
    senderEphemeralPublicKey, senderEphemeralPrivateKey, err := utils.CreateNewKeysX25519()
    if err != nil {
        panic(err)
    }

    // Encrypt the secret
    nonce, encryptedContent, err := senderEphemeralPrivateKey.Seal(content, recipientPublicKey)
    if err != nil {
        panic(err)
    }

    // Send a version 1 of a secret on Katena blockchain
    transactionStatus, err := transactor.SendSecretV1(certificateUuid, senderEphemeralPublicKey, nonce, encryptedContent)
    if err != nil {
        panic(err)
    }

    fmt.Println("Transaction status")
    fmt.Println(fmt.Sprintf("  Code    : %d", transactionStatus.Code))
    fmt.Println(fmt.Sprintf("  Message : %s", transactionStatus.Message))

}
