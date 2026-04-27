# Oauth

## Overview

### Available Operations

* [Generate](#generate) - Generate tokens
* [Revoke](#revoke) - Revoke tokens

## Generate

Exchange the authorization code you received from the [Authorize endpoint](oauth-authorize) for an 'access token'
API credential, with which you can communicate with the Mollie API on behalf of the consenting merchant.

This endpoint can only be accessed using **OAuth client credentials**.

### Example Usage

<!-- UsageSnippet language="go" operationID="oauth-generate-tokens" method="post" path="/oauth2/tokens" example="oauth-generate-tokens" -->
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
            OAuth: client.Pointer(os.Getenv("CLIENT_O_AUTH")),
        }),
    )

    res, err := s.Oauth.Generate(ctx, client.Pointer("123e4567-e89b-12d3-a456-426"), &operations.OauthGenerateTokensRequestBody{
        GrantType: components.OauthGrantTypeAuthorizationCode,
        Code: client.Pointer("auth_..."),
        RefreshToken: client.Pointer("refresh_..."),
        RedirectURI: client.Pointer("https://example.com/redirect"),
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

| Parameter                                                                                               | Type                                                                                                    | Required                                                                                                | Description                                                                                             | Example                                                                                                 |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                   | :heavy_check_mark:                                                                                      | The context to use for the request.                                                                     |                                                                                                         |
| `idempotencyKey`                                                                                        | `*string`                                                                                               | :heavy_minus_sign:                                                                                      | A unique key to ensure idempotent requests. This key should be a UUID v4 string.                        | 123e4567-e89b-12d3-a456-426                                                                             |
| `requestBody`                                                                                           | [*operations.OauthGenerateTokensRequestBody](../../models/operations/oauthgeneratetokensrequestbody.md) | :heavy_minus_sign:                                                                                      | N/A                                                                                                     |                                                                                                         |
| `opts`                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                | :heavy_minus_sign:                                                                                      | The options for this request.                                                                           |                                                                                                         |

### Response

**[*operations.OauthGenerateTokensResponse](../../models/operations/oauthgeneratetokensresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Revoke

Revoke an access token or refresh token. Once revoked, the token can no longer be used.

Revoking a refresh token revokes all access tokens that were created using the same authorization.

This endpoint can only be accessed using **OAuth client credentials**.

### Example Usage

<!-- UsageSnippet language="go" operationID="oauth-revoke-tokens" method="delete" path="/oauth2/tokens" -->
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
            OAuth: client.Pointer(os.Getenv("CLIENT_O_AUTH")),
        }),
    )

    res, err := s.Oauth.Revoke(ctx, client.Pointer("123e4567-e89b-12d3-a456-426"), &operations.OauthRevokeTokensRequestBody{
        TokenTypeHint: components.OauthTokenTypeHintAccessToken,
        Token: "access_...",
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

| Parameter                                                                                           | Type                                                                                                | Required                                                                                            | Description                                                                                         | Example                                                                                             |
| --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                               | :heavy_check_mark:                                                                                  | The context to use for the request.                                                                 |                                                                                                     |
| `idempotencyKey`                                                                                    | `*string`                                                                                           | :heavy_minus_sign:                                                                                  | A unique key to ensure idempotent requests. This key should be a UUID v4 string.                    | 123e4567-e89b-12d3-a456-426                                                                         |
| `requestBody`                                                                                       | [*operations.OauthRevokeTokensRequestBody](../../models/operations/oauthrevoketokensrequestbody.md) | :heavy_minus_sign:                                                                                  | N/A                                                                                                 |                                                                                                     |
| `opts`                                                                                              | [][operations.Option](../../models/operations/option.md)                                            | :heavy_minus_sign:                                                                                  | The options for this request.                                                                       |                                                                                                     |

### Response

**[*operations.OauthRevokeTokensResponse](../../models/operations/oauthrevoketokensresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |