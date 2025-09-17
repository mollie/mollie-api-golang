# Onboarding
(*Onboarding*)

## Overview

### Available Operations

* [Get](#get) - Get onboarding status
* [Submit](#submit) - Submit onboarding data

## Get

Retrieve the onboarding status of the currently authenticated organization.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-onboarding-status" method="get" path="/onboarding/me" -->
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

    res, err := s.Onboarding.Get(ctx, client.Pointer("123e4567-e89b-12d3-a456-426"))
    if err != nil {
        log.Fatal(err)
    }
    if res.EntityOnboardingStatus != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |                                                                                  |
| `idempotencyKey`                                                                 | **string*                                                                        | :heavy_minus_sign:                                                               | A unique key to ensure idempotent requests. This key should be a UUID v4 string. | 123e4567-e89b-12d3-a456-426                                                      |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |                                                                                  |

### Response

**[*operations.GetOnboardingStatusResponse](../../models/operations/getonboardingstatusresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Submit

**⚠️ We no longer recommend implementing this endpoint. Please refer to the Client Links API instead to kick off the
onboarding process for your merchants.**

Submit data that will be prefilled in the merchant's onboarding. The data you submit will only be processed when the
onboarding status is `needs-data`.  
Information that the merchant has entered in their dashboard will not be overwritten.

### Example Usage

<!-- UsageSnippet language="go" operationID="submit-onboarding-data" method="post" path="/onboarding/me" -->
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

    res, err := s.Onboarding.Submit(ctx, client.Pointer("123e4567-e89b-12d3-a456-426"), &operations.SubmitOnboardingDataRequestBody{
        Organization: &operations.Organization{
            Name: client.Pointer("Mollie B.V."),
            RegistrationNumber: client.Pointer("30204462"),
            VatNumber: client.Pointer("NL815839091B01"),
            VatRegulation: operations.VatRegulationDutch.ToPointer(),
        },
        Profile: &operations.Profile{
            Name: client.Pointer("Mollie"),
            URL: client.Pointer("https://www.mollie.com"),
            Email: client.Pointer("info@mollie.com"),
            Phone: client.Pointer("+31208202070"),
            Description: client.Pointer("Payment service provider"),
            BusinessCategory: client.Pointer("MONEY_SERVICES"),
        },
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
| `idempotencyKey`                                                                                          | **string*                                                                                                 | :heavy_minus_sign:                                                                                        | A unique key to ensure idempotent requests. This key should be a UUID v4 string.                          | 123e4567-e89b-12d3-a456-426                                                                               |
| `requestBody`                                                                                             | [*operations.SubmitOnboardingDataRequestBody](../../models/operations/submitonboardingdatarequestbody.md) | :heavy_minus_sign:                                                                                        | N/A                                                                                                       |                                                                                                           |
| `opts`                                                                                                    | [][operations.Option](../../models/operations/option.md)                                                  | :heavy_minus_sign:                                                                                        | The options for this request.                                                                             |                                                                                                           |

### Response

**[*operations.SubmitOnboardingDataResponse](../../models/operations/submitonboardingdataresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |