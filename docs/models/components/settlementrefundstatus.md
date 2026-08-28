# SettlementRefundStatus

The refund's status. Settlement refunds are normally `refunded`, but can be `failed` if the refund
could not be processed.

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
| `SettlementRefundStatusFailed`   | failed                           |