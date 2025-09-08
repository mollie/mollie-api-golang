# SalesInvoices
(*SalesInvoices*)

## Overview

### Available Operations

* [Create](#create) - Create sales invoice
* [List](#list) - List sales invoices
* [Get](#get) - Get sales invoice
* [Update](#update) - Update sales invoice
* [Delete](#delete) - Delete sales invoice

## Create

> 🚧 Beta feature
>
> This feature is currently in beta testing, and the final specification may still change.

With the Sales Invoice API you can generate sales invoices to send to your customers.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-sales-invoice" method="post" path="/sales-invoices" -->
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

    res, err := s.SalesInvoices.Create(ctx, &components.EntitySalesInvoice{
        ID: client.String("invoice_4Y0eZitmBnQ6IDoMqZQKh"),
        Testmode: client.Bool(false),
        ProfileID: client.String("pfl_QkEhN94Ba"),
        Status: components.SalesInvoiceStatusDraft.ToPointer(),
        VatScheme: components.SalesInvoiceVatSchemeStandard.ToPointer(),
        VatMode: components.SalesInvoiceVatModeExclusive.ToPointer(),
        Memo: client.String("This is a memo!"),
        PaymentTerm: components.SalesInvoicePaymentTermThirtydays.ToPointer(),
        PaymentDetails: &components.SalesInvoicePaymentDetails{
            Source: components.SalesInvoicePaymentDetailsSourcePaymentLink,
            SourceReference: client.String("pl_d9fQur83kFdhH8hIhaZfq"),
        },
        EmailDetails: &components.SalesInvoiceEmailDetails{
            Subject: "Your invoice is available",
            Body: "Please find your invoice enclosed.",
        },
        CustomerID: client.String("cst_8wmqcHMN4U"),
        MandateID: client.String("mdt_pWUnw6pkBN"),
        RecipientIdentifier: client.String("customer-xyz-0123"),
        Recipient: &components.SalesInvoiceRecipient{
            Type: components.SalesInvoiceRecipientTypeConsumer,
            Title: client.String("Mrs."),
            GivenName: client.String("Jane"),
            FamilyName: client.String("Doe"),
            OrganizationName: client.String("Organization Corp."),
            OrganizationNumber: client.String("12345678"),
            VatNumber: client.String("NL123456789B01"),
            Email: "example@email.com",
            Phone: client.String("+0123456789"),
            StreetAndNumber: "Keizersgracht 126",
            StreetAdditional: client.String("4th floor"),
            PostalCode: "5678AB",
            City: "Amsterdam",
            Region: client.String("Noord-Holland"),
            Country: "NL",
            Locale: components.SalesInvoiceRecipientLocaleNlNl,
        },
        Lines: []components.SalesInvoiceLineItem{},
        Discount: &components.SalesInvoiceDiscount{
            Type: components.SalesInvoiceDiscountTypeAmount,
            Value: "10.00",
        },
        AmountDue: &components.Amount{
            Currency: "EUR",
            Value: "10.00",
        },
        SubtotalAmount: &components.Amount{
            Currency: "EUR",
            Value: "10.00",
        },
        TotalAmount: &components.Amount{
            Currency: "EUR",
            Value: "10.00",
        },
        TotalVatAmount: &components.Amount{
            Currency: "EUR",
            Value: "10.00",
        },
        DiscountedSubtotalAmount: &components.Amount{
            Currency: "EUR",
            Value: "10.00",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EntitySalesInvoiceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.EntitySalesInvoice](../../models/components/entitysalesinvoice.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.CreateSalesInvoiceResponse](../../models/operations/createsalesinvoiceresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 422                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## List

> 🚧 Beta feature
>
> This feature is currently in beta testing, and the final specification may still change.

Retrieve a list of all sales invoices created through the API.

The results are paginated.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-sales-invoices" method="get" path="/sales-invoices" -->
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

    res, err := s.SalesInvoices.List(ctx, client.String("invoice_4Y0eZitmBnQ6IDoMqZQKh"), client.Int64(50), client.Bool(false))
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
| `from`                                                                                                                                                                                                                                                                                                                                                                                 | **string*                                                                                                                                                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | Provide an ID to start the result set from the item with the given ID and onwards. This allows you to paginate the<br/>result set.                                                                                                                                                                                                                                                     |                                                                                                                                                                                                                                                                                                                                                                                        |
| `limit`                                                                                                                                                                                                                                                                                                                                                                                | **int64*                                                                                                                                                                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | The maximum number of items to return. Defaults to 50 items.                                                                                                                                                                                                                                                                                                                           | 50                                                                                                                                                                                                                                                                                                                                                                                     |
| `testmode`                                                                                                                                                                                                                                                                                                                                                                             | **bool*                                                                                                                                                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | Most API credentials are specifically created for either live mode or test mode. In those cases the `testmode` query<br/>parameter can be omitted. For organization-level credentials such as OAuth access tokens, you can enable test mode by<br/>setting the `testmode` query parameter to `true`.<br/><br/>Test entities cannot be retrieved when the endpoint is set to live mode, and vice versa. | false                                                                                                                                                                                                                                                                                                                                                                                  |
| `opts`                                                                                                                                                                                                                                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | The options for this request.                                                                                                                                                                                                                                                                                                                                                          |                                                                                                                                                                                                                                                                                                                                                                                        |

### Response

**[*operations.ListSalesInvoicesResponse](../../models/operations/listsalesinvoicesresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Get

> 🚧 Beta feature
>
> This feature is currently in beta testing, and the final specification may still change.

Retrieve a single sales invoice by its ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-sales-invoice" method="get" path="/sales-invoices/{id}" -->
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

    res, err := s.SalesInvoices.Get(ctx, "invoice_4Y0eZitmBnQ6IDoMqZQKh", client.Bool(false))
    if err != nil {
        log.Fatal(err)
    }
    if res.EntitySalesInvoiceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                              | Type                                                                                                                                                                                                                                                                                                                                                                                   | Required                                                                                                                                                                                                                                                                                                                                                                               | Description                                                                                                                                                                                                                                                                                                                                                                            | Example                                                                                                                                                                                                                                                                                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                     | The context to use for the request.                                                                                                                                                                                                                                                                                                                                                    |                                                                                                                                                                                                                                                                                                                                                                                        |
| `id`                                                                                                                                                                                                                                                                                                                                                                                   | *string*                                                                                                                                                                                                                                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                     | Provide the ID of the item you want to perform this operation on.                                                                                                                                                                                                                                                                                                                      |                                                                                                                                                                                                                                                                                                                                                                                        |
| `testmode`                                                                                                                                                                                                                                                                                                                                                                             | **bool*                                                                                                                                                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | Most API credentials are specifically created for either live mode or test mode. In those cases the `testmode` query<br/>parameter can be omitted. For organization-level credentials such as OAuth access tokens, you can enable test mode by<br/>setting the `testmode` query parameter to `true`.<br/><br/>Test entities cannot be retrieved when the endpoint is set to live mode, and vice versa. | false                                                                                                                                                                                                                                                                                                                                                                                  |
| `opts`                                                                                                                                                                                                                                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | The options for this request.                                                                                                                                                                                                                                                                                                                                                          |                                                                                                                                                                                                                                                                                                                                                                                        |

### Response

**[*operations.GetSalesInvoiceResponse](../../models/operations/getsalesinvoiceresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Update

> 🚧 Beta feature
>
> This feature is currently in beta testing, and the final specification may still change.

Certain details of an existing sales invoice can be updated. For `draft` it is all values listed below, but for
statuses `paid` and `issued` there are certain additional requirements (`paymentDetails` and `emailDetails`,
respectively).

### Example Usage

<!-- UsageSnippet language="go" operationID="update-sales-invoice" method="patch" path="/sales-invoices/{id}" -->
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

    res, err := s.SalesInvoices.Update(ctx, "invoice_4Y0eZitmBnQ6IDoMqZQKh", &components.UpdateValuesSalesInvoice{
        Testmode: client.Bool(false),
        Status: components.SalesInvoiceStatusDraft.ToPointer(),
        Memo: client.String("An updated memo!"),
        PaymentTerm: components.SalesInvoicePaymentTermThirtydays.ToPointer(),
        PaymentDetails: &components.SalesInvoicePaymentDetails{
            Source: components.SalesInvoicePaymentDetailsSourcePaymentLink,
            SourceReference: client.String("pl_d9fQur83kFdhH8hIhaZfq"),
        },
        EmailDetails: &components.SalesInvoiceEmailDetails{
            Subject: "Your invoice is available",
            Body: "Please find your invoice enclosed.",
        },
        RecipientIdentifier: client.String("customer-xyz-0123"),
        Recipient: &components.SalesInvoiceRecipient{
            Type: components.SalesInvoiceRecipientTypeConsumer,
            Title: client.String("Mrs."),
            GivenName: client.String("Jane"),
            FamilyName: client.String("Doe"),
            OrganizationName: client.String("Organization Corp."),
            OrganizationNumber: client.String("12345678"),
            VatNumber: client.String("NL123456789B01"),
            Email: "example@email.com",
            Phone: client.String("+0123456789"),
            StreetAndNumber: "Keizersgracht 126",
            StreetAdditional: client.String("4th floor"),
            PostalCode: "5678AB",
            City: "Amsterdam",
            Region: client.String("Noord-Holland"),
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
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EntitySalesInvoiceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                   | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `ctx`                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                       | :heavy_check_mark:                                                                          | The context to use for the request.                                                         |
| `id`                                                                                        | *string*                                                                                    | :heavy_check_mark:                                                                          | Provide the ID of the item you want to perform this operation on.                           |
| `updateValuesSalesInvoice`                                                                  | [*components.UpdateValuesSalesInvoice](../../models/components/updatevaluessalesinvoice.md) | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `opts`                                                                                      | [][operations.Option](../../models/operations/option.md)                                    | :heavy_minus_sign:                                                                          | The options for this request.                                                               |

### Response

**[*operations.UpdateSalesInvoiceResponse](../../models/operations/updatesalesinvoiceresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 422                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Delete

> 🚧 Beta feature
>
> This feature is currently in beta testing, and the final specification may still change.

Sales invoices which are in status `draft` can be deleted. For all other statuses, please use the
[Update sales invoice](update-sales-invoice) endpoint instead.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-sales-invoice" method="delete" path="/sales-invoices/{id}" -->
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

    res, err := s.SalesInvoices.Delete(ctx, "invoice_4Y0eZitmBnQ6IDoMqZQKh", &components.DeleteValuesSalesInvoice{
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

| Parameter                                                                                   | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `ctx`                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                       | :heavy_check_mark:                                                                          | The context to use for the request.                                                         |
| `id`                                                                                        | *string*                                                                                    | :heavy_check_mark:                                                                          | Provide the ID of the item you want to perform this operation on.                           |
| `deleteValuesSalesInvoice`                                                                  | [*components.DeleteValuesSalesInvoice](../../models/components/deletevaluessalesinvoice.md) | :heavy_minus_sign:                                                                          | N/A                                                                                         |
| `opts`                                                                                      | [][operations.Option](../../models/operations/option.md)                                    | :heavy_minus_sign:                                                                          | The options for this request.                                                               |

### Response

**[*operations.DeleteSalesInvoiceResponse](../../models/operations/deletesalesinvoiceresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 422                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |