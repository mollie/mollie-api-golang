# SettlementPaymentStatus

The payment's status. Settlement payments always have a status of `paid`.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.SettlementPaymentStatusPaid

// Open enum: custom values can be created with a direct type cast
custom := components.SettlementPaymentStatus("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `SettlementPaymentStatusPaid` | paid                          |