# UnmatchedCreditTransfers

## Overview

### Available Operations

* [List](#list) - List unmatched credit transfers
* [Get](#get) - Get unmatched credit transfer
* [Match](#match) - Match unmatched credit transfer
* [Return](#return) - Return unmatched credit transfer

## List

> 🚧 Beta feature
>
> This feature is currently in private beta, and the final specification may still change.

Retrieves a list of unmatched credit transfers for the profile.

The results are paginated.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-unmatched-credit-transfers" method="get" path="/v2/unmatched-credit-transfers" example="list-unmatched-credit-transfers-200-1" -->
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
            APIKey: client.Pointer(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.UnmatchedCreditTransfers.List(ctx, operations.ListUnmatchedCreditTransfersRequest{
        Limit: client.Pointer[int64](50),
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.ListUnmatchedCreditTransfersRequest](../../models/operations/listunmatchedcredittransfersrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.ListUnmatchedCreditTransfersResponse](../../models/operations/listunmatchedcredittransfersresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 429                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Get

> 🚧 Beta feature
>
> This feature is currently in private beta, and the final specification may still change.

Retrieves a single unmatched credit transfer by its identifier.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-unmatched-credit-transfer" method="get" path="/v2/unmatched-credit-transfers/{unmatchedCreditTransferId}" example="get-unmatched-credit-transfer-200-1" -->
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
            APIKey: client.Pointer(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.UnmatchedCreditTransfers.Get(ctx, "uct_abcDEFghij123456789", client.Pointer("123e4567-e89b-12d3-a456-426"), client.Pointer("MEUCIQDx8sJvK3z..."), client.Pointer("2025-01-01T12:00:00Z"))
    if err != nil {
        log.Fatal(err)
    }
    if res.EntityUnmatchedCreditTransfer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                               | Type                                                                                                                                                                                    | Required                                                                                                                                                                                | Description                                                                                                                                                                             | Example                                                                                                                                                                                 |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                   | :heavy_check_mark:                                                                                                                                                                      | The context to use for the request.                                                                                                                                                     |                                                                                                                                                                                         |
| `unmatchedCreditTransferID`                                                                                                                                                             | `string`                                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                      | Provide the ID of the related unmatched credit transfer.                                                                                                                                | uct_abcDEFghij123456789                                                                                                                                                                 |
| `idempotencyKey`                                                                                                                                                                        | `*string`                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                      | A unique key to ensure idempotent requests. This key should be a UUID v4 string.                                                                                                        | 123e4567-e89b-12d3-a456-426                                                                                                                                                             |
| `xClientSignature`                                                                                                                                                                      | `*string`                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                      | A cryptographic signature of the request, used to verify that it was sent by you and hasn't been tampered with. Required if you've enabled request signing for your integration.        | MEUCIQDx8sJvK3z...                                                                                                                                                                      |
| `xClientSignedAt`                                                                                                                                                                       | `*string`                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                      | The timestamp, in ISO 8601 format, at which you generated the `X-Client-Signature`. Required together with `X-Client-Signature` if you've enabled request signing for your integration. | 2025-01-01T12:00:00Z                                                                                                                                                                    |
| `opts`                                                                                                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                      | The options for this request.                                                                                                                                                           |                                                                                                                                                                                         |

### Response

**[*operations.GetUnmatchedCreditTransferResponse](../../models/operations/getunmatchedcredittransferresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 429                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Match

> 🚧
>
> To get access to this functionality, a virtual IBAN must be assigned to your webshop. We are currently rolling out virtual IBANs to our merchant base. Contact our support team or your account/sales manager to request access today.

Matches an unmatched credit transfer to one or more payments, settling the funds accordingly.

### Example Usage

<!-- UsageSnippet language="go" operationID="match-unmatched-credit-transfer" method="post" path="/v2/unmatched-credit-transfers/{unmatchedCreditTransferId}/match" example="match-unmatched-credit-transfer-201-1" -->
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
            APIKey: client.Pointer(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.UnmatchedCreditTransfers.Match(ctx, operations.MatchUnmatchedCreditTransferRequest{
        UnmatchedCreditTransferID: "uct_abcDEFghij123456789",
        IdempotencyKey: client.Pointer("123e4567-e89b-12d3-a456-426"),
        XClientSignature: client.Pointer("MEUCIQDx8sJvK3z..."),
        XClientSignedAt: client.Pointer("2025-01-01T12:00:00Z"),
        UnmatchedCreditTransferMatchRequest: &components.UnmatchedCreditTransferMatchRequest{
            PaymentIds: []string{
                "tr_5B8cwPMGnU",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UnmatchedCreditTransferActionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.MatchUnmatchedCreditTransferRequest](../../models/operations/matchunmatchedcredittransferrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.MatchUnmatchedCreditTransferResponse](../../models/operations/matchunmatchedcredittransferresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 422, 429           | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Return

> 🚧 Beta feature
>
> This feature is currently in private beta, and the final specification may still change.

Returns an unmatched credit transfer, sending the funds back to the original sender.

### Example Usage

<!-- UsageSnippet language="go" operationID="return-unmatched-credit-transfer" method="post" path="/v2/unmatched-credit-transfers/{unmatchedCreditTransferId}/return" example="return-unmatched-credit-transfer-201-1" -->
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
            APIKey: client.Pointer(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.UnmatchedCreditTransfers.Return(ctx, "uct_abcDEFghij123456789", client.Pointer("123e4567-e89b-12d3-a456-426"), client.Pointer("MEUCIQDx8sJvK3z..."), client.Pointer("2025-01-01T12:00:00Z"))
    if err != nil {
        log.Fatal(err)
    }
    if res.UnmatchedCreditTransferActionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                               | Type                                                                                                                                                                                    | Required                                                                                                                                                                                | Description                                                                                                                                                                             | Example                                                                                                                                                                                 |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                   | :heavy_check_mark:                                                                                                                                                                      | The context to use for the request.                                                                                                                                                     |                                                                                                                                                                                         |
| `unmatchedCreditTransferID`                                                                                                                                                             | `string`                                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                      | Provide the ID of the related unmatched credit transfer.                                                                                                                                | uct_abcDEFghij123456789                                                                                                                                                                 |
| `idempotencyKey`                                                                                                                                                                        | `*string`                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                      | A unique key to ensure idempotent requests. This key should be a UUID v4 string.                                                                                                        | 123e4567-e89b-12d3-a456-426                                                                                                                                                             |
| `xClientSignature`                                                                                                                                                                      | `*string`                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                      | A cryptographic signature of the request, used to verify that it was sent by you and hasn't been tampered with. Required if you've enabled request signing for your integration.        | MEUCIQDx8sJvK3z...                                                                                                                                                                      |
| `xClientSignedAt`                                                                                                                                                                       | `*string`                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                      | The timestamp, in ISO 8601 format, at which you generated the `X-Client-Signature`. Required together with `X-Client-Signature` if you've enabled request signing for your integration. | 2025-01-01T12:00:00Z                                                                                                                                                                    |
| `opts`                                                                                                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                      | The options for this request.                                                                                                                                                           |                                                                                                                                                                                         |

### Response

**[*operations.ReturnUnmatchedCreditTransferResponse](../../models/operations/returnunmatchedcredittransferresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 429                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |