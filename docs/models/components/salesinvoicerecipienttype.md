# SalesInvoiceRecipientType

The type of recipient, either `consumer` or `business`. This will determine what further fields are
required on the `recipient` object.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.SalesInvoiceRecipientTypeConsumer
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `SalesInvoiceRecipientTypeConsumer` | consumer                            |
| `SalesInvoiceRecipientTypeBusiness` | business                            |