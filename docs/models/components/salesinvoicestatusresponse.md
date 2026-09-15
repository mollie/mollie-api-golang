# SalesInvoiceStatusResponse

The current status of the invoice.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.SalesInvoiceStatusResponseDraft

// Open enum: custom values can be created with a direct type cast
custom := components.SalesInvoiceStatusResponse("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `SalesInvoiceStatusResponseDraft`           | draft                                       |
| `SalesInvoiceStatusResponseIssuing`         | issuing                                     |
| `SalesInvoiceStatusResponseIssued`          | issued                                      |
| `SalesInvoiceStatusResponsePendingPayment`  | pending-payment                             |
| `SalesInvoiceStatusResponsePaid`            | paid                                        |
| `SalesInvoiceStatusResponseOverdue`         | overdue                                     |
| `SalesInvoiceStatusResponsePaymentReversed` | payment_reversed                            |
| `SalesInvoiceStatusResponseCancelled`       | cancelled                                   |
| `SalesInvoiceStatusResponseExpired`         | expired                                     |
| `SalesInvoiceStatusResponseFailed`          | failed                                      |