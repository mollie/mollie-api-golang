# SalesInvoicePaymentTermResponse

The payment term to be set on the invoice.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.SalesInvoicePaymentTermResponseSevendays

// Open enum: custom values can be created with a direct type cast
custom := components.SalesInvoicePaymentTermResponse("custom_value")
```


## Values

| Name                                                     | Value                                                    |
| -------------------------------------------------------- | -------------------------------------------------------- |
| `SalesInvoicePaymentTermResponseSevendays`               | 7 days                                                   |
| `SalesInvoicePaymentTermResponseFourteendays`            | 14 days                                                  |
| `SalesInvoicePaymentTermResponseThirtydays`              | 30 days                                                  |
| `SalesInvoicePaymentTermResponseFortyFivedays`           | 45 days                                                  |
| `SalesInvoicePaymentTermResponseSixtydays`               | 60 days                                                  |
| `SalesInvoicePaymentTermResponseNinetydays`              | 90 days                                                  |
| `SalesInvoicePaymentTermResponseOneHundredAndTwentydays` | 120 days                                                 |