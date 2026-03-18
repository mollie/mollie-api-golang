# EntitySettlementStatus

The status of the settlement.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.EntitySettlementStatusOpen

// Open enum: custom values can be created with a direct type cast
custom := components.EntitySettlementStatus("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `EntitySettlementStatusOpen`    | open                            |
| `EntitySettlementStatusPending` | pending                         |
| `EntitySettlementStatusPaidout` | paidout                         |
| `EntitySettlementStatusFailed`  | failed                          |