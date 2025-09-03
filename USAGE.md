<!-- Start SDK Example Usage [usage] -->
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