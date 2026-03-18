# ListEntitySettlementStatus

The status of the settlement.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/components"
)

value := components.ListEntitySettlementStatusOpen

// Open enum: custom values can be created with a direct type cast
custom := components.ListEntitySettlementStatus("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `ListEntitySettlementStatusOpen`    | open                                |
| `ListEntitySettlementStatusPending` | pending                             |
| `ListEntitySettlementStatusPaidout` | paidout                             |
| `ListEntitySettlementStatusFailed`  | failed                              |