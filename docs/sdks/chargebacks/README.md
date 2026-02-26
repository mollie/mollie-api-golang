# Chargebacks

## Overview

### Available Operations

* [List](#list) - List payment chargebacks
* [Get](#get) - Get payment chargeback
* [All](#all) - List all chargebacks

## List

Retrieve the chargebacks initiated for a specific payment.

The results are paginated.

### Example Usage: list-chargeback-200-1

<!-- UsageSnippet language="go" operationID="list-chargebacks" method="get" path="/payments/{paymentId}/chargebacks" example="list-chargeback-200-1" -->
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

    res, err := s.Chargebacks.List(ctx, operations.ListChargebacksRequest{
        PaymentID: "tr_5B8cwPMGnU",
        From: client.Pointer("chb_xFzwUN4ci8HAmSGUACS4J"),
        Limit: client.Pointer[int64](50),
        Embed: client.Pointer("payment"),
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
### Example Usage: list-chargeback-200-2

<!-- UsageSnippet language="go" operationID="list-chargebacks" method="get" path="/payments/{paymentId}/chargebacks" example="list-chargeback-200-2" -->
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

    res, err := s.Chargebacks.List(ctx, operations.ListChargebacksRequest{
        PaymentID: "tr_5B8cwPMGnU",
        From: client.Pointer("chb_xFzwUN4ci8HAmSGUACS4J"),
        Limit: client.Pointer[int64](50),
        Embed: client.Pointer("payment"),
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
### Example Usage: list-chargeback-200-3

<!-- UsageSnippet language="go" operationID="list-chargebacks" method="get" path="/payments/{paymentId}/chargebacks" example="list-chargeback-200-3" -->
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

    res, err := s.Chargebacks.List(ctx, operations.ListChargebacksRequest{
        PaymentID: "tr_5B8cwPMGnU",
        From: client.Pointer("chb_xFzwUN4ci8HAmSGUACS4J"),
        Limit: client.Pointer[int64](50),
        Embed: client.Pointer("payment"),
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListChargebacksRequest](../../models/operations/listchargebacksrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.ListChargebacksResponse](../../models/operations/listchargebacksresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Get

Retrieve a single payment chargeback by its ID and the ID of its parent payment.

### Example Usage: get-chargeback-200-1

<!-- UsageSnippet language="go" operationID="get-chargeback" method="get" path="/payments/{paymentId}/chargebacks/{chargebackId}" example="get-chargeback-200-1" -->
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

    res, err := s.Chargebacks.Get(ctx, operations.GetChargebackRequest{
        PaymentID: "tr_5B8cwPMGnU",
        ChargebackID: "chb_xFzwUN4ci8HAmSGUACS4J",
        Embed: client.Pointer("payment"),
        IdempotencyKey: client.Pointer("123e4567-e89b-12d3-a456-426"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EntityChargeback != nil {
        // handle response
    }
}
```
### Example Usage: get-chargeback-200-2

<!-- UsageSnippet language="go" operationID="get-chargeback" method="get" path="/payments/{paymentId}/chargebacks/{chargebackId}" example="get-chargeback-200-2" -->
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

    res, err := s.Chargebacks.Get(ctx, operations.GetChargebackRequest{
        PaymentID: "tr_5B8cwPMGnU",
        ChargebackID: "chb_xFzwUN4ci8HAmSGUACS4J",
        Embed: client.Pointer("payment"),
        IdempotencyKey: client.Pointer("123e4567-e89b-12d3-a456-426"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EntityChargeback != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetChargebackRequest](../../models/operations/getchargebackrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.GetChargebackResponse](../../models/operations/getchargebackresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## All

Retrieve all chargebacks initiated for all your payments.

The results are paginated.

### Example Usage: list-all-chargebacks-200-1

<!-- UsageSnippet language="go" operationID="list-all-chargebacks" method="get" path="/chargebacks" example="list-all-chargebacks-200-1" -->
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
        client.WithProfileID("pfl_5B8cwPMGnU"),
        client.WithTestmode(false),
        client.WithSecurity(components.Security{
            APIKey: client.Pointer(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.Chargebacks.All(ctx, operations.ListAllChargebacksRequest{
        From: client.Pointer("chb_xFzwUN4ci8HAmSGUACS4J"),
        Limit: client.Pointer[int64](50),
        Embed: client.Pointer("payment"),
        Sort: components.SortingDesc.ToPointer(),
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
### Example Usage: list-all-chargebacks-200-2

<!-- UsageSnippet language="go" operationID="list-all-chargebacks" method="get" path="/chargebacks" example="list-all-chargebacks-200-2" -->
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
        client.WithProfileID("pfl_5B8cwPMGnU"),
        client.WithTestmode(false),
        client.WithSecurity(components.Security{
            APIKey: client.Pointer(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.Chargebacks.All(ctx, operations.ListAllChargebacksRequest{
        From: client.Pointer("chb_xFzwUN4ci8HAmSGUACS4J"),
        Limit: client.Pointer[int64](50),
        Embed: client.Pointer("payment"),
        Sort: components.SortingDesc.ToPointer(),
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
### Example Usage: list-all-chargebacks-200-3

<!-- UsageSnippet language="go" operationID="list-all-chargebacks" method="get" path="/chargebacks" example="list-all-chargebacks-200-3" -->
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
        client.WithProfileID("pfl_5B8cwPMGnU"),
        client.WithTestmode(false),
        client.WithSecurity(components.Security{
            APIKey: client.Pointer(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.Chargebacks.All(ctx, operations.ListAllChargebacksRequest{
        From: client.Pointer("chb_xFzwUN4ci8HAmSGUACS4J"),
        Limit: client.Pointer[int64](50),
        Embed: client.Pointer("payment"),
        Sort: components.SortingDesc.ToPointer(),
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
| `request`                                                                                    | [operations.ListAllChargebacksRequest](../../models/operations/listallchargebacksrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.ListAllChargebacksResponse](../../models/operations/listallchargebacksresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |