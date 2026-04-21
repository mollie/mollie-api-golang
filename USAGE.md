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
			OAuth: client.Pointer(os.Getenv("CLIENT_O_AUTH")),
		}),
	)

	res, err := s.Oauth.Generate(ctx, client.Pointer("123e4567-e89b-12d3-a456-426"), nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.Body != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->