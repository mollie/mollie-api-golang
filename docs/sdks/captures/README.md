# Captures
(*Captures*)

## Overview

### Available Operations

* [Create](#create) - Create capture
* [List](#list) - List captures
* [Get](#get) - Get capture

## Create

Capture an *authorized* payment.

Some payment methods allow you to first collect a customer's authorization,
and capture the amount at a later point.

By default, Mollie captures payments automatically. If however you
configured your payment with `captureMode: manual`, you can capture the payment using this endpoint after
having collected the customer's authorization.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-capture" method="post" path="/payments/{paymentId}/captures" -->
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

    res, err := s.Captures.Create(ctx, "tr_5B8cwPMGnU", client.Pointer("123e4567-e89b-12d3-a456-426"), &components.EntityCapture{
        ID: client.Pointer("cpt_vytxeTZskVKR7C7WgdSP3d"),
        Description: client.Pointer("Capture for cart #12345"),
        Amount: &components.AmountNullable{
            Currency: "EUR",
            Value: "10.00",
        },
        SettlementAmount: &components.AmountNullable{
            Currency: "EUR",
            Value: "10.00",
        },
        Status: components.CaptureStatusSucceeded.ToPointer(),
        PaymentID: client.Pointer("tr_5B8cwPMGnU"),
        ShipmentID: client.Pointer("shp_5x4xQJDWGNcY3tKGL7X5J"),
        SettlementID: client.Pointer("stl_5B8cwPMGnU"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CaptureResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |                                                                                  |
| `paymentID`                                                                      | *string*                                                                         | :heavy_check_mark:                                                               | Provide the ID of the related payment.                                           | tr_5B8cwPMGnU                                                                    |
| `idempotencyKey`                                                                 | **string*                                                                        | :heavy_minus_sign:                                                               | A unique key to ensure idempotent requests. This key should be a UUID v4 string. | 123e4567-e89b-12d3-a456-426                                                      |
| `entityCapture`                                                                  | [*components.EntityCapture](../../models/components/entitycapture.md)            | :heavy_minus_sign:                                                               | N/A                                                                              |                                                                                  |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |                                                                                  |

### Response

**[*operations.CreateCaptureResponse](../../models/operations/createcaptureresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 422                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## List

Retrieve a list of all captures created for a specific payment.

The results are paginated.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-captures" method="get" path="/payments/{paymentId}/captures" -->
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
        client.WithTestmode(false),
        client.WithSecurity(components.Security{
            APIKey: client.Pointer(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.Captures.List(ctx, operations.ListCapturesRequest{
        PaymentID: "tr_5B8cwPMGnU",
        From: client.Pointer("cpt_vytxeTZskVKR7C7WgdSP3d"),
        Limit: client.Pointer[int64](50),
        Embed: client.Pointer("payment"),
        IdempotencyKey: client.Pointer("123e4567-e89b-12d3-a456-426"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListCapturesRequest](../../models/operations/listcapturesrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.ListCapturesResponse](../../models/operations/listcapturesresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Get

Retrieve a single payment capture by its ID and the ID of its parent
payment.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-capture" method="get" path="/payments/{paymentId}/captures/{captureId}" -->
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
        client.WithTestmode(false),
        client.WithSecurity(components.Security{
            APIKey: client.Pointer(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.Captures.Get(ctx, operations.GetCaptureRequest{
        PaymentID: "tr_5B8cwPMGnU",
        CaptureID: "cpt_gVMhHKqSSRYJyPsuoPNFH",
        Embed: client.Pointer("payment"),
        IdempotencyKey: client.Pointer("123e4567-e89b-12d3-a456-426"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CaptureResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetCaptureRequest](../../models/operations/getcapturerequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.GetCaptureResponse](../../models/operations/getcaptureresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |