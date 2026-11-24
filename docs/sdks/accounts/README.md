# Accounts

## Overview

### Available Operations

* [CreateAccount](#createaccount) - Create business account
* [ListAccounts](#listaccounts) - List business accounts
* [GetAccount](#getaccount) - Get business account
* [List](#list) - List transactions
* [Get](#get) - Get transaction

## CreateAccount

> 🚧 Beta feature
>
> This feature is currently in beta testing, and the final specification may still change.

> 🔐 **Request signing required**
>
> Signature enforcement is mandatory for Direct Customer API users for `live` mode requests.
>
> Sign every request with the `X-Client-Signature` and `X-Client-Signed-At` headers.
>
> [Learn how to sign requests](https://docs.mollie.com/reference/request-signing)

Create a new business account for the authenticated organization.

The response contains the newly created business account.

Each organization can have up to 5 active business accounts. Need more? [Contact us](https://www.mollie.com/contact)
and we'll help you out.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-business-account" method="post" path="/v2/business-accounts/accounts" example="create-business-account-201" -->
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
            OAuth: client.Pointer(os.Getenv("CLIENT_O_AUTH")),
        }),
    )

    res, err := s.Accounts.CreateAccount(ctx, client.Pointer("123e4567-e89b-12d3-a456-426614174000"), client.Pointer("MEUCIQDx8sJvK3z..."), client.Pointer("2025-01-01T12:00:00Z"), &operations.CreateBusinessAccountRequestBody{
        Name: "Main Checking Account",
        Currency: components.AccountCurrencyEur,
        Testmode: client.Pointer(false),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BusinessAccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                                                                           | Type                                                                                                                                                                                                                                                                                                                                                                                                                                | Required                                                                                                                                                                                                                                                                                                                                                                                                                            | Description                                                                                                                                                                                                                                                                                                                                                                                                                         | Example                                                                                                                                                                                                                                                                                                                                                                                                                             |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                  | The context to use for the request.                                                                                                                                                                                                                                                                                                                                                                                                 |                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| `idempotencyKey`                                                                                                                                                                                                                                                                                                                                                                                                                    | `*string`                                                                                                                                                                                                                                                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                  | A unique value that identifies this request, so you can safely retry it without creating a second account.<br/>Retrying with the same key within one hour returns the original response. We recommend a UUIDv4.<br/><br/>Keys are tied to the credential that sent the original request, so a refreshed OAuth token won't match<br/>earlier keys. View the [public documentation](https://docs.mollie.com/reference/api-idempotency) to learn<br/>more. | 00000000-0000-0000-0000-000000000000                                                                                                                                                                                                                                                                                                                                                                                                |
| `xClientSignature`                                                                                                                                                                                                                                                                                                                                                                                                                  | `*string`                                                                                                                                                                                                                                                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                  | A cryptographic signature of the request, used to verify that it was sent by you and hasn't been tampered with. Required if you've enabled request signing for your integration.                                                                                                                                                                                                                                                    | MEUCIQDx8sJvK3z...                                                                                                                                                                                                                                                                                                                                                                                                                  |
| `xClientSignedAt`                                                                                                                                                                                                                                                                                                                                                                                                                   | `*string`                                                                                                                                                                                                                                                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                  | The timestamp, in ISO 8601 format, at which you generated the `X-Client-Signature`. Required together with `X-Client-Signature` if you've enabled request signing for your integration.                                                                                                                                                                                                                                             | 2025-01-01T12:00:00Z                                                                                                                                                                                                                                                                                                                                                                                                                |
| `requestBody`                                                                                                                                                                                                                                                                                                                                                                                                                       | [*operations.CreateBusinessAccountRequestBody](../../models/operations/createbusinessaccountrequestbody.md)                                                                                                                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                  | N/A                                                                                                                                                                                                                                                                                                                                                                                                                                 |                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| `opts`                                                                                                                                                                                                                                                                                                                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                  | The options for this request.                                                                                                                                                                                                                                                                                                                                                                                                       |                                                                                                                                                                                                                                                                                                                                                                                                                                     |

### Response

**[*operations.CreateBusinessAccountResponse](../../models/operations/createbusinessaccountresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 409, 422, 429      | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## ListAccounts

> 🔐 **Request signing required**
>
> Signature enforcement is mandatory for Direct Customer API users for `live` mode requests.
>
> Sign every request with the `X-Client-Signature` and `X-Client-Signed-At` headers.
>
> [Learn how to sign requests](https://docs.mollie.com/reference/request-signing)

Retrieve all business accounts for the authenticated organization.

The results are paginated.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-business-accounts" method="get" path="/v2/business-accounts/accounts" example="list-business-accounts-200" -->
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
        client.WithTestmode(true),
        client.WithSecurity(components.Security{
            OAuth: client.Pointer(os.Getenv("CLIENT_O_AUTH")),
        }),
    )

    res, err := s.Accounts.ListAccounts(ctx, operations.ListBusinessAccountsRequest{
        Iban: client.Pointer("NL95MLLE1234567890"),
        From: client.Pointer("ba_nopqrstuvwxyz23456789A"),
        Limit: client.Pointer[int64](50),
        Sort: components.SortingDesc.ToPointer(),
        IdempotencyKey: client.Pointer("123e4567-e89b-12d3-a456-426"),
        XClientSignature: client.Pointer("MEUCIQDx8sJvK3z..."),
        XClientSignedAt: client.Pointer("2025-01-01T12:00:00Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        for {
            // handle items

            res, err = res.Next()

            if err != nil {
                // handle error
            }

            if res == nil {
                break
            }
        }
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListBusinessAccountsRequest](../../models/operations/listbusinessaccountsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListBusinessAccountsResponse](../../models/operations/listbusinessaccountsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 429                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## GetAccount

> 🔐 **Request signing required**
>
> Signature enforcement is mandatory for Direct Customer API users for `live` mode requests.
>
> Sign every request with the `X-Client-Signature` and `X-Client-Signed-At` headers.
>
> [Learn how to sign requests](https://docs.mollie.com/reference/request-signing)

Retrieve a single business account object by its account ID. This allows you to check the current status,
balance, and account details.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-business-account" method="get" path="/v2/business-accounts/accounts/{businessAccountId}" example="get-business-account-200" -->
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
        client.WithTestmode(true),
        client.WithSecurity(components.Security{
            OAuth: client.Pointer(os.Getenv("CLIENT_O_AUTH")),
        }),
    )

    res, err := s.Accounts.GetAccount(ctx, operations.GetBusinessAccountRequest{
        BusinessAccountID: "ba_nopqrstuvwxyz23456789A",
        IdempotencyKey: client.Pointer("123e4567-e89b-12d3-a456-426"),
        XClientSignature: client.Pointer("MEUCIQDx8sJvK3z..."),
        XClientSignedAt: client.Pointer("2025-01-01T12:00:00Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BusinessAccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetBusinessAccountRequest](../../models/operations/getbusinessaccountrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetBusinessAccountResponse](../../models/operations/getbusinessaccountresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 429                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## List

> 🔐 **Request signing required**
>
> Signature enforcement is mandatory for Direct Customer API users for `live` mode requests.
>
> Sign every request with the `X-Client-Signature` and `X-Client-Signed-At` headers.
>
> [Learn how to sign requests](https://docs.mollie.com/reference/request-signing)

Retrieve all transactions for a specific business account.

The results are paginated.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-business-account-transactions" method="get" path="/v2/business-accounts/accounts/{businessAccountId}/transactions" example="list-transactions-200" -->
```go
package main

import(
	"context"
	"os"
	"github.com/mollie/mollie-api-golang/models/components"
	client "github.com/mollie/mollie-api-golang"
	"github.com/mollie/mollie-api-golang/types"
	"github.com/mollie/mollie-api-golang/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := client.New(
        client.WithTestmode(true),
        client.WithSecurity(components.Security{
            OAuth: client.Pointer(os.Getenv("CLIENT_O_AUTH")),
        }),
    )

    res, err := s.Accounts.List(ctx, operations.ListBusinessAccountTransactionsRequest{
        BusinessAccountID: "ba_nopqrstuvwxyz23456789A",
        CreditDebitIndicator: components.CreditDebitIndicatorDebit.ToPointer(),
        ProcessedAfter: types.MustNewTimeFromString("2025-02-01T00:00:00+00:00"),
        ProcessedBefore: types.MustNewTimeFromString("2025-02-26T23:59:59+00:00"),
        From: client.Pointer("batr_87GByBuj4UCcUTEbs6aGJ"),
        Limit: client.Pointer[int64](50),
        Sort: components.SortingDesc.ToPointer(),
        IdempotencyKey: client.Pointer("123e4567-e89b-12d3-a456-426"),
        XClientSignature: client.Pointer("MEUCIQDx8sJvK3z..."),
        XClientSignedAt: client.Pointer("2025-01-01T12:00:00Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        for {
            // handle items

            res, err = res.Next()

            if err != nil {
                // handle error
            }

            if res == nil {
                break
            }
        }
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.ListBusinessAccountTransactionsRequest](../../models/operations/listbusinessaccounttransactionsrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.ListBusinessAccountTransactionsResponse](../../models/operations/listbusinessaccounttransactionsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 429                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Get

> 🔐 **Request signing required**
>
> Signature enforcement is mandatory for Direct Customer API users for `live` mode requests.
>
> Sign every request with the `X-Client-Signature` and `X-Client-Signed-At` headers.
>
> [Learn how to sign requests](https://docs.mollie.com/reference/request-signing)

Retrieve a single transaction object by its transaction ID. This allows you to check the details,
amount, counterparty, and balance impact of a specific transaction.

### Example Usage: get-transaction-200

<!-- UsageSnippet language="go" operationID="get-business-account-transaction" method="get" path="/v2/business-accounts/accounts/{businessAccountId}/transactions/{transactionId}" example="get-transaction-200" -->
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
        client.WithTestmode(true),
        client.WithSecurity(components.Security{
            OAuth: client.Pointer(os.Getenv("CLIENT_O_AUTH")),
        }),
    )

    res, err := s.Accounts.Get(ctx, operations.GetBusinessAccountTransactionRequest{
        BusinessAccountID: "ba_nopqrstuvwxyz23456789A",
        TransactionID: "batr_87GByBuj4UCcUTEbs6aGJ",
        IdempotencyKey: client.Pointer("123e4567-e89b-12d3-a456-426"),
        XClientSignature: client.Pointer("MEUCIQDx8sJvK3z..."),
        XClientSignedAt: client.Pointer("2025-01-01T12:00:00Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TransactionResponse != nil {
        // handle response
    }
}
```
### Example Usage: get-transaction-200-credit

<!-- UsageSnippet language="go" operationID="get-business-account-transaction" method="get" path="/v2/business-accounts/accounts/{businessAccountId}/transactions/{transactionId}" example="get-transaction-200-credit" -->
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
        client.WithTestmode(true),
        client.WithSecurity(components.Security{
            OAuth: client.Pointer(os.Getenv("CLIENT_O_AUTH")),
        }),
    )

    res, err := s.Accounts.Get(ctx, operations.GetBusinessAccountTransactionRequest{
        BusinessAccountID: "ba_nopqrstuvwxyz23456789A",
        TransactionID: "batr_87GByBuj4UCcUTEbs6aGJ",
        IdempotencyKey: client.Pointer("123e4567-e89b-12d3-a456-426"),
        XClientSignature: client.Pointer("MEUCIQDx8sJvK3z..."),
        XClientSignedAt: client.Pointer("2025-01-01T12:00:00Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TransactionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetBusinessAccountTransactionRequest](../../models/operations/getbusinessaccounttransactionrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.GetBusinessAccountTransactionResponse](../../models/operations/getbusinessaccounttransactionresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 429                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |