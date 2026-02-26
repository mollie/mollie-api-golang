# BalanceTransfers

## Overview

### Available Operations

* [Create](#create) - Create a Connect balance transfer
* [List](#list) - List all Connect balance transfers
* [Get](#get) - Get a Connect balance transfer

## Create

This API endpoint allows you to create a balance transfer from your organization's balance to a connected organization's balance, or vice versa.
You can also create a balance transfer between two connected organizations.
To create a balance transfer, you must be authenticated as the source organization, and the destination organization must be a connected organization
that has authorized the `balance-transfers.write` scope for your organization.

### Example Usage: create-balance-transfer-200-1

<!-- UsageSnippet language="go" operationID="create-connect-balance-transfer" method="post" path="/connect/balance-transfers" example="create-balance-transfer-200-1" -->
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

    res, err := s.BalanceTransfers.Create(ctx, client.Pointer("123e4567-e89b-12d3-a456-426"), &components.EntityBalanceTransfer{
        Amount: components.Amount{
            Currency: "EUR",
            Value: "10.00",
        },
        Source: components.EntityBalanceTransferParty{
            Type: components.BalanceTransferPartyTypeOrganization,
            ID: "org_1234567",
            Description: "Invoice fee",
        },
        Destination: components.EntityBalanceTransferParty{
            Type: components.BalanceTransferPartyTypeOrganization,
            ID: "org_1234567",
            Description: "Invoice fee",
        },
        Description: "Invoice fee",
        Category: components.BalanceTransferCategoryInvoiceCollection.ToPointer(),
        Metadata: map[string]any{
            "order_id": 12345,
            "customer_id": 9876,
        },
        Testmode: client.Pointer(false),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EntityBalanceTransferResponse != nil {
        // handle response
    }
}
```
### Example Usage: create-balance-transfer-422-1

<!-- UsageSnippet language="go" operationID="create-connect-balance-transfer" method="post" path="/connect/balance-transfers" example="create-balance-transfer-422-1" -->
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

    res, err := s.BalanceTransfers.Create(ctx, client.Pointer("123e4567-e89b-12d3-a456-426"), &components.EntityBalanceTransfer{
        Amount: components.Amount{
            Currency: "EUR",
            Value: "10.00",
        },
        Source: components.EntityBalanceTransferParty{
            Type: components.BalanceTransferPartyTypeOrganization,
            ID: "org_1234567",
            Description: "Invoice fee",
        },
        Destination: components.EntityBalanceTransferParty{
            Type: components.BalanceTransferPartyTypeOrganization,
            ID: "org_1234567",
            Description: "Invoice fee",
        },
        Description: "Invoice fee",
        Category: components.BalanceTransferCategoryInvoiceCollection.ToPointer(),
        Metadata: map[string]any{
            "order_id": 12345,
            "customer_id": 9876,
        },
        Testmode: client.Pointer(false),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.EntityBalanceTransferResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                             | Type                                                                                  | Required                                                                              | Description                                                                           | Example                                                                               |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `ctx`                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                 | :heavy_check_mark:                                                                    | The context to use for the request.                                                   |                                                                                       |
| `idempotencyKey`                                                                      | **string*                                                                             | :heavy_minus_sign:                                                                    | A unique key to ensure idempotent requests. This key should be a UUID v4 string.      | 123e4567-e89b-12d3-a456-426                                                           |
| `entityBalanceTransfer`                                                               | [*components.EntityBalanceTransfer](../../models/components/entitybalancetransfer.md) | :heavy_minus_sign:                                                                    | N/A                                                                                   |                                                                                       |
| `opts`                                                                                | [][operations.Option](../../models/operations/option.md)                              | :heavy_minus_sign:                                                                    | The options for this request.                                                         |                                                                                       |

### Response

**[*operations.CreateConnectBalanceTransferResponse](../../models/operations/createconnectbalancetransferresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 422                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## List

Returns a paginated list of balance transfers associated with your organization. These may be a balance transfer that was received or sent from your balance, or a balance transfer that you initiated on behalf of your clients. If no balance transfers are available, the resulting array will be empty. This request should never throw an error.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-connect-balance-transfers" method="get" path="/connect/balance-transfers" example="list-balance-transfer-200-1" -->
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

    res, err := s.BalanceTransfers.List(ctx, operations.ListConnectBalanceTransfersRequest{
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.ListConnectBalanceTransfersRequest](../../models/operations/listconnectbalancetransfersrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.ListConnectBalanceTransfersResponse](../../models/operations/listconnectbalancetransfersresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## Get

Retrieve a single Connect balance transfer object by its ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-connect-balance-transfer" method="get" path="/connect/balance-transfers/{balanceTransferId}" example="get-balance-transfer-200-1" -->
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

    res, err := s.BalanceTransfers.Get(ctx, "cbtr_j8NvRAM2WNZtsykpLEX8J", client.Pointer("123e4567-e89b-12d3-a456-426"))
    if err != nil {
        log.Fatal(err)
    }
    if res.EntityBalanceTransferResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                               | Type                                                                                                                                                                    | Required                                                                                                                                                                | Description                                                                                                                                                             | Example                                                                                                                                                                 |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                   | :heavy_check_mark:                                                                                                                                                      | The context to use for the request.                                                                                                                                     |                                                                                                                                                                         |
| `balanceTransferID`                                                                                                                                                     | *string*                                                                                                                                                                | :heavy_check_mark:                                                                                                                                                      | Provide the ID of the related balance transfer.                                                                                                                         | cbtr_j8NvRAM2WNZtsykpLEX8J                                                                                                                                              |
| `testmode`                                                                                                                                                              | **bool*                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                      | You can enable test mode by setting the `testmode` query parameter to `true`.<br/><br/>Test entities cannot be retrieved when the endpoint is set to live mode, and vice versa. |                                                                                                                                                                         |
| `idempotencyKey`                                                                                                                                                        | **string*                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                      | A unique key to ensure idempotent requests. This key should be a UUID v4 string.                                                                                        | 123e4567-e89b-12d3-a456-426                                                                                                                                             |
| `opts`                                                                                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                                                                                | :heavy_minus_sign:                                                                                                                                                      | The options for this request.                                                                                                                                           |                                                                                                                                                                         |

### Response

**[*operations.GetConnectBalanceTransferResponse](../../models/operations/getconnectbalancetransferresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 404                     | application/hal+json    |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |