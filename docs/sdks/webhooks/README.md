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
            APIKey: client.String(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.Webhooks.Create(ctx, &operations.CreateWebhookRequest{
        Name: "Webhook #1",
        URL: "https://mollie.com/",
        EventTypes: operations.CreateWebhookEventTypesPaymentLinkPaid,
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.CreateWebhookRequest](../../models/operations/createwebhookrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.CreateWebhookResponse](../../models/operations/createwebhookresponse.md), error**

### Errors

| Error Type                          | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| apierrors.CreateWebhookHalJSONError | 422                                 | application/hal+json                |
| apierrors.APIError                  | 4XX, 5XX                            | \*/\*                               |

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
        client.WithSecurity(components.Security{
            APIKey: client.String(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.Webhooks.List(ctx, operations.ListWebhooksRequest{
        From: client.String("hook_B2EyhTH5N4KWUnoYPcgiH"),
        Limit: client.Int64(50),
        Sort: operations.ListWebhooksSortDesc.ToPointer(),
        EventTypes: operations.ListWebhooksEventTypesPaymentLinkPaid.ToPointer(),
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListWebhooksRequest](../../models/operations/listwebhooksrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.ListWebhooksResponse](../../models/operations/listwebhooksresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| apierrors.ListWebhooksHalJSONError | 400                                | application/hal+json               |
| apierrors.APIError                 | 4XX, 5XX                           | \*/\*                              |

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
            APIKey: client.String(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.Webhooks.Update(ctx, "hook_B2EyhTH5N4KWUnoYPcgiH", &operations.UpdateWebhookRequestBody{
        Name: client.String("Webhook #1"),
        URL: client.String("https://mollie.com/"),
        EventTypes: operations.UpdateWebhookEventTypesPaymentLinkPaid.ToPointer(),
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

| Parameter                                                                                   | Type                                                                                        | Required                                                                                    | Description                                                                                 | Example                                                                                     |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `ctx`                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                       | :heavy_check_mark:                                                                          | The context to use for the request.                                                         |                                                                                             |
| `id`                                                                                        | *string*                                                                                    | :heavy_check_mark:                                                                          | Provide the ID of the item you want to perform this operation on.                           | hook_B2EyhTH5N4KWUnoYPcgiH                                                                  |
| `requestBody`                                                                               | [*operations.UpdateWebhookRequestBody](../../models/operations/updatewebhookrequestbody.md) | :heavy_minus_sign:                                                                          | N/A                                                                                         |                                                                                             |
| `opts`                                                                                      | [][operations.Option](../../models/operations/option.md)                                    | :heavy_minus_sign:                                                                          | The options for this request.                                                               |                                                                                             |

### Response

**[*operations.UpdateWebhookResponse](../../models/operations/updatewebhookresponse.md), error**

### Errors

| Error Type                                             | Status Code                                            | Content Type                                           |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| apierrors.UpdateWebhookNotFoundHalJSONError            | 404                                                    | application/hal+json                                   |
| apierrors.UpdateWebhookUnprocessableEntityHalJSONError | 422                                                    | application/hal+json                                   |
| apierrors.APIError                                     | 4XX, 5XX                                               | \*/\*                                                  |

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
        client.WithSecurity(components.Security{
            APIKey: client.String(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.Webhooks.Get(ctx, "hook_B2EyhTH5N4KWUnoYPcgiH", client.Bool(false))
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
| `id`                                                                                                                                                                                                                                                                                                                                                                                   | *string*                                                                                                                                                                                                                                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                     | Provide the ID of the item you want to perform this operation on.                                                                                                                                                                                                                                                                                                                      | hook_B2EyhTH5N4KWUnoYPcgiH                                                                                                                                                                                                                                                                                                                                                             |
| `testmode`                                                                                                                                                                                                                                                                                                                                                                             | **bool*                                                                                                                                                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | Most API credentials are specifically created for either live mode or test mode. In those cases the `testmode` query<br/>parameter can be omitted. For organization-level credentials such as OAuth access tokens, you can enable test mode by<br/>setting the `testmode` query parameter to `true`.<br/><br/>Test entities cannot be retrieved when the endpoint is set to live mode, and vice versa. | false                                                                                                                                                                                                                                                                                                                                                                                  |
| `opts`                                                                                                                                                                                                                                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | The options for this request.                                                                                                                                                                                                                                                                                                                                                          |                                                                                                                                                                                                                                                                                                                                                                                        |

### Response

**[*operations.GetWebhookResponse](../../models/operations/getwebhookresponse.md), error**

### Errors

| Error Type                                          | Status Code                                         | Content Type                                        |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| apierrors.GetWebhookNotFoundHalJSONError            | 404                                                 | application/hal+json                                |
| apierrors.GetWebhookUnprocessableEntityHalJSONError | 422                                                 | application/hal+json                                |
| apierrors.APIError                                  | 4XX, 5XX                                            | \*/\*                                               |

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
            APIKey: client.String(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.Webhooks.Delete(ctx, "hook_B2EyhTH5N4KWUnoYPcgiH", &operations.DeleteWebhookRequestBody{
        Testmode: client.Bool(false),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Any != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                   | Type                                                                                        | Required                                                                                    | Description                                                                                 | Example                                                                                     |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `ctx`                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                       | :heavy_check_mark:                                                                          | The context to use for the request.                                                         |                                                                                             |
| `id`                                                                                        | *string*                                                                                    | :heavy_check_mark:                                                                          | Provide the ID of the item you want to perform this operation on.                           | hook_B2EyhTH5N4KWUnoYPcgiH                                                                  |
| `requestBody`                                                                               | [*operations.DeleteWebhookRequestBody](../../models/operations/deletewebhookrequestbody.md) | :heavy_minus_sign:                                                                          | N/A                                                                                         |                                                                                             |
| `opts`                                                                                      | [][operations.Option](../../models/operations/option.md)                                    | :heavy_minus_sign:                                                                          | The options for this request.                                                               |                                                                                             |

### Response

**[*operations.DeleteWebhookResponse](../../models/operations/deletewebhookresponse.md), error**

### Errors

| Error Type                                             | Status Code                                            | Content Type                                           |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| apierrors.DeleteWebhookNotFoundHalJSONError            | 404                                                    | application/hal+json                                   |
| apierrors.DeleteWebhookUnprocessableEntityHalJSONError | 422                                                    | application/hal+json                                   |
| apierrors.APIError                                     | 4XX, 5XX                                               | \*/\*                                                  |

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
            APIKey: client.String(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.Webhooks.Test(ctx, "hook_B2EyhTH5N4KWUnoYPcgiH", &operations.TestWebhookRequestBody{
        Testmode: client.Bool(false),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Any != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                               | Type                                                                                    | Required                                                                                | Description                                                                             | Example                                                                                 |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `ctx`                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                   | :heavy_check_mark:                                                                      | The context to use for the request.                                                     |                                                                                         |
| `id`                                                                                    | *string*                                                                                | :heavy_check_mark:                                                                      | Provide the ID of the item you want to perform this operation on.                       | hook_B2EyhTH5N4KWUnoYPcgiH                                                              |
| `requestBody`                                                                           | [*operations.TestWebhookRequestBody](../../models/operations/testwebhookrequestbody.md) | :heavy_minus_sign:                                                                      | N/A                                                                                     |                                                                                         |
| `opts`                                                                                  | [][operations.Option](../../models/operations/option.md)                                | :heavy_minus_sign:                                                                      | The options for this request.                                                           |                                                                                         |

### Response

**[*operations.TestWebhookResponse](../../models/operations/testwebhookresponse.md), error**

### Errors

| Error Type                                           | Status Code                                          | Content Type                                         |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| apierrors.TestWebhookNotFoundHalJSONError            | 404                                                  | application/hal+json                                 |
| apierrors.TestWebhookUnprocessableEntityHalJSONError | 422                                                  | application/hal+json                                 |
| apierrors.APIError                                   | 4XX, 5XX                                             | \*/\*                                                |