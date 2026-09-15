# GetNextSettlementStatus

The status of the settlement.

## Example Usage

```go
import (
	"github.com/mollie/mollie-api-golang/models/operations"
)

value := operations.GetNextSettlementStatusOpen

// Open enum: custom values can be created with a direct type cast
custom := operations.GetNextSettlementStatus("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `GetNextSettlementStatusOpen`             | open                                      |
| `GetNextSettlementStatusPending`          | pending                                   |
| `GetNextSettlementStatusProcessing`       | processing                                |
| `GetNextSettlementStatusProcessingAtBank` | processing-at-bank                        |
| `GetNextSettlementStatusPaidout`          | paidout                                   |
| `GetNextSettlementStatusFailed`           | failed                                    |