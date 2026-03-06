# SalesInvoicePaymentDetailsSource

The way through which the invoice is to be set to paid.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.SalesInvoicePaymentDetailsSourceManual
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `SalesInvoicePaymentDetailsSourceManual`      | manual                                        |
| `SalesInvoicePaymentDetailsSourcePaymentLink` | payment-link                                  |
| `SalesInvoicePaymentDetailsSourcePayment`     | payment                                       |