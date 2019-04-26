# SDK Go Client

## Requirements

- Go >= 1.10

## Install

```bash
go get -u github.com/katena-chain/sdk-go-client/...
```

## Usage

To rapidly interact with our API, you can use our `Transactor` helper. It handles all the steps needed to correctly
format, sign and send a transaction.

Feel free to explore and modify its code to meet your expectations.

Here is a snippet to demonstrate its usage:

```go
package main

import (
    "fmt"
    "strconv"
    "strings"

    "github.com/katena-chain/sdk-go-client/client"
)

func main() {

    apiUrl := "https://api.demo.katena.transchain.io"
    apiUrlSuffix := "api/v1"
    chainID := "katena-chain"
    privateKeyBase64 := "7C67DeoLnhI6jvsp3eMksU2Z6uzj8sqZbpgwZqfIyuCZbfoPcitCiCsSp2EzCfkY52Mx58xDOyQLb1OhC7cL5A=="
    companyChainID := "abcdef"

    transactor, err := client.NewTransactor(apiUrl, apiUrlSuffix, chainID, privateKeyBase64, companyChainID)
    if err != nil {
        panic(err)
    }

    uuidv4 := "2075c941-6876-405b-87d5-13791c0dc53a"
    dataSignature := "document_signature_value"
    dataSigner := "document_signer_value"
    apiResponse, err := transactor.SendCertificate(uuidv4, dataSignature, dataSigner)
    if err != nil {
        panic(err)
    }

    fmt.Println("API status code : " + strconv.Itoa(apiResponse.StatusCode))
    fmt.Println("API body        : " + strings.Replace(string(apiResponse.Body), "\n", "", -1))

}
```

## Katena documentation

For more information, check the [katena documentation](https://doc.katena.transchain.io)