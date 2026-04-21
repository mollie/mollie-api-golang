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
	"log"
)

func main() {
    ctx := context.Background()

    s := client.New(
        client.WithSecurity(components.Security{
            APIKey: client.Pointer(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.UnmatchedCreditTransfers.List(ctx, nil, client.Pointer[int64](50), client.Pointer("123e4567-e89b-12d3-a456-426"))
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

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    | Example                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |                                                                                                                                |
| `from`                                                                                                                         | `*string`                                                                                                                      | :heavy_minus_sign:                                                                                                             | Provide an ID to start the result set from the item with the given ID and onwards. This allows you to paginate the<br/>result set. |                                                                                                                                |
| `limit`                                                                                                                        | `*int64`                                                                                                                       | :heavy_minus_sign:                                                                                                             | The maximum number of items to return. Defaults to 50 items.                                                                   | 50                                                                                                                             |
| `idempotencyKey`                                                                                                               | `*string`                                                                                                                      | :heavy_minus_sign:                                                                                                             | A unique key to ensure idempotent requests. This key should be a UUID v4 string.                                               | 123e4567-e89b-12d3-a456-426                                                                                                    |
| `opts`                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |                                                                                                                                |

### Response

**[*operations.ListUnmatchedCreditTransfersResponse](../../models/operations/listunmatchedcredittransfersresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400                     | application/hal+json    |
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

    res, err := s.UnmatchedCreditTransfers.Get(ctx, "uct_abcDEFghij123456789", client.Pointer("123e4567-e89b-12d3-a456-426"))
    if err != nil {
        log.Fatal(err)
    }
    if res.EntityUnmatchedCreditTransfer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |                                                                                  |
| `unmatchedCreditTransferID`                                                      | `string`                                                                         | :heavy_check_mark:                                                               | Provide the ID of the related unmatched credit transfer.                         | uct_abcDEFghij123456789                                                          |
| `idempotencyKey`                                                                 | `*string`                                                                        | :heavy_minus_sign:                                                               | A unique key to ensure idempotent requests. This key should be a UUID v4 string. | 123e4567-e89b-12d3-a456-426                                                      |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |                                                                                  |

### Response

**[*operations.GetUnmatchedCreditTransferResponse](../../models/operations/getunmatchedcredittransferresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Match

> 🚧 Beta feature
>
> This feature is currently in private beta, and the final specification may still change.

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
	"log"
)

func main() {
    ctx := context.Background()

    s := client.New(
        client.WithSecurity(components.Security{
            APIKey: client.Pointer(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.UnmatchedCreditTransfers.Match(ctx, "uct_abcDEFghij123456789", client.Pointer("123e4567-e89b-12d3-a456-426"), &components.UnmatchedCreditTransferMatchRequest{
        PaymentIds: []string{
            "tr_5B8cwPMGnU",
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

| Parameter                                                                                                         | Type                                                                                                              | Required                                                                                                          | Description                                                                                                       | Example                                                                                                           |
| ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                                             | :heavy_check_mark:                                                                                                | The context to use for the request.                                                                               |                                                                                                                   |
| `unmatchedCreditTransferID`                                                                                       | `string`                                                                                                          | :heavy_check_mark:                                                                                                | Provide the ID of the related unmatched credit transfer.                                                          | uct_abcDEFghij123456789                                                                                           |
| `idempotencyKey`                                                                                                  | `*string`                                                                                                         | :heavy_minus_sign:                                                                                                | A unique key to ensure idempotent requests. This key should be a UUID v4 string.                                  | 123e4567-e89b-12d3-a456-426                                                                                       |
| `unmatchedCreditTransferMatchRequest`                                                                             | [*components.UnmatchedCreditTransferMatchRequest](../../models/components/unmatchedcredittransfermatchrequest.md) | :heavy_minus_sign:                                                                                                | N/A                                                                                                               |                                                                                                                   |
| `opts`                                                                                                            | [][operations.Option](../../models/operations/option.md)                                                          | :heavy_minus_sign:                                                                                                | The options for this request.                                                                                     |                                                                                                                   |

### Response

**[*operations.MatchUnmatchedCreditTransferResponse](../../models/operations/matchunmatchedcredittransferresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 422                | application/hal+json    |
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

    res, err := s.UnmatchedCreditTransfers.Return(ctx, "uct_abcDEFghij123456789", client.Pointer("123e4567-e89b-12d3-a456-426"))
    if err != nil {
        log.Fatal(err)
    }
    if res.UnmatchedCreditTransferActionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |                                                                                  |
| `unmatchedCreditTransferID`                                                      | `string`                                                                         | :heavy_check_mark:                                                               | Provide the ID of the related unmatched credit transfer.                         | uct_abcDEFghij123456789                                                          |
| `idempotencyKey`                                                                 | `*string`                                                                        | :heavy_minus_sign:                                                               | A unique key to ensure idempotent requests. This key should be a UUID v4 string. | 123e4567-e89b-12d3-a456-426                                                      |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |                                                                                  |

### Response

**[*operations.ReturnUnmatchedCreditTransferResponse](../../models/operations/returnunmatchedcredittransferresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |