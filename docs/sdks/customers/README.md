# Customers
(*Customers*)

## Overview

### Available Operations

* [Create](#create) - Create customer
* [List](#list) - List customers
* [Get](#get) - Get customer
* [Update](#update) - Update customer
* [Delete](#delete) - Delete customer
* [CreatePayment](#createpayment) - Create customer payment
* [ListPayments](#listpayments) - List customer payments

## Create

Creates a simple minimal representation of a customer. Payments, recurring mandates, and subscriptions can be linked
to this customer object, which simplifies management of recurring payments.

Once registered, customers will also appear in your Mollie dashboard.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-customer" method="post" path="/customers" -->
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

    res, err := s.Customers.Create(ctx, client.Pointer("123e4567-e89b-12d3-a456-426"), &components.EntityCustomer{
        ID: client.Pointer("cst_5B8cwPMGnU"),
        Name: client.Pointer("John Doe"),
        Email: client.Pointer("example@email.com"),
        Locale: components.LocaleResponseEnUs.ToPointer(),
        Testmode: client.Pointer(false),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |                                                                                  |
| `idempotencyKey`                                                                 | **string*                                                                        | :heavy_minus_sign:                                                               | A unique key to ensure idempotent requests. This key should be a UUID v4 string. | 123e4567-e89b-12d3-a456-426                                                      |
| `entityCustomer`                                                                 | [*components.EntityCustomer](../../models/components/entitycustomer.md)          | :heavy_minus_sign:                                                               | N/A                                                                              |                                                                                  |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |                                                                                  |

### Response

**[*operations.CreateCustomerResponse](../../models/operations/createcustomerresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## List

Retrieve a list of all customers.

The results are paginated.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-customers" method="get" path="/customers" -->
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

    res, err := s.Customers.List(ctx, operations.ListCustomersRequest{
        From: client.Pointer("cst_5B8cwPMGnU"),
        Limit: client.Pointer[int64](50),
        Sort: components.ListSortDesc.ToPointer(),
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListCustomersRequest](../../models/operations/listcustomersrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.ListCustomersResponse](../../models/operations/listcustomersresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Get

Retrieve a single customer by its ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-customer" method="get" path="/customers/{customerId}" -->
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

    res, err := s.Customers.Get(ctx, "cst_5B8cwPMGnU", client.Pointer("events"), client.Pointer(false), client.Pointer("123e4567-e89b-12d3-a456-426"))
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
| `customerID`                                                                                                                                                                                                                                                                                                                                                                           | *string*                                                                                                                                                                                                                                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                     | Provide the ID of the related customer.                                                                                                                                                                                                                                                                                                                                                | cst_5B8cwPMGnU                                                                                                                                                                                                                                                                                                                                                                         |
| `include`                                                                                                                                                                                                                                                                                                                                                                              | **string*                                                                                                                                                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | This endpoint allows you to include additional information via the `include` query string parameter.                                                                                                                                                                                                                                                                                   |                                                                                                                                                                                                                                                                                                                                                                                        |
| `testmode`                                                                                                                                                                                                                                                                                                                                                                             | **bool*                                                                                                                                                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | Most API credentials are specifically created for either live mode or test mode. In those cases the `testmode` query<br/>parameter can be omitted. For organization-level credentials such as OAuth access tokens, you can enable test mode by<br/>setting the `testmode` query parameter to `true`.<br/><br/>Test entities cannot be retrieved when the endpoint is set to live mode, and vice versa. | false                                                                                                                                                                                                                                                                                                                                                                                  |
| `idempotencyKey`                                                                                                                                                                                                                                                                                                                                                                       | **string*                                                                                                                                                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | A unique key to ensure idempotent requests. This key should be a UUID v4 string.                                                                                                                                                                                                                                                                                                       | 123e4567-e89b-12d3-a456-426                                                                                                                                                                                                                                                                                                                                                            |
| `opts`                                                                                                                                                                                                                                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                     | The options for this request.                                                                                                                                                                                                                                                                                                                                                          |                                                                                                                                                                                                                                                                                                                                                                                        |

### Response

**[*operations.GetCustomerResponse](../../models/operations/getcustomerresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Update

Update an existing customer.

For an in-depth explanation of each parameter, refer to the [Create customer](create-customer) endpoint.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-customer" method="patch" path="/customers/{customerId}" -->
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

    res, err := s.Customers.Update(ctx, "cst_5B8cwPMGnU", client.Pointer("123e4567-e89b-12d3-a456-426"), &components.EntityCustomer{
        ID: client.Pointer("cst_5B8cwPMGnU"),
        Name: client.Pointer("John Doe"),
        Email: client.Pointer("example@email.com"),
        Locale: components.LocaleResponseEnUs.ToPointer(),
        Testmode: client.Pointer(false),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CustomerResponse != nil {
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
| `entityCustomer`                                                                 | [*components.EntityCustomer](../../models/components/entitycustomer.md)          | :heavy_minus_sign:                                                               | N/A                                                                              |                                                                                  |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |                                                                                  |

### Response

**[*operations.UpdateCustomerResponse](../../models/operations/updatecustomerresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Delete

Delete a customer. All mandates and subscriptions created for this customer will be canceled as well.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-customer" method="delete" path="/customers/{customerId}" -->
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

    res, err := s.Customers.Delete(ctx, "cst_5B8cwPMGnU", client.Pointer("123e4567-e89b-12d3-a456-426"), &operations.DeleteCustomerRequestBody{
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

| Parameter                                                                                     | Type                                                                                          | Required                                                                                      | Description                                                                                   | Example                                                                                       |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `ctx`                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                         | :heavy_check_mark:                                                                            | The context to use for the request.                                                           |                                                                                               |
| `customerID`                                                                                  | *string*                                                                                      | :heavy_check_mark:                                                                            | Provide the ID of the related customer.                                                       | cst_5B8cwPMGnU                                                                                |
| `idempotencyKey`                                                                              | **string*                                                                                     | :heavy_minus_sign:                                                                            | A unique key to ensure idempotent requests. This key should be a UUID v4 string.              | 123e4567-e89b-12d3-a456-426                                                                   |
| `requestBody`                                                                                 | [*operations.DeleteCustomerRequestBody](../../models/operations/deletecustomerrequestbody.md) | :heavy_minus_sign:                                                                            | N/A                                                                                           |                                                                                               |
| `opts`                                                                                        | [][operations.Option](../../models/operations/option.md)                                      | :heavy_minus_sign:                                                                            | The options for this request.                                                                 |                                                                                               |

### Response

**[*operations.DeleteCustomerResponse](../../models/operations/deletecustomerresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## CreatePayment

Creates a payment for the customer.

Linking customers to payments enables you to:

* Keep track of payment preferences for your customers
* Allow your customers to charge a previously used credit card with a single click in our hosted checkout
* Improve payment insights in the Mollie dashboard
* Use recurring payments

This endpoint is effectively an alias of the [Create payment endpoint](create-payment) with the `customerId`
parameter predefined.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-customer-payment" method="post" path="/customers/{customerId}/payments" -->
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
            APIKey: client.Pointer(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.Customers.CreatePayment(ctx, "cst_5B8cwPMGnU", client.Pointer("123e4567-e89b-12d3-a456-426"), &components.PaymentRequest{
        ID: client.Pointer("tr_5B8cwPMGnU"),
        Description: client.Pointer("Chess Board"),
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
        RedirectURL: client.Pointer("https://example.org/redirect"),
        CancelURL: client.Pointer("https://example.org/cancel"),
        WebhookURL: client.Pointer("https://example.org/webhooks"),
        Lines: []components.PaymentRequestLine{
            components.PaymentRequestLine{
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
                Categories: []components.PaymentRequestCategory{
                    components.PaymentRequestCategoryMeal,
                    components.PaymentRequestCategoryEco,
                },
                ImageURL: client.Pointer("https://..."),
                ProductURL: client.Pointer("https://..."),
                Recurring: &components.RecurringLineItem{
                    Description: client.Pointer("Gym subscription"),
                    Interval: "12 months",
                    Amount: &components.Amount{
                        Currency: "EUR",
                        Value: "10.00",
                    },
                    Times: client.Pointer[int64](1),
                    StartDate: client.Pointer("2024-12-12"),
                },
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
        Locale: components.LocaleEnUs.ToPointer(),
        Method: components.MethodIdeal.ToPointer(),
        Issuer: client.Pointer("ideal_INGBNL2A"),
        RestrictPaymentMethodsToCountry: client.Pointer("NL"),
        CaptureMode: components.CaptureModeManual.ToPointer(),
        CaptureDelay: client.Pointer("8 hours"),
        ApplicationFee: &components.PaymentRequestApplicationFee{
            Amount: &components.Amount{
                Currency: "EUR",
                Value: "10.00",
            },
            Description: client.Pointer("10"),
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
                ReleaseDate: client.Pointer("2024-12-12"),
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
        SubscriptionID: client.Pointer("sub_5B8cwPMGnU"),
        MandateID: client.Pointer("mdt_5B8cwPMGnU"),
        CustomerID: client.Pointer("cst_5B8cwPMGnU"),
        ProfileID: client.Pointer("pfl_5B8cwPMGnU"),
        SettlementID: client.Pointer("stl_5B8cwPMGnU"),
        OrderID: client.Pointer("ord_5B8cwPMGnU"),
        DueDate: client.Pointer("2025-01-01"),
        Testmode: client.Pointer(false),
        ApplePayPaymentToken: client.Pointer("{\"paymentData\": {\"version\": \"EC_v1\", \"data\": \"vK3BbrCbI/....\"}}"),
        Company: &components.Company{
            RegistrationNumber: client.Pointer("12345678"),
            VatNumber: client.Pointer("NL123456789B01"),
            EntityType: nil,
        },
        CardToken: client.Pointer("tkn_12345"),
        VoucherNumber: client.Pointer("1234567890"),
        VoucherPin: client.Pointer("1234"),
        ConsumerDateOfBirth: types.MustNewDateFromString("2000-01-01"),
        SessionID: nil,
        DigitalGoods: client.Pointer(true),
        CustomerReference: client.Pointer("1234567890"),
        TerminalID: client.Pointer("term_1234567890"),
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |                                                                                  |
| `customerID`                                                                     | *string*                                                                         | :heavy_check_mark:                                                               | Provide the ID of the related customer.                                          | cst_5B8cwPMGnU                                                                   |
| `idempotencyKey`                                                                 | **string*                                                                        | :heavy_minus_sign:                                                               | A unique key to ensure idempotent requests. This key should be a UUID v4 string. | 123e4567-e89b-12d3-a456-426                                                      |
| `paymentRequest`                                                                 | [*components.PaymentRequest](../../models/components/paymentrequest.md)          | :heavy_minus_sign:                                                               | N/A                                                                              |                                                                                  |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |                                                                                  |

### Response

**[*operations.CreateCustomerPaymentResponse](../../models/operations/createcustomerpaymentresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 422                     | application/hal+json    |
| apierrors.ErrorResponse | 503                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## ListPayments

Retrieve all payments linked to the customer.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-customer-payments" method="get" path="/customers/{customerId}/payments" -->
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

    res, err := s.Customers.ListPayments(ctx, operations.ListCustomerPaymentsRequest{
        CustomerID: "cst_5B8cwPMGnU",
        From: client.Pointer("tr_5B8cwPMGnU"),
        Limit: client.Pointer[int64](50),
        Sort: components.ListSortDesc.ToPointer(),
        ProfileID: client.Pointer("pfl_5B8cwPMGnU"),
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ListCustomerPaymentsRequest](../../models/operations/listcustomerpaymentsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ListCustomerPaymentsResponse](../../models/operations/listcustomerpaymentsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |