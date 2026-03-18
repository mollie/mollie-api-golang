# ListPaymentResponseStatus

The payment's status. Refer to the [documentation regarding statuses](https://docs.mollie.com/docs/handling-payment-status) for more info about which
statuses occur at what point.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.ListPaymentResponseStatusOpen

// Open enum: custom values can be created with a direct type cast
custom := components.ListPaymentResponseStatus("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `ListPaymentResponseStatusOpen`       | open                                  |
| `ListPaymentResponseStatusPending`    | pending                               |
| `ListPaymentResponseStatusAuthorized` | authorized                            |
| `ListPaymentResponseStatusPaid`       | paid                                  |
| `ListPaymentResponseStatusCanceled`   | canceled                              |
| `ListPaymentResponseStatusExpired`    | expired                               |
| `ListPaymentResponseStatusFailed`     | failed                                |