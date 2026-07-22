# SessionRequiredCustomerDetailsResponse

Customer details that should be collected during checkout.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.SessionRequiredCustomerDetailsResponseEmail

// Open enum: custom values can be created with a direct type cast
custom := components.SessionRequiredCustomerDetailsResponse("custom_value")
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `SessionRequiredCustomerDetailsResponseEmail`           | email                                                   |
| `SessionRequiredCustomerDetailsResponseBillingAddress`  | billing-address                                         |
| `SessionRequiredCustomerDetailsResponseShippingAddress` | shipping-address                                        |