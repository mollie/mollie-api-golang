# SalesInvoices

## Overview

### Available Operations

* [Create](#create) - Create sales invoice
* [List](#list) - List sales invoices
* [Get](#get) - Get sales invoice
* [Update](#update) - Update sales invoice
* [Delete](#delete) - Delete sales invoice

## Create

With the Sales Invoice API you can generate sales invoices to send to your customers.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-sales-invoice" method="post" path="/v2/sales-invoices" -->
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

    res, err := s.SalesInvoices.Create(ctx, client.Pointer("123e4567-e89b-12d3-a456-426"), client.Pointer("MEUCIQDx8sJvK3z..."), client.Pointer("2025-01-01T12:00:00Z"), &components.SalesInvoiceRequest{
        Testmode: client.Pointer(false),
        ProfileID: client.Pointer("pfl_QkEhN94Ba"),
        Status: components.SalesInvoiceStatusDraft,
        VatScheme: components.SalesInvoiceVatSchemeStandard.ToPointer(),
        VatMode: components.SalesInvoiceVatModeExclusive.ToPointer(),
        Memo: client.Pointer("This is a memo!"),
        PaymentTerm: components.SalesInvoicePaymentTermThirtydays.ToPointer(),
        PaymentDetails: &components.SalesInvoicePaymentDetails{
            Source: components.SalesInvoicePaymentDetailsSourcePaymentLink,
            SourceReference: client.Pointer("pl_d9fQur83kFdhH8hIhaZfq"),
        },
        EmailDetails: &components.SalesInvoiceEmailDetails{
            Subject: "Your invoice is available",
            Body: "Please find your invoice enclosed.",
        },
        CustomerID: client.Pointer("cst_8wmqcHMN4U"),
        MandateID: client.Pointer("mdt_pWUnw6pkBN"),
        RecipientIdentifier: "customer-xyz-0123",
        Recipient: &components.SalesInvoiceRecipient{
            Type: components.SalesInvoiceRecipientTypeConsumer,
            Title: client.Pointer("Mrs."),
            GivenName: client.Pointer("Jane"),
            FamilyName: client.Pointer("Doe"),
            OrganizationName: client.Pointer("Organization Corp."),
            OrganizationNumber: client.Pointer("12345678"),
            VatNumber: client.Pointer("NL123456789B01"),
            Email: "example@email.com",
            Phone: client.Pointer("+0123456789"),
            StreetAndNumber: "Keizersgracht 126",
            StreetAdditional: client.Pointer("4th floor"),
            PostalCode: "5678AB",
            City: "Amsterdam",
            Region: client.Pointer("Noord-Holland"),
            Country: "NL",
            Locale: components.SalesInvoiceRecipientLocaleNlNl,
        },
        Lines: []components.SalesInvoiceLineItem{},
        Discount: &components.SalesInvoiceDiscount{
            Type: components.SalesInvoiceDiscountTypeAmount,
            Value: "10.00",
        },
        IsEInvoice: client.Pointer(false),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SalesInvoiceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                               | Type                                                                                                                                                                                    | Required                                                                                                                                                                                | Description                                                                                                                                                                             | Example                                                                                                                                                                                 |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                   | :heavy_check_mark:                                                                                                                                                                      | The context to use for the request.                                                                                                                                                     |                                                                                                                                                                                         |
| `idempotencyKey`                                                                                                                                                                        | `*string`                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                      | A unique key to ensure idempotent requests. This key should be a UUID v4 string.                                                                                                        | 123e4567-e89b-12d3-a456-426                                                                                                                                                             |
| `xClientSignature`                                                                                                                                                                      | `*string`                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                      | A cryptographic signature of the request, used to verify that it was sent by you and hasn't been tampered with. Required if you've enabled request signing for your integration.        | MEUCIQDx8sJvK3z...                                                                                                                                                                      |
| `xClientSignedAt`                                                                                                                                                                       | `*string`                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                      | The timestamp, in ISO 8601 format, at which you generated the `X-Client-Signature`. Required together with `X-Client-Signature` if you've enabled request signing for your integration. | 2025-01-01T12:00:00Z                                                                                                                                                                    |
| `salesInvoiceRequest`                                                                                                                                                                   | [*components.SalesInvoiceRequest](../../models/components/salesinvoicerequest.md)                                                                                                       | :heavy_minus_sign:                                                                                                                                                                      | N/A                                                                                                                                                                                     |                                                                                                                                                                                         |
| `opts`                                                                                                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                      | The options for this request.                                                                                                                                                           |                                                                                                                                                                                         |

### Response

**[*operations.CreateSalesInvoiceResponse](../../models/operations/createsalesinvoiceresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 422, 429           | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## List

Retrieve a list of all sales invoices created through the API.

The results are paginated.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-sales-invoices" method="get" path="/v2/sales-invoices" -->
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

    res, err := s.SalesInvoices.List(ctx, operations.ListSalesInvoicesRequest{
        From: client.Pointer("invoice_4Y0eZitmBnQ6IDoMqZQKh"),
        Limit: client.Pointer[int64](50),
        IdempotencyKey: client.Pointer("123e4567-e89b-12d3-a456-426"),
        XClientSignature: client.Pointer("MEUCIQDx8sJvK3z..."),
        XClientSignedAt: client.Pointer("2025-01-01T12:00:00Z"),
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListSalesInvoicesRequest](../../models/operations/listsalesinvoicesrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListSalesInvoicesResponse](../../models/operations/listsalesinvoicesresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 429                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Get

Retrieve a single sales invoice by its ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-sales-invoice" method="get" path="/v2/sales-invoices/{salesInvoiceId}" -->
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

    res, err := s.SalesInvoices.Get(ctx, operations.GetSalesInvoiceRequest{
        SalesInvoiceID: "invoice_4Y0eZitmBnQ6IDoMqZQKh",
        IdempotencyKey: client.Pointer("123e4567-e89b-12d3-a456-426"),
        XClientSignature: client.Pointer("MEUCIQDx8sJvK3z..."),
        XClientSignedAt: client.Pointer("2025-01-01T12:00:00Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SalesInvoiceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetSalesInvoiceRequest](../../models/operations/getsalesinvoicerequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetSalesInvoiceResponse](../../models/operations/getsalesinvoiceresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 429                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Update

Certain details of an existing sales invoice can be updated. For `draft` it is all values listed below, but for
statuses `paid` and `issued` there are certain additional requirements (`paymentDetails` and `emailDetails`,
respectively).

### Example Usage

<!-- UsageSnippet language="go" operationID="update-sales-invoice" method="patch" path="/v2/sales-invoices/{salesInvoiceId}" -->
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

    res, err := s.SalesInvoices.Update(ctx, operations.UpdateSalesInvoiceRequest{
        SalesInvoiceID: "invoice_4Y0eZitmBnQ6IDoMqZQKh",
        IdempotencyKey: client.Pointer("123e4567-e89b-12d3-a456-426"),
        XClientSignature: client.Pointer("MEUCIQDx8sJvK3z..."),
        XClientSignedAt: client.Pointer("2025-01-01T12:00:00Z"),
        RequestBody: &operations.UpdateSalesInvoiceRequestBody{
            Testmode: client.Pointer(false),
            Status: components.SalesInvoiceStatusUpdateDraft.ToPointer(),
            Memo: client.Pointer("An updated memo!"),
            PaymentTerm: components.SalesInvoicePaymentTermThirtydays.ToPointer(),
            PaymentDetails: &components.SalesInvoicePaymentDetails{
                Source: components.SalesInvoicePaymentDetailsSourcePaymentLink,
                SourceReference: client.Pointer("pl_d9fQur83kFdhH8hIhaZfq"),
            },
            EmailDetails: &components.SalesInvoiceEmailDetails{
                Subject: "Your invoice is available",
                Body: "Please find your invoice enclosed.",
            },
            RecipientIdentifier: client.Pointer("customer-xyz-0123"),
            Recipient: &components.SalesInvoiceRecipient{
                Type: components.SalesInvoiceRecipientTypeConsumer,
                Title: client.Pointer("Mrs."),
                GivenName: client.Pointer("Jane"),
                FamilyName: client.Pointer("Doe"),
                OrganizationName: client.Pointer("Organization Corp."),
                OrganizationNumber: client.Pointer("12345678"),
                VatNumber: client.Pointer("NL123456789B01"),
                Email: "example@email.com",
                Phone: client.Pointer("+0123456789"),
                StreetAndNumber: "Keizersgracht 126",
                StreetAdditional: client.Pointer("4th floor"),
                PostalCode: "5678AB",
                City: "Amsterdam",
                Region: client.Pointer("Noord-Holland"),
                Country: "NL",
                Locale: components.SalesInvoiceRecipientLocaleNlNl,
            },
            Lines: []components.SalesInvoiceLineItem{
                components.SalesInvoiceLineItem{
                    Description: "LEGO 4440 Forest Police Station",
                    Quantity: 1,
                    VatRate: "21.00",
                    UnitPrice: components.Amount{
                        Currency: "EUR",
                        Value: "10.00",
                    },
                    Discount: &components.SalesInvoiceDiscount{
                        Type: components.SalesInvoiceDiscountTypeAmount,
                        Value: "10.00",
                    },
                },
            },
            Discount: &components.SalesInvoiceDiscount{
                Type: components.SalesInvoiceDiscountTypeAmount,
                Value: "10.00",
            },
            IsEInvoice: client.Pointer(false),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SalesInvoiceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateSalesInvoiceRequest](../../models/operations/updatesalesinvoicerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.UpdateSalesInvoiceResponse](../../models/operations/updatesalesinvoiceresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 422, 429           | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Delete

Sales invoices which are in status `draft` can be deleted. For all other statuses, please use the
[Update sales invoice](update-sales-invoice) endpoint instead.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-sales-invoice" method="delete" path="/v2/sales-invoices/{salesInvoiceId}" -->
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

    res, err := s.SalesInvoices.Delete(ctx, operations.DeleteSalesInvoiceRequest{
        SalesInvoiceID: "invoice_4Y0eZitmBnQ6IDoMqZQKh",
        IdempotencyKey: client.Pointer("123e4567-e89b-12d3-a456-426"),
        XClientSignature: client.Pointer("MEUCIQDx8sJvK3z..."),
        XClientSignedAt: client.Pointer("2025-01-01T12:00:00Z"),
        DeleteValuesSalesInvoice: &components.DeleteValuesSalesInvoice{
            Testmode: client.Pointer(false),
        },
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.DeleteSalesInvoiceRequest](../../models/operations/deletesalesinvoicerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.DeleteSalesInvoiceResponse](../../models/operations/deletesalesinvoiceresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 422, 429           | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |