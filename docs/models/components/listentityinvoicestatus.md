# ListEntityInvoiceStatus

Status of the invoice.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.ListEntityInvoiceStatusOpen

// Open enum: custom values can be created with a direct type cast
custom := components.ListEntityInvoiceStatus("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `ListEntityInvoiceStatusOpen`    | open                             |
| `ListEntityInvoiceStatusPaid`    | paid                             |
| `ListEntityInvoiceStatusOverdue` | overdue                          |