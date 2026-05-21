# Terminals

## Overview

### Available Operations

* [List](#list) - List terminals
* [Get](#get) - Get terminal
* [TerminalsRequestPairingCode](#terminalsrequestpairingcode) - Request terminal pairing code
* [TerminalsListPairingCodes](#terminalslistpairingcodes) - List terminal pairing codes
* [TerminalsGetPairingCode](#terminalsgetpairingcode) - Get terminal pairing code
* [TerminalsRevokePairingCode](#terminalsrevokepairingcode) - Revoke terminal pairing code

## List

Retrieve a list of all physical point-of-sale devices.

The results are paginated.

### Example Usage: list-terminals-200-1

<!-- UsageSnippet language="go" operationID="list-terminals" method="get" path="/v2/terminals" example="list-terminals-200-1" -->
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

    res, err := s.Terminals.List(ctx, operations.ListTerminalsRequest{
        From: client.Pointer("term_vytxeTZskVKR7C7WgdSP3d"),
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
### Example Usage: list-terminals-200-2

<!-- UsageSnippet language="go" operationID="list-terminals" method="get" path="/v2/terminals" example="list-terminals-200-2" -->
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

    res, err := s.Terminals.List(ctx, operations.ListTerminalsRequest{
        From: client.Pointer("term_vytxeTZskVKR7C7WgdSP3d"),
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListTerminalsRequest](../../models/operations/listterminalsrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.ListTerminalsResponse](../../models/operations/listterminalsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 429                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Get

Retrieve a single terminal by its ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-terminal" method="get" path="/v2/terminals/{terminalId}" example="get-terminal-200-1" -->
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

    res, err := s.Terminals.Get(ctx, "term_vytxeTZskVKR7C7WgdSP3d", client.Pointer("123e4567-e89b-12d3-a456-426"))
    if err != nil {
        log.Fatal(err)
    }
    if res.EntityTerminal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                                | Type                                                                                                                                                                                                                                                                                                                                                                                     | Required                                                                                                                                                                                                                                                                                                                                                                                 | Description                                                                                                                                                                                                                                                                                                                                                                              | Example                                                                                                                                                                                                                                                                                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                                    | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                       | The context to use for the request.                                                                                                                                                                                                                                                                                                                                                      |                                                                                                                                                                                                                                                                                                                                                                                          |
| `terminalID`                                                                                                                                                                                                                                                                                                                                                                             | `string`                                                                                                                                                                                                                                                                                                                                                                                 | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                       | Provide the ID of the related terminal.                                                                                                                                                                                                                                                                                                                                                  | term_vytxeTZskVKR7C7WgdSP3d                                                                                                                                                                                                                                                                                                                                                              |
| `testmode`                                                                                                                                                                                                                                                                                                                                                                               | `*bool`                                                                                                                                                                                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                       | Most API credentials are specifically created for either live mode or test mode. In those cases the `testmode` query<br/>parameter must not be sent. For organization-level credentials such as OAuth access tokens, you can enable test mode by<br/>setting the `testmode` query parameter to `true`.<br/><br/>Test entities cannot be retrieved when the endpoint is set to live mode, and vice versa. |                                                                                                                                                                                                                                                                                                                                                                                          |
| `idempotencyKey`                                                                                                                                                                                                                                                                                                                                                                         | `*string`                                                                                                                                                                                                                                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                       | A unique key to ensure idempotent requests. This key should be a UUID v4 string.                                                                                                                                                                                                                                                                                                         | 123e4567-e89b-12d3-a456-426                                                                                                                                                                                                                                                                                                                                                              |
| `opts`                                                                                                                                                                                                                                                                                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                       | The options for this request.                                                                                                                                                                                                                                                                                                                                                            |                                                                                                                                                                                                                                                                                                                                                                                          |

### Response

**[*operations.GetTerminalResponse](../../models/operations/getterminalresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 429                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## TerminalsRequestPairingCode

> ℹ️ **Test mode**
>
> This endpoint currently does not support test mode yet.

Request a pairing code to onboard a point-of-sale terminal.

The response includes a human-readable `code` for manual entry on the terminal, and a QR Code as a
base64 encoded SVG data URI for scanning if you specify the query parameter `include` with value `details.qrCode`.

Pairing codes expire after 90 days (see `expiresAt`) and can be used multiple times.

### Example Usage

<!-- UsageSnippet language="go" operationID="terminals-request-pairing-code" method="post" path="/v2/terminals/pairing-codes" -->
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

    res, err := s.Terminals.TerminalsRequestPairingCode(ctx, nil, client.Pointer("123e4567-e89b-12d3-a456-426"), &operations.TerminalsRequestPairingCodeRequestBody{
        ProfileID: "pfl_jA9bC4DkFj3G",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EntityPairingCode != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                               | Type                                                                                                                    | Required                                                                                                                | Description                                                                                                             | Example                                                                                                                 |
| ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                                   | :heavy_check_mark:                                                                                                      | The context to use for the request.                                                                                     |                                                                                                                         |
| `include`                                                                                                               | `*string`                                                                                                               | :heavy_minus_sign:                                                                                                      | Include additional information in the response.                                                                         |                                                                                                                         |
| `idempotencyKey`                                                                                                        | `*string`                                                                                                               | :heavy_minus_sign:                                                                                                      | A unique key to ensure idempotent requests. This key should be a UUID v4 string.                                        | 123e4567-e89b-12d3-a456-426                                                                                             |
| `requestBody`                                                                                                           | [*operations.TerminalsRequestPairingCodeRequestBody](../../models/operations/terminalsrequestpairingcoderequestbody.md) | :heavy_minus_sign:                                                                                                      | N/A                                                                                                                     |                                                                                                                         |
| `opts`                                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                                | :heavy_minus_sign:                                                                                                      | The options for this request.                                                                                           |                                                                                                                         |

### Response

**[*operations.TerminalsRequestPairingCodeResponse](../../models/operations/terminalsrequestpairingcoderesponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 422, 429                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## TerminalsListPairingCodes

> ℹ️ **Test mode**
>
> This endpoint currently does not support test mode yet.

Returns all pairing codes: `active`, `expired`, and `revoked`. Results are paginated.

### Example Usage

<!-- UsageSnippet language="go" operationID="terminals-list-pairing-codes" method="get" path="/v2/terminals/pairing-codes" -->
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
        client.WithProfileID("<id>"),
        client.WithSecurity(components.Security{
            APIKey: client.Pointer(os.Getenv("CLIENT_API_KEY")),
        }),
    )

    res, err := s.Terminals.TerminalsListPairingCodes(ctx, operations.TerminalsListPairingCodesRequest{
        Limit: client.Pointer[int64](50),
        Sort: components.SortingDesc.ToPointer(),
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.TerminalsListPairingCodesRequest](../../models/operations/terminalslistpairingcodesrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.TerminalsListPairingCodesResponse](../../models/operations/terminalslistpairingcodesresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 429                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## TerminalsGetPairingCode

> ℹ️ **Test mode**
>
> This endpoint currently does not support test mode yet.

Get a pairing code to onboard a point-of-sale terminal.

The response includes a human-readable `code` for manual entry on the terminal and, optionally, a QR Code as a
base64 encoded SVG data URI when you use the `include` query parameter with value `details.qrCode`.

### Example Usage

<!-- UsageSnippet language="go" operationID="terminals-get-pairing-code" method="get" path="/v2/terminals/pairing-codes/{pairingCodeId}" -->
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

    res, err := s.Terminals.TerminalsGetPairingCode(ctx, "termpc_R7gX5Ea9bC4DkFj3G", nil, client.Pointer("123e4567-e89b-12d3-a456-426"))
    if err != nil {
        log.Fatal(err)
    }
    if res.EntityPairingCode != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |                                                                                  |
| `pairingCodeID`                                                                  | `string`                                                                         | :heavy_check_mark:                                                               | Provide the ID of the terminal pairing code.                                     | termpc_R7gX5Ea9bC4DkFj3G                                                         |
| `include`                                                                        | `*string`                                                                        | :heavy_minus_sign:                                                               | Include additional information in the response.                                  |                                                                                  |
| `idempotencyKey`                                                                 | `*string`                                                                        | :heavy_minus_sign:                                                               | A unique key to ensure idempotent requests. This key should be a UUID v4 string. | 123e4567-e89b-12d3-a456-426                                                      |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |                                                                                  |

### Response

**[*operations.TerminalsGetPairingCodeResponse](../../models/operations/terminalsgetpairingcoderesponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 429                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## TerminalsRevokePairingCode

> ℹ️ **Test mode**
>
> This endpoint currently does not support test mode yet.

Revoke a pairing code, preventing the onboarding of new point-of-sale terminals.

Terminals that have already paired with this code are not affected.

### Example Usage

<!-- UsageSnippet language="go" operationID="terminals-revoke-pairing-code" method="delete" path="/v2/terminals/pairing-codes/{pairingCodeId}" -->
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

    res, err := s.Terminals.TerminalsRevokePairingCode(ctx, "termpc_R7gX5Ea9bC4DkFj3G", client.Pointer("123e4567-e89b-12d3-a456-426"))
    if err != nil {
        log.Fatal(err)
    }
    if res.EntityPairingCode != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |                                                                                  |
| `pairingCodeID`                                                                  | `string`                                                                         | :heavy_check_mark:                                                               | Provide the ID of the terminal pairing code.                                     | termpc_R7gX5Ea9bC4DkFj3G                                                         |
| `idempotencyKey`                                                                 | `*string`                                                                        | :heavy_minus_sign:                                                               | A unique key to ensure idempotent requests. This key should be a UUID v4 string. | 123e4567-e89b-12d3-a456-426                                                      |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |                                                                                  |

### Response

**[*operations.TerminalsRevokePairingCodeResponse](../../models/operations/terminalsrevokepairingcoderesponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 422, 429           | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |