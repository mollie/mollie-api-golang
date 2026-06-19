# SalesInvoiceEInvoiceStatus

The e-invoice submission status for the invoice, if it was configured to be an e-invoice.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.SalesInvoiceEInvoiceStatusIssuing

// Open enum: custom values can be created with a direct type cast
custom := components.SalesInvoiceEInvoiceStatus("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `SalesInvoiceEInvoiceStatusIssuing` | issuing                             |
| `SalesInvoiceEInvoiceStatusIssued`  | issued                              |
| `SalesInvoiceEInvoiceStatusFailed`  | failed                              |