# Invoices
(*Invoices*)

## Overview

### Available Operations

* [List](#list) - List invoices
* [Get](#get) - Get invoice

## List

Retrieve a list of all your invoices, optionally filtered by year or by
invoice reference.

The results are paginated.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-invoices" method="get" path="/invoices" -->
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

    res, err := s.Invoices.List(ctx, operations.ListInvoicesRequest{
        Reference: client.String("2024.10000"),
        Year: client.String("2024"),
        Month: client.String("01"),
        From: client.String("inv_xBEbP9rvAq"),
        Limit: client.Int64(50),
        Sort: operations.ListInvoicesSortDesc.ToPointer(),
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
| `request`                                                                        | [operations.ListInvoicesRequest](../../models/operations/listinvoicesrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.ListInvoicesResponse](../../models/operations/listinvoicesresponse.md), error**

### Errors

| Error Type                                   | Status Code                                  | Content Type                                 |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| apierrors.ListInvoicesBadRequestHalJSONError | 400                                          | application/hal+json                         |
| apierrors.ListInvoicesNotFoundHalJSONError   | 404                                          | application/hal+json                         |
| apierrors.APIError                           | 4XX, 5XX                                     | \*/\*                                        |

## Get

Retrieve a single invoice by its ID.

If you want to retrieve the details of an invoice by its invoice number,
call the [List invoices](list-invoices) endpoint with the `reference` parameter.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-invoice" method="get" path="/invoices/{id}" -->
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

    res, err := s.Invoices.Get(ctx, "inv_FrvewDA3Pr")
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                         | Type                                                              | Required                                                          | Description                                                       | Example                                                           |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `ctx`                                                             | [context.Context](https://pkg.go.dev/context#Context)             | :heavy_check_mark:                                                | The context to use for the request.                               |                                                                   |
| `id`                                                              | *string*                                                          | :heavy_check_mark:                                                | Provide the ID of the item you want to perform this operation on. | inv_FrvewDA3Pr                                                    |
| `opts`                                                            | [][operations.Option](../../models/operations/option.md)          | :heavy_minus_sign:                                                | The options for this request.                                     |                                                                   |

### Response

**[*operations.GetInvoiceResponse](../../models/operations/getinvoiceresponse.md), error**

### Errors

| Error Type                       | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| apierrors.GetInvoiceHalJSONError | 404                              | application/hal+json             |
| apierrors.APIError               | 4XX, 5XX                         | \*/\*                            |