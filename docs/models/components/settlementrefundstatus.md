# SettlementRefundStatus

The refund's status. Settlement refunds always have a status of `refunded`.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.SettlementRefundStatusRefunded

// Open enum: custom values can be created with a direct type cast
custom := components.SettlementRefundStatus("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `SettlementRefundStatusRefunded` | refunded                         |