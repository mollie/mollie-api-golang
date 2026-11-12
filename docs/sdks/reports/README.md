# Reports

## Overview

### Available Operations

* [Create](#create) - Create report
* [Get](#get) - Get report

## Create

> 🚧 Beta feature
>
> This feature is currently in private beta, and the final specification may still change. To
> join the beta program, reach out to [partners@mollie.com](mailto:partners@mollie.com).

Create a report. Mollie generates it in the background, so you get a `pending` status back right away.

Poll [Get report](get-report) to check on progress, and grab a download link once it's ready.

Creating a report with the exact same parameters while one is already pending or completed doesn't create
a new one. You get the existing report back instead, whatever status it currently has.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-report" method="post" path="/v2/reports" -->
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
            AdvancedAccessToken: client.Pointer(os.Getenv("CLIENT_ADVANCED_ACCESS_TOKEN")),
        }),
    )

    res, err := s.Reports.Create(ctx, components.CreateReportRequest{
        ReportType: components.ReportTypeTransactionMargin,
        FromDate: types.MustDateFromString("2026-08-01"),
        UntilDate: types.MustDateFromString("2026-08-31"),
    }, client.Pointer("123e4567-e89b-12d3-a456-426"))
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
| `createReportRequest`                                                            | [components.CreateReportRequest](../../models/components/createreportrequest.md) | :heavy_check_mark:                                                               | N/A                                                                              |                                                                                  |
| `idempotencyKey`                                                                 | `*string`                                                                        | :heavy_minus_sign:                                                               | A unique key to ensure idempotent requests. This key should be a UUID v4 string. | 123e4567-e89b-12d3-a456-426                                                      |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |                                                                                  |

### Response

**[*operations.CreateReportResponse](../../models/operations/createreportresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 403, 422, 429      | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Get

> 🚧 Beta feature
>
> This feature is currently in private beta, and the final specification may still change. To
> join the beta program, reach out to [partners@mollie.com](mailto:partners@mollie.com).

Retrieve a single report by its ID. Poll this to check on a report you created, and to grab a download
link once it's ready.

The download link is pre-signed and freshly generated on every call, and expires 15 minutes after you get
it. If it expires before you download the file, just call this endpoint again for a new one. Expiry of
the link doesn't affect the report itself.

### Example Usage: get-report-200-completed

<!-- UsageSnippet language="go" operationID="get-report" method="get" path="/v2/reports/{reportId}" example="get-report-200-completed" -->
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

    res, err := s.Reports.Get(ctx, "rpt_XRAG3v9EUyCDcX99pJ2Vc", client.Pointer("123e4567-e89b-12d3-a456-426"))
    if err != nil {
        log.Fatal(err)
    }
    if res.EntityReport != nil {
        // handle response
    }
}
```
### Example Usage: get-report-200-failed

<!-- UsageSnippet language="go" operationID="get-report" method="get" path="/v2/reports/{reportId}" example="get-report-200-failed" -->
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

    res, err := s.Reports.Get(ctx, "rpt_XRAG3v9EUyCDcX99pJ2Vc", client.Pointer("123e4567-e89b-12d3-a456-426"))
    if err != nil {
        log.Fatal(err)
    }
    if res.EntityReport != nil {
        // handle response
    }
}
```
### Example Usage: get-report-200-pending

<!-- UsageSnippet language="go" operationID="get-report" method="get" path="/v2/reports/{reportId}" example="get-report-200-pending" -->
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

    res, err := s.Reports.Get(ctx, "rpt_XRAG3v9EUyCDcX99pJ2Vc", client.Pointer("123e4567-e89b-12d3-a456-426"))
    if err != nil {
        log.Fatal(err)
    }
    if res.EntityReport != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |                                                                                  |
| `reportID`                                                                       | `string`                                                                         | :heavy_check_mark:                                                               | Provide the ID of the related report.                                            | rpt_XRAG3v9EUyCDcX99pJ2Vc                                                        |
| `idempotencyKey`                                                                 | `*string`                                                                        | :heavy_minus_sign:                                                               | A unique key to ensure idempotent requests. This key should be a UUID v4 string. | 123e4567-e89b-12d3-a456-426                                                      |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |                                                                                  |

### Response

**[*operations.GetReportResponse](../../models/operations/getreportresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 403, 404, 429      | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |