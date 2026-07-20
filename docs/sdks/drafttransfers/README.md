# DraftTransfers

## Overview

### Available Operations

* [Create](#create) - Create draft transfer
* [List](#list) - List draft transfers
* [Get](#get) - Get draft transfer
* [Cancel](#cancel) - Cancel draft transfer

## Create

> 🚧 Beta feature
>
> This feature is currently in beta testing, and the final specification may still change.

Creates a draft transfer. The draft transfer immediately enters `awaiting-initiation` and appears in the
initiator's queue in Mollie Apps. It carries no legal weight and moves no funds until a human initiator
approves it there.

### Test mode

| Action | Test-mode behavior |
|---|---|
| Create | Always returns a synthetic draft in `awaiting-initiation`, regardless of most input values |
| Get / List | Always returns synthetic `awaiting-initiation` draft(s), not your real data |
| Cancel | Always returns a synthetic `declined` draft |

There is currently no way to synthesize an `initiated` draft via test mode alone. Initiation only happens when
a real initiator approves in Mollie Apps, which test mode doesn't touch.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-draft-transfer" method="post" path="/v2/business-accounts/draft-transfers" example="create-draft-transfer-201" -->
```go
package main

import(
	"context"
	"os"
	"github.com/mollie/mollie-api-golang/models/components"
	client "github.com/mollie/mollie-api-golang"
	"github.com/mollie/mollie-api-golang/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := client.New(
        client.WithSecurity(components.Security{
            AdvancedAccessToken: client.Pointer(os.Getenv("CLIENT_ADVANCED_ACCESS_TOKEN")),
        }),
    )

    res, err := s.DraftTransfers.Create(ctx, client.Pointer("123e4567-e89b-12d3-a456-426"), &components.CreateDraftTransferRequest{
        DebtorIban: client.Pointer("NL55MLLE0123456789"),
        Creditor: components.DraftTransferParty{
            FullName: "Jan Jansen",
            Account: components.DraftTransferPartyAccount{
                Iban: "NL02ABNA0123456789",
            },
        },
        Amount: components.Amount{
            Currency: "EUR",
            Value: "10.00",
        },
        Description: client.Pointer("Invoice 12345"),
        ScheduledFor: types.MustNewDateFromString("2025-03-01"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DraftTransferResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                       | Type                                                                                            | Required                                                                                        | Description                                                                                     | Example                                                                                         |
| ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| `ctx`                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                           | :heavy_check_mark:                                                                              | The context to use for the request.                                                             |                                                                                                 |
| `idempotencyKey`                                                                                | `*string`                                                                                       | :heavy_minus_sign:                                                                              | A unique key to ensure idempotent requests. This key should be a UUID v4 string.                | 123e4567-e89b-12d3-a456-426                                                                     |
| `createDraftTransferRequest`                                                                    | [*components.CreateDraftTransferRequest](../../models/components/createdrafttransferrequest.md) | :heavy_minus_sign:                                                                              | N/A                                                                                             |                                                                                                 |
| `opts`                                                                                          | [][operations.Option](../../models/operations/option.md)                                        | :heavy_minus_sign:                                                                              | The options for this request.                                                                   |                                                                                                 |

### Response

**[*operations.CreateDraftTransferResponse](../../models/operations/createdrafttransferresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 422, 429                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## List

> 🚧 Beta feature
>
> This feature is currently in beta testing, and the final specification may still change.

Retrieves a list of draft transfers for the organization including ones created in Mollie Apps, not just
ones created via this API.

The results are paginated.

In test mode, this always returns a synthetic `awaiting-initiation` list. See [Create draft
transfer](create-draft-transfer) for full test-mode behavior.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-draft-transfers" method="get" path="/v2/business-accounts/draft-transfers" example="list-draft-transfers-200" -->
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
            AdvancedAccessToken: client.Pointer(os.Getenv("CLIENT_ADVANCED_ACCESS_TOKEN")),
        }),
    )

    res, err := s.DraftTransfers.List(ctx, operations.ListDraftTransfersRequest{
        Limit: client.Pointer[int64](50),
        Status: components.DraftTransferStatusAwaitingInitiation.ToPointer(),
        Source: components.DraftTransferSourceAPI.ToPointer(),
        IdempotencyKey: client.Pointer("123e4567-e89b-12d3-a456-426"),
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListDraftTransfersRequest](../../models/operations/listdrafttransfersrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListDraftTransfersResponse](../../models/operations/listdrafttransfersresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 429                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Get

> 🚧 Beta feature
>
> This feature is currently in beta testing, and the final specification may still change.

Retrieves a single draft transfer by its identifier.

Only draft transfers created via this API are visible via this endpoint. Draft transfers created in Mollie
Apps return a `404`, even though they appear in the [list endpoint](list-draft-transfers).

In test mode, this always returns a synthetic `awaiting-initiation` draft. See [Create draft
transfer](create-draft-transfer) for full test-mode behavior.

### Example Usage: get-draft-transfer-200

<!-- UsageSnippet language="go" operationID="get-draft-transfer" method="get" path="/v2/business-accounts/draft-transfers/{draftTransferId}" example="get-draft-transfer-200" -->
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
            AdvancedAccessToken: client.Pointer(os.Getenv("CLIENT_ADVANCED_ACCESS_TOKEN")),
        }),
    )

    res, err := s.DraftTransfers.Get(ctx, "badrt_87GByBuj4UCcUTEbs6aGJ", client.Pointer("123e4567-e89b-12d3-a456-426"))
    if err != nil {
        log.Fatal(err)
    }
    if res.DraftTransferResponse != nil {
        // handle response
    }
}
```
### Example Usage: initiated-draft-transfer

<!-- UsageSnippet language="go" operationID="get-draft-transfer" method="get" path="/v2/business-accounts/draft-transfers/{draftTransferId}" example="initiated-draft-transfer" -->
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
            AdvancedAccessToken: client.Pointer(os.Getenv("CLIENT_ADVANCED_ACCESS_TOKEN")),
        }),
    )

    res, err := s.DraftTransfers.Get(ctx, "badrt_87GByBuj4UCcUTEbs6aGJ", client.Pointer("123e4567-e89b-12d3-a456-426"))
    if err != nil {
        log.Fatal(err)
    }
    if res.DraftTransferResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |                                                                                  |
| `draftTransferID`                                                                | `string`                                                                         | :heavy_check_mark:                                                               | Provide the ID of the related draft transfer.                                    | badrt_87GByBuj4UCcUTEbs6aGJ                                                      |
| `idempotencyKey`                                                                 | `*string`                                                                        | :heavy_minus_sign:                                                               | A unique key to ensure idempotent requests. This key should be a UUID v4 string. | 123e4567-e89b-12d3-a456-426                                                      |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |                                                                                  |

### Response

**[*operations.GetDraftTransferResponse](../../models/operations/getdrafttransferresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 429                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Cancel

> 🚧 Beta feature
>
> This feature is currently in beta testing, and the final specification may still change.

Cancels a draft transfer created via this API. Transitions the draft transfer to `declined` with
`statusReason` set to `"Declined by creator"`.

Only draft transfers created via this API, and still in `awaiting-initiation`, can be cancelled this way. A
`422` is returned if the initiator has already started approving it.

In test mode, this always returns a synthetic `declined` draft. See [Create draft
transfer](create-draft-transfer) for full test-mode behavior.

### Example Usage

<!-- UsageSnippet language="go" operationID="cancel-draft-transfer" method="delete" path="/v2/business-accounts/draft-transfers/{draftTransferId}" example="cancel-draft-transfer-200" -->
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
            AdvancedAccessToken: client.Pointer(os.Getenv("CLIENT_ADVANCED_ACCESS_TOKEN")),
        }),
    )

    res, err := s.DraftTransfers.Cancel(ctx, "badrt_87GByBuj4UCcUTEbs6aGJ", client.Pointer("123e4567-e89b-12d3-a456-426"))
    if err != nil {
        log.Fatal(err)
    }
    if res.DraftTransferResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |                                                                                  |
| `draftTransferID`                                                                | `string`                                                                         | :heavy_check_mark:                                                               | Provide the ID of the related draft transfer.                                    | badrt_87GByBuj4UCcUTEbs6aGJ                                                      |
| `idempotencyKey`                                                                 | `*string`                                                                        | :heavy_minus_sign:                                                               | A unique key to ensure idempotent requests. This key should be a UUID v4 string. | 123e4567-e89b-12d3-a456-426                                                      |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |                                                                                  |

### Response

**[*operations.CancelDraftTransferResponse](../../models/operations/canceldrafttransferresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 422, 429           | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |