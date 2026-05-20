<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	client "github.com/mollie/mollie-api-golang"
	"github.com/mollie/mollie-api-golang/models/components"
	"github.com/mollie/mollie-api-golang/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := client.New()

	res, err := s.Oauth.Generate(ctx, operations.OauthGenerateTokensSecurity{
		Username: "",
		Password: "",
	}, client.Pointer("123e4567-e89b-12d3-a456-426"), &operations.OauthGenerateTokensRequestBody{
		GrantType:    components.OauthGrantTypeAuthorizationCode,
		Code:         client.Pointer("auth_..."),
		RefreshToken: client.Pointer("refresh_..."),
		RedirectURI:  client.Pointer("https://example.com/redirect"),
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