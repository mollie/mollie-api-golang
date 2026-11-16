# Permissions

## Overview

### Available Operations

* [List](#list) - List permissions
* [Get](#get) - Get permission

## List

Retrieve a list of all permissions available to the current access token.

The results are **not** paginated.

### Example Usage: list-permissions-200-1

<!-- UsageSnippet language="go" operationID="list-permissions" method="get" path="/v2/permissions" example="list-permissions-200-1" -->
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

    res, err := s.Permissions.List(ctx, client.Pointer("123e4567-e89b-12d3-a456-426"), client.Pointer("MEUCIQDx8sJvK3z..."), client.Pointer("2025-01-01T12:00:00Z"))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```
### Example Usage: list-permissions-200-2

<!-- UsageSnippet language="go" operationID="list-permissions" method="get" path="/v2/permissions" example="list-permissions-200-2" -->
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

    res, err := s.Permissions.List(ctx, client.Pointer("123e4567-e89b-12d3-a456-426"), client.Pointer("MEUCIQDx8sJvK3z..."), client.Pointer("2025-01-01T12:00:00Z"))
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

**[*operations.ListPermissionsResponse](../../models/operations/listpermissionsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 429                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Get

Retrieve a single permission by its ID, and see if the permission is granted to the current access token.

### Example Usage: get-permission-200-1

<!-- UsageSnippet language="go" operationID="get-permission" method="get" path="/v2/permissions/{permissionId}" example="get-permission-200-1" -->
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

    res, err := s.Permissions.Get(ctx, operations.GetPermissionRequest{
        PermissionID: "payments.read",
        IdempotencyKey: client.Pointer("123e4567-e89b-12d3-a456-426"),
        XClientSignature: client.Pointer("MEUCIQDx8sJvK3z..."),
        XClientSignedAt: client.Pointer("2025-01-01T12:00:00Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EntityPermission != nil {
        // handle response
    }
}
```
### Example Usage: get-permission-200-2

<!-- UsageSnippet language="go" operationID="get-permission" method="get" path="/v2/permissions/{permissionId}" example="get-permission-200-2" -->
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

    res, err := s.Permissions.Get(ctx, operations.GetPermissionRequest{
        PermissionID: "payments.read",
        IdempotencyKey: client.Pointer("123e4567-e89b-12d3-a456-426"),
        XClientSignature: client.Pointer("MEUCIQDx8sJvK3z..."),
        XClientSignedAt: client.Pointer("2025-01-01T12:00:00Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EntityPermission != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetPermissionRequest](../../models/operations/getpermissionrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.GetPermissionResponse](../../models/operations/getpermissionresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 429                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |