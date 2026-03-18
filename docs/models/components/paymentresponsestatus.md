# PaymentResponseStatus

The payment's status. Refer to the [documentation regarding statuses](https://docs.mollie.com/docs/handling-payment-status) for more info about which
statuses occur at what point.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.PaymentResponseStatusOpen

// Open enum: custom values can be created with a direct type cast
custom := components.PaymentResponseStatus("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `PaymentResponseStatusOpen`       | open                              |
| `PaymentResponseStatusPending`    | pending                           |
| `PaymentResponseStatusAuthorized` | authorized                        |
| `PaymentResponseStatusPaid`       | paid                              |
| `PaymentResponseStatusCanceled`   | canceled                          |
| `PaymentResponseStatusExpired`    | expired                           |
| `PaymentResponseStatusFailed`     | failed                            |