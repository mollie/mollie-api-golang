# VerifyPayee

## Overview

### Available Operations

* [Create](#create) - Verify Payee

## Create

> 🚧 Beta feature
>
> This feature is currently in beta testing, and the final specification may still change.

Perform a Verification of Payee (VoP) check. This allows you to verify the account holder name against the
records held by the receiving bank before initiating a transfer.

The verification result indicates whether the provided name matches, closely matches, or does not match the
name on file at the receiving bank. This helps prevent misdirected payments.

### Simulating verification scenarios in test mode

In test mode, you can simulate various verification outcomes by adjusting the creditor name in the
`creditorBankAccount.accountHolderName` property. This allows you to test all possible Verification of Payee
results without needing special properties. The names are case insensitive.

| Account holder name                    | Scenario                                      | Verification result | Suggested name |
|----------------------------------------|-----------------------------------------------|---------------------|----------------|
| `John Close Match`                     | Name closely matches the bank records          | `close-match`       | `John Match`   |
| `John No Match`                        | Name does not match the bank records           | `no-match`          | —              |
| `John Unavailable`                     | Verification is not available                  | `not-available`     | —              |
| Any other name                         | Default: name matches the bank records         | `match`             | —              |

### Example Usage: verify-payee-200-close-match

<!-- UsageSnippet language="go" operationID="verify-payee" method="post" path="/v2/business-accounts/payee-verifications" example="verify-payee-200-close-match" -->
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
        client.WithSecurity(components.Security{
            OrganizationAccessToken: client.Pointer(os.Getenv("CLIENT_ORGANIZATION_ACCESS_TOKEN")),
        }),
    )

    res, err := s.VerifyPayee.Create(ctx, client.Pointer("123e4567-e89b-12d3-a456-426"), &components.VerificationOfPayeeRequest{
        CreditorBankAccount: components.CreditorBankAccount{
            AccountHolderName: "Jan Jansen",
            Format: components.AccountNumberFormatIban,
            AccountNumber: "NL02ABNA0123456789",
        },
        Testmode: client.Pointer(false),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VerificationOfPayeeResponse != nil {
        // handle response
    }
}
```
### Example Usage: verify-payee-200-match

<!-- UsageSnippet language="go" operationID="verify-payee" method="post" path="/v2/business-accounts/payee-verifications" example="verify-payee-200-match" -->
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
        client.WithSecurity(components.Security{
            OrganizationAccessToken: client.Pointer(os.Getenv("CLIENT_ORGANIZATION_ACCESS_TOKEN")),
        }),
    )

    res, err := s.VerifyPayee.Create(ctx, client.Pointer("123e4567-e89b-12d3-a456-426"), &components.VerificationOfPayeeRequest{
        CreditorBankAccount: components.CreditorBankAccount{
            AccountHolderName: "Jan Jansen",
            Format: components.AccountNumberFormatIban,
            AccountNumber: "NL02ABNA0123456789",
        },
        Testmode: client.Pointer(false),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VerificationOfPayeeResponse != nil {
        // handle response
    }
}
```
### Example Usage: verify-payee-200-no-match

<!-- UsageSnippet language="go" operationID="verify-payee" method="post" path="/v2/business-accounts/payee-verifications" example="verify-payee-200-no-match" -->
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
        client.WithSecurity(components.Security{
            OrganizationAccessToken: client.Pointer(os.Getenv("CLIENT_ORGANIZATION_ACCESS_TOKEN")),
        }),
    )

    res, err := s.VerifyPayee.Create(ctx, client.Pointer("123e4567-e89b-12d3-a456-426"), &components.VerificationOfPayeeRequest{
        CreditorBankAccount: components.CreditorBankAccount{
            AccountHolderName: "Jan Jansen",
            Format: components.AccountNumberFormatIban,
            AccountNumber: "NL02ABNA0123456789",
        },
        Testmode: client.Pointer(false),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VerificationOfPayeeResponse != nil {
        // handle response
    }
}
```
### Example Usage: verify-payee-200-not-available

<!-- UsageSnippet language="go" operationID="verify-payee" method="post" path="/v2/business-accounts/payee-verifications" example="verify-payee-200-not-available" -->
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
        client.WithSecurity(components.Security{
            OrganizationAccessToken: client.Pointer(os.Getenv("CLIENT_ORGANIZATION_ACCESS_TOKEN")),
        }),
    )

    res, err := s.VerifyPayee.Create(ctx, client.Pointer("123e4567-e89b-12d3-a456-426"), &components.VerificationOfPayeeRequest{
        CreditorBankAccount: components.CreditorBankAccount{
            AccountHolderName: "Jan Jansen",
            Format: components.AccountNumberFormatIban,
            AccountNumber: "NL02ABNA0123456789",
        },
        Testmode: client.Pointer(false),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.VerificationOfPayeeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                       | Type                                                                                            | Required                                                                                        | Description                                                                                     | Example                                                                                         |
| ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| `ctx`                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                           | :heavy_check_mark:                                                                              | The context to use for the request.                                                             |                                                                                                 |
| `idempotencyKey`                                                                                | `*string`                                                                                       | :heavy_minus_sign:                                                                              | A unique key to ensure idempotent requests. This key should be a UUID v4 string.                | 123e4567-e89b-12d3-a456-426                                                                     |
| `verificationOfPayeeRequest`                                                                    | [*components.VerificationOfPayeeRequest](../../models/components/verificationofpayeerequest.md) | :heavy_minus_sign:                                                                              | N/A                                                                                             |                                                                                                 |
| `opts`                                                                                          | [][operations.Option](../../models/operations/option.md)                                        | :heavy_minus_sign:                                                                              | The options for this request.                                                                   |                                                                                                 |

### Response

**[*operations.VerifyPayeeResponse](../../models/operations/verifypayeeresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 422, 429                | application/hal+json    |
| apierrors.ErrorResponse | 503                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |