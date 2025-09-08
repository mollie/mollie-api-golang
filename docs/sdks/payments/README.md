# Payments
(*Payments*)

## Overview

### Available Operations

* [Create](#create) - Create payment
* [List](#list) - List payments
* [Get](#get) - Get payment
* [Update](#update) - Update payment
* [Cancel](#cancel) - Cancel payment
* [ReleaseAuthorization](#releaseauthorization) - Release payment authorization

## Create

Payment creation is elemental to the Mollie API: this is where most payment
implementations start off.

Once you have created a payment, you should redirect your customer to the
URL in the `_links.checkout` property from the response.

To wrap your head around the payment process, an explanation and flow charts
can be found in the 'Accepting payments' guide.

If you specify the `method` parameter when creating a payment, optional
additional parameters may be available for the payment method that are not listed below. Please refer to the
guide on [method-specific parameters](extra-payment-parameters).

### Example Usage

<!-- UsageSnippet language="go" operationID="create-payment" method="post" path="/payments" -->
```go
package main

import(
	"context"
	"os"
	"github.com/mollie/mollie-api-golang/models/components"
	client "github.com/mollie/mollie-api-golang"
	"github.com/mollie/mollie-api-golang/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := client.New(
        client.WithSecurity(components.Security{
            APIKey: client.String(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.Payments.Create(ctx, client.String("details.qrCode"), &components.PaymentRequest{
        ID: client.String("tr_5B8cwPMGnU"),
        Description: client.String("Chess Board"),
        Amount: &components.Amount{
            Currency: "EUR",
            Value: "10.00",
        },
        AmountRefunded: &components.Amount{
            Currency: "EUR",
            Value: "10.00",
        },
        AmountRemaining: &components.Amount{
            Currency: "EUR",
            Value: "10.00",
        },
        AmountCaptured: &components.Amount{
            Currency: "EUR",
            Value: "10.00",
        },
        AmountChargedBack: &components.Amount{
            Currency: "EUR",
            Value: "10.00",
        },
        SettlementAmount: &components.Amount{
            Currency: "EUR",
            Value: "10.00",
        },
        RedirectURL: client.String("https://example.org/redirect"),
        CancelURL: client.String("https://example.org/cancel"),
        WebhookURL: client.String("https://example.org/webhooks"),
        Lines: []components.PaymentRequestLine{
            components.PaymentRequestLine{
                Type: components.PaymentLineTypePhysical.ToPointer(),
                Description: "LEGO 4440 Forest Police Station",
                Quantity: 1,
                QuantityUnit: client.String("pcs"),
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
                VatRate: client.String("21.00"),
                VatAmount: &components.Amount{
                    Currency: "EUR",
                    Value: "10.00",
                },
                Sku: client.String("9780241661628"),
                Categories: []components.PaymentRequestCategory{
                    components.PaymentRequestCategoryMeal,
                    components.PaymentRequestCategoryEco,
                },
                ImageURL: client.String("https://..."),
                ProductURL: client.String("https://..."),
                Recurring: &components.RecurringLineItem{
                    Description: client.String("Gym subscription"),
                    Interval: "12 months",
                    Amount: &components.Amount{
                        Currency: "EUR",
                        Value: "10.00",
                    },
                    Times: client.Int64(1),
                    StartDate: client.String("2024-12-12"),
                },
            },
        },
        BillingAddress: &components.PaymentAddress{
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
        ShippingAddress: &components.PaymentAddress{
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
        Locale: components.LocaleEnUs.ToPointer(),
        Method: components.MethodIdeal.ToPointer(),
        Issuer: client.String("ideal_INGBNL2A"),
        RestrictPaymentMethodsToCountry: client.String("NL"),
        CaptureMode: components.CaptureModeManual.ToPointer(),
        CaptureDelay: client.String("8 hours"),
        ApplicationFee: &components.PaymentRequestApplicationFee{
            Amount: &components.Amount{
                Currency: "EUR",
                Value: "10.00",
            },
            Description: client.String("10"),
        },
        Routing: []components.EntityPaymentRoute{
            components.EntityPaymentRoute{
                ID: "rt_5B8cwPMGnU",
                Amount: components.Amount{
                    Currency: "EUR",
                    Value: "10.00",
                },
                Destination: components.EntityPaymentRouteDestination{
                    Type: components.RouteDestinationTypeOrganization,
                    OrganizationID: "org_1234567",
                },
                ReleaseDate: client.String("2024-12-12"),
                Links: components.EntityPaymentRouteLinks{
                    Self: components.URLObj{
                        Href: "https://...",
                        Type: "application/hal+json",
                    },
                    Payment: components.URLObj{
                        Href: "https://...",
                        Type: "application/hal+json",
                    },
                },
            },
        },
        SequenceType: components.SequenceTypeOneoff.ToPointer(),
        SubscriptionID: client.String("sub_5B8cwPMGnU"),
        MandateID: client.String("mdt_5B8cwPMGnU"),
        CustomerID: client.String("cst_5B8cwPMGnU"),
        ProfileID: client.String("pfl_5B8cwPMGnU"),
        SettlementID: client.String("stl_5B8cwPMGnU"),
        OrderID: client.String("ord_5B8cwPMGnU"),
        DueDate: client.String("2025-01-01"),
        Testmode: client.Bool(false),
        ApplePayPaymentToken: client.String("{\"paymentData\": {\"version\": \"EC_v1\", \"data\": \"vK3BbrCbI/....\"}}"),
        Company: &components.Company{
            RegistrationNumber: client.String("12345678"),
            VatNumber: client.String("NL123456789B01"),
            EntityType: nil,
        },
        CardToken: client.String("tkn_12345"),
        VoucherNumber: client.String("1234567890"),
        VoucherPin: client.String("1234"),
        ConsumerDateOfBirth: types.MustNewDateFromString("2000-01-01"),
        SessionID: nil,
        DigitalGoods: client.Bool(true),
        CustomerReference: client.String("1234567890"),
        TerminalID: client.String("term_1234567890"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `include`                                                                                            | **string*                                                                                            | :heavy_minus_sign:                                                                                   | This endpoint allows you to include additional information via the `include` query string parameter. |
| `paymentRequest`                                                                                     | [*components.PaymentRequest](../../models/components/paymentrequest.md)                              | :heavy_minus_sign:                                                                                   | N/A                                                                                                  |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.CreatePaymentResponse](../../models/operations/createpaymentresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 422                     | application/hal+json    |
| apierrors.ErrorResponse | 503                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## List

Retrieve all payments created with the current website profile.

The results are paginated.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-payments" method="get" path="/payments" -->
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

    res, err := s.Payments.List(ctx, operations.ListPaymentsRequest{
        From: client.String("tr_5B8cwPMGnU"),
        Limit: client.Int64(50),
        Sort: components.ListSortDesc.ToPointer(),
        ProfileID: client.String("pfl_5B8cwPMGnU"),
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
| `request`                                                                        | [operations.ListPaymentsRequest](../../models/operations/listpaymentsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.ListPaymentsResponse](../../models/operations/listpaymentsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Get

Retrieve a single payment object by its payment ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-payment" method="get" path="/payments/{paymentId}" -->
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

    res, err := s.Payments.Get(ctx, "tr_5B8cwPMGnU", client.String("details.qrCode"), client.String("captures"), client.Bool(false))
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                              | Type                                                                                                                                                                                                                                                                                                                                                                                   | Required                                                                                                                                                                                                                                                                                                                                                                               | Description                                                                                                                                                                                                                                                                                                                                                                            | Example                                                                                                                                                                                                                                                                                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                     | The context to use for the request.                                                                                                                                                                                                                                                                                                                                                    |                                                                                                                                                                                                                                                                                                                                                                                        |
| `paymentID`                                                                                                                                                                                                                                                                                                                                                                            | *string*                                                                                                                                                                                                                                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                     | Provide the ID of the related payment.                                                                                                                                                                                                                                                                                                                                                 | tr_5B8cwPMGnU                                                                                                                                                                                                                                                                                                                                                                          |
| `include`                                                                                                                                                                                                                                                                                                                                                                              | **string*                                                                                                                                                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | This endpoint allows you to include additional information via the `include` query string parameter.                                                                                                                                                                                                                                                                                   |                                                                                                                                                                                                                                                                                                                                                                                        |
| `embed`                                                                                                                                                                                                                                                                                                                                                                                | **string*                                                                                                                                                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | This endpoint allows embedding related API items by appending the following values via the `embed` query string<br/>parameter.                                                                                                                                                                                                                                                         |                                                                                                                                                                                                                                                                                                                                                                                        |
| `testmode`                                                                                                                                                                                                                                                                                                                                                                             | **bool*                                                                                                                                                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | Most API credentials are specifically created for either live mode or test mode. In those cases the `testmode` query<br/>parameter can be omitted. For organization-level credentials such as OAuth access tokens, you can enable test mode by<br/>setting the `testmode` query parameter to `true`.<br/><br/>Test entities cannot be retrieved when the endpoint is set to live mode, and vice versa. | false                                                                                                                                                                                                                                                                                                                                                                                  |
| `opts`                                                                                                                                                                                                                                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | The options for this request.                                                                                                                                                                                                                                                                                                                                                          |                                                                                                                                                                                                                                                                                                                                                                                        |

### Response

**[*operations.GetPaymentResponse](../../models/operations/getpaymentresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Update

Certain details of an existing payment can be updated.

Updating the payment details will not result in a webhook call.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-payment" method="patch" path="/payments/{paymentId}" -->
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

    res, err := s.Payments.Update(ctx, "tr_5B8cwPMGnU", &operations.UpdatePaymentRequestBody{
        Description: client.String("Chess Board"),
        RedirectURL: client.String("https://example.org/redirect"),
        CancelURL: client.String("https://example.org/cancel"),
        WebhookURL: client.String("https://example.org/webhooks"),
        Method: components.MethodIdeal.ToPointer(),
        Locale: components.LocaleEnUs.ToPointer(),
        DueDate: client.String("2025-01-01"),
        RestrictPaymentMethodsToCountry: client.String("NL"),
        Testmode: client.Bool(false),
        Issuer: client.String("ideal_INGBNL2A"),
        BillingAddress: &components.PaymentAddress{
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
        ShippingAddress: &components.PaymentAddress{
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
        BillingEmail: client.String("test@example.com"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                   | Type                                                                                        | Required                                                                                    | Description                                                                                 | Example                                                                                     |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `ctx`                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                       | :heavy_check_mark:                                                                          | The context to use for the request.                                                         |                                                                                             |
| `paymentID`                                                                                 | *string*                                                                                    | :heavy_check_mark:                                                                          | Provide the ID of the related payment.                                                      | tr_5B8cwPMGnU                                                                               |
| `requestBody`                                                                               | [*operations.UpdatePaymentRequestBody](../../models/operations/updatepaymentrequestbody.md) | :heavy_minus_sign:                                                                          | N/A                                                                                         |                                                                                             |
| `opts`                                                                                      | [][operations.Option](../../models/operations/option.md)                                    | :heavy_minus_sign:                                                                          | The options for this request.                                                               |                                                                                             |

### Response

**[*operations.UpdatePaymentResponse](../../models/operations/updatepaymentresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 422                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Cancel

Depending on the payment method, you may be able to cancel a payment for a certain amount of time — usually until
the next business day or as long as the payment status is open.

Payments may also be canceled manually from the Mollie Dashboard.

The `isCancelable` property on the [Payment object](get-payment) will indicate if the payment can be canceled.

### Example Usage

<!-- UsageSnippet language="go" operationID="cancel-payment" method="delete" path="/payments/{paymentId}" -->
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

    res, err := s.Payments.Cancel(ctx, "tr_5B8cwPMGnU", &operations.CancelPaymentRequestBody{
        Testmode: client.Bool(false),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PaymentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                   | Type                                                                                        | Required                                                                                    | Description                                                                                 | Example                                                                                     |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `ctx`                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                       | :heavy_check_mark:                                                                          | The context to use for the request.                                                         |                                                                                             |
| `paymentID`                                                                                 | *string*                                                                                    | :heavy_check_mark:                                                                          | Provide the ID of the related payment.                                                      | tr_5B8cwPMGnU                                                                               |
| `requestBody`                                                                               | [*operations.CancelPaymentRequestBody](../../models/operations/cancelpaymentrequestbody.md) | :heavy_minus_sign:                                                                          | N/A                                                                                         |                                                                                             |
| `opts`                                                                                      | [][operations.Option](../../models/operations/option.md)                                    | :heavy_minus_sign:                                                                          | The options for this request.                                                               |                                                                                             |

### Response

**[*operations.CancelPaymentResponse](../../models/operations/cancelpaymentresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 422                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## ReleaseAuthorization

Releases the full remaining authorized amount. Call this endpoint when you will not be making any additional
captures. Payment authorizations may also be released manually from the Mollie Dashboard.

Mollie will do its best to process release requests, but it is not guaranteed that it will succeed. It is up to
the issuing bank if and when the hold will be released.

If the request does succeed, the payment status will change to `canceled` for payments without captures.
If there is a successful capture, the payment will transition to `paid`.

### Example Usage

<!-- UsageSnippet language="go" operationID="release-authorization" method="post" path="/payments/{paymentId}/release-authorization" -->
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

    res, err := s.Payments.ReleaseAuthorization(ctx, "tr_5B8cwPMGnU", &operations.ReleaseAuthorizationRequestBody{
        ProfileID: client.String("pfl_5B8cwPMGnU"),
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

| Parameter                                                                                                 | Type                                                                                                      | Required                                                                                                  | Description                                                                                               | Example                                                                                                   |
| --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                     | :heavy_check_mark:                                                                                        | The context to use for the request.                                                                       |                                                                                                           |
| `paymentID`                                                                                               | *string*                                                                                                  | :heavy_check_mark:                                                                                        | Provide the ID of the related payment.                                                                    | tr_5B8cwPMGnU                                                                                             |
| `requestBody`                                                                                             | [*operations.ReleaseAuthorizationRequestBody](../../models/operations/releaseauthorizationrequestbody.md) | :heavy_minus_sign:                                                                                        | N/A                                                                                                       |                                                                                                           |
| `opts`                                                                                                    | [][operations.Option](../../models/operations/option.md)                                                  | :heavy_minus_sign:                                                                                        | The options for this request.                                                                             |                                                                                                           |

### Response

**[*operations.ReleaseAuthorizationResponse](../../models/operations/releaseauthorizationresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 422                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |