<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	client "github.com/mollie/mollie-api-golang"
	"github.com/mollie/mollie-api-golang/models/components"
	"github.com/mollie/mollie-api-golang/models/operations"
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

	res, err := s.Balances.List(ctx, operations.ListBalancesRequest{
		Currency:       client.String("EUR"),
		From:           client.String("bal_gVMhHKqSSRYJyPsuoPNFH"),
		Limit:          client.Int64(50),
		Testmode:       client.Bool(false),
		IdempotencyKey: client.String("123e4567-e89b-12d3-a456-426"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Object != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->