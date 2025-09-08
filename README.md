# client

Developer-friendly & type-safe Go SDK specifically catered to leverage *client* API.

<div align="left">
    <a href="https://www.speakeasy.com/?utm_source=client&utm_campaign=go"><img src="https://www.speakeasy.com/assets/badges/built-by-speakeasy.svg" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>

<!-- Start Summary [summary] -->
## Summary


<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [client](#client)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
  * [Special Types](#special-types)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/mollie/mollie-api-golang
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	client "github.com/mollie/mollie-api-golang"
	"github.com/mollie/mollie-api-golang/models/components"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := client.New(
		client.WithSecurity(components.Security{
			APIKey: client.String(os.Getenv("CLIENT_API_KEY")),
		}),
	)

	res, err := s.Balances.List(ctx, client.String("EUR"), client.String("bal_gVMhHKqSSRYJyPsuoPNFH"), client.Int64(50), client.Bool(false))
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security schemes globally:

| Name     | Type   | Scheme       | Environment Variable |
| -------- | ------ | ------------ | -------------------- |
| `APIKey` | http   | HTTP Bearer  | `CLIENT_API_KEY`     |
| `OAuth`  | oauth2 | OAuth2 token | `CLIENT_O_AUTH`      |

You can set the security parameters through the `WithSecurity` option when initializing the SDK client instance. The selected scheme will be used by default to authenticate with the API for all operations that support it. For example:
```go
package main

import (
	"context"
	client "github.com/mollie/mollie-api-golang"
	"github.com/mollie/mollie-api-golang/models/components"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := client.New(
		client.WithSecurity(components.Security{
			APIKey: client.String(os.Getenv("CLIENT_API_KEY")),
		}),
	)

	res, err := s.Balances.List(ctx, client.String("EUR"), client.String("bal_gVMhHKqSSRYJyPsuoPNFH"), client.Int64(50), client.Bool(false))
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Balances](docs/sdks/balances/README.md)

* [List](docs/sdks/balances/README.md#list) - List balances
* [Get](docs/sdks/balances/README.md#get) - Get balance
* [GetPrimary](docs/sdks/balances/README.md#getprimary) - Get primary balance
* [GetReport](docs/sdks/balances/README.md#getreport) - Get balance report
* [ListTransactions](docs/sdks/balances/README.md#listtransactions) - List balance transactions

### [Capabilities](docs/sdks/capabilities/README.md)

* [List](docs/sdks/capabilities/README.md#list) - List capabilities

### [Captures](docs/sdks/captures/README.md)

* [Create](docs/sdks/captures/README.md#create) - Create capture
* [List](docs/sdks/captures/README.md#list) - List captures
* [Get](docs/sdks/captures/README.md#get) - Get capture

### [Chargebacks](docs/sdks/chargebacks/README.md)

* [List](docs/sdks/chargebacks/README.md#list) - List payment chargebacks
* [Get](docs/sdks/chargebacks/README.md#get) - Get payment chargeback
* [All](docs/sdks/chargebacks/README.md#all) - List all chargebacks


### [ClientLinks](docs/sdks/clientlinks/README.md)

* [Create](docs/sdks/clientlinks/README.md#create) - Create client link

### [Clients](docs/sdks/clients/README.md)

* [List](docs/sdks/clients/README.md#list) - List clients
* [Get](docs/sdks/clients/README.md#get) - Get client

### [Customers](docs/sdks/customers/README.md)

* [Create](docs/sdks/customers/README.md#create) - Create customer
* [List](docs/sdks/customers/README.md#list) - List customers
* [Get](docs/sdks/customers/README.md#get) - Get customer
* [Update](docs/sdks/customers/README.md#update) - Update customer
* [Delete](docs/sdks/customers/README.md#delete) - Delete customer
* [CreatePayment](docs/sdks/customers/README.md#createpayment) - Create customer payment
* [ListPayments](docs/sdks/customers/README.md#listpayments) - List customer payments

### [DelayedRouting](docs/sdks/delayedrouting/README.md)

* [Create](docs/sdks/delayedrouting/README.md#create) - Create a delayed route
* [List](docs/sdks/delayedrouting/README.md#list) - List payment routes

### [Invoices](docs/sdks/invoices/README.md)

* [List](docs/sdks/invoices/README.md#list) - List invoices
* [Get](docs/sdks/invoices/README.md#get) - Get invoice

### [Mandates](docs/sdks/mandates/README.md)

* [Create](docs/sdks/mandates/README.md#create) - Create mandate
* [List](docs/sdks/mandates/README.md#list) - List mandates
* [Get](docs/sdks/mandates/README.md#get) - Get mandate
* [Revoke](docs/sdks/mandates/README.md#revoke) - Revoke mandate

### [Methods](docs/sdks/methods/README.md)

* [List](docs/sdks/methods/README.md#list) - List payment methods
* [All](docs/sdks/methods/README.md#all) - List all payment methods
* [Get](docs/sdks/methods/README.md#get) - Get payment method

### [Onboarding](docs/sdks/onboarding/README.md)

* [Get](docs/sdks/onboarding/README.md#get) - Get onboarding status
* [Submit](docs/sdks/onboarding/README.md#submit) - Submit onboarding data

### [Organizations](docs/sdks/organizations/README.md)

* [Get](docs/sdks/organizations/README.md#get) - Get organization
* [GetCurrent](docs/sdks/organizations/README.md#getcurrent) - Get current organization
* [GetPartner](docs/sdks/organizations/README.md#getpartner) - Get partner status

### [PaymentLinks](docs/sdks/paymentlinks/README.md)

* [Create](docs/sdks/paymentlinks/README.md#create) - Create payment link
* [List](docs/sdks/paymentlinks/README.md#list) - List payment links
* [Get](docs/sdks/paymentlinks/README.md#get) - Get payment link
* [Update](docs/sdks/paymentlinks/README.md#update) - Update payment link
* [Delete](docs/sdks/paymentlinks/README.md#delete) - Delete payment link
* [ListPayments](docs/sdks/paymentlinks/README.md#listpayments) - Get payment link payments

### [Payments](docs/sdks/payments/README.md)

* [Create](docs/sdks/payments/README.md#create) - Create payment
* [List](docs/sdks/payments/README.md#list) - List payments
* [Get](docs/sdks/payments/README.md#get) - Get payment
* [Update](docs/sdks/payments/README.md#update) - Update payment
* [Cancel](docs/sdks/payments/README.md#cancel) - Cancel payment
* [ReleaseAuthorization](docs/sdks/payments/README.md#releaseauthorization) - Release payment authorization

### [Permissions](docs/sdks/permissions/README.md)

* [List](docs/sdks/permissions/README.md#list) - List permissions
* [Get](docs/sdks/permissions/README.md#get) - Get permission

### [Profiles](docs/sdks/profiles/README.md)

* [Create](docs/sdks/profiles/README.md#create) - Create profile
* [List](docs/sdks/profiles/README.md#list) - List profiles
* [Get](docs/sdks/profiles/README.md#get) - Get profile
* [Update](docs/sdks/profiles/README.md#update) - Update profile
* [Delete](docs/sdks/profiles/README.md#delete) - Delete profile
* [GetCurrent](docs/sdks/profiles/README.md#getcurrent) - Get current profile

### [Refunds](docs/sdks/refunds/README.md)

* [Create](docs/sdks/refunds/README.md#create) - Create payment refund
* [List](docs/sdks/refunds/README.md#list) - List payment refunds
* [Get](docs/sdks/refunds/README.md#get) - Get payment refund
* [Cancel](docs/sdks/refunds/README.md#cancel) - Cancel payment refund
* [All](docs/sdks/refunds/README.md#all) - List all refunds

### [SalesInvoices](docs/sdks/salesinvoices/README.md)

* [Create](docs/sdks/salesinvoices/README.md#create) - Create sales invoice
* [List](docs/sdks/salesinvoices/README.md#list) - List sales invoices
* [Get](docs/sdks/salesinvoices/README.md#get) - Get sales invoice
* [Update](docs/sdks/salesinvoices/README.md#update) - Update sales invoice
* [Delete](docs/sdks/salesinvoices/README.md#delete) - Delete sales invoice

### [Settlements](docs/sdks/settlements/README.md)

* [List](docs/sdks/settlements/README.md#list) - List settlements
* [Get](docs/sdks/settlements/README.md#get) - Get settlement
* [GetOpen](docs/sdks/settlements/README.md#getopen) - Get open settlement
* [GetNext](docs/sdks/settlements/README.md#getnext) - Get next settlement
* [ListPayments](docs/sdks/settlements/README.md#listpayments) - List settlement payments
* [ListCaptures](docs/sdks/settlements/README.md#listcaptures) - List settlement captures
* [ListRefunds](docs/sdks/settlements/README.md#listrefunds) - List settlement refunds
* [ListChargebacks](docs/sdks/settlements/README.md#listchargebacks) - List settlement chargebacks

### [Subscriptions](docs/sdks/subscriptions/README.md)

* [Create](docs/sdks/subscriptions/README.md#create) - Create subscription
* [List](docs/sdks/subscriptions/README.md#list) - List customer subscriptions
* [Get](docs/sdks/subscriptions/README.md#get) - Get subscription
* [Update](docs/sdks/subscriptions/README.md#update) - Update subscription
* [Cancel](docs/sdks/subscriptions/README.md#cancel) - Cancel subscription
* [All](docs/sdks/subscriptions/README.md#all) - List all subscriptions
* [ListPayments](docs/sdks/subscriptions/README.md#listpayments) - List subscription payments

### [Terminals](docs/sdks/terminals/README.md)

* [List](docs/sdks/terminals/README.md#list) - List terminals
* [Get](docs/sdks/terminals/README.md#get) - Get terminal

### [Wallets](docs/sdks/wallets/README.md)

* [RequestApplePaySession](docs/sdks/wallets/README.md#requestapplepaysession) - Request Apple Pay payment session

### [WebhookEvents](docs/sdks/webhookevents/README.md)

* [Get](docs/sdks/webhookevents/README.md#get) - Get a Webhook Event

### [Webhooks](docs/sdks/webhooks/README.md)

* [Create](docs/sdks/webhooks/README.md#create) - Create a webhook
* [List](docs/sdks/webhooks/README.md#list) - List all webhooks
* [Update](docs/sdks/webhooks/README.md#update) - Update a webhook
* [Get](docs/sdks/webhooks/README.md#get) - Get a webhook
* [Delete](docs/sdks/webhooks/README.md#delete) - Delete a webhook
* [Test](docs/sdks/webhooks/README.md#test) - Test a webhook

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	client "github.com/mollie/mollie-api-golang"
	"github.com/mollie/mollie-api-golang/models/components"
	"github.com/mollie/mollie-api-golang/retry"
	"log"
	"models/operations"
	"os"
)

func main() {
	ctx := context.Background()

	s := client.New(
		client.WithSecurity(components.Security{
			APIKey: client.String(os.Getenv("CLIENT_API_KEY")),
		}),
	)

	res, err := s.Balances.List(ctx, client.String("EUR"), client.String("bal_gVMhHKqSSRYJyPsuoPNFH"), client.Int64(50), client.Bool(false), operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	client "github.com/mollie/mollie-api-golang"
	"github.com/mollie/mollie-api-golang/models/components"
	"github.com/mollie/mollie-api-golang/retry"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := client.New(
		client.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		client.WithSecurity(components.Security{
			APIKey: client.String(os.Getenv("CLIENT_API_KEY")),
		}),
	)

	res, err := s.Balances.List(ctx, client.String("EUR"), client.String("bal_gVMhHKqSSRYJyPsuoPNFH"), client.Int64(50), client.Bool(false))
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `List` function may return the following errors:

| Error Type              | Status Code | Content Type         |
| ----------------------- | ----------- | -------------------- |
| apierrors.ErrorResponse | 400, 404    | application/hal+json |
| apierrors.APIError      | 4XX, 5XX    | \*/\*                |

### Example

```go
package main

import (
	"context"
	"errors"
	client "github.com/mollie/mollie-api-golang"
	"github.com/mollie/mollie-api-golang/models/apierrors"
	"github.com/mollie/mollie-api-golang/models/components"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := client.New(
		client.WithSecurity(components.Security{
			APIKey: client.String(os.Getenv("CLIENT_API_KEY")),
		}),
	)

	res, err := s.Balances.List(ctx, client.String("EUR"), client.String("bal_gVMhHKqSSRYJyPsuoPNFH"), client.Int64(50), client.Bool(false))
	if err != nil {

		var e *apierrors.ErrorResponse
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Override Server URL Per-Client

The default server can be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	client "github.com/mollie/mollie-api-golang"
	"github.com/mollie/mollie-api-golang/models/components"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := client.New(
		client.WithServerURL("https://api.mollie.com/v2"),
		client.WithSecurity(components.Security{
			APIKey: client.String(os.Getenv("CLIENT_API_KEY")),
		}),
	)

	res, err := s.Balances.List(ctx, client.String("EUR"), client.String("bal_gVMhHKqSSRYJyPsuoPNFH"), client.Int64(50), client.Bool(false))
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"

	"github.com/mollie/mollie-api-golang"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = client.New(client.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Special Types [types] -->
## Special Types

This SDK defines the following custom types to assist with marshalling and unmarshalling data.

### Date

`types.Date` is a wrapper around time.Time that allows for JSON marshaling a date string formatted as "2006-01-02".

#### Usage

```go
d1 := types.NewDate(time.Now()) // returns *types.Date

d2 := types.DateFromTime(time.Now()) // returns types.Date

d3, err := types.NewDateFromString("2019-01-01") // returns *types.Date, error

d4, err := types.DateFromString("2019-01-01") // returns types.Date, error

d5 := types.MustNewDateFromString("2019-01-01") // returns *types.Date and panics on error

d6 := types.MustDateFromString("2019-01-01") // returns types.Date and panics on error
```
<!-- End Special Types [types] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=client&utm_campaign=go)
