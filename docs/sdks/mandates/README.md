# Mandates

## Overview

### Available Operations

* [Create](#create) - Create mandate
* [List](#list) - List mandates
* [Get](#get) - Get mandate
* [Revoke](#revoke) - Revoke mandate

## Create

Create a mandate for a specific customer. Mandates allow you to charge a customer's card, PayPal account or bank
account recurrently.

It is only possible to create mandates for IBANs and PayPal billing agreements with this endpoint. To create
mandates for cards, your customers need to perform a 'first payment' with their card.

### Example Usage: create-mandate-201-1

<!-- UsageSnippet language="go" operationID="create-mandate" method="post" path="/customers/{customerId}/mandates" example="create-mandate-201-1" -->
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

    res, err := s.Mandates.Create(ctx, "cst_5B8cwPMGnU", client.Pointer("123e4567-e89b-12d3-a456-426"), &components.MandateRequest{
        ID: client.Pointer("mdt_5B8cwPMGnU"),
        Method: components.MandateMethodDirectdebit,
        ConsumerName: "John Doe",
        ConsumerAccount: client.Pointer("NL55INGB0000000000"),
        ConsumerBic: client.Pointer("BANKBIC"),
        ConsumerEmail: client.Pointer("example@email.com"),
        SignatureDate: client.Pointer("2025-01-01"),
        MandateReference: client.Pointer("ID-1023892"),
        PaypalBillingAgreementID: client.Pointer("B-12A34567B8901234CD"),
        PayPalVaultID: client.Pointer("8kk8451t"),
        Testmode: client.Pointer(false),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MandateResponse != nil {
        // handle response
    }
}
```
### Example Usage: create-mandate-201-2

<!-- UsageSnippet language="go" operationID="create-mandate" method="post" path="/customers/{customerId}/mandates" example="create-mandate-201-2" -->
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

    res, err := s.Mandates.Create(ctx, "cst_5B8cwPMGnU", client.Pointer("123e4567-e89b-12d3-a456-426"), &components.MandateRequest{
        ID: client.Pointer("mdt_5B8cwPMGnU"),
        Method: components.MandateMethodDirectdebit,
        ConsumerName: "John Doe",
        ConsumerAccount: client.Pointer("NL55INGB0000000000"),
        ConsumerBic: client.Pointer("BANKBIC"),
        ConsumerEmail: client.Pointer("example@email.com"),
        SignatureDate: client.Pointer("2025-01-01"),
        MandateReference: client.Pointer("ID-1023892"),
        PaypalBillingAgreementID: client.Pointer("B-12A34567B8901234CD"),
        PayPalVaultID: client.Pointer("8kk8451t"),
        Testmode: client.Pointer(false),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MandateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |                                                                                  |
| `customerID`                                                                     | *string*                                                                         | :heavy_check_mark:                                                               | Provide the ID of the related customer.                                          | cst_5B8cwPMGnU                                                                   |
| `idempotencyKey`                                                                 | **string*                                                                        | :heavy_minus_sign:                                                               | A unique key to ensure idempotent requests. This key should be a UUID v4 string. | 123e4567-e89b-12d3-a456-426                                                      |
| `mandateRequest`                                                                 | [*components.MandateRequest](../../models/components/mandaterequest.md)          | :heavy_minus_sign:                                                               | N/A                                                                              |                                                                                  |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |                                                                                  |

### Response

**[*operations.CreateMandateResponse](../../models/operations/createmandateresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## List

Retrieve a list of all mandates.

The results are paginated.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-mandates" method="get" path="/customers/{customerId}/mandates" example="list-mandates-200-1" -->
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

    res, err := s.Mandates.List(ctx, operations.ListMandatesRequest{
        CustomerID: "cst_5B8cwPMGnU",
        From: client.Pointer("mdt_5B8cwPMGnU"),
        Limit: client.Pointer[int64](50),
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListMandatesRequest](../../models/operations/listmandatesrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.ListMandatesResponse](../../models/operations/listmandatesresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Get

Retrieve a single mandate by its ID. Depending on the type of mandate, the object will contain the customer's bank
account details, card details, or PayPal account details.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-mandate" method="get" path="/customers/{customerId}/mandates/{mandateId}" example="get-mandate-200-1" -->
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

    res, err := s.Mandates.Get(ctx, "cst_5B8cwPMGnU", "mdt_5B8cwPMGnU", client.Pointer("123e4567-e89b-12d3-a456-426"))
    if err != nil {
        log.Fatal(err)
    }
    if res.MandateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                                | Type                                                                                                                                                                                                                                                                                                                                                                                     | Required                                                                                                                                                                                                                                                                                                                                                                                 | Description                                                                                                                                                                                                                                                                                                                                                                              | Example                                                                                                                                                                                                                                                                                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                       | The context to use for the request.                                                                                                                                                                                                                                                                                                                                                      |                                                                                                                                                                                                                                                                                                                                                                                          |
| `customerID`                                                                                                                                                                                                                                                                                                                                                                             | *string*                                                                                                                                                                                                                                                                                                                                                                                 | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                       | Provide the ID of the related customer.                                                                                                                                                                                                                                                                                                                                                  | cst_5B8cwPMGnU                                                                                                                                                                                                                                                                                                                                                                           |
| `mandateID`                                                                                                                                                                                                                                                                                                                                                                              | *string*                                                                                                                                                                                                                                                                                                                                                                                 | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                       | Provide the ID of the related mandate.                                                                                                                                                                                                                                                                                                                                                   | mdt_5B8cwPMGnU                                                                                                                                                                                                                                                                                                                                                                           |
| `testmode`                                                                                                                                                                                                                                                                                                                                                                               | **bool*                                                                                                                                                                                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                       | Most API credentials are specifically created for either live mode or test mode. In those cases the `testmode` query<br/>parameter must not be sent. For organization-level credentials such as OAuth access tokens, you can enable test mode by<br/>setting the `testmode` query parameter to `true`.<br/><br/>Test entities cannot be retrieved when the endpoint is set to live mode, and vice versa. |                                                                                                                                                                                                                                                                                                                                                                                          |
| `idempotencyKey`                                                                                                                                                                                                                                                                                                                                                                         | **string*                                                                                                                                                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                       | A unique key to ensure idempotent requests. This key should be a UUID v4 string.                                                                                                                                                                                                                                                                                                         | 123e4567-e89b-12d3-a456-426                                                                                                                                                                                                                                                                                                                                                              |
| `opts`                                                                                                                                                                                                                                                                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                       | The options for this request.                                                                                                                                                                                                                                                                                                                                                            |                                                                                                                                                                                                                                                                                                                                                                                          |

### Response

**[*operations.GetMandateResponse](../../models/operations/getmandateresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Revoke

Revoke a customer's mandate. You will no longer be able to charge the customer's bank account or card with this
mandate, and all connected subscriptions will be canceled.

### Example Usage

<!-- UsageSnippet language="go" operationID="revoke-mandate" method="delete" path="/customers/{customerId}/mandates/{mandateId}" -->
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

    res, err := s.Mandates.Revoke(ctx, "cst_5B8cwPMGnU", "mdt_5B8cwPMGnU", client.Pointer("123e4567-e89b-12d3-a456-426"), &operations.RevokeMandateRequestBody{
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
| `customerID`                                                                                | *string*                                                                                    | :heavy_check_mark:                                                                          | Provide the ID of the related customer.                                                     | cst_5B8cwPMGnU                                                                              |
| `mandateID`                                                                                 | *string*                                                                                    | :heavy_check_mark:                                                                          | Provide the ID of the related mandate.                                                      | mdt_5B8cwPMGnU                                                                              |
| `idempotencyKey`                                                                            | **string*                                                                                   | :heavy_minus_sign:                                                                          | A unique key to ensure idempotent requests. This key should be a UUID v4 string.            | 123e4567-e89b-12d3-a456-426                                                                 |
| `requestBody`                                                                               | [*operations.RevokeMandateRequestBody](../../models/operations/revokemandaterequestbody.md) | :heavy_minus_sign:                                                                          | N/A                                                                                         |                                                                                             |
| `opts`                                                                                      | [][operations.Option](../../models/operations/option.md)                                    | :heavy_minus_sign:                                                                          | The options for this request.                                                               |                                                                                             |

### Response

**[*operations.RevokeMandateResponse](../../models/operations/revokemandateresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |