# SalesInvoicePaymentDetailsSourceResponse

The way through which the invoice is to be set to paid.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.SalesInvoicePaymentDetailsSourceResponseManual

// Open enum: custom values can be created with a direct type cast
custom := components.SalesInvoicePaymentDetailsSourceResponse("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `SalesInvoicePaymentDetailsSourceResponseManual`      | manual                                                |
| `SalesInvoicePaymentDetailsSourceResponsePaymentLink` | payment-link                                          |
| `SalesInvoicePaymentDetailsSourceResponsePayment`     | payment                                               |