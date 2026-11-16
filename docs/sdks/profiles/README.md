# Profiles

## Overview

### Available Operations

* [Create](#create) - Create profile
* [List](#list) - List profiles
* [Get](#get) - Get profile
* [Update](#update) - Update profile
* [Delete](#delete) - Delete profile
* [GetCurrent](#getcurrent) - Get current profile

## Create

Create a profile to process payments on.

Profiles are required for payment processing. Normally they are created via the Mollie dashboard. Alternatively, you
can use this endpoint to automate profile creation.

### Example Usage: create-profile-201-1

<!-- UsageSnippet language="go" operationID="create-profile" method="post" path="/v2/profiles" example="create-profile-201-1" -->
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

    res, err := s.Profiles.Create(ctx, components.ProfileRequest{
        Name: "My website name",
        Website: "https://example.com",
        Email: "test@mollie.com",
        Phone: "+31208202070",
        Description: client.Pointer("My website description"),
        CountriesOfActivity: []string{
            "NL",
            "GB",
        },
        BusinessCategory: client.Pointer("OTHER_MERCHANDISE"),
    }, client.Pointer("123e4567-e89b-12d3-a456-426"), client.Pointer("MEUCIQDx8sJvK3z..."), client.Pointer("2025-01-01T12:00:00Z"))
    if err != nil {
        log.Fatal(err)
    }
    if res.ProfileResponse != nil {
        // handle response
    }
}
```
### Example Usage: create-profile-201-2

<!-- UsageSnippet language="go" operationID="create-profile" method="post" path="/v2/profiles" example="create-profile-201-2" -->
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

    res, err := s.Profiles.Create(ctx, components.ProfileRequest{
        Name: "My website name",
        Website: "https://example.com",
        Email: "test@mollie.com",
        Phone: "+31208202070",
        Description: client.Pointer("My website description"),
        CountriesOfActivity: []string{
            "NL",
            "GB",
        },
        BusinessCategory: client.Pointer("OTHER_MERCHANDISE"),
    }, client.Pointer("123e4567-e89b-12d3-a456-426"), client.Pointer("MEUCIQDx8sJvK3z..."), client.Pointer("2025-01-01T12:00:00Z"))
    if err != nil {
        log.Fatal(err)
    }
    if res.ProfileResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                               | Type                                                                                                                                                                                    | Required                                                                                                                                                                                | Description                                                                                                                                                                             | Example                                                                                                                                                                                 |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                   | :heavy_check_mark:                                                                                                                                                                      | The context to use for the request.                                                                                                                                                     |                                                                                                                                                                                         |
| `profileRequest`                                                                                                                                                                        | [components.ProfileRequest](../../models/components/profilerequest.md)                                                                                                                  | :heavy_check_mark:                                                                                                                                                                      | N/A                                                                                                                                                                                     |                                                                                                                                                                                         |
| `idempotencyKey`                                                                                                                                                                        | `*string`                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                      | A unique key to ensure idempotent requests. This key should be a UUID v4 string.                                                                                                        | 123e4567-e89b-12d3-a456-426                                                                                                                                                             |
| `xClientSignature`                                                                                                                                                                      | `*string`                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                      | A cryptographic signature of the request, used to verify that it was sent by you and hasn't been tampered with. Required if you've enabled request signing for your integration.        | MEUCIQDx8sJvK3z...                                                                                                                                                                      |
| `xClientSignedAt`                                                                                                                                                                       | `*string`                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                      | The timestamp, in ISO 8601 format, at which you generated the `X-Client-Signature`. Required together with `X-Client-Signature` if you've enabled request signing for your integration. | 2025-01-01T12:00:00Z                                                                                                                                                                    |
| `opts`                                                                                                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                      | The options for this request.                                                                                                                                                           |                                                                                                                                                                                         |

### Response

**[*operations.CreateProfileResponse](../../models/operations/createprofileresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 403, 422, 429           | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## List

Retrieve a list of all of your profiles.

The results are paginated.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-profiles" method="get" path="/v2/profiles" example="list-profiles-200-1" -->
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
            AdvancedAccessToken: client.Pointer(os.Getenv("CLIENT_ADVANCED_ACCESS_TOKEN")),
        }),
    )

    res, err := s.Profiles.List(ctx, operations.ListProfilesRequest{
        From: client.Pointer("pfl_QkEhN94Ba"),
        Limit: client.Pointer[int64](50),
        IdempotencyKey: client.Pointer("123e4567-e89b-12d3-a456-426"),
        XClientSignature: client.Pointer("MEUCIQDx8sJvK3z..."),
        XClientSignedAt: client.Pointer("2025-01-01T12:00:00Z"),
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListProfilesRequest](../../models/operations/listprofilesrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.ListProfilesResponse](../../models/operations/listprofilesresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 429                | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Get

Retrieve a single profile by its ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-profile" method="get" path="/v2/profiles/{profileId}" example="get-profile-200-1" -->
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

    res, err := s.Profiles.Get(ctx, operations.GetProfileRequest{
        ProfileID: "pfl_5B8cwPMGnU",
        IdempotencyKey: client.Pointer("123e4567-e89b-12d3-a456-426"),
        XClientSignature: client.Pointer("MEUCIQDx8sJvK3z..."),
        XClientSignedAt: client.Pointer("2025-01-01T12:00:00Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ProfileResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetProfileRequest](../../models/operations/getprofilerequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.GetProfileResponse](../../models/operations/getprofileresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 410, 429           | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Update

Update an existing profile.

Profiles are required for payment processing. Normally they are created and updated via the Mollie dashboard.
Alternatively, you can use this endpoint to automate profile management.

### Example Usage: update-profile-200-1

<!-- UsageSnippet language="go" operationID="update-profile" method="patch" path="/v2/profiles/{profileId}" example="update-profile-200-1" -->
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
            AdvancedAccessToken: client.Pointer(os.Getenv("CLIENT_ADVANCED_ACCESS_TOKEN")),
        }),
    )

    res, err := s.Profiles.Update(ctx, operations.UpdateProfileRequest{
        ProfileID: "pfl_5B8cwPMGnU",
        IdempotencyKey: client.Pointer("123e4567-e89b-12d3-a456-426"),
        XClientSignature: client.Pointer("MEUCIQDx8sJvK3z..."),
        XClientSignedAt: client.Pointer("2025-01-01T12:00:00Z"),
        RequestBody: operations.UpdateProfileRequestBody{
            Name: client.Pointer("My new website name"),
            Website: client.Pointer("https://example.com"),
            Email: client.Pointer("test@mollie.com"),
            Phone: client.Pointer("+31208202071"),
            Description: client.Pointer("My website description"),
            CountriesOfActivity: []string{
                "NL",
                "GB",
            },
            BusinessCategory: client.Pointer("OTHER_MERCHANDISE"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ProfileResponse != nil {
        // handle response
    }
}
```
### Example Usage: update-profile-200-2

<!-- UsageSnippet language="go" operationID="update-profile" method="patch" path="/v2/profiles/{profileId}" example="update-profile-200-2" -->
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
            AdvancedAccessToken: client.Pointer(os.Getenv("CLIENT_ADVANCED_ACCESS_TOKEN")),
        }),
    )

    res, err := s.Profiles.Update(ctx, operations.UpdateProfileRequest{
        ProfileID: "pfl_5B8cwPMGnU",
        IdempotencyKey: client.Pointer("123e4567-e89b-12d3-a456-426"),
        XClientSignature: client.Pointer("MEUCIQDx8sJvK3z..."),
        XClientSignedAt: client.Pointer("2025-01-01T12:00:00Z"),
        RequestBody: operations.UpdateProfileRequestBody{
            Name: client.Pointer("My new website name"),
            Website: client.Pointer("https://example.com"),
            Email: client.Pointer("test@mollie.com"),
            Phone: client.Pointer("+31208202071"),
            Description: client.Pointer("My website description"),
            CountriesOfActivity: []string{
                "NL",
                "GB",
            },
            BusinessCategory: client.Pointer("OTHER_MERCHANDISE"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ProfileResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UpdateProfileRequest](../../models/operations/updateprofilerequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.UpdateProfileResponse](../../models/operations/updateprofileresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 403, 404, 410, 422, 429 | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Delete

Delete a profile. A deleted profile and its related credentials can no longer be used for accepting payments.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-profile" method="delete" path="/v2/profiles/{profileId}" -->
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

    res, err := s.Profiles.Delete(ctx, "pfl_5B8cwPMGnU", client.Pointer("123e4567-e89b-12d3-a456-426"), client.Pointer("MEUCIQDx8sJvK3z..."), client.Pointer("2025-01-01T12:00:00Z"))
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                               | Type                                                                                                                                                                                    | Required                                                                                                                                                                                | Description                                                                                                                                                                             | Example                                                                                                                                                                                 |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                   | :heavy_check_mark:                                                                                                                                                                      | The context to use for the request.                                                                                                                                                     |                                                                                                                                                                                         |
| `profileID`                                                                                                                                                                             | `string`                                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                                      | Provide the ID of the related profile.                                                                                                                                                  | pfl_5B8cwPMGnU                                                                                                                                                                          |
| `idempotencyKey`                                                                                                                                                                        | `*string`                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                      | A unique key to ensure idempotent requests. This key should be a UUID v4 string.                                                                                                        | 123e4567-e89b-12d3-a456-426                                                                                                                                                             |
| `xClientSignature`                                                                                                                                                                      | `*string`                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                      | A cryptographic signature of the request, used to verify that it was sent by you and hasn't been tampered with. Required if you've enabled request signing for your integration.        | MEUCIQDx8sJvK3z...                                                                                                                                                                      |
| `xClientSignedAt`                                                                                                                                                                       | `*string`                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                      | The timestamp, in ISO 8601 format, at which you generated the `X-Client-Signature`. Required together with `X-Client-Signature` if you've enabled request signing for your integration. | 2025-01-01T12:00:00Z                                                                                                                                                                    |
| `opts`                                                                                                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                                                                                                | :heavy_minus_sign:                                                                                                                                                                      | The options for this request.                                                                                                                                                           |                                                                                                                                                                                         |

### Response

**[*operations.DeleteProfileResponse](../../models/operations/deleteprofileresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404, 410, 429           | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## GetCurrent

Retrieve the currently authenticated profile. A convenient alias of the [Get profile](get-profile)
endpoint.

For a complete reference of the profile object, refer to the [Get profile](get-profile) endpoint
documentation.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-current-profile" method="get" path="/v2/profiles/me" example="get-current-profile-200-1" -->
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

    res, err := s.Profiles.GetCurrent(ctx, client.Pointer("123e4567-e89b-12d3-a456-426"), client.Pointer("MEUCIQDx8sJvK3z..."), client.Pointer("2025-01-01T12:00:00Z"))
    if err != nil {
        log.Fatal(err)
    }
    if res.ProfileResponse != nil {
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

**[*operations.GetCurrentProfileResponse](../../models/operations/getcurrentprofileresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 429                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |