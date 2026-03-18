# SalesInvoiceRecipientTypeResponse

The type of recipient, either `consumer` or `business`. This will determine what further fields are
required on the `recipient` object.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.SalesInvoiceRecipientTypeResponseConsumer

// Open enum: custom values can be created with a direct type cast
custom := components.SalesInvoiceRecipientTypeResponse("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `SalesInvoiceRecipientTypeResponseConsumer` | consumer                                    |
| `SalesInvoiceRecipientTypeResponseBusiness` | business                                    |