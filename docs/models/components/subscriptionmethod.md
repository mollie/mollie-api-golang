# SubscriptionMethod

The payment method used for this subscription. If omitted, any of the customer's valid mandates may be used.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.SubscriptionMethodCreditcard
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `SubscriptionMethodCreditcard`  | creditcard                      |
| `SubscriptionMethodDirectdebit` | directdebit                     |
| `SubscriptionMethodPaypal`      | paypal                          |