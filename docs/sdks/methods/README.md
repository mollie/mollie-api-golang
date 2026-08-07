# Methods

## Overview

### Available Operations

* [List](#list) - List payment methods
* [All](#all) - List all payment methods
* [Get](#get) - Get payment method
* [Enable](#enable) - Enable payment method
* [Disable](#disable) - Disable payment method
* [EnableIssuer](#enableissuer) - Enable payment method issuer
* [DisableIssuer](#disableissuer) - Disable payment method issuer

## List

Retrieve all enabled payment methods. The results of this endpoint are
**not** paginated — unlike most other list endpoints in our API.

For test mode, all pending and enabled payment methods are returned. If no
payment methods are requested yet, the most popular payment methods are returned in the test mode. For live
mode, only fully enabled payment methods are returned.

Payment methods can be requested and enabled via the Mollie Dashboard, or
via the [Enable payment method endpoint](enable-method) of the Profiles API.

The list can optionally be filtered using a number of parameters described
below.

By default, only payment methods for the Euro currency are returned. If you
wish to retrieve payment methods which exclusively support other currencies (e.g. Twint), you need to use the
`amount` parameters.

ℹ️ **Note:** This endpoint only returns **online** payment methods. If you wish to retrieve the information about
a non-online payment method, you can use the [Get payment method endpoint](get-method).

### Example Usage: list-method-200-3

<!-- UsageSnippet language="go" operationID="list-methods" method="get" path="/v2/methods" example="list-method-200-3" -->
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
        client.WithProfileID("pfl_5B8cwPMGnU"),
        client.WithTestmode(false),
        client.WithSecurity(components.Security{
            APIKey: client.Pointer(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.Methods.List(ctx, operations.ListMethodsRequest{
        SequenceType: components.SequenceTypeOneoff.ToPointer(),
        Locale: components.LocaleEnUs.ToPointer(),
        Amount: &components.Amount{
            Currency: "EUR",
            Value: "10.00",
        },
        BillingCountry: client.Pointer("DE"),
        IncludeWallets: components.MethodIncludeWalletsParameterApplepay.ToPointer(),
        OrderLineCategories: components.LineCategoriesEco.ToPointer(),
        Include: client.Pointer("issuers"),
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
### Example Usage: list-methods-200-1

<!-- UsageSnippet language="go" operationID="list-methods" method="get" path="/v2/methods" example="list-methods-200-1" -->
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
        client.WithProfileID("pfl_5B8cwPMGnU"),
        client.WithTestmode(false),
        client.WithSecurity(components.Security{
            APIKey: client.Pointer(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.Methods.List(ctx, operations.ListMethodsRequest{
        SequenceType: components.SequenceTypeOneoff.ToPointer(),
        Locale: components.LocaleEnUs.ToPointer(),
        Amount: &components.Amount{
            Currency: "EUR",
            Value: "10.00",
        },
        BillingCountry: client.Pointer("DE"),
        IncludeWallets: components.MethodIncludeWalletsParameterApplepay.ToPointer(),
        OrderLineCategories: components.LineCategoriesEco.ToPointer(),
        Include: client.Pointer("issuers"),
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
### Example Usage: list-methods-200-2

<!-- UsageSnippet language="go" operationID="list-methods" method="get" path="/v2/methods" example="list-methods-200-2" -->
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
        client.WithProfileID("pfl_5B8cwPMGnU"),
        client.WithTestmode(false),
        client.WithSecurity(components.Security{
            APIKey: client.Pointer(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.Methods.List(ctx, operations.ListMethodsRequest{
        SequenceType: components.SequenceTypeOneoff.ToPointer(),
        Locale: components.LocaleEnUs.ToPointer(),
        Amount: &components.Amount{
            Currency: "EUR",
            Value: "10.00",
        },
        BillingCountry: client.Pointer("DE"),
        IncludeWallets: components.MethodIncludeWalletsParameterApplepay.ToPointer(),
        OrderLineCategories: components.LineCategoriesEco.ToPointer(),
        Include: client.Pointer("issuers"),
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ListMethodsRequest](../../models/operations/listmethodsrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.ListMethodsResponse](../../models/operations/listmethodsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 429                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## All

Retrieve all payment methods that Mollie offers, regardless of the eligibility of the organization for the specific
method. The results of this endpoint are **not** paginated — unlike most other list endpoints in our API.

The list can optionally be filtered using a number of parameters described below.

ℹ️ **Note:** This endpoint only returns **online** payment methods. If you wish to retrieve the information about
a non-online payment method, you can use the [Get payment method endpoint](get-method).

### Example Usage: list-all-methods-200-1

<!-- UsageSnippet language="go" operationID="list-all-methods" method="get" path="/v2/methods/all" example="list-all-methods-200-1" -->
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
        client.WithProfileID("pfl_5B8cwPMGnU"),
        client.WithTestmode(false),
        client.WithSecurity(components.Security{
            APIKey: client.Pointer(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.Methods.All(ctx, operations.ListAllMethodsRequest{
        Locale: components.LocaleEnUs.ToPointer(),
        Amount: &components.Amount{
            Currency: "EUR",
            Value: "10.00",
        },
        Include: client.Pointer("issuers"),
        SequenceType: components.SequenceTypeOneoff.ToPointer(),
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
### Example Usage: list-all-methods-200-2

<!-- UsageSnippet language="go" operationID="list-all-methods" method="get" path="/v2/methods/all" example="list-all-methods-200-2" -->
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
        client.WithProfileID("pfl_5B8cwPMGnU"),
        client.WithTestmode(false),
        client.WithSecurity(components.Security{
            APIKey: client.Pointer(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.Methods.All(ctx, operations.ListAllMethodsRequest{
        Locale: components.LocaleEnUs.ToPointer(),
        Amount: &components.Amount{
            Currency: "EUR",
            Value: "10.00",
        },
        Include: client.Pointer("issuers"),
        SequenceType: components.SequenceTypeOneoff.ToPointer(),
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
### Example Usage: list-all-methods-200-3

<!-- UsageSnippet language="go" operationID="list-all-methods" method="get" path="/v2/methods/all" example="list-all-methods-200-3" -->
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
        client.WithProfileID("pfl_5B8cwPMGnU"),
        client.WithTestmode(false),
        client.WithSecurity(components.Security{
            APIKey: client.Pointer(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.Methods.All(ctx, operations.ListAllMethodsRequest{
        Locale: components.LocaleEnUs.ToPointer(),
        Amount: &components.Amount{
            Currency: "EUR",
            Value: "10.00",
        },
        Include: client.Pointer("issuers"),
        SequenceType: components.SequenceTypeOneoff.ToPointer(),
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
### Example Usage: list-all-methods-200-4

<!-- UsageSnippet language="go" operationID="list-all-methods" method="get" path="/v2/methods/all" example="list-all-methods-200-4" -->
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
        client.WithProfileID("pfl_5B8cwPMGnU"),
        client.WithTestmode(false),
        client.WithSecurity(components.Security{
            APIKey: client.Pointer(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.Methods.All(ctx, operations.ListAllMethodsRequest{
        Locale: components.LocaleEnUs.ToPointer(),
        Amount: &components.Amount{
            Currency: "EUR",
            Value: "10.00",
        },
        Include: client.Pointer("issuers"),
        SequenceType: components.SequenceTypeOneoff.ToPointer(),
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListAllMethodsRequest](../../models/operations/listallmethodsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.ListAllMethodsResponse](../../models/operations/listallmethodsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 429                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Get

Retrieve a single payment method by its ID.

If a method is not available on this profile, a `404 Not Found` response is
returned. If the method is available but not enabled yet, a status `403 Forbidden` is returned. You can enable
payments methods via the [Enable payment method endpoint](enable-method) of the Profiles API, or via
the Mollie Dashboard.

If you do not know the method's ID, you can use the [methods list
endpoint](list-methods) to retrieve all payment methods that are available.

Additionally, it is possible to check if wallet methods such as Apple Pay
are enabled by passing the wallet ID (`applepay`) as the method ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-method" method="get" path="/v2/methods/{methodId}" example="get-method-200-1" -->
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
        client.WithProfileID("pfl_5B8cwPMGnU"),
        client.WithTestmode(false),
        client.WithSecurity(components.Security{
            APIKey: client.Pointer(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.Methods.Get(ctx, operations.GetMethodRequest{
        MethodID: components.MethodEnumIdeal.ToPointer(),
        Locale: components.LocaleEnUs.ToPointer(),
        Currency: client.Pointer("EUR"),
        Include: client.Pointer("issuers"),
        SequenceType: components.SequenceTypeOneoff.ToPointer(),
        IdempotencyKey: client.Pointer("123e4567-e89b-12d3-a456-426"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EntityMethodGet != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetMethodRequest](../../models/operations/getmethodrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.GetMethodResponse](../../models/operations/getmethodresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404, 429           | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Enable

Enable a payment method on a specific profile.

When using a profile-specific API credential, the alias `me` can be used
instead of the profile ID to refer to the current profile.

Some payment methods require extra steps in order to be activated. In cases
where a step at the payment method provider needs to be completed first, the status will be set to
`pending-external` and the response will contain a link to complete the activation at the provider.

To enable voucher or gift card issuers, refer to the [Enable payment method issuer](enable-method-issuer) endpoint.

### Example Usage: enable-method-200-1

<!-- UsageSnippet language="go" operationID="enable-method" method="post" path="/v2/profiles/{profileId}/methods/{methodId}" example="enable-method-200-1" -->
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

    res, err := s.Methods.Enable(ctx, "pfl_5B8cwPMGnU", components.MethodEnumIdeal.ToPointer(), client.Pointer("123e4567-e89b-12d3-a456-426"))
    if err != nil {
        log.Fatal(err)
    }
    if res.EntityMethodGet != nil {
        // handle response
    }
}
```
### Example Usage: enable-method-200-2

<!-- UsageSnippet language="go" operationID="enable-method" method="post" path="/v2/profiles/{profileId}/methods/{methodId}" example="enable-method-200-2" -->
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

    res, err := s.Methods.Enable(ctx, "pfl_5B8cwPMGnU", components.MethodEnumIdeal.ToPointer(), client.Pointer("123e4567-e89b-12d3-a456-426"))
    if err != nil {
        log.Fatal(err)
    }
    if res.EntityMethodGet != nil {
        // handle response
    }
}
```
### Example Usage: enable-method-200-3

<!-- UsageSnippet language="go" operationID="enable-method" method="post" path="/v2/profiles/{profileId}/methods/{methodId}" example="enable-method-200-3" -->
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

    res, err := s.Methods.Enable(ctx, "pfl_5B8cwPMGnU", components.MethodEnumIdeal.ToPointer(), client.Pointer("123e4567-e89b-12d3-a456-426"))
    if err != nil {
        log.Fatal(err)
    }
    if res.EntityMethodGet != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |                                                                                  |
| `profileID`                                                                      | `string`                                                                         | :heavy_check_mark:                                                               | Provide the ID of the related profile.                                           | pfl_5B8cwPMGnU                                                                   |
| `methodID`                                                                       | [*components.MethodEnum](../../models/components/methodenum.md)                  | :heavy_check_mark:                                                               | Provide the ID of the related payment method.                                    | ideal                                                                            |
| `idempotencyKey`                                                                 | `*string`                                                                        | :heavy_minus_sign:                                                               | A unique key to ensure idempotent requests. This key should be a UUID v4 string. | 123e4567-e89b-12d3-a456-426                                                      |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |                                                                                  |

### Response

**[*operations.EnableMethodResponse](../../models/operations/enablemethodresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 429                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Disable

Disable a payment method on a specific profile.

When using a profile-specific API credential, the alias `me` can be used
instead of the profile ID to refer to the current profile.

### Example Usage

<!-- UsageSnippet language="go" operationID="disable-method" method="delete" path="/v2/profiles/{profileId}/methods/{methodId}" -->
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

    res, err := s.Methods.Disable(ctx, "pfl_5B8cwPMGnU", components.MethodEnumIdeal.ToPointer(), client.Pointer("123e4567-e89b-12d3-a456-426"))
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |                                                                                  |
| `profileID`                                                                      | `string`                                                                         | :heavy_check_mark:                                                               | Provide the ID of the related profile.                                           | pfl_5B8cwPMGnU                                                                   |
| `methodID`                                                                       | [*components.MethodEnum](../../models/components/methodenum.md)                  | :heavy_check_mark:                                                               | Provide the ID of the related payment method.                                    | ideal                                                                            |
| `idempotencyKey`                                                                 | `*string`                                                                        | :heavy_minus_sign:                                                               | A unique key to ensure idempotent requests. This key should be a UUID v4 string. | 123e4567-e89b-12d3-a456-426                                                      |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |                                                                                  |

### Response

**[*operations.DisableMethodResponse](../../models/operations/disablemethodresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 429                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## EnableIssuer

Enable an issuer for a payment method on a specific profile.

Currently only the payment methods `voucher` and `giftcard` are supported.

When using a profile-specific API credential, the alias `me` can be used instead of the profile ID to refer to the
current profile.

### Example Usage

<!-- UsageSnippet language="go" operationID="enable-method-issuer" method="post" path="/v2/profiles/{profileId}/methods/{methodId}/issuers/{issuerId}" example="enable-method-issuer-200-1" -->
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

    res, err := s.Methods.EnableIssuer(ctx, operations.EnableMethodIssuerRequest{
        ProfileID: "pfl_5B8cwPMGnU",
        MethodID: components.MethodIDWithIssuerVoucher,
        IssuerID: "edenred-france-sports",
        IdempotencyKey: client.Pointer("123e4567-e89b-12d3-a456-426"),
        RequestBody: &operations.EnableMethodIssuerRequestBody{
            ContractID: client.Pointer("ideal"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OneOf != nil {
        switch res.OneOf.Type {
            case operations.EnableMethodIssuerResponseBodyTypeGiftcard:
                // res.OneOf.Giftcard is populated
            case operations.EnableMethodIssuerResponseBodyTypeVoucher:
                // res.OneOf.Voucher is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.EnableMethodIssuerRequest](../../models/operations/enablemethodissuerrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.EnableMethodIssuerResponse](../../models/operations/enablemethodissuerresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 429                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## DisableIssuer

Disable an issuer for a payment method on a specific profile.

Currently only the payment methods `voucher` and `giftcard` are supported.

When using a profile-specific API credential, the alias `me` can be used instead of the profile ID to refer to the
current profile.

### Example Usage

<!-- UsageSnippet language="go" operationID="disable-method-issuer" method="delete" path="/v2/profiles/{profileId}/methods/{methodId}/issuers/{issuerId}" -->
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

    res, err := s.Methods.DisableIssuer(ctx, "pfl_5B8cwPMGnU", components.MethodIDWithIssuerVoucher, "edenred-france-sports", client.Pointer("123e4567-e89b-12d3-a456-426"))
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |                                                                                  |
| `profileID`                                                                      | `string`                                                                         | :heavy_check_mark:                                                               | Provide the ID of the related profile.                                           | pfl_5B8cwPMGnU                                                                   |
| `methodID`                                                                       | [components.MethodIDWithIssuer](../../models/components/methodidwithissuer.md)   | :heavy_check_mark:                                                               | Provide the ID of the related payment method.                                    | voucher                                                                          |
| `issuerID`                                                                       | `string`                                                                         | :heavy_check_mark:                                                               | Provide the ID of the related issuer.                                            | edenred-france-sports                                                            |
| `idempotencyKey`                                                                 | `*string`                                                                        | :heavy_minus_sign:                                                               | A unique key to ensure idempotent requests. This key should be a UUID v4 string. | 123e4567-e89b-12d3-a456-426                                                      |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |                                                                                  |

### Response

**[*operations.DisableMethodIssuerResponse](../../models/operations/disablemethodissuerresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 429                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |