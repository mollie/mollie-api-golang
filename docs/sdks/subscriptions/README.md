# Subscriptions
(*Subscriptions*)

## Overview

### Available Operations

* [Create](#create) - Create subscription
* [List](#list) - List customer subscriptions
* [Get](#get) - Get subscription
* [Update](#update) - Update subscription
* [Cancel](#cancel) - Cancel subscription
* [All](#all) - List all subscriptions
* [ListPayments](#listpayments) - List subscription payments

## Create

With subscriptions, you can schedule recurring payments to take place at regular intervals.

For example, by simply specifying an `amount` and an `interval`, you can create an endless subscription to charge a
monthly fee, until you cancel the subscription.

Or, you could use the times parameter to only charge a limited number of times, for example to split a big
transaction in multiple parts.

A few example usages:

`amount[currency]="EUR"` `amount[value]="5.00"` `interval="2 weeks"`
Your customer will be charged €5 once every two weeks.

`amount[currency]="EUR"` `amount[value]="20.00"` `interval="1 day" times=5`
Your customer will be charged €20 every day, for five consecutive days.

`amount[currency]="EUR"` `amount[value]="10.00"` `interval="1 month"`
`startDate="2018-04-30"`
Your customer will be charged €10 on the last day of each month, starting in April 2018.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-subscription" method="post" path="/customers/{customerId}/subscriptions" -->
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
            APIKey: client.String(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.Subscriptions.Create(ctx, "cst_5B8cwPMGnU", client.String("123e4567-e89b-12d3-a456-426"), &components.SubscriptionRequest{
        ID: client.String("sub_5B8cwPMGnU"),
        Status: components.SubscriptionStatusActive.ToPointer(),
        Amount: &components.Amount{
            Currency: "EUR",
            Value: "10.00",
        },
        Times: client.Int64(6),
        Interval: client.String("2 days"),
        StartDate: client.String("2025-01-01"),
        Description: client.String("Subscription of streaming channel"),
        Method: components.SubscriptionMethodPaypal.ToPointer(),
        ApplicationFee: &components.SubscriptionRequestApplicationFee{
            Amount: components.Amount{
                Currency: "EUR",
                Value: "10.00",
            },
            Description: "Platform fee",
        },
        WebhookURL: client.String("https://example.com/webhook"),
        CustomerID: client.String("cst_5B8cwPMGnU"),
        MandateID: client.String("mdt_5B8cwPMGnU"),
        Testmode: client.Bool(false),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SubscriptionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                         | Type                                                                              | Required                                                                          | Description                                                                       | Example                                                                           |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `ctx`                                                                             | [context.Context](https://pkg.go.dev/context#Context)                             | :heavy_check_mark:                                                                | The context to use for the request.                                               |                                                                                   |
| `customerID`                                                                      | *string*                                                                          | :heavy_check_mark:                                                                | Provide the ID of the related customer.                                           | cst_5B8cwPMGnU                                                                    |
| `idempotencyKey`                                                                  | **string*                                                                         | :heavy_minus_sign:                                                                | A unique key to ensure idempotent requests. This key should be a UUID v4 string.  | 123e4567-e89b-12d3-a456-426                                                       |
| `subscriptionRequest`                                                             | [*components.SubscriptionRequest](../../models/components/subscriptionrequest.md) | :heavy_minus_sign:                                                                | N/A                                                                               |                                                                                   |
| `opts`                                                                            | [][operations.Option](../../models/operations/option.md)                          | :heavy_minus_sign:                                                                | The options for this request.                                                     |                                                                                   |

### Response

**[*operations.CreateSubscriptionResponse](../../models/operations/createsubscriptionresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## List

Retrieve all subscriptions of a customer.

The results are paginated.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-subscriptions" method="get" path="/customers/{customerId}/subscriptions" -->
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

    res, err := s.Subscriptions.List(ctx, operations.ListSubscriptionsRequest{
        CustomerID: "cst_5B8cwPMGnU",
        From: client.String("sub_5B8cwPMGnU"),
        Limit: client.Int64(50),
        Sort: components.ListSortDesc.ToPointer(),
        Testmode: client.Bool(false),
        IdempotencyKey: client.String("123e4567-e89b-12d3-a456-426"),
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListSubscriptionsRequest](../../models/operations/listsubscriptionsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListSubscriptionsResponse](../../models/operations/listsubscriptionsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Get

Retrieve a single subscription by its ID and the ID of its parent customer.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-subscription" method="get" path="/customers/{customerId}/subscriptions/{subscriptionId}" -->
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
            APIKey: client.String(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.Subscriptions.Get(ctx, "cst_5B8cwPMGnU", "sub_5B8cwPMGnU", client.Bool(false), client.String("123e4567-e89b-12d3-a456-426"))
    if err != nil {
        log.Fatal(err)
    }
    if res.SubscriptionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                              | Type                                                                                                                                                                                                                                                                                                                                                                                   | Required                                                                                                                                                                                                                                                                                                                                                                               | Description                                                                                                                                                                                                                                                                                                                                                                            | Example                                                                                                                                                                                                                                                                                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                     | The context to use for the request.                                                                                                                                                                                                                                                                                                                                                    |                                                                                                                                                                                                                                                                                                                                                                                        |
| `customerID`                                                                                                                                                                                                                                                                                                                                                                           | *string*                                                                                                                                                                                                                                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                     | Provide the ID of the related customer.                                                                                                                                                                                                                                                                                                                                                | cst_5B8cwPMGnU                                                                                                                                                                                                                                                                                                                                                                         |
| `subscriptionID`                                                                                                                                                                                                                                                                                                                                                                       | *string*                                                                                                                                                                                                                                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                     | Provide the ID of the related subscription.                                                                                                                                                                                                                                                                                                                                            | sub_5B8cwPMGnU                                                                                                                                                                                                                                                                                                                                                                         |
| `testmode`                                                                                                                                                                                                                                                                                                                                                                             | **bool*                                                                                                                                                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | Most API credentials are specifically created for either live mode or test mode. In those cases the `testmode` query<br/>parameter can be omitted. For organization-level credentials such as OAuth access tokens, you can enable test mode by<br/>setting the `testmode` query parameter to `true`.<br/><br/>Test entities cannot be retrieved when the endpoint is set to live mode, and vice versa. | false                                                                                                                                                                                                                                                                                                                                                                                  |
| `idempotencyKey`                                                                                                                                                                                                                                                                                                                                                                       | **string*                                                                                                                                                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | A unique key to ensure idempotent requests. This key should be a UUID v4 string.                                                                                                                                                                                                                                                                                                       | 123e4567-e89b-12d3-a456-426                                                                                                                                                                                                                                                                                                                                                            |
| `opts`                                                                                                                                                                                                                                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | The options for this request.                                                                                                                                                                                                                                                                                                                                                          |                                                                                                                                                                                                                                                                                                                                                                                        |

### Response

**[*operations.GetSubscriptionResponse](../../models/operations/getsubscriptionresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Update

Update an existing subscription.

Canceled subscriptions cannot be updated.

For an in-depth explanation of each parameter, refer to the [Create subscription](create-subscription) endpoint.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-subscription" method="patch" path="/customers/{customerId}/subscriptions/{subscriptionId}" -->
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

    res, err := s.Subscriptions.Update(ctx, "cst_5B8cwPMGnU", "sub_5B8cwPMGnU", client.String("123e4567-e89b-12d3-a456-426"), &operations.UpdateSubscriptionRequestBody{
        Amount: &components.Amount{
            Currency: "EUR",
            Value: "10.00",
        },
        Description: client.String("Subscription of streaming channel"),
        Interval: client.String("1 months"),
        StartDate: client.String("2025-01-01"),
        Times: client.Int64(6),
        WebhookURL: client.String("https://example.com/webhook"),
        MandateID: client.String("mdt_5B8cwPMGnU"),
        Testmode: client.Bool(false),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SubscriptionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                             | Type                                                                                                  | Required                                                                                              | Description                                                                                           | Example                                                                                               |
| ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                 | :heavy_check_mark:                                                                                    | The context to use for the request.                                                                   |                                                                                                       |
| `customerID`                                                                                          | *string*                                                                                              | :heavy_check_mark:                                                                                    | Provide the ID of the related customer.                                                               | cst_5B8cwPMGnU                                                                                        |
| `subscriptionID`                                                                                      | *string*                                                                                              | :heavy_check_mark:                                                                                    | Provide the ID of the related subscription.                                                           | sub_5B8cwPMGnU                                                                                        |
| `idempotencyKey`                                                                                      | **string*                                                                                             | :heavy_minus_sign:                                                                                    | A unique key to ensure idempotent requests. This key should be a UUID v4 string.                      | 123e4567-e89b-12d3-a456-426                                                                           |
| `requestBody`                                                                                         | [*operations.UpdateSubscriptionRequestBody](../../models/operations/updatesubscriptionrequestbody.md) | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |                                                                                                       |
| `opts`                                                                                                | [][operations.Option](../../models/operations/option.md)                                              | :heavy_minus_sign:                                                                                    | The options for this request.                                                                         |                                                                                                       |

### Response

**[*operations.UpdateSubscriptionResponse](../../models/operations/updatesubscriptionresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Cancel

Cancel an existing subscription. Canceling a subscription has no effect on the mandates of the customer.

### Example Usage

<!-- UsageSnippet language="go" operationID="cancel-subscription" method="delete" path="/customers/{customerId}/subscriptions/{subscriptionId}" -->
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

    res, err := s.Subscriptions.Cancel(ctx, "cst_5B8cwPMGnU", "sub_5B8cwPMGnU", client.String("123e4567-e89b-12d3-a456-426"), &operations.CancelSubscriptionRequestBody{
        Testmode: client.Bool(false),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SubscriptionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                             | Type                                                                                                  | Required                                                                                              | Description                                                                                           | Example                                                                                               |
| ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                 | :heavy_check_mark:                                                                                    | The context to use for the request.                                                                   |                                                                                                       |
| `customerID`                                                                                          | *string*                                                                                              | :heavy_check_mark:                                                                                    | Provide the ID of the related customer.                                                               | cst_5B8cwPMGnU                                                                                        |
| `subscriptionID`                                                                                      | *string*                                                                                              | :heavy_check_mark:                                                                                    | Provide the ID of the related subscription.                                                           | sub_5B8cwPMGnU                                                                                        |
| `idempotencyKey`                                                                                      | **string*                                                                                             | :heavy_minus_sign:                                                                                    | A unique key to ensure idempotent requests. This key should be a UUID v4 string.                      | 123e4567-e89b-12d3-a456-426                                                                           |
| `requestBody`                                                                                         | [*operations.CancelSubscriptionRequestBody](../../models/operations/cancelsubscriptionrequestbody.md) | :heavy_minus_sign:                                                                                    | N/A                                                                                                   |                                                                                                       |
| `opts`                                                                                                | [][operations.Option](../../models/operations/option.md)                                              | :heavy_minus_sign:                                                                                    | The options for this request.                                                                         |                                                                                                       |

### Response

**[*operations.CancelSubscriptionResponse](../../models/operations/cancelsubscriptionresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## All

Retrieve all subscriptions initiated across all your customers.

The results are paginated.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-all-subscriptions" method="get" path="/subscriptions" -->
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

    res, err := s.Subscriptions.All(ctx, operations.ListAllSubscriptionsRequest{
        From: client.String("sub_rVKGtNd6s3"),
        Limit: client.Int64(50),
        ProfileID: client.String("pfl_QkEhN94Ba"),
        Testmode: client.Bool(false),
        IdempotencyKey: client.String("123e4567-e89b-12d3-a456-426"),
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListAllSubscriptionsRequest](../../models/operations/listallsubscriptionsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListAllSubscriptionsResponse](../../models/operations/listallsubscriptionsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## ListPayments

Retrieve all payments of a specific subscription.

The results are paginated.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-subscription-payments" method="get" path="/customers/{customerId}/subscriptions/{subscriptionId}/payments" -->
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

    res, err := s.Subscriptions.ListPayments(ctx, operations.ListSubscriptionPaymentsRequest{
        CustomerID: "cst_5B8cwPMGnU",
        SubscriptionID: "sub_5B8cwPMGnU",
        From: client.String("tr_5B8cwPMGnU"),
        Limit: client.Int64(50),
        Sort: components.ListSortDesc.ToPointer(),
        ProfileID: client.String("pfl_5B8cwPMGnU"),
        Testmode: client.Bool(false),
        IdempotencyKey: client.String("123e4567-e89b-12d3-a456-426"),
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ListSubscriptionPaymentsRequest](../../models/operations/listsubscriptionpaymentsrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.ListSubscriptionPaymentsResponse](../../models/operations/listsubscriptionpaymentsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |