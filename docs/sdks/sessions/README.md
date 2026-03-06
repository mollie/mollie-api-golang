# Sessions

## Overview

### Available Operations

* [Create](#create) - Create session [BETA]
* [Get](#get) - Get session

## Create

> 🚧 Beta feature
>
> This feature is currently in private beta, and the final specification may still change.

Create a session to start a checkout process with Mollie Components.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-session" method="post" path="/sessions" example="create-session-201-1" -->
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

    res, err := s.Sessions.Create(ctx, client.Pointer("123e4567-e89b-12d3-a456-426"), &components.SessionRequest{
        Amount: components.Amount{
            Currency: "EUR",
            Value: "10.00",
        },
        Description: "Order #12345",
        RedirectURL: "https://example.org/redirect",
        BillingAddress: &components.PaymentAddress{
            Title: client.Pointer("Mr."),
            GivenName: client.Pointer("Piet"),
            FamilyName: client.Pointer("Mondriaan"),
            OrganizationName: client.Pointer("Mollie B.V."),
            StreetAndNumber: client.Pointer("Keizersgracht 126"),
            StreetAdditional: client.Pointer("Apt. 1"),
            PostalCode: client.Pointer("1234AB"),
            Email: client.Pointer("piet@example.org"),
            Phone: client.Pointer("31208202070"),
            City: client.Pointer("Amsterdam"),
            Region: client.Pointer("Noord-Holland"),
            Country: client.Pointer("NL"),
        },
        ShippingAddress: &components.PaymentAddress{
            Title: client.Pointer("Mr."),
            GivenName: client.Pointer("Piet"),
            FamilyName: client.Pointer("Mondriaan"),
            OrganizationName: client.Pointer("Mollie B.V."),
            StreetAndNumber: client.Pointer("Keizersgracht 126"),
            StreetAdditional: client.Pointer("Apt. 1"),
            PostalCode: client.Pointer("1234AB"),
            Email: client.Pointer("piet@example.org"),
            Phone: client.Pointer("31208202070"),
            City: client.Pointer("Amsterdam"),
            Region: client.Pointer("Noord-Holland"),
            Country: client.Pointer("NL"),
        },
        CustomerID: client.Pointer("cst_5B8cwPMGnU"),
        SequenceType: components.SessionSequenceTypeOneoff.ToPointer(),
        Payment: &components.SessionRequestPayment{
            WebhookURL: client.Pointer("https://example.org/webhook"),
        },
        Lines: []components.SessionLineItem{},
        ProfileID: client.Pointer("pfl_5B8cwPMGnU"),
        Testmode: client.Pointer(false),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SessionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |                                                                                  |
| `idempotencyKey`                                                                 | **string*                                                                        | :heavy_minus_sign:                                                               | A unique key to ensure idempotent requests. This key should be a UUID v4 string. | 123e4567-e89b-12d3-a456-426                                                      |
| `sessionRequest`                                                                 | [*components.SessionRequest](../../models/components/sessionrequest.md)          | :heavy_minus_sign:                                                               | N/A                                                                              |                                                                                  |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |                                                                                  |

### Response

**[*operations.CreateSessionResponse](../../models/operations/createsessionresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 422                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Get

> 🚧 Beta feature
>
> This feature is currently in private beta, and the final specification may still change.

Retrieve a session to view its details and status to inform your customers about the checkout process.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-session" method="get" path="/sessions/{sessionId}" example="get-session-200-1" -->
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

    res, err := s.Sessions.Get(ctx, "sess_82jFYDTrLcCQV68NLDvMJ", client.Pointer("123e4567-e89b-12d3-a456-426"))
    if err != nil {
        log.Fatal(err)
    }
    if res.SessionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |                                                                                  |
| `sessionID`                                                                      | *string*                                                                         | :heavy_check_mark:                                                               | Provide the ID of the related session.                                           | sess_82jFYDTrLcCQV68NLDvMJ                                                       |
| `idempotencyKey`                                                                 | **string*                                                                        | :heavy_minus_sign:                                                               | A unique key to ensure idempotent requests. This key should be a UUID v4 string. | 123e4567-e89b-12d3-a456-426                                                      |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |                                                                                  |

### Response

**[*operations.GetSessionResponse](../../models/operations/getsessionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |