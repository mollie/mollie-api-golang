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
            APIKey: client.String(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.PaymentLinks.Create(ctx, &operations.CreatePaymentLinkRequest{
        Description: "Chess Board",
        Amount: &operations.CreatePaymentLinkAmountRequest{
            Currency: "EUR",
            Value: "10.00",
        },
        MinimumAmount: &operations.CreatePaymentLinkMinimumAmountRequest{
            Currency: "EUR",
            Value: "10.00",
        },
        RedirectURL: client.String("https://webshop.example.org/payment-links/redirect/"),
        WebhookURL: client.String("https://webshop.example.org/payment-links/webhook/"),
        Lines: []operations.CreatePaymentLinkLineRequest{
            operations.CreatePaymentLinkLineRequest{
                Type: operations.CreatePaymentLinkTypeRequestPhysical.ToPointer(),
                Description: "LEGO 4440 Forest Police Station",
                Quantity: 1,
                QuantityUnit: client.String("pcs"),
                UnitPrice: operations.CreatePaymentLinkUnitPriceRequest{
                    Currency: "EUR",
                    Value: "10.00",
                },
                DiscountAmount: &operations.CreatePaymentLinkDiscountAmountRequest{
                    Currency: "EUR",
                    Value: "10.00",
                },
                TotalAmount: operations.CreatePaymentLinkTotalAmountRequest{
                    Currency: "EUR",
                    Value: "10.00",
                },
                VatRate: client.String("21.00"),
                VatAmount: &operations.CreatePaymentLinkVatAmountRequest{
                    Currency: "EUR",
                    Value: "10.00",
                },
                Sku: client.String("9780241661628"),
                Categories: []operations.CreatePaymentLinkCategoryRequest{
                    operations.CreatePaymentLinkCategoryRequestMeal,
                    operations.CreatePaymentLinkCategoryRequestEco,
                },
                ImageURL: client.String("https://..."),
                ProductURL: client.String("https://..."),
            },
        },
        BillingAddress: &operations.CreatePaymentLinkBillingAddressRequest{
            Title: client.String("Mr."),
            GivenName: client.String("Piet"),
            FamilyName: client.String("Mondriaan"),
            OrganizationName: client.String("Mollie B.V."),
            StreetAndNumber: client.String("Keizersgracht 126"),
            StreetAdditional: client.String("Apt. 1"),
            PostalCode: client.String("1234AB"),
            Email: client.String("piet@example.org"),
            Phone: client.String("31208202070"),
            City: client.String("Amsterdam"),
            Region: client.String("Noord-Holland"),
            Country: client.String("NL"),
        },
        ShippingAddress: &operations.CreatePaymentLinkShippingAddressRequest{
            Title: client.String("Mr."),
            GivenName: client.String("Piet"),
            FamilyName: client.String("Mondriaan"),
            OrganizationName: client.String("Mollie B.V."),
            StreetAndNumber: client.String("Keizersgracht 126"),
            StreetAdditional: client.String("Apt. 1"),
            PostalCode: client.String("1234AB"),
            Email: client.String("piet@example.org"),
            Phone: client.String("31208202070"),
            City: client.String("Amsterdam"),
            Region: client.String("Noord-Holland"),
            Country: client.String("NL"),
        },
        ProfileID: client.String("pfl_QkEhN94Ba"),
        Reusable: client.Bool(false),
        ExpiresAt: client.String("2025-12-24T11:00:16+00:00"),
        AllowedMethods: nil,
        ApplicationFee: &operations.CreatePaymentLinkApplicationFeeRequest{
            Amount: operations.CreatePaymentLinkApplicationFeeAmountRequest{
                Currency: "EUR",
                Value: "10.00",
            },
            Description: "Platform fee",
        },
        SequenceType: operations.CreatePaymentLinkSequenceTypeRequestOneoff.ToPointer(),
        CustomerID: client.String("cst_XimFHuaEzd"),
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreatePaymentLinkRequest](../../models/operations/createpaymentlinkrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.CreatePaymentLinkResponse](../../models/operations/createpaymentlinkresponse.md), error**

### Errors

| Error Type                                                 | Status Code                                                | Content Type                                               |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| apierrors.CreatePaymentLinkNotFoundHalJSONError            | 404                                                        | application/hal+json                                       |
| apierrors.CreatePaymentLinkUnprocessableEntityHalJSONError | 422                                                        | application/hal+json                                       |
| apierrors.APIError                                         | 4XX, 5XX                                                   | \*/\*                                                      |

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
            APIKey: client.String(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.PaymentLinks.List(ctx, client.String("pl_d9fQur83kFdhH8hIhaZfq"), client.Int64(50), client.Bool(false))
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
| `opts`                                                                                                                                                                                                                                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | The options for this request.                                                                                                                                                                                                                                                                                                                                                          |                                                                                                                                                                                                                                                                                                                                                                                        |

### Response

**[*operations.ListPaymentLinksResponse](../../models/operations/listpaymentlinksresponse.md), error**

### Errors

| Error Type                             | Status Code                            | Content Type                           |
| -------------------------------------- | -------------------------------------- | -------------------------------------- |
| apierrors.ListPaymentLinksHalJSONError | 400                                    | application/hal+json                   |
| apierrors.APIError                     | 4XX, 5XX                               | \*/\*                                  |

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
            APIKey: client.String(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.PaymentLinks.Get(ctx, "pl_d9fQur83kFdhH8hIhaZfq", client.Bool(false))
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
| `paymentLinkID`                                                                                                                                                                                                                                                                                                                                                                        | *string*                                                                                                                                                                                                                                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                     | Provide the ID of the related payment link.                                                                                                                                                                                                                                                                                                                                            | pl_d9fQur83kFdhH8hIhaZfq                                                                                                                                                                                                                                                                                                                                                               |
| `testmode`                                                                                                                                                                                                                                                                                                                                                                             | **bool*                                                                                                                                                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | Most API credentials are specifically created for either live mode or test mode. In those cases the `testmode` query<br/>parameter can be omitted. For organization-level credentials such as OAuth access tokens, you can enable test mode by<br/>setting the `testmode` query parameter to `true`.<br/><br/>Test entities cannot be retrieved when the endpoint is set to live mode, and vice versa. | false                                                                                                                                                                                                                                                                                                                                                                                  |
| `opts`                                                                                                                                                                                                                                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | The options for this request.                                                                                                                                                                                                                                                                                                                                                          |                                                                                                                                                                                                                                                                                                                                                                                        |

### Response

**[*operations.GetPaymentLinkResponse](../../models/operations/getpaymentlinkresponse.md), error**

### Errors

| Error Type                           | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| apierrors.GetPaymentLinkHalJSONError | 404                                  | application/hal+json                 |
| apierrors.APIError                   | 4XX, 5XX                             | \*/\*                                |

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
            APIKey: client.String(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.PaymentLinks.Update(ctx, "pl_d9fQur83kFdhH8hIhaZfq", &operations.UpdatePaymentLinkRequestBody{
        Description: client.String("Chess Board"),
        MinimumAmount: &operations.UpdatePaymentLinkMinimumAmountRequest{
            Currency: "EUR",
            Value: "10.00",
        },
        Archived: client.Bool(false),
        AllowedMethods: []string{
            "ideal",
        },
        Lines: nil,
        BillingAddress: &operations.UpdatePaymentLinkBillingAddressRequest{
            Title: client.String("Mr."),
            GivenName: client.String("Piet"),
            FamilyName: client.String("Mondriaan"),
            OrganizationName: client.String("Mollie B.V."),
            StreetAndNumber: client.String("Keizersgracht 126"),
            StreetAdditional: client.String("Apt. 1"),
            PostalCode: client.String("1234AB"),
            Email: client.String("piet@example.org"),
            Phone: client.String("31208202070"),
            City: client.String("Amsterdam"),
            Region: client.String("Noord-Holland"),
            Country: client.String("NL"),
        },
        ShippingAddress: &operations.UpdatePaymentLinkShippingAddressRequest{
            Title: client.String("Mr."),
            GivenName: client.String("Piet"),
            FamilyName: client.String("Mondriaan"),
            OrganizationName: client.String("Mollie B.V."),
            StreetAndNumber: client.String("Keizersgracht 126"),
            StreetAdditional: client.String("Apt. 1"),
            PostalCode: client.String("1234AB"),
            Email: client.String("piet@example.org"),
            Phone: client.String("31208202070"),
            City: client.String("Amsterdam"),
            Region: client.String("Noord-Holland"),
            Country: client.String("NL"),
        },
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

| Parameter                                                                                           | Type                                                                                                | Required                                                                                            | Description                                                                                         | Example                                                                                             |
| --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                               | :heavy_check_mark:                                                                                  | The context to use for the request.                                                                 |                                                                                                     |
| `paymentLinkID`                                                                                     | *string*                                                                                            | :heavy_check_mark:                                                                                  | Provide the ID of the related payment link.                                                         | pl_d9fQur83kFdhH8hIhaZfq                                                                            |
| `requestBody`                                                                                       | [*operations.UpdatePaymentLinkRequestBody](../../models/operations/updatepaymentlinkrequestbody.md) | :heavy_minus_sign:                                                                                  | N/A                                                                                                 |                                                                                                     |
| `opts`                                                                                              | [][operations.Option](../../models/operations/option.md)                                            | :heavy_minus_sign:                                                                                  | The options for this request.                                                                       |                                                                                                     |

### Response

**[*operations.UpdatePaymentLinkResponse](../../models/operations/updatepaymentlinkresponse.md), error**

### Errors

| Error Type                                                 | Status Code                                                | Content Type                                               |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| apierrors.UpdatePaymentLinkNotFoundHalJSONError            | 404                                                        | application/hal+json                                       |
| apierrors.UpdatePaymentLinkUnprocessableEntityHalJSONError | 422                                                        | application/hal+json                                       |
| apierrors.APIError                                         | 4XX, 5XX                                                   | \*/\*                                                      |

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
            APIKey: client.String(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.PaymentLinks.Delete(ctx, "pl_d9fQur83kFdhH8hIhaZfq", &operations.DeletePaymentLinkRequestBody{
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

| Parameter                                                                                           | Type                                                                                                | Required                                                                                            | Description                                                                                         | Example                                                                                             |
| --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                               | :heavy_check_mark:                                                                                  | The context to use for the request.                                                                 |                                                                                                     |
| `paymentLinkID`                                                                                     | *string*                                                                                            | :heavy_check_mark:                                                                                  | Provide the ID of the related payment link.                                                         | pl_d9fQur83kFdhH8hIhaZfq                                                                            |
| `requestBody`                                                                                       | [*operations.DeletePaymentLinkRequestBody](../../models/operations/deletepaymentlinkrequestbody.md) | :heavy_minus_sign:                                                                                  | N/A                                                                                                 |                                                                                                     |
| `opts`                                                                                              | [][operations.Option](../../models/operations/option.md)                                            | :heavy_minus_sign:                                                                                  | The options for this request.                                                                       |                                                                                                     |

### Response

**[*operations.DeletePaymentLinkResponse](../../models/operations/deletepaymentlinkresponse.md), error**

### Errors

| Error Type                                                 | Status Code                                                | Content Type                                               |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| apierrors.DeletePaymentLinkNotFoundHalJSONError            | 404                                                        | application/hal+json                                       |
| apierrors.DeletePaymentLinkUnprocessableEntityHalJSONError | 422                                                        | application/hal+json                                       |
| apierrors.APIError                                         | 4XX, 5XX                                                   | \*/\*                                                      |

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
            APIKey: client.String(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.PaymentLinks.ListPayments(ctx, operations.GetPaymentLinkPaymentsRequest{
        PaymentLinkID: "pl_d9fQur83kFdhH8hIhaZfq",
        From: client.String("tr_5B8cwPMGnU"),
        Limit: client.Int64(50),
        Sort: operations.GetPaymentLinkPaymentsSortDesc.ToPointer(),
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetPaymentLinkPaymentsRequest](../../models/operations/getpaymentlinkpaymentsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.GetPaymentLinkPaymentsResponse](../../models/operations/getpaymentlinkpaymentsresponse.md), error**

### Errors

| Error Type                                   | Status Code                                  | Content Type                                 |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| apierrors.GetPaymentLinkPaymentsHalJSONError | 400                                          | application/hal+json                         |
| apierrors.APIError                           | 4XX, 5XX                                     | \*/\*                                        |