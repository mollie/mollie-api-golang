# EntityInvoiceStatus

Status of the invoice.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.EntityInvoiceStatusOpen

// Open enum: custom values can be created with a direct type cast
custom := components.EntityInvoiceStatus("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `EntityInvoiceStatusOpen`    | open                         |
| `EntityInvoiceStatusPaid`    | paid                         |
| `EntityInvoiceStatusOverdue` | overdue                      |