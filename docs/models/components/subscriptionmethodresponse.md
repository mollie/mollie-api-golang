# SubscriptionMethodResponse

The payment method used for this subscription. If omitted, any of the customer's valid mandates may be used.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.SubscriptionMethodResponseCreditcard

// Open enum: custom values can be created with a direct type cast
custom := components.SubscriptionMethodResponse("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `SubscriptionMethodResponseCreditcard`  | creditcard                              |
| `SubscriptionMethodResponseDirectdebit` | directdebit                             |
| `SubscriptionMethodResponsePaypal`      | paypal                                  |