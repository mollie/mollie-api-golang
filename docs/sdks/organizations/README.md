# Organizations

## Overview

### Available Operations

* [Get](#get) - Get organization
* [GetCurrent](#getcurrent) - Get current organization
* [GetPartner](#getpartner) - Get partner status

## Get

Retrieve a single organization by its ID.

You can normally only retrieve the currently authenticated organization with this endpoint. This is primarily useful
for OAuth apps. See also [Get current organization](get-current-organization).

If you have a *partner account*', you can retrieve organization details of connected organizations.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-organization" method="get" path="/v2/organizations/{organizationId}" example="get-organization-200-1" -->
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

    res, err := s.Organizations.Get(ctx, operations.GetOrganizationRequest{
        OrganizationID: "org_1234567",
        IdempotencyKey: client.Pointer("123e4567-e89b-12d3-a456-426"),
        XClientSignature: client.Pointer("MEUCIQDx8sJvK3z..."),
        XClientSignedAt: client.Pointer("2025-01-01T12:00:00Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EntityOrganization != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetOrganizationRequest](../../models/operations/getorganizationrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.GetOrganizationResponse](../../models/operations/getorganizationresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 429                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## GetCurrent

Retrieve the currently authenticated organization. A convenient alias of the [Get organization](get-organization)
endpoint.

For a complete reference of the organization object, refer to the [Get organization](get-organization) endpoint
documentation.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-current-organization" method="get" path="/v2/organizations/me" example="get-current-organization-200-1" -->
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
            AdvancedAccessToken: client.Pointer(os.Getenv("CLIENT_ADVANCED_ACCESS_TOKEN")),
        }),
    )

    res, err := s.Organizations.GetCurrent(ctx, client.Pointer("123e4567-e89b-12d3-a456-426"), client.Pointer("MEUCIQDx8sJvK3z..."), client.Pointer("2025-01-01T12:00:00Z"))
    if err != nil {
        log.Fatal(err)
    }
    if res.EntityOrganization != nil {
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
| `opts`                                                                                                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                      | The options for this request.                                                                                                                                                           |                                                                                                                                                                                         |

### Response

**[*operations.GetCurrentOrganizationResponse](../../models/operations/getcurrentorganizationresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 429                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## GetPartner

Retrieve partnership details about the currently authenticated organization. Only relevant for so-called *partner
accounts*.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-partner-status" method="get" path="/v2/organizations/me/partner" example="get-partner-status-200-1" -->
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
            AdvancedAccessToken: client.Pointer(os.Getenv("CLIENT_ADVANCED_ACCESS_TOKEN")),
        }),
    )

    res, err := s.Organizations.GetPartner(ctx, client.Pointer("123e4567-e89b-12d3-a456-426"), client.Pointer("MEUCIQDx8sJvK3z..."), client.Pointer("2025-01-01T12:00:00Z"))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
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
| `opts`                                                                                                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                      | The options for this request.                                                                                                                                                           |                                                                                                                                                                                         |

### Response

**[*operations.GetPartnerStatusResponse](../../models/operations/getpartnerstatusresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 429                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |