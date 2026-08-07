# GetOpenSettlementStatus

The status of the settlement.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/operations"
)

value := operations.GetOpenSettlementStatusOpen

// Open enum: custom values can be created with a direct type cast
custom := operations.GetOpenSettlementStatus("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `GetOpenSettlementStatusOpen`             | open                                      |
| `GetOpenSettlementStatusPending`          | pending                                   |
| `GetOpenSettlementStatusProcessingAtBank` | processing-at-bank                        |
| `GetOpenSettlementStatusPaidout`          | paidout                                   |
| `GetOpenSettlementStatusFailed`           | failed                                    |