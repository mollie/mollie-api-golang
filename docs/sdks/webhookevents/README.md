# WebhookEvents

## Overview

### Available Operations

* [Get](#get) - Get a Webhook Event

## Get

Retrieve a single webhook event object by its event ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-webhook-event" method="get" path="/v2/events/{webhookEventId}" example="get-webhook-event-200" -->
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
            AdvancedAccessToken: client.Pointer(os.Getenv("CLIENT_ADVANCED_ACCESS_TOKEN")),
        }),
    )

    res, err := s.WebhookEvents.Get(ctx, operations.GetWebhookEventRequest{
        WebhookEventID: "event_1234567890",
        IdempotencyKey: client.Pointer("123e4567-e89b-12d3-a456-426"),
        XClientSignature: client.Pointer("MEUCIQDx8sJvK3z..."),
        XClientSignedAt: client.Pointer("2025-01-01T12:00:00Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EntityWebhookEvent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetWebhookEventRequest](../../models/operations/getwebhookeventrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetWebhookEventResponse](../../models/operations/getwebhookeventresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 429                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |