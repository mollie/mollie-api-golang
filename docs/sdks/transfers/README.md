# Transfers

## Overview

### Available Operations

* [Create](#create) - Create transfer
* [Get](#get) - Get transfer

## Create

> 🚧 Beta feature
>
> This feature is currently in beta testing, and the final specification may still change.

Create a SEPA Credit Transfer from your Mollie Business Account.

To initiate a transfer, you must provide the transfer scheme, the amount, the debtor IBAN (your Mollie Business
Account IBAN), and the creditor (recipient) details.

Each request must include an `Idempotency-Key` header to prevent duplicate transfers, and must be signed using the
`X-Client-Signature` and `X-Client-Signed-At` headers.

### Simulating transfer scenarios in test mode

In test mode, you can simulate various transfer scenarios by adjusting the transfer amount. This allows you to
mimic the typical status progression of a real-world transfer. Note that a transfer's progression will stop once
it reaches a final status: `blocked`, `failed`, or `processed`.

| Amount  | Scenario                                           | Webhook sequence                                                                                                                                                   |
|---------|----------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `11.00` | Transfer initiated, pending review by Mollie       | `business-account-transfer.requested` → `business-account-transfer.initiated` → `business-account-transfer.pending-review`                                         |
| `12.00` | Transfer initiated, blocked by Mollie              | `business-account-transfer.requested` → `business-account-transfer.initiated` → `business-account-transfer.pending-review` → `business-account-transfer.blocked`   |
| `13.00` | Transfer initiated, failed on scheme submission    | `business-account-transfer.requested` → `business-account-transfer.initiated` → `business-account-transfer.failed`                                                 |
| Other   | Default: transfer is processed                     | `business-account-transfer.requested` → `business-account-transfer.initiated` → `business-account-transfer.processed`                                              |

### Example Usage

<!-- UsageSnippet language="go" operationID="create-transfer" method="post" path="/business-accounts/transfers" example="create-transfer-201" -->
```go
package main

import(
	"context"
	"os"
	"github.com/mollie/mollie-api-golang/models/components"
	client "github.com/mollie/mollie-api-golang"
	"github.com/mollie/mollie-api-golang/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := client.New(
        client.WithSecurity(components.Security{
            OrganizationAccessToken: client.Pointer(os.Getenv("CLIENT_ORGANIZATION_ACCESS_TOKEN")),
        }),
    )

    res, err := s.Transfers.Create(ctx, operations.CreateTransferRequest{
        XClientSignature: "<value>",
        XClientSignedAt: "2025-01-01T12:00:00Z",
        IdempotencyKey: "aa84d3c0-1484-4f45-8a8d-4674a0147f3f",
        IdempotencyKey1: client.Pointer("123e4567-e89b-12d3-a456-426"),
        TransferRequest: &components.TransferRequest{
            DebtorIban: "NL55MLLE0123456789",
            Creditor: components.TransferParty{
                FullName: "Jan Jansen",
                Account: components.TransferPartyAccount{
                    Iban: "NL02ABNA0123456789",
                },
            },
            Amount: components.Amount{
                Currency: "EUR",
                Value: "10.00",
            },
            Description: client.Pointer("Invoice 12345"),
            TransferScheme: components.TransferScheme{
                Type: components.TransferSchemeTypeSepaCreditInst,
            },
            Testmode: client.Pointer(false),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TransferResponse != nil {
        switch res.TransferResponse.Metadata.Type {
            case components.MetadataTypeStr:
                // res.TransferResponse.Metadata.Str is populated
            case components.MetadataTypeNumber:
                // res.TransferResponse.Metadata.Number is populated
            case components.MetadataTypeMapOfAny:
                // res.TransferResponse.Metadata.MapOfAny is populated
            case components.MetadataTypeArrayOfStr:
                // res.TransferResponse.Metadata.ArrayOfStr is populated
        }

    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateTransferRequest](../../models/operations/createtransferrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.CreateTransferResponse](../../models/operations/createtransferresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 422                     | application/hal+json    |
| apierrors.ErrorResponse | 503                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Get

> 🚧 Beta feature
>
> This feature is currently in beta testing, and the final specification may still change.

Retrieve a single transfer object by its transfer ID. This allows you to check the current status
and details of a previously created transfer.

### Example Usage: get-transfer-200

<!-- UsageSnippet language="go" operationID="get-transfer" method="get" path="/business-accounts/transfers/{businessAccountsTransferId}" example="get-transfer-200" -->
```go
package main

import(
	"context"
	"os"
	"github.com/mollie/mollie-api-golang/models/components"
	client "github.com/mollie/mollie-api-golang"
	"log"
)

func main() {
    ctx := context.Background()

    s := client.New(
        client.WithTestmode(false),
        client.WithSecurity(components.Security{
            OrganizationAccessToken: client.Pointer(os.Getenv("CLIENT_ORGANIZATION_ACCESS_TOKEN")),
        }),
    )

    res, err := s.Transfers.Get(ctx, "batrf_87GByBuj4UCcUTEbs6aGJ", client.Pointer("123e4567-e89b-12d3-a456-426"))
    if err != nil {
        log.Fatal(err)
    }
    if res.TransferResponse != nil {
        switch res.TransferResponse.Metadata.Type {
            case components.MetadataTypeStr:
                // res.TransferResponse.Metadata.Str is populated
            case components.MetadataTypeNumber:
                // res.TransferResponse.Metadata.Number is populated
            case components.MetadataTypeMapOfAny:
                // res.TransferResponse.Metadata.MapOfAny is populated
            case components.MetadataTypeArrayOfStr:
                // res.TransferResponse.Metadata.ArrayOfStr is populated
        }

    }
}
```
### Example Usage: processed-transfer

<!-- UsageSnippet language="go" operationID="get-transfer" method="get" path="/business-accounts/transfers/{businessAccountsTransferId}" example="processed-transfer" -->
```go
package main

import(
	"context"
	"os"
	"github.com/mollie/mollie-api-golang/models/components"
	client "github.com/mollie/mollie-api-golang"
	"log"
)

func main() {
    ctx := context.Background()

    s := client.New(
        client.WithTestmode(false),
        client.WithSecurity(components.Security{
            OrganizationAccessToken: client.Pointer(os.Getenv("CLIENT_ORGANIZATION_ACCESS_TOKEN")),
        }),
    )

    res, err := s.Transfers.Get(ctx, "batrf_87GByBuj4UCcUTEbs6aGJ", client.Pointer("123e4567-e89b-12d3-a456-426"))
    if err != nil {
        log.Fatal(err)
    }
    if res.TransferResponse != nil {
        switch res.TransferResponse.Metadata.Type {
            case components.MetadataTypeStr:
                // res.TransferResponse.Metadata.Str is populated
            case components.MetadataTypeNumber:
                // res.TransferResponse.Metadata.Number is populated
            case components.MetadataTypeMapOfAny:
                // res.TransferResponse.Metadata.MapOfAny is populated
            case components.MetadataTypeArrayOfStr:
                // res.TransferResponse.Metadata.ArrayOfStr is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                                | Type                                                                                                                                                                                                                                                                                                                                                                                     | Required                                                                                                                                                                                                                                                                                                                                                                                 | Description                                                                                                                                                                                                                                                                                                                                                                              | Example                                                                                                                                                                                                                                                                                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                       | The context to use for the request.                                                                                                                                                                                                                                                                                                                                                      |                                                                                                                                                                                                                                                                                                                                                                                          |
| `businessAccountsTransferID`                                                                                                                                                                                                                                                                                                                                                             | `string`                                                                                                                                                                                                                                                                                                                                                                                 | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                       | Provide the ID of the related transfer.                                                                                                                                                                                                                                                                                                                                                  | batrf_87GByBuj4UCcUTEbs6aGJ                                                                                                                                                                                                                                                                                                                                                              |
| `testmode`                                                                                                                                                                                                                                                                                                                                                                               | `*bool`                                                                                                                                                                                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                       | Most API credentials are specifically created for either live mode or test mode. In those cases the `testmode` query<br/>parameter must not be sent. For organization-level credentials such as OAuth access tokens, you can enable test mode by<br/>setting the `testmode` query parameter to `true`.<br/><br/>Test entities cannot be retrieved when the endpoint is set to live mode, and vice versa. |                                                                                                                                                                                                                                                                                                                                                                                          |
| `idempotencyKey`                                                                                                                                                                                                                                                                                                                                                                         | `*string`                                                                                                                                                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                       | A unique key to ensure idempotent requests. This key should be a UUID v4 string.                                                                                                                                                                                                                                                                                                         | 123e4567-e89b-12d3-a456-426                                                                                                                                                                                                                                                                                                                                                              |
| `opts`                                                                                                                                                                                                                                                                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                       | The options for this request.                                                                                                                                                                                                                                                                                                                                                            |                                                                                                                                                                                                                                                                                                                                                                                          |

### Response

**[*operations.GetTransferResponse](../../models/operations/gettransferresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |