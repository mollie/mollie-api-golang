# PaymentLinks
(*PaymentLinks*)

## Overview

### Available Operations

* [Create](#create) - Create payment link
* [List](#list) - List payment links
* [Get](#get) - Get payment link
* [Update](#update) - Update payment link
* [Delete](#delete) - Delete payment link
* [ListPayments](#listpayments) - Get payment link payments

## Create

With the Payment links API you can generate payment links that by default, unlike regular payments, do not expire.
The payment link can be shared with your customers and will redirect them to them the payment page where they can
complete the payment. A [payment](get-payment) will only be created once the customer initiates the payment.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-payment-link" method="post" path="/payment-links" -->
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

    res, err := s.PaymentLinks.Create(ctx, client.Pointer("123e4567-e89b-12d3-a456-426"), &operations.CreatePaymentLinkRequestBody{
        ID: client.Pointer("pl_d9fQur83kFdhH8hIhaZfq"),
        Description: "Chess Board",
        Amount: &components.AmountNullable{
            Currency: "EUR",
            Value: "10.00",
        },
        MinimumAmount: &components.AmountNullable{
            Currency: "EUR",
            Value: "10.00",
        },
        RedirectURL: client.Pointer("https://webshop.example.org/payment-links/redirect/"),
        WebhookURL: client.Pointer("https://webshop.example.org/payment-links/webhook/"),
        Lines: []components.PaymentLineItem{
            components.PaymentLineItem{
                Type: components.PaymentLineTypePhysical.ToPointer(),
                Description: "LEGO 4440 Forest Police Station",
                Quantity: 1,
                QuantityUnit: client.Pointer("pcs"),
                UnitPrice: components.Amount{
                    Currency: "EUR",
                    Value: "10.00",
                },
                DiscountAmount: &components.Amount{
                    Currency: "EUR",
                    Value: "10.00",
                },
                TotalAmount: components.Amount{
                    Currency: "EUR",
                    Value: "10.00",
                },
                VatRate: client.Pointer("21.00"),
                VatAmount: &components.Amount{
                    Currency: "EUR",
                    Value: "10.00",
                },
                Sku: client.Pointer("9780241661628"),
                Categories: []components.LineCategories{
                    components.LineCategoriesMeal,
                    components.LineCategoriesEco,
                },
                ImageURL: client.Pointer("https://..."),
                ProductURL: client.Pointer("https://..."),
            },
        },
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
        ProfileID: client.Pointer("pfl_QkEhN94Ba"),
        Reusable: client.Pointer(false),
        ExpiresAt: client.Pointer("2025-12-24T11:00:16+00:00"),
        AllowedMethods: nil,
        ApplicationFee: &operations.ApplicationFee{
            Amount: components.Amount{
                Currency: "EUR",
                Value: "10.00",
            },
            Description: "Platform fee",
        },
        SequenceType: components.PaymentLinkSequenceTypeOneoff.ToPointer(),
        CustomerID: client.Pointer("cst_XimFHuaEzd"),
        Testmode: client.Pointer(false),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentLinkResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                           | Type                                                                                                | Required                                                                                            | Description                                                                                         | Example                                                                                             |
| --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                               | :heavy_check_mark:                                                                                  | The context to use for the request.                                                                 |                                                                                                     |
| `idempotencyKey`                                                                                    | **string*                                                                                           | :heavy_minus_sign:                                                                                  | A unique key to ensure idempotent requests. This key should be a UUID v4 string.                    | 123e4567-e89b-12d3-a456-426                                                                         |
| `requestBody`                                                                                       | [*operations.CreatePaymentLinkRequestBody](../../models/operations/createpaymentlinkrequestbody.md) | :heavy_minus_sign:                                                                                  | N/A                                                                                                 |                                                                                                     |
| `opts`                                                                                              | [][operations.Option](../../models/operations/option.md)                                            | :heavy_minus_sign:                                                                                  | The options for this request.                                                                       |                                                                                                     |

### Response

**[*operations.CreatePaymentLinkResponse](../../models/operations/createpaymentlinkresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 422                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## List

Retrieve a list of all payment links.

The results are paginated.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-payment-links" method="get" path="/payment-links" -->
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

    res, err := s.PaymentLinks.List(ctx, client.Pointer("pl_d9fQur83kFdhH8hIhaZfq"), client.Pointer[int64](50), client.Pointer(false), client.Pointer("123e4567-e89b-12d3-a456-426"))
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
| `from`                                                                                                                                                                                                                                                                                                                                                                                 | **string*                                                                                                                                                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | Provide an ID to start the result set from the item with the given ID and onwards. This allows you to paginate the<br/>result set.                                                                                                                                                                                                                                                     | pl_d9fQur83kFdhH8hIhaZfq                                                                                                                                                                                                                                                                                                                                                               |
| `limit`                                                                                                                                                                                                                                                                                                                                                                                | **int64*                                                                                                                                                                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | The maximum number of items to return. Defaults to 50 items.                                                                                                                                                                                                                                                                                                                           | 50                                                                                                                                                                                                                                                                                                                                                                                     |
| `testmode`                                                                                                                                                                                                                                                                                                                                                                             | **bool*                                                                                                                                                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | Most API credentials are specifically created for either live mode or test mode. In those cases the `testmode` query<br/>parameter can be omitted. For organization-level credentials such as OAuth access tokens, you can enable test mode by<br/>setting the `testmode` query parameter to `true`.<br/><br/>Test entities cannot be retrieved when the endpoint is set to live mode, and vice versa. | false                                                                                                                                                                                                                                                                                                                                                                                  |
| `idempotencyKey`                                                                                                                                                                                                                                                                                                                                                                       | **string*                                                                                                                                                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | A unique key to ensure idempotent requests. This key should be a UUID v4 string.                                                                                                                                                                                                                                                                                                       | 123e4567-e89b-12d3-a456-426                                                                                                                                                                                                                                                                                                                                                            |
| `opts`                                                                                                                                                                                                                                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | The options for this request.                                                                                                                                                                                                                                                                                                                                                          |                                                                                                                                                                                                                                                                                                                                                                                        |

### Response

**[*operations.ListPaymentLinksResponse](../../models/operations/listpaymentlinksresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Get

Retrieve a single payment link by its ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-payment-link" method="get" path="/payment-links/{paymentLinkId}" -->
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

    res, err := s.PaymentLinks.Get(ctx, "pl_d9fQur83kFdhH8hIhaZfq", client.Pointer(false), client.Pointer("123e4567-e89b-12d3-a456-426"))
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentLinkResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                              | Type                                                                                                                                                                                                                                                                                                                                                                                   | Required                                                                                                                                                                                                                                                                                                                                                                               | Description                                                                                                                                                                                                                                                                                                                                                                            | Example                                                                                                                                                                                                                                                                                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                     | The context to use for the request.                                                                                                                                                                                                                                                                                                                                                    |                                                                                                                                                                                                                                                                                                                                                                                        |
| `paymentLinkID`                                                                                                                                                                                                                                                                                                                                                                        | *string*                                                                                                                                                                                                                                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                     | Provide the ID of the related payment link.                                                                                                                                                                                                                                                                                                                                            | pl_d9fQur83kFdhH8hIhaZfq                                                                                                                                                                                                                                                                                                                                                               |
| `testmode`                                                                                                                                                                                                                                                                                                                                                                             | **bool*                                                                                                                                                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | Most API credentials are specifically created for either live mode or test mode. In those cases the `testmode` query<br/>parameter can be omitted. For organization-level credentials such as OAuth access tokens, you can enable test mode by<br/>setting the `testmode` query parameter to `true`.<br/><br/>Test entities cannot be retrieved when the endpoint is set to live mode, and vice versa. | false                                                                                                                                                                                                                                                                                                                                                                                  |
| `idempotencyKey`                                                                                                                                                                                                                                                                                                                                                                       | **string*                                                                                                                                                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | A unique key to ensure idempotent requests. This key should be a UUID v4 string.                                                                                                                                                                                                                                                                                                       | 123e4567-e89b-12d3-a456-426                                                                                                                                                                                                                                                                                                                                                            |
| `opts`                                                                                                                                                                                                                                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | The options for this request.                                                                                                                                                                                                                                                                                                                                                          |                                                                                                                                                                                                                                                                                                                                                                                        |

### Response

**[*operations.GetPaymentLinkResponse](../../models/operations/getpaymentlinkresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Update

Certain details of an existing payment link can be updated.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-payment-link" method="patch" path="/payment-links/{paymentLinkId}" -->
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

    res, err := s.PaymentLinks.Update(ctx, "pl_d9fQur83kFdhH8hIhaZfq", client.Pointer("123e4567-e89b-12d3-a456-426"), &operations.UpdatePaymentLinkRequestBody{
        Description: client.Pointer("Chess Board"),
        MinimumAmount: &components.Amount{
            Currency: "EUR",
            Value: "10.00",
        },
        Archived: client.Pointer(false),
        AllowedMethods: []string{
            "ideal",
        },
        Lines: nil,
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
        Testmode: client.Pointer(false),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentLinkResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                           | Type                                                                                                | Required                                                                                            | Description                                                                                         | Example                                                                                             |
| --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                               | :heavy_check_mark:                                                                                  | The context to use for the request.                                                                 |                                                                                                     |
| `paymentLinkID`                                                                                     | *string*                                                                                            | :heavy_check_mark:                                                                                  | Provide the ID of the related payment link.                                                         | pl_d9fQur83kFdhH8hIhaZfq                                                                            |
| `idempotencyKey`                                                                                    | **string*                                                                                           | :heavy_minus_sign:                                                                                  | A unique key to ensure idempotent requests. This key should be a UUID v4 string.                    | 123e4567-e89b-12d3-a456-426                                                                         |
| `requestBody`                                                                                       | [*operations.UpdatePaymentLinkRequestBody](../../models/operations/updatepaymentlinkrequestbody.md) | :heavy_minus_sign:                                                                                  | N/A                                                                                                 |                                                                                                     |
| `opts`                                                                                              | [][operations.Option](../../models/operations/option.md)                                            | :heavy_minus_sign:                                                                                  | The options for this request.                                                                       |                                                                                                     |

### Response

**[*operations.UpdatePaymentLinkResponse](../../models/operations/updatepaymentlinkresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 422                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Delete

Payment links which have not been opened and no payments have been made yet can be deleted entirely.
This can be useful for removing payment links that have been incorrectly configured or that are no longer relevant.

Once deleted, the payment link will no longer show up in the API or Mollie dashboard.

To simply disable a payment link without fully deleting it, you can use the `archived` parameter on the
[Update payment link](update-payment-link) endpoint instead.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-payment-link" method="delete" path="/payment-links/{paymentLinkId}" -->
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

    res, err := s.PaymentLinks.Delete(ctx, "pl_d9fQur83kFdhH8hIhaZfq", client.Pointer("123e4567-e89b-12d3-a456-426"), &operations.DeletePaymentLinkRequestBody{
        Testmode: client.Pointer(false),
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

| Parameter                                                                                           | Type                                                                                                | Required                                                                                            | Description                                                                                         | Example                                                                                             |
| --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                               | :heavy_check_mark:                                                                                  | The context to use for the request.                                                                 |                                                                                                     |
| `paymentLinkID`                                                                                     | *string*                                                                                            | :heavy_check_mark:                                                                                  | Provide the ID of the related payment link.                                                         | pl_d9fQur83kFdhH8hIhaZfq                                                                            |
| `idempotencyKey`                                                                                    | **string*                                                                                           | :heavy_minus_sign:                                                                                  | A unique key to ensure idempotent requests. This key should be a UUID v4 string.                    | 123e4567-e89b-12d3-a456-426                                                                         |
| `requestBody`                                                                                       | [*operations.DeletePaymentLinkRequestBody](../../models/operations/deletepaymentlinkrequestbody.md) | :heavy_minus_sign:                                                                                  | N/A                                                                                                 |                                                                                                     |
| `opts`                                                                                              | [][operations.Option](../../models/operations/option.md)                                            | :heavy_minus_sign:                                                                                  | The options for this request.                                                                       |                                                                                                     |

### Response

**[*operations.DeletePaymentLinkResponse](../../models/operations/deletepaymentlinkresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 422                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## ListPayments

Retrieve the list of payments for a specific payment link.

The results are paginated.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-payment-link-payments" method="get" path="/payment-links/{paymentLinkId}/payments" -->
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

    res, err := s.PaymentLinks.ListPayments(ctx, operations.GetPaymentLinkPaymentsRequest{
        PaymentLinkID: "pl_d9fQur83kFdhH8hIhaZfq",
        From: client.Pointer("tr_5B8cwPMGnU"),
        Limit: client.Pointer[int64](50),
        Sort: components.SortingDesc.ToPointer(),
        Testmode: client.Pointer(false),
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetPaymentLinkPaymentsRequest](../../models/operations/getpaymentlinkpaymentsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.GetPaymentLinkPaymentsResponse](../../models/operations/getpaymentlinkpaymentsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |