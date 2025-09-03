# Chargebacks
(*Chargebacks*)

## Overview

### Available Operations

* [List](#list) - List payment chargebacks
* [Get](#get) - Get payment chargeback
* [All](#all) - List all chargebacks

## List

Retrieve the chargebacks initiated for a specific payment.

The results are paginated.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-chargebacks" method="get" path="/payments/{paymentId}/chargebacks" -->
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
            APIKey: client.String(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.Chargebacks.List(ctx, operations.ListChargebacksRequest{
        PaymentID: "tr_5B8cwPMGnU",
        From: client.String("chb_xFzwUN4ci8HAmSGUACS4J"),
        Limit: client.Int64(50),
        Embed: operations.ListChargebacksEmbedPayment.ToPointer(),
        Testmode: client.Bool(false),
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListChargebacksRequest](../../models/operations/listchargebacksrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.ListChargebacksResponse](../../models/operations/listchargebacksresponse.md), error**

### Errors

| Error Type                                      | Status Code                                     | Content Type                                    |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| apierrors.ListChargebacksBadRequestHalJSONError | 400                                             | application/hal+json                            |
| apierrors.ListChargebacksNotFoundHalJSONError   | 404                                             | application/hal+json                            |
| apierrors.APIError                              | 4XX, 5XX                                        | \*/\*                                           |

## Get

Retrieve a single payment chargeback by its ID and the ID of its parent payment.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-chargeback" method="get" path="/payments/{paymentId}/chargebacks/{chargebackId}" -->
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
            APIKey: client.String(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.Chargebacks.Get(ctx, "tr_5B8cwPMGnU", "chb_xFzwUN4ci8HAmSGUACS4J", operations.GetChargebackEmbedPayment.ToPointer(), client.Bool(false))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                              | Type                                                                                                                                                                                                                                                                                                                                                                                   | Required                                                                                                                                                                                                                                                                                                                                                                               | Description                                                                                                                                                                                                                                                                                                                                                                            | Example                                                                                                                                                                                                                                                                                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                     | The context to use for the request.                                                                                                                                                                                                                                                                                                                                                    |                                                                                                                                                                                                                                                                                                                                                                                        |
| `paymentID`                                                                                                                                                                                                                                                                                                                                                                            | *string*                                                                                                                                                                                                                                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                     | Provide the ID of the related payment.                                                                                                                                                                                                                                                                                                                                                 | tr_5B8cwPMGnU                                                                                                                                                                                                                                                                                                                                                                          |
| `chargebackID`                                                                                                                                                                                                                                                                                                                                                                         | *string*                                                                                                                                                                                                                                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                     | Provide the ID of the related chargeback.                                                                                                                                                                                                                                                                                                                                              | chb_xFzwUN4ci8HAmSGUACS4J                                                                                                                                                                                                                                                                                                                                                              |
| `embed`                                                                                                                                                                                                                                                                                                                                                                                | [*operations.GetChargebackEmbed](../../models/operations/getchargebackembed.md)                                                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | This endpoint allows you to embed additional information via the `embed` query string parameter.                                                                                                                                                                                                                                                                                       | payment                                                                                                                                                                                                                                                                                                                                                                                |
| `testmode`                                                                                                                                                                                                                                                                                                                                                                             | **bool*                                                                                                                                                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | Most API credentials are specifically created for either live mode or test mode. In those cases the `testmode` query<br/>parameter can be omitted. For organization-level credentials such as OAuth access tokens, you can enable test mode by<br/>setting the `testmode` query parameter to `true`.<br/><br/>Test entities cannot be retrieved when the endpoint is set to live mode, and vice versa. | false                                                                                                                                                                                                                                                                                                                                                                                  |
| `opts`                                                                                                                                                                                                                                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | The options for this request.                                                                                                                                                                                                                                                                                                                                                          |                                                                                                                                                                                                                                                                                                                                                                                        |

### Response

**[*operations.GetChargebackResponse](../../models/operations/getchargebackresponse.md), error**

### Errors

| Error Type                          | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| apierrors.GetChargebackHalJSONError | 404                                 | application/hal+json                |
| apierrors.APIError                  | 4XX, 5XX                            | \*/\*                               |

## All

Retrieve all chargebacks initiated for all your payments.

The results are paginated.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-all-chargebacks" method="get" path="/chargebacks" -->
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
            APIKey: client.String(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.Chargebacks.All(ctx, operations.ListAllChargebacksRequest{
        From: client.String("chb_xFzwUN4ci8HAmSGUACS4J"),
        Limit: client.Int64(50),
        Embed: operations.ListAllChargebacksEmbedPayment.ToPointer(),
        Sort: operations.ListAllChargebacksSortDesc.ToPointer(),
        ProfileID: client.String("pfl_5B8cwPMGnU"),
        Testmode: client.Bool(false),
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListAllChargebacksRequest](../../models/operations/listallchargebacksrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListAllChargebacksResponse](../../models/operations/listallchargebacksresponse.md), error**

### Errors

| Error Type                                         | Status Code                                        | Content Type                                       |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| apierrors.ListAllChargebacksBadRequestHalJSONError | 400                                                | application/hal+json                               |
| apierrors.ListAllChargebacksNotFoundHalJSONError   | 404                                                | application/hal+json                               |
| apierrors.APIError                                 | 4XX, 5XX                                           | \*/\*                                              |