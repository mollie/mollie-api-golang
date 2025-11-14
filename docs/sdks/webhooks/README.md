# Webhooks
(*Webhooks*)

## Overview

### Available Operations

* [Create](#create) - Create a webhook
* [List](#list) - List all webhooks
* [Update](#update) - Update a webhook
* [Get](#get) - Get a webhook
* [Delete](#delete) - Delete a webhook
* [Test](#test) - Test a webhook

## Create

A webhook must have a name, an url and a list of event types. You can also create webhooks in the webhooks settings section of the Dashboard.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-webhook" method="post" path="/webhooks" -->
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

    res, err := s.Webhooks.Create(ctx, client.Pointer("123e4567-e89b-12d3-a456-426"), &operations.CreateWebhookRequestBody{
        Name: "Webhook #1",
        URL: "https://mollie.com/",
        WebhookEventTypes: components.WebhookEventTypesPaymentLinkPaid,
        Testmode: client.Pointer(false),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateWebhook != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                   | Type                                                                                        | Required                                                                                    | Description                                                                                 | Example                                                                                     |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `ctx`                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                       | :heavy_check_mark:                                                                          | The context to use for the request.                                                         |                                                                                             |
| `idempotencyKey`                                                                            | **string*                                                                                   | :heavy_minus_sign:                                                                          | A unique key to ensure idempotent requests. This key should be a UUID v4 string.            | 123e4567-e89b-12d3-a456-426                                                                 |
| `requestBody`                                                                               | [*operations.CreateWebhookRequestBody](../../models/operations/createwebhookrequestbody.md) | :heavy_minus_sign:                                                                          | N/A                                                                                         |                                                                                             |
| `opts`                                                                                      | [][operations.Option](../../models/operations/option.md)                                    | :heavy_minus_sign:                                                                          | The options for this request.                                                               |                                                                                             |

### Response

**[*operations.CreateWebhookResponse](../../models/operations/createwebhookresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 422                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## List

Returns a paginated list of your webhooks. If no webhook endpoints are available, the resulting array will be empty. This request should never throw an error.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-webhooks" method="get" path="/webhooks" -->
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

    res, err := s.Webhooks.List(ctx, operations.ListWebhooksRequest{
        From: client.Pointer("hook_B2EyhTH5N4KWUnoYPcgiH"),
        Limit: client.Pointer[int64](50),
        Sort: components.SortingDesc.ToPointer(),
        EventTypes: components.WebhookEventTypesPaymentLinkPaid.ToPointer(),
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
| `request`                                                                        | [operations.ListWebhooksRequest](../../models/operations/listwebhooksrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.ListWebhooksResponse](../../models/operations/listwebhooksresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Update

Updates the webhook. You may edit the name, url and the list of subscribed event types.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-webhook" method="patch" path="/webhooks/{id}" -->
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

    res, err := s.Webhooks.Update(ctx, "hook_B2EyhTH5N4KWUnoYPcgiH", client.Pointer("123e4567-e89b-12d3-a456-426"), &operations.UpdateWebhookRequestBody{
        Name: client.Pointer("Webhook #1"),
        URL: client.Pointer("https://mollie.com/"),
        WebhookEventTypes: components.WebhookEventTypesPaymentLinkPaid.ToPointer(),
        Testmode: client.Pointer(false),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EntityWebhook != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                   | Type                                                                                        | Required                                                                                    | Description                                                                                 | Example                                                                                     |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `ctx`                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                       | :heavy_check_mark:                                                                          | The context to use for the request.                                                         |                                                                                             |
| `id`                                                                                        | *string*                                                                                    | :heavy_check_mark:                                                                          | Provide the ID of the item you want to perform this operation on.                           |                                                                                             |
| `idempotencyKey`                                                                            | **string*                                                                                   | :heavy_minus_sign:                                                                          | A unique key to ensure idempotent requests. This key should be a UUID v4 string.            | 123e4567-e89b-12d3-a456-426                                                                 |
| `requestBody`                                                                               | [*operations.UpdateWebhookRequestBody](../../models/operations/updatewebhookrequestbody.md) | :heavy_minus_sign:                                                                          | N/A                                                                                         |                                                                                             |
| `opts`                                                                                      | [][operations.Option](../../models/operations/option.md)                                    | :heavy_minus_sign:                                                                          | The options for this request.                                                               |                                                                                             |

### Response

**[*operations.UpdateWebhookResponse](../../models/operations/updatewebhookresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 422                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Get

Retrieve a single webhook object by its ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-webhook" method="get" path="/webhooks/{id}" -->
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
            APIKey: client.Pointer(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.Webhooks.Get(ctx, "hook_B2EyhTH5N4KWUnoYPcgiH", client.Pointer("123e4567-e89b-12d3-a456-426"))
    if err != nil {
        log.Fatal(err)
    }
    if res.EntityWebhook != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                              | Type                                                                                                                                                                                                                                                                                                                                                                                   | Required                                                                                                                                                                                                                                                                                                                                                                               | Description                                                                                                                                                                                                                                                                                                                                                                            | Example                                                                                                                                                                                                                                                                                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                     | The context to use for the request.                                                                                                                                                                                                                                                                                                                                                    |                                                                                                                                                                                                                                                                                                                                                                                        |
| `id`                                                                                                                                                                                                                                                                                                                                                                                   | *string*                                                                                                                                                                                                                                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                     | Provide the ID of the item you want to perform this operation on.                                                                                                                                                                                                                                                                                                                      |                                                                                                                                                                                                                                                                                                                                                                                        |
| `testmode`                                                                                                                                                                                                                                                                                                                                                                             | **bool*                                                                                                                                                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | Most API credentials are specifically created for either live mode or test mode. In those cases the `testmode` query<br/>parameter can be omitted. For organization-level credentials such as OAuth access tokens, you can enable test mode by<br/>setting the `testmode` query parameter to `true`.<br/><br/>Test entities cannot be retrieved when the endpoint is set to live mode, and vice versa. |                                                                                                                                                                                                                                                                                                                                                                                        |
| `idempotencyKey`                                                                                                                                                                                                                                                                                                                                                                       | **string*                                                                                                                                                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | A unique key to ensure idempotent requests. This key should be a UUID v4 string.                                                                                                                                                                                                                                                                                                       | 123e4567-e89b-12d3-a456-426                                                                                                                                                                                                                                                                                                                                                            |
| `opts`                                                                                                                                                                                                                                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | The options for this request.                                                                                                                                                                                                                                                                                                                                                          |                                                                                                                                                                                                                                                                                                                                                                                        |

### Response

**[*operations.GetWebhookResponse](../../models/operations/getwebhookresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 422                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Delete

Delete a single webhook object by its webhook ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-webhook" method="delete" path="/webhooks/{id}" -->
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

    res, err := s.Webhooks.Delete(ctx, "hook_B2EyhTH5N4KWUnoYPcgiH", client.Pointer("123e4567-e89b-12d3-a456-426"), &operations.DeleteWebhookRequestBody{
        Testmode: client.Pointer(false),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                   | Type                                                                                        | Required                                                                                    | Description                                                                                 | Example                                                                                     |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `ctx`                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                       | :heavy_check_mark:                                                                          | The context to use for the request.                                                         |                                                                                             |
| `id`                                                                                        | *string*                                                                                    | :heavy_check_mark:                                                                          | Provide the ID of the item you want to perform this operation on.                           |                                                                                             |
| `idempotencyKey`                                                                            | **string*                                                                                   | :heavy_minus_sign:                                                                          | A unique key to ensure idempotent requests. This key should be a UUID v4 string.            | 123e4567-e89b-12d3-a456-426                                                                 |
| `requestBody`                                                                               | [*operations.DeleteWebhookRequestBody](../../models/operations/deletewebhookrequestbody.md) | :heavy_minus_sign:                                                                          | N/A                                                                                         |                                                                                             |
| `opts`                                                                                      | [][operations.Option](../../models/operations/option.md)                                    | :heavy_minus_sign:                                                                          | The options for this request.                                                               |                                                                                             |

### Response

**[*operations.DeleteWebhookResponse](../../models/operations/deletewebhookresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 422                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Test

Sends a test event to the webhook to verify the endpoint is working as expected.

### Example Usage

<!-- UsageSnippet language="go" operationID="test-webhook" method="post" path="/webhooks/{id}/ping" -->
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

    res, err := s.Webhooks.Test(ctx, "hook_B2EyhTH5N4KWUnoYPcgiH", client.Pointer("123e4567-e89b-12d3-a456-426"), &operations.TestWebhookRequestBody{
        Testmode: client.Pointer(false),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                               | Type                                                                                    | Required                                                                                | Description                                                                             | Example                                                                                 |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `ctx`                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                   | :heavy_check_mark:                                                                      | The context to use for the request.                                                     |                                                                                         |
| `id`                                                                                    | *string*                                                                                | :heavy_check_mark:                                                                      | Provide the ID of the item you want to perform this operation on.                       |                                                                                         |
| `idempotencyKey`                                                                        | **string*                                                                               | :heavy_minus_sign:                                                                      | A unique key to ensure idempotent requests. This key should be a UUID v4 string.        | 123e4567-e89b-12d3-a456-426                                                             |
| `requestBody`                                                                           | [*operations.TestWebhookRequestBody](../../models/operations/testwebhookrequestbody.md) | :heavy_minus_sign:                                                                      | N/A                                                                                     |                                                                                         |
| `opts`                                                                                  | [][operations.Option](../../models/operations/option.md)                                | :heavy_minus_sign:                                                                      | The options for this request.                                                           |                                                                                         |

### Response

**[*operations.TestWebhookResponse](../../models/operations/testwebhookresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 422                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |